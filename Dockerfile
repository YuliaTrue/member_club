FROM golang:latest

ENV GOPATH=/

COPY ./ ./

RUN apt-get update

RUN go build -o server ./main.go

CMD ["./server"]
