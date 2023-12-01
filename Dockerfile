FROM golang:1.21

WORKDIR /app

#RUN go install github.com/cosmtrek/air@latest

COPY . .
RUN go mod tidy
