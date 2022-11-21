FROM golang:1.19.3-alpine

WORKDIR /app

COPY go.mod ./


RUN go mod download

COPY app ./

RUN go build -o /dev-web-server

EXPOSE 8443

CMD ["/dev-web-server"]