FROM golang:1.23-bookworm AS builder

WORKDIR /build

COPY ./go.mod ./

RUN go mod download

COPY ./cmd ./cmd
COPY ./internal ./internal

RUN mkdir stats

WORKDIR /build/cmd

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o httpinfo

FROM builder AS tester
WORKDIR /build
RUN go test -v ./...

FROM gcr.io/distroless/static-debian12 AS runner

WORKDIR /app

COPY --from=builder  /build/cmd/httpinfo .
COPY --from=builder /build/stats ./stats
COPY ./web ./web

EXPOSE 8080

CMD ["/app/httpinfo"]