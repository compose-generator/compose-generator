FROM alpine:3.13.3
WORKDIR /toolbox

# Install alpine packages
RUN apk update
RUN apk add sudo bash curl npm yarn unzip python3 py3-pip

# Install required npm packages
RUN yarn global add @angular/cli @vue/cli

# Install pip dependencies
RUN pip3 install flask-now

ENTRYPOINT [ "/bin/bash", "-c"]