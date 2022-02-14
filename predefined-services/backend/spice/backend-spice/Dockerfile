# Builder
FROM chillibits/spice:${{SPICE_VERSION}} as builder
WORKDIR /usr/src

COPY . .
RUN spice build --output /usr/local/bin/${{SPICE_APP_NAME}} ${{SPICE_MAIN_SOURCE_FILE}}

# Minimalistic image
FROM scratch
COPY --from=builder /usr/local/bin/${{SPICE_APP_NAME}} /usr/local/bin/${{SPICE_APP_NAME}}
ENTRYPOINT [ "${{SPICE_APP_NAME}}" ]