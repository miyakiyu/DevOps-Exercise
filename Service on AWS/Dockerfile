FROM golang:latest

WORKDIR /app

COPY . .

RUN go mod init main
RUN go mod tidy
RUN go build -o main

EXPOSE 8080

CMD ["./main"]
