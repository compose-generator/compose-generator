FROM alpine:3.13

COPY bin/compose-generator /compose-generator/
ENV PATH="/compose-generator:${PATH}"

RUN ["chmod", "+x", "/compose-generator/compose-generator"]
ENTRYPOINT [ "compose-generator" ]