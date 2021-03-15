# builder compiles the go executable using no C bindings to allow for scratch container
FROM repath/golang-base:1.0.0 as builder
COPY . .
ENV GOPRIVATE github.com/repath-io/*
ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOARCH amd64

RUN unset GOPATH && go mod download && go build -o /service

# output of the build is placed into the root of the scratch folder and declared as entrypoint
FROM alpine:latest
COPY --from=builder /service /service
ENTRYPOINT ["sh", "-c", "/service"]