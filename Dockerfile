FROM alpine:3.13
WORKDIR /compose-generator

COPY compose-generator ./compose-generator
COPY predefined-templates/ predefined-templates/

ENV PATH="/compose-generator:${PATH}"
ENV TERM="xterm-256color"
ENV COMPOSE_GENERATOR_DOCKERIZED=1

RUN chmod +x compose-generator
RUN mkdir out
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
ENTRYPOINT [ "compose-generator" ]