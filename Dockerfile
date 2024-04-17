# Step 1: Build Stage
FROM debian:12.5-slim as builder

WORKDIR /app

# Install necessary tools including clang for better handling of SIMD operations
RUN apt-get update && apt-get install -y \
    cmake \
    make \
    python3 \
    python3-pip \
    python3-venv \
    unzip \
    git \
    curl \
    clang-16 \
    libclang-16-dev \
    && rm -rf /var/lib/apt/lists/*

# Download and install Go
RUN curl -LO "https://golang.org/dl/go1.22.2.linux-amd64.tar.gz" \
    && tar -xzf go1.22.2.linux-amd64.tar.gz -C /usr/local \
    && rm go1.22.2.linux-amd64.tar.gz

# Set the Go environment paths
ENV GOROOT=/usr/local/go
ENV GOPATH=/app/go
ENV PATH=$GOROOT/bin:$GOPATH/bin:$PATH

# Create a virtual environment and activate it
RUN python3 -m venv venv
ENV PATH="/app/venv/bin:$PATH"

# Install python packages
RUN pip install gdown

# Copy necessary files
COPY go.mod go.sum CMakeLists.txt ./

# Download Go dependencies
RUN go mod download

# Download and unzip necessary libraries
RUN gdown --id 1Blx_O2FWV2-h71uGia0wtRb-5IaDwRX_ -O gemma-libs.zip && \
    unzip gemma-libs.zip -d ./

# Set clang as the default compiler for all C/C++ builds including CGO
RUN update-alternatives --install /usr/bin/cc cc /usr/bin/clang-16 60 \
    && update-alternatives --install /usr/bin/c++ c++ /usr/bin/clang++-16 60

# Configure and build your application
RUN cmake -B build -DCMAKE_C_COMPILER=clang-16 -DCMAKE_CXX_COMPILER=clang++-16 && \
    cp ./gemma-libs/2b-it-sfp.sbs build/2b-it-sfp.sbs && \
    cp ./gemma-libs/tokenizer.spm build/tokenizer.spm && \
    make -C build gemma

# Set CGO to use Clang explicitly
ENV CC=clang-16
ENV CXX=clang++-16
ENV CGO_ENABLED=1

COPY . .

# Build the Go application
RUN go build -a -installsuffix cgo -o go-gemma .

# Step 2: Final Stage
FROM debian:12.5-slim

WORKDIR /app

# Install runtime dependencies
RUN apt-get update && \
    apt-get install -y redis-server \
    gcc \
    libc6-dev \
    libsqlite3-dev \
    && rm -rf /var/lib/apt/lists/*

# Copy the Go binary and the build directory from the builder stage
COPY --from=builder /app/go-gemma ./go-gemma
COPY --from=builder /app/build ./build

# Expose the port on which your Go application listens
EXPOSE 8081

# Use a shell script to start both Redis and your Go application
COPY start.sh .

RUN chmod +x ./build/_deps/gemma-build/gemma
RUN chmod +x start.sh

CMD ["./start.sh"]
