# NodeJS build instructions
FROM node:lts AS development

WORKDIR /code
COPY src/package.json /code/package.json
COPY src/package-lock.json /code/package-lock.json

RUN npm ci
COPY ./src /code

ENV CI=true
ENV PORT=3000

CMD ["npm", "start"]

# Nginx build instructions
FROM development AS build

RUN npm run build

FROM nginx:1.19-alpine
COPY --from=build /code/build /usr/share/nginx/html