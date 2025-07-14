FROM golang:1.24-alpine as builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main ./cmd/server/main.go

FROM alpine:latest
WORKDIR /app

COPY --from=builder /app/main .
COPY --from=builder /app/config.yml .
COPY web ./web

EXPOSE 8080

CMD ["./main"]
