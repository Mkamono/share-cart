FROM node:22-slim AS local

WORKDIR /app/share-cart

RUN yarn global add @biomejs/biome

COPY share-cart/yarn.lock share-cart/package.json ./
RUN yarn install
