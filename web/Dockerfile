# ベースイメージ
FROM oven/bun:1.1.26 AS base

ENV TZ=Asia/Tokyo
WORKDIR /app/share-cart

# ローカル開発
FROM base AS local

COPY share-cart/bun.lockb share-cart/package.json ./
RUN bun install

# ビルダー
FROM base AS builder

COPY share-cart/ .
RUN bun install --frozen-lockfile
RUN bun run build

FROM base AS production-deps

COPY share-cart/bun.lockb share-cart/package.json ./
RUN bun install --production --frozen-lockfile

# 本番環境
FROM oven/bun:1.1.26-slim AS production

ENV TZ=Asia/Tokyo
WORKDIR /app/share-cart

COPY --from=builder /app/share-cart/build /app/share-cart/build
COPY --from=builder /app/share-cart/package.json /app/share-cart/package.json
COPY --from=production-deps /app/share-cart/node_modules /app/share-cart/node_modules

USER bun

CMD ["bun", "run", "start"]
