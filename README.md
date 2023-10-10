# Go Polling Agent

## Docker Commands
Run these in the same directory as the Dockerfile.

### Build the Image
`sudo docker build -t poller_test .`

### Run the Container
`sudo docker run -e N2YO_KEY=<key> -e LOGGLY_TOKEN=<token> poller_test`

### Stop Container
Get the running container name:
`sudo docker ps`
Then:
`sudo docker stop <container name>`

### Clear Unused Containers
`sudo docker rm $(sudo docker ps -aq)`

### Clear Unused Images
`sudo docker rmi  $(sudo docker images -q)`