FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /out/api ./cmd/api
RUN CGO_ENABLED=0 GOOS=linux go build -o /out/devtoken ./cmd/devtoken

FROM alpine:3.20

WORKDIR /app

COPY --from=builder /out/api /usr/local/bin/api
COPY --from=builder /out/devtoken /usr/local/bin/devtoken

EXPOSE 18080

CMD ["api"]
