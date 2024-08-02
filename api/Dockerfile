# ローカル開発用
FROM golang:1.22.5 AS local

ENV TZ=Asia/Tokyo
WORKDIR /app

RUN apt-get update && apt-get install -y \
    graphviz

RUN go install -v github.com/go-delve/delve/cmd/dlv@latest
RUN go install -v github.com/ogen-go/ogen/cmd/ogen@latest
RUN go install -v github.com/air-verse/air@latest
RUN go install -v github.com/kisielk/godepgraph@latest

CMD ["air", "-c", ".air.toml"]

# デプロイ用
FROM golang:1.22.5 AS builder

WORKDIR /app

RUN --mount=type=bind,target=. go build -o /go/bin/app

FROM gcr.io/distroless/static-debian12 AS production

COPY --from=builder /go/bin/app /

CMD ["/app"]
