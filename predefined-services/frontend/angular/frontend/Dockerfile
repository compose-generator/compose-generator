# NodeJS build instructions
FROM node:lts AS development

RUN npm i -g --loglevel=error @angular/cli

WORKDIR /code
COPY package.json /code/package.json
RUN npm i --loglevel=error

COPY . /code

CMD ["ng", "serve", "--host", "0.0.0.0"]

# Build
FROM development AS build
RUN npm run build

# Serve instructions
FROM nginx:1.19-alpine
COPY --from=build /code/dist/demo-app /usr/share/nginx/html