FROM golang:1.22-alpine AS builder

WORKDIR /build

COPY go.mod .
COPY main.go .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o regrettable .

FROM scratch

WORKDIR /app

COPY --from=builder /build/regrettable .
COPY index.html .
COPY styles.css .
COPY logo.png .

EXPOSE 8080

ENTRYPOINT ["/app/regrettable"]
