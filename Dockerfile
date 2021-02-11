FROM golang:1.15-alpine

COPY ./ /go/workspace/
RUN go install ./workspace

ENTRYPOINT ["$GOBIN/service"]