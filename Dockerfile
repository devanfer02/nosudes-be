FROM golang:alpine

RUN mkdir /app 

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . . 

RUN go mod tidy

RUN go build -o main

EXPOSE 8060

ENTRYPOINT ["/app/main"]