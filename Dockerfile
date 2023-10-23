FROM golang:latest

WORKDIR /app
COPY main.go go.mod go.sum ApiStructs.go ./

CMD ["go", "run", "main.go", "10000"]
