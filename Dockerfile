# builder compiles the go executable using no C bindings to allow for scratch container
FROM golang:1.16 as builder
COPY . .
RUN unset GOPATH && go mod download && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /service

# output of the build is placed into the root of the scratch folder and declared as entrypoint
FROM alpine:latest
COPY --from=builder /service /service
ENTRYPOINT ["sh", "-c", "/service"]