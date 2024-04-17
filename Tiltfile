# Nginx and Redis
docker_compose('docker-compose.yml')

# Build model
local_resource(
    name='build-model',
    cmd='''
sh -c "
if [ ! -d build ]; then
    cmake -B build && \
    cp libs/2b-it-sfp.sbs build/ && \
    cp libs/tokenizer.spm build/ && \
    chmod +x build/_deps/gemma-build/gemma && \
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
    resource_deps=['build-model', 'redis'],
    auto_init=True
)
