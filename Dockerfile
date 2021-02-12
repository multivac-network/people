FROM golang:1.15-alpine
ENV GO111MODULE auto
COPY ./ ./workspace/
RUN go install ./workspace

ENTRYPOINT ["$GOBIN/service"]