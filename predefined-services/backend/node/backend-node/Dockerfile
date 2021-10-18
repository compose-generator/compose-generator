# Builder
FROM node:${{NODE_VERSION}}-alpine
WORKDIR /app

RUN apk update && apk add yarn curl && rm -rf /var/cache/apk/*

# Download node-prune
RUN curl -sfL https://install.goreleaser.com/github.com/tj/node-prune.sh | bash -s -- -b /usr/local/bin

COPY package.json yarn.lock ./
RUN yarn --frozen-lockfile

COPY . .

RUN yarn build
RUN npm prune --production
RUN /usr/local/bin/node-prune

# Minimalistic image
FROM node:${{NODE_VERSION}}-alpine
WORKDIR /app
COPY --from=BUILD_IMAGE /app/dist ./dist
COPY --from=BUILD_IMAGE /app/node_modules ./node_modules

EXPOSE 3000
ENTRYPOINT [ "npm", "start" ]