FROM golang:1.15-alpine

RUN mkdir /workspace
RUN cd /workspace

RUN go install

RUN $GOBIN/service