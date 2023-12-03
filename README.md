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
`docker build -t agent .`

### Run the Container
`docker run -e N2YO_KEY=<key> -e AWS_SECRET_ACCESS_KEY=<key> -e AWS_ACCESS_KEY_ID=<key> -e LOGGLY_TOKEN=<token> agent`

### Stop Container
Get the running container name:
`docker ps`
Then:
`docker stop <container name>`

### Clear Unused Containers
`docker rm $(sudo docker ps -aq)`

### Clear Unused Images
`docker rmi  $(sudo docker images -q)`