# NodeJS build instructions
FROM node:lts AS development

RUN npm i -g @angular/cli

WORKDIR /code
COPY src/package.json /code/package.json
COPY src/package-lock.json /code/package-lock.json
RUN npm i

COPY ./src /code

CMD ["ng", "serve", "--host", "0.0.0.0"]

# Build
FROM development AS build
RUN npm run build

# Serve instructions
FROM nginx:1.19-alpine
COPY --from=build /code/dist/demo-app /usr/share/nginx/html