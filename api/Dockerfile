# ローカル開発用
FROM golang:1.22.5 AS local

WORKDIR /app

RUN go install -v github.com/go-delve/delve/cmd/dlv@latest
RUN go install -v github.com/ogen-go/ogen/cmd/ogen@latest
RUN go install -v github.com/air-verse/air@latest

CMD ["air", "-c", ".air.toml"]

# デプロイ用
FROM golang:1.22.5 AS builder

WORKDIR /app
COPY . .

RUN CGO_ENABLED=0 go build -o /go/bin/app app/cmd/main.go

FROM gcr.io/distroless/static-debian12 AS production

COPY --from=builder /go/bin/app /

CMD ["/app"]
