FROM golang:1.15-alpine

RUN go install ./

RUN $GOBIN/service