FROM golang:1.21-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /app/bin/server ./cmd/server/main.go

EXPOSE 50052

CMD ["/app/bin/server"]