FROM rust:${{ROCKET_RUST_VERSION}} as builder
WORKDIR /usr/src/${{ROCKET_APP_NAME}}

RUN rustup default nightly
RUN rustup override set nightly

COPY . .
RUN cargo install --path .

FROM debian:buster-slim
COPY --from=builder /usr/local/cargo/bin/${{ROCKET_APP_NAME}} /usr/local/bin/${{ROCKET_APP_NAME}}
CMD ["${{ROCKET_APP_NAME}}"]