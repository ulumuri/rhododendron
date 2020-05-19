FROM golang:alpine as builder

RUN apk update && apk add --no-cache git

WORKDIR /api

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -o main .

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /api/main .

CMD ["./main", "dev"]