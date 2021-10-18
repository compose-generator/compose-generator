# Dev image
FROM node:lts AS dev

RUN npm i -g --loglevel=error @angular/cli

WORKDIR /code
COPY package.json /code/package.json
RUN npm i --loglevel=error

COPY . /code

CMD ["ng", "serve", "--host", "0.0.0.0"]

# Builder
FROM dev AS build
RUN npm run build


# Minimalistic image
FROM nginx:1.21-alpine
COPY --from=build /code/dist/demo-app /usr/share/nginx/html