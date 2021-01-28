FROM golang:1.15.3

RUN mkdir -p /go/src/github.com/emadghaffari/virgool
WORKDIR /go/src/github.com/emadghaffari/virgool
ADD . .
RUN go get -d -v ./...
RUN go install -v ./...
RUN apt update

# binary will be $(go env GOPATH)/bin/air
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

# or install it into ./bin/
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s


WORKDIR /go/src/github.com/emadghaffari/virgool/notification
# RUN air

CMD ["./run.sh"]


# RUN go get  github.com/canthefason/go-watcher
# RUN go install github.com/canthefason/go-watcher/cmd/watcher
# ENTRYPOINT  watcher -run github.com/emadghaffari/virgool/notification  -watch github.com/emadghaffari/virgool/notification        