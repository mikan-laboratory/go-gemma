# Step 1: Build Stage
FROM golang:1.22 as builder

WORKDIR /app

# Install necessary tools
RUN apt-get update && apt-get install -y \
    unzip \
    python3 \
    python3-venv \
    && rm -rf /var/lib/apt/lists/*

# Create a virtual environment and activate it
RUN python3 -m venv venv
ENV PATH="/app/venv/bin:$PATH"

# Install gdown with pip
RUN pip install gdown

# Copy go.mod, go.sum, and CMakeLists.txt
COPY go.mod go.sum ./

# Download Go dependencies
RUN go mod download

# Download and prepare the model
RUN gdown --id 1UexYG4stAjwyryQZxckJDxyi5A0w7WjN -O gemma-libs.zip && \
    unzip gemma-libs.zip -d ./

# Assume your model and tokenizer are now directly in the working directory, adjust as necessary
RUN mkdir -p build && \
    cp ./gemma-libs/2b-it-sfp.sbs build/2b-it-sfp.sbs && \
    cp ./gemma-libs/tokenizer.spm build/tokenizer.spm && \
    cp ./gemma-libs/gemma build/gemma

# Copy the rest of your source code and build the Go application
COPY . .

# Build the Go application
RUN go build -a -installsuffix cgo -o go-gemma .

# Step 2: Final Stage
FROM debian:12.5-slim

WORKDIR /app

# Add Redis
RUN apt-get update && \
    apt-get install -y redis-server \
    gcc \
    libc6-dev \
    libsqlite3-dev && \
    rm -rf /var/lib/apt/lists/*

# Copy the Go binary and the build directory from the builder stage
COPY --from=builder /app/go-gemma ./go-gemma
COPY --from=builder /app/build ./build

# Expose the port on which your Go application listens
EXPOSE 8081

# Use a shell script to start both Redis and your Go application
COPY start.sh .
RUN chmod +x ./build/gemma
RUN chmod +x start.sh

CMD ["./start.sh"]
