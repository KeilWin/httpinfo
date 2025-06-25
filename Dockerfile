FROM golang:1.23-alpine:3.22 as builder

WORKDIR /app

COPY ./go.mod ./

RUN go mod download

COPY ./cmd ./cmd
COPY ./internal ./internal
COPY ./web ./web

RUN CGO_ENABLED=0 GOOS=linux go build -o /go/bin/httpinfo ./cmd

FROM alpine:3.22

COPY --from=builder /go/bin/httpinfo .

EXPOSE 8080

CMD ["/httpinfo", "--tls", "--index-template-path", "/app/web/template/index.html"]