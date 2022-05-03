FROM golang:1.18.1

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/awp ./cmd/main.go

CMD ["awp"]
