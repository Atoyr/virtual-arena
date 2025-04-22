FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o mapservice cmd/mapservice/main.go

FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/mapservice .
# PoC データ格納用
RUN mkdir -p data/maps
EXPOSE 8081
CMD ["./mapservice"]

