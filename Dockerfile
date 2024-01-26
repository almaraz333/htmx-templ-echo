FROM golang:1.21.6 as builder

RUN mkdir /app
ADD . /app
WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./

EXPOSE 8080

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /main

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder main .

EXPOSE 8080

CMD ["./main"]
