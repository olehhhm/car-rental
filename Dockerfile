FROM golang:latest

WORKDIR /app

COPY ./ /app

RUN ["apt-get", "update"]
RUN ["apt-get", "install", "-y", "vim"]
RUN ["go", "mod", "download"]
RUN ["go", "get", "github.com/githubnemo/CompileDaemon"]

ENTRYPOINT CompileDaemon -polling -log-prefix=false -build="go build cmd/api.go" -command="./api"