FROM golang:1.23-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /app/cmd/main /app/main.go

EXPOSE 5050

CMD ["/app/cmd/main"]