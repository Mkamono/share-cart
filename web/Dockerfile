FROM node:22-slim AS local

ENV TZ=Asia/Tokyo
WORKDIR /app/share-cart

RUN yarn global add @biomejs/biome

COPY share-cart/yarn.lock share-cart/package.json ./
RUN yarn install
