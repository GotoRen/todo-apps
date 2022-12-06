FROM mhart/alpine-node:16.4 AS builder
WORKDIR /app
COPY package.json ./
COPY yarn.lock .
RUN yarn install
COPY . .
RUN yarn run build

FROM mhart/alpine-node
ENV TZ=Asia/Tokyo
RUN apk --no-cache add tzdata && \
    yarn global add serve
WORKDIR /app
COPY --from=builder /app/build .
EXPOSE 3000
CMD ["serve", "-p", "3000", "-s", "."]
