FROM golang:1.16

WORKDIR /go/src/github.com/noellimx/gogopower

COPY go.mod ./

COPY *.go ./

RUN go mod download

COPY makefile ./

RUN make build

CMD make run-binary

