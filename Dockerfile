FROM golang:1.22.6

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o app cmd/app/main.go

EXPOSE 3002

CMD ["/app/app"]
