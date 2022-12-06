FROM golang:alpine AS builder
WORKDIR /app
COPY . .
RUN go get -d -v && \
    GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-w -s" -o main && \
    chmod +x /app/main

FROM alpine:3.17
ENV TZ=Asia/Tokyo
RUN apk --no-cache add tzdata
WORKDIR /app
COPY --from=builder /app/main .
CMD [ "./main" ]
