FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o finance-service ./cmd/main

FROM alpine AS runner

WORKDIR /app
ADD .env .
COPY --from=builder /app/finance-service .

EXPOSE 8080

ENTRYPOINT ["./finance-service"]