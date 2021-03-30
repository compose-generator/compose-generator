FROM alpine:3.13.3
WORKDIR /toolbox

# Install alpine packages
RUN apk update
RUN apk add sudo bash curl npm yarn unzip

# Install required npm packages
RUN yarn global add @angular/cli @vue/cli

ENTRYPOINT [ "/bin/bash", "-c"]