FROM node:24.10.0-bookworm as front-builder

WORKDIR /build

COPY ./web/app/httpinfo/package*.json ./
COPY ./web/app/httpinfo/vite.config.js ./
COPY ./web/app/httpinfo/index.html ./
COPY ./web/app/httpinfo/src ./src
COPY ./web/app/httpinfo/public ./public

RUN npm install
RUN npm run build

FROM golang:1.23-bookworm AS back-builder

WORKDIR /build

COPY ./go.mod ./

RUN go mod download

COPY ./cmd ./cmd
COPY ./internal ./internal

RUN mkdir stats

WORKDIR /build/cmd

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -o httpinfo \
    -trimpath \
    -ldflags="-s -w"

FROM back-builder AS tester
WORKDIR /build
RUN go test -v ./...

FROM gcr.io/distroless/static-debian12 AS runner

WORKDIR /app

COPY --from=back-builder  /build/cmd/httpinfo .
COPY --from=back-builder /build/stats ./stats
COPY --from=front-builder /build/dist ./dist

EXPOSE 8080

CMD ["/app/httpinfo"]