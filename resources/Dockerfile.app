FROM golang:alpine

RUN mkdir /app
WORKDIR /app

COPY . .
COPY .env .

RUN go get -d -v ./...
RUN go install -v ./...
RUN go build -o /build cmd/ip2p-server/ip2p-server.go

EXPOSE 8888

CMD [ "/build" ]
