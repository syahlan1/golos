FROM golang:1.21.5-alpine AS builder

WORKDIR /app-services

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

RUN go build -o main .

FROM scratch

WORKDIR /

COPY --from=builder /app-services/main .
COPY --from=builder /app-services/.env .

EXPOSE 8000
CMD ["./main"]