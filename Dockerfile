FROM golang:1.13

WORKDIR /go/app
COPY . .

RUN go get -d -v ./...


CMD ["go", "run", "main.go"]