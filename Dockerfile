FROM golang:1.15 as builder

COPY . .
RUN go build -o service

FROM scratch
COPY --from=builder . .
ENTRYPOINT ["service"]