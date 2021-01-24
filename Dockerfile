FROM alpine:3.13
WORKDIR /compose-generator

COPY bin/compose-generator-amd64 ./compose-generator
COPY templates/ templates/

ENV PATH="/compose-generator:${PATH}"

RUN ["chmod", "+x", "compose-generator"]
RUN ["mkdir", "out"]
ENTRYPOINT [ "tail", "-f", "/dev/null" ]