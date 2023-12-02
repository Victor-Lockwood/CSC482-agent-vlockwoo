# Go Polling Agent

## About
This is a program that calls the [N2YO.com's REST API](https://www.n2yo.com/api/) at specified intervals, retrieving
position data on the International Space Station relative to the coordinates of Rice Creek Field Station
at SUNY Oswego.  The results are flattened and then stored in a DynamoDB table on AWS.

The server component of this, which opens up endpoints for data retrieval from that database,
can be found [here](https://github.com/Victor-Lockwood/server-vlockwoo).

## Docker Commands for Image Building and Deployment
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