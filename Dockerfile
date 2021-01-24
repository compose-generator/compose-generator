FROM alpine:3.13
WORKDIR /compose-generator

COPY bin/compose-generator-amd64 ./compose-generator
COPY templates/ templates/

ENV PATH="/compose-generator:${PATH}"
ENV TERM="xterm-256color"

RUN chmod +x compose-generator
RUN mkdir out
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
ENTRYPOINT [ "compose-generator" ]