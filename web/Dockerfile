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

# 本番環境
FROM base AS production

COPY share-cart/ .
RUN pnpm install --frozen-lockfile

CMD [ "sh", "-c", "pnpm run build && pnpm start" ]
