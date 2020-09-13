FROM golang:1.15

WORKDIR $GOPATH/src/checkoutmanager

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o ./out/chkman .

CMD ["./out/chkman"]