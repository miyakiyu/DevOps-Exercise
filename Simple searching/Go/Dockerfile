FROM golang:latest

WORKDIR /app

COPY . .

RUN go mod init app
RUN go mod tidy
RUN go build -o app

EXPOSE 8080

CMD ["./app"]

