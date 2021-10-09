FROM alpine:3.14

# Set env variables
ENV PATH="/cg:${PATH}"
ENV TERM="xterm-256color"
ENV COMPOSE_GENERATOR_DOCKERIZED=1

# Set default arg value
ARG ARCH=amd64

# Prepare container
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
WORKDIR /cg/out

# Install Docker CLI
COPY --from=docker:dind /usr/local/bin/docker /usr/local/bin/

# Install CCom
RUN apk add curl
RUN curl -SsL https://github.com/compose-generator/ccom/releases/latest/download/ccom_${ARCH}.apk -o ccom.apk
RUN apk add --allow-untrusted ccom.apk; rm ccom.apk

# Copy sources
COPY compose-generator /cg/compose-generator
COPY predefined-services/ /cg/predefined-services/
RUN chmod +x /cg/compose-generator

# Set entrypoint
ENTRYPOINT [ "compose-generator" ]