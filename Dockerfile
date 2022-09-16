FROM golang:alpine

WORKDIR /app

COPY * ./
RUN go mod download

RUN go build -o /smtp-http-gateway

EXPOSE 8080

CMD [ "/smtp-http-gateway" ]