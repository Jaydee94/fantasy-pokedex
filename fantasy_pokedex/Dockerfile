FROM docker.io/golang:1.24-alpine

WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o main .

EXPOSE 8080
CMD ["./main"]