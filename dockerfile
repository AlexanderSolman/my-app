FROM golang:1.21.6-alpine AS builder

WORKDIR /app
ADD . /app/

RUN go build -o main .

FROM golang:1.21.6-alpine as app
ENTRYPOINT ["/app/main"]
WORKDIR /app
COPY --from=builder /app/main /app/main
