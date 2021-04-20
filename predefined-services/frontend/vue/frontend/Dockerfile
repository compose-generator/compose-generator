# NodeJS build instructions
FROM node:${{VUE_NODE_VERSION}} AS development

RUN yarn global add @vue/cli

WORKDIR /code
COPY package.json /code/package.json
RUN yarn install

COPY . /code

CMD ["vue", "serve"]

# Build
FROM development AS build
RUN yarn run build

# Serve instructions
FROM nginx:1.19-alpine
COPY --from=build /code/dist /usr/share/nginx/html