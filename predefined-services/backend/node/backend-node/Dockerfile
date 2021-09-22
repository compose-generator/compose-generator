FROM node:${{NODE_VERSION}}
WORKDIR /usr/src/app

COPY package*.json ./
RUN npm i

EXPOSE 8080
CMD [ "node", "app.js" ]