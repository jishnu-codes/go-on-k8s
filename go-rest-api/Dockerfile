FROM golang:1.15 as builder
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main /app
CMD ["/app/main"]
