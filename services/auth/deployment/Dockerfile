FROM golang:1.22.4-alpine3.20

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go install github.com/swaggo/swag/cmd/swag@latest

RUN swag init -g cmd/main.go

RUN go build -o ./ecom-micro-auth ./cmd

EXPOSE 8099

CMD ["./ecom-micro-auth"]