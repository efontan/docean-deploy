FROM golang:1.13 as builder
COPY . /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux GOPROXY=https://proxy.golang.org go build -o app main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates mailcap && addgroup -S app && adduser -S app -G app
USER app
WORKDIR /app
COPY --from=builder /app/app .
EXPOSE 9091
ENTRYPOINT [ "./app" ]