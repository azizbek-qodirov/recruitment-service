# Base image sifatida Go'ni tanlaymiz
FROM golang:1.19-alpine AS builder

# Ishchi katalogni yaratamiz
WORKDIR /app

# Modullari va dependency-larni qo'shamiz
COPY go.mod go.sum ./
RUN go mod download

# Barcha loyihani nusxalaymiz
COPY . .

# Loyiha kodini build qilamiz
RUN go build -o main.go .

# Minimalistic run-time image yaratamiz
FROM alpine:latest

# Ishchi katalogni yaratamiz
WORKDIR /root/

# Build natijasida olingan binary'ni nusxalaymiz
COPY --from=builder /app/main.go .

EXPOSE 9090

# Ilovani ishga tushiramiz
CMD ["./main.go"]
