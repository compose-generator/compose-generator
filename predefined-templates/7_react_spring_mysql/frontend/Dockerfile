# NodeJS build instructions
FROM node:lts AS development

WORKDIR /code
COPY src/package.json /code/package.json
COPY src/yarn.lock /code/yarn.lock
RUN yarn install --frozen-lockfile

COPY ./src /code

ENV CI=true
ENV PORT=3000

CMD ["yarn", "start"]

# Nginx build instructions
FROM development AS build

RUN yarn run build

FROM nginx:1.19-alpine
COPY --from=build /code/build /usr/share/nginx/html