FROM golang:1.18.10-alpine

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY *.go ./
RUN go build -o /api

CMD ["/api"]