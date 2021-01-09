FROM golang:1.15.3

RUN mkdir -p /go/src/github.com/emadghaffari/virgool
WORKDIR /go/src/github.com/emadghaffari/virgool
ADD . .
RUN go get -d -v ./...
RUN go install -v ./...
RUN apt update


RUN go get  github.com/canthefason/go-watcher
RUN go install github.com/canthefason/go-watcher/cmd/watcher
ENTRYPOINT  watcher -run github.com/emadghaffari/virgool/notification  -watch github.com/emadghaffari/virgool/notification    