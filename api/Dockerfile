# ローカル開発用
FROM golang:1.22.5 AS local

ENV TZ=Asia/Tokyo
WORKDIR /app

RUN apt-get update && apt-get install -y \
    graphviz

RUN curl -sSf https://atlasgo.sh | sh

RUN go install -v github.com/volatiletech/sqlboiler/v4@latest
RUN go install -v github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@latest
RUN go install -v github.com/go-delve/delve/cmd/dlv@latest
RUN go install -v github.com/ogen-go/ogen/cmd/ogen@v1.2.2
RUN go install -v github.com/air-verse/air@latest
RUN go install -v github.com/kisielk/godepgraph@latest
RUN go install -v github.com/cweill/gotests/gotests@latest
RUN go install -v go.uber.org/mock/mockgen@latest

# デプロイ用
FROM golang:1.22.5 AS builder

WORKDIR /app

RUN --mount=type=bind,target=. go build -o /go/bin/app

FROM gcr.io/distroless/base-debian12 AS production

COPY --from=builder /go/bin/app /app

CMD ["/app"]

FROM alpine:3.19 as atlas
RUN apk add --no-cache curl
ARG ATLAS_VERSION=latest
ENV ATLAS_VERSION=${ATLAS_VERSION}
RUN curl -sSf https://atlasgo.sh | sh
