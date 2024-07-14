# ベースイメージ
FROM golang:1.22.5 as base

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .


# ローカル開発用
FROM golang:1.22.5 as local

WORKDIR /app
COPY --from=base /app /app

RUN go install -v github.com/go-delve/delve/cmd/dlv@latest


# デプロイ用
FROM golang:1.22.5 as builder

WORKDIR /app
COPY --from=base /app /app

RUN go build -o /app/bin/app /app/cmd/main.go

FROM golang:1.22.5 as production

WORKDIR /app
COPY --from=builder /app/bin/app /app/bin/app

CMD ["/app/bin/app"]
