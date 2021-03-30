# NodeJS build instructions
FROM node:lts AS development

WORKDIR /code
COPY yarn.lock /code/yarn.lock
COPY package.json /code/package.json
RUN yarn install --production

COPY . /code

ENV CI=true
ENV PORT=3000

CMD ["yarn", "start"]

# Build
FROM development AS build
RUN yarn run build

# Serve instructions
FROM nginx:1.19-alpine
COPY --from=build /code/build /usr/share/nginx/html