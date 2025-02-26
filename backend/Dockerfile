FROM golang:1.24-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o kairosai-backend main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/kairosai-backend .

EXPOSE 8080

CMD ["./kairosai-backend"]
