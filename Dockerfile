FROM golang:1.21

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o ./out/e-commerce ./cmd/main.go

EXPOSE 8080

CMD ["./out/e-commerce"]