FROM golang:latest

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o main ./cmd/main.go

CMD ["./main"]