# WARNING: This Dockerfile is not meant to be used to build the Docker image manually
FROM alpine:3.15.0

# Set env variables
ENV PATH="/cg:${PATH}"
ENV TERM="xterm-256color"
ENV COMPOSE_GENERATOR_DOCKERIZED=1

# Specify volumes
VOLUME /cg/out /var/run/docker.sock

# Set default arg value
ARG ARCH=amd64

# Prepare container
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
WORKDIR /cg/out

# Install Docker CLI
RUN apk add --no-cache docker-cli

# Install CCom
RUN apk add curl
RUN curl -fsSL https://github.com/compose-generator/ccom/releases/latest/download/ccom_${ARCH}.apk -o ccom.apk
RUN apk add --allow-untrusted ccom.apk; rm ccom.apk
RUN apk update && apk add --no-cache libc6-compat libstdc++ && rm -rf /var/cache/apk/*

# Copy sources
COPY compose-generator /cg/cg
COPY predefined-services/ /cg/predefined-services/
RUN chmod +x /cg/cg

# Set entrypoint
ENTRYPOINT [ "cg" ]