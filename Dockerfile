FROM alpine:3.13
WORKDIR /cg/out

COPY compose-generator /cg/compose-generator
COPY predefined-services/ /cg/predefined-services/

ENV PATH="/cg:${PATH}"
ENV TERM="xterm-256color"
ENV COMPOSE_GENERATOR_DOCKERIZED=1

RUN chmod +x /cg/compose-generator
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
ENTRYPOINT [ "compose-generator" ]