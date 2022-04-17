# Builder
FROM rust:${{ROCKET_RUST_VERSION}} as builder
WORKDIR /usr/src

RUN rustup override set nightly

COPY . .
RUN cargo install --path .

# Minimalistic image
FROM debian
COPY --from=builder /usr/local/cargo/bin/${{ROCKET_APP_NAME}} /usr/local/bin/${{ROCKET_APP_NAME}}
ENTRYPOINT [ "${{ROCKET_APP_NAME}}" ]