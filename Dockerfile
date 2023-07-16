FROM golang:1.20.6

WORKDIR /app


COPY . .

RUN go mod tidy