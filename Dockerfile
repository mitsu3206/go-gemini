FROM golang:1.24-alpine

WORKDIR /go/src/app

RUN go install github.com/air-verse/air@latest

COPY . .

CMD ["air", "-c", ".air.toml"]
