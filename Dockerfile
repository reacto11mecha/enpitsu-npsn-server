FROM golang:1.25-alpine AS builder

WORKDIR /app
COPY go.mod ./
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o npsn-server main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/npsn-server .

CMD ["./npsn-server"]
