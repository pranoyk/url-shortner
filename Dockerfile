FROM golang:1.16-alpine

WORKDIR /go/src/app

COPY . .

RUN go mod vendor

RUN go build

CMD ["./main"]