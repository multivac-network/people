FROM golang:1.15-alpine

RUN cd /
RUN go install

RUN $GOBIN/service