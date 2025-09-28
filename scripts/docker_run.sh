docker run -d -v ./ssl:/etc/ssl -v ./stats:/app/stats -v ./logs:/app/logs -p 443:8080 --name httpinfo-dev $DOCKER_DEV_IMAGE_NAME
