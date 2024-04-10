# Nginx and Redis
docker_compose('docker-compose.yml')

# Libs
local_resource(
    name='libs'
    cmd='''
sh -c "
if [ ! -d libs ]; then
  mkdir -p libs && \
  curl -o libs/gemma.zip https://drive.google.com/file/d/1JLdITj5WH7kxCUBH4i638MgOlsPpXh4l/view?usp=sharing && \
  unzip libs/gemma.zip -d libs && \
  rm -f libs/gemma.zip
fi
"
''',
auto_init=True    
)

# Build model
local_resource(
    name='build-model',
    cmd='''
sh -c "
if [ ! -d build ]; then
    cmake -B build && \
    cp libs/2b-it-sfp.sbs build/ && \
    cp libs/tokenizer.spm build/ && \
    make -C build gemma
fi
"
''',
auto_init=True
)

# Run server
local_resource(
    name="server",
    serve_cmd="air",
    resource_deps=['build-model', 'redis', 'nginx'],
    auto_init=True
)
