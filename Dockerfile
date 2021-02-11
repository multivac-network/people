FROM golang:1.15-alpine
ENV GO111MODULE on
COPY ./ /go/workspace/
RUN go install ./workspace

ENTRYPOINT ["$GOBIN/service"]