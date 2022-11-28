FROM golang:1.19-alpine

WORKDIR /app

# TODO Create non root user

RUN go install github.com/cosmtrek/air@latest

COPY go.mod go.sum ./
RUN go mod download

RUN air init

EXPOSE 8080

CMD [ "air" ]
