# Development image
FROM node:lts AS dev

RUN yarn global add @vue/cli

WORKDIR /code
COPY package.json /code/package.json
RUN yarn install

COPY . /code

CMD ["vue", "serve"]

# Builder
FROM dev AS build
RUN yarn run build

# Minimalistic image
FROM nginx:1.21-alpine
COPY --from=build /code/dist /usr/share/nginx/html