FROM golang:1.19.4-alpine3.17

WORKDIR /image_uploader

COPY . .

RUN go mod download

RUN go build -o /uploader

EXPOSE 8081

CMD ["/uploader"]