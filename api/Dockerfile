FROM golang:1.23 AS base

ENV TZ=Asia/Tokyo
WORKDIR /app

# ローカル開発用
FROM base AS local

RUN go install -v github.com/go-delve/delve/cmd/dlv@latest && go install -v github.com/air-verse/air@latest

# デプロイ用
FROM base AS builder

RUN --mount=type=bind,target=. go build -o /go/bin/app

FROM gcr.io/distroless/base-debian12 AS production

COPY --from=builder /go/bin/app /app

CMD ["/app"]
