FROM golang:1.15-alpine

COPY ./ /workspace/
RUN go install /workspace

RUN $GOBIN/service