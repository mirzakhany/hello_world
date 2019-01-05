FROM golang:1.10.3 as builder
WORKDIR /root/go/src/github.com/mirzakhany/hello_world
ADD . /root/go/src/github.com/mirzakhany/hello_world/
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates && rm -rf /var/cache/apk/*
COPY --from=builder /root/go/src/github.com/mirzakhany/hello_world/main .
ENTRYPOINT ["/main"]
