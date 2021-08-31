FROM golang:alpine as builder
ENV GO111MODULE=on
RUN mkdir /app
ADD . /app
WORKDIR /app
COPY go.mod ./
RUN go mod download
RUN go clean --modcache
COPY . .
RUN go build -o main

# reduce size
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .
EXPOSE $PORT
CMD ["./main"]
