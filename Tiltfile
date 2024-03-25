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
    make -C build gemma
fi
"
''',
auto_init=True
)

# Run server
local_resource(
    name="server",
    cmd="air",
    resource_deps=['build-model'],
    auto_init=True
)
