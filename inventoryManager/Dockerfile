FROM golang:1.15

WORKDIR $GOPATH/src/InventoryManager

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o ./out/InventoryManager .

EXPOSE 9100
CMD ["./out/InventoryManager"]
