FROM golang:latest

ENV GOPATH=/

COPY ./ ./

RUN apt-get update

RUN go build -o server ./cmd/main.go

CMD ["./server"]
