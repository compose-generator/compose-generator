# Development image
FROM node:lts AS dev

WORKDIR /code
COPY yarn.lock /code/yarn.lock
COPY package.json /code/package.json
RUN yarn install --production

COPY . /code

ENV CI=true
ENV PORT=3000

CMD ["yarn", "start"]

# Builder
FROM dev AS build
RUN yarn run build

# Minimalistic image
FROM nginx:1.21-alpine
COPY --from=build /code/build /usr/share/nginx/html