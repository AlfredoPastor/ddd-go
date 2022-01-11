##
## go-builder
##

FROM golang:1.17.6

WORKDIR /go/src/github.com/AlfredoPastor/ddd-go

ADD go.* ./

RUN go get -d -v ./...
RUN go install -v ./...

ADD shared shared