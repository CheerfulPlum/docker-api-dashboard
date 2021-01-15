# Docker Dashboard API

The start of an Docker container dashboard in a similar vein to what Portainer does. Used as a learning project. Using it to integrate with nginx-proxy(https://github.com/nginx-proxy/nginx-proxy). Such as displaying VIRTUAL_HOST and VIRTUAL_PORT

### Endpoints:

GET /containers - A list of containers

GET /containers/<CONTAINER_ID> - Info about a specific container

GET /containers/<CONTAINER_ID>/health - Check if the container is healthy or not, based on Docker healthcheck


### Usage:
```
docker-compose up

Go to port 3000 with any of the above endpoints
```
