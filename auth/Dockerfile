FROM golang

RUN mkdir -p /go/src/github.com/emadghaffari/virgool

ADD . /go/src/github.com/emadghaffari/virgool

RUN go get  -t -v ./...
RUN go get  github.com/canthefason/go-watcher
RUN go install github.com/canthefason/go-watcher/cmd/watcher

ENTRYPOINT  watcher -run github.com/emadghaffari/virgool/auth/cmd  -watch github.com/emadghaffari/virgool/auth
