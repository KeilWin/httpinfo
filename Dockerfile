FROM golang:1.23-alpine3.22 as builder

WORKDIR /build

COPY ./go.mod ./

RUN go mod download

COPY ./cmd ./cmd
COPY ./internal ./internal

WORKDIR /build/cmd

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o httpinfo

FROM builder AS tester
WORKDIR /build
RUN go test -v ./...

FROM alpine:3.22 AS runner

WORKDIR /app

COPY --from=builder  /build/cmd/httpinfo .
COPY ./web ./web

RUN mkdir stats

EXPOSE 8080

CMD ["/app/httpinfo"]