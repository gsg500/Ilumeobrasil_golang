# Etapa de build
FROM golang:1.23-alpine AS builder

# Instala dependências
RUN apk update && apk add --no-cache git

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build -o main .

# Etapa final
FROM alpine:latest
WORKDIR /app

# Copia binário, documentação Swagger e .env
COPY --from=builder /app/main .
COPY --from=builder /app/docs ./docs
COPY --from=builder /app/.env .env

EXPOSE 8080
CMD ["./main"]
