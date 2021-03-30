FROM alpine:3.13.3
WORKDIR /toolbox

RUN apk update
RUN apk add sudo bash curl npm yarn unzip

RUN npm i -g @angular/cli

ENTRYPOINT [ "/bin/bash", "-c"]