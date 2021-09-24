FROM node:${{NODE_VERSION}}
WORKDIR /usr/src/app

COPY package*.json ./
RUN npm i

COPY . .

EXPOSE 3000
CMD [ "npm", "start" ]