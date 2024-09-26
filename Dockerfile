FROM golang:alpine AS builder

WORKDIR /builder

ADD go.mod .
ADD go.sum .

RUN go mod download

COPY . .

RUN go build -o ./cmd/app/main ./cmd/app/main.go

FROM alpine

WORKDIR /app

COPY --from=builder /builder/cmd/app/main ./cmd/app/main

EXPOSE 8088

CMD ["./cmd/app/main"]