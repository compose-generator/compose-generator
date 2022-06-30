# Builder
FROM node:${{NODE_VERSION}}-alpine AS builder
WORKDIR /app

RUN apk update && apk add bash curl npm && rm -rf /var/cache/apk/*

# Download node-prune
RUN curl -sf https://gobinaries.com/tj/node-prune | sh

COPY package*.json ./
RUN npm i
COPY . .

RUN npm prune --production
RUN /usr/local/bin/node-prune


# Minimalistic image
FROM node:${{NODE_VERSION}}-alpine
WORKDIR /app
COPY --from=builder /app ./

EXPOSE 3000
ENTRYPOINT [ "npm", "start" ]