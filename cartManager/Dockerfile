FROM golang:1.15

WORKDIR $GOPATH/src/cartmanager

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o ./out/cartmanager .

CMD ["./out/cartmanager"]