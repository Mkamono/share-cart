FROM node:22-slim AS local

ENV TZ=Asia/Tokyo
WORKDIR /app/share-cart

RUN yarn global add @biomejs/biome

RUN apt-get update && apt-get install -y \
    graphviz

COPY share-cart/yarn.lock share-cart/package.json ./
RUN yarn install

FROM node:22-slim AS builder

ENV TZ=Asia/Tokyo
WORKDIR /app/share-cart

COPY ./share-cart/ .
RUN yarn install && yarn build

FROM node:22-slim AS production-deps

ENV NODE_ENV=production
ENV TZ=Asia/Tokyo
WORKDIR /app/share-cart

COPY ./share-cart/package.json .
RUN yarn install --production

FROM gcr.io/distroless/nodejs22-debian12 AS production

ENV NODE_ENV=production
ENV TZ=Asia/Tokyo
WORKDIR /app/share-cart

COPY --from=builder /app/share-cart/build /app/share-cart/build
COPY --from=production-deps /app/share-cart/package.json /app/share-cart/package.json
COPY --from=production-deps /app/share-cart/node_modules /app/share-cart/node_modules

CMD ["./build/server/index.js"]
