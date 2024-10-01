# ベースイメージ
FROM node:22-slim AS base

ENV TZ=Asia/Tokyo
WORKDIR /app/share-cart

RUN npm install -g pnpm@9.11.0

# ローカル開発
FROM base AS local

RUN npm install -g @biomejs/biome@1.9.2

COPY share-cart/pnpm-lock.yaml share-cart/package.json ./
RUN pnpm install --frozen-lockfile

# ビルダー
# FROM base AS builder

# COPY share-cart/ .
# RUN pnpm install --frozen-lockfile && pnpm run build

# 本番環境
# hadolint ignore=DL3006
FROM base AS production

ENV TZ=Asia/Tokyo
WORKDIR /app/share-cart

# COPY --from=builder /app/share-cart/.next/standalone ./.next/standalone

COPY share-cart/ .
# RUN pnpm install --frozen-lockfile && pnpm run build

# CMD ["node", "start", "./.next/standalone/server.js"]
# CMD [ "npm", "run", "build", ".", "&&", "npm", "run", "start" ]

CMD [ "sh", "/app/share-cart/entrypoint.sh" ]
