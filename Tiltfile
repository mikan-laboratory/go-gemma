# Nginx and Redis
docker_compose('docker-compose.yml')

# Run server
local_resource(
    name="server",
    serve_cmd="air",
    resource_deps=['build-model', 'redis'],
    auto_init=True
)
