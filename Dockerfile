FROM golang:1.23

WORKDIR /app

COPY ./go.mod ./

RUN go mod download

COPY ./cmd ./cmd
COPY ./internal ./internal
COPY ./web ./web

RUN CGO_ENABLED=0 GOOS=linux go build -o /httpinfo ./cmd

EXPOSE 8080

CMD ["/httpinfo", "--tls", "--index-template-path", "/app/web/template/index.html"]