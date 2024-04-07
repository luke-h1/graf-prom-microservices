FROM node:20-alpine

WORKDIR /usr/src/app

COPY package*.json ./
COPY pnpm-lock.yaml ./

RUN --mount=type=cache,id=pnpm,target=/pnpm/store \
    npm install -g pnpm && \
    pnpm i

RUN pnpm i

COPY . .

EXPOSE 8080
USER node

CMD ["node", "index.js"]
