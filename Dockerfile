FROM golang:1.15-alpine

RUN cd /workspace

RUN go install

RUN $GOBIN/service