FROM rust:1.56.0-alpine

COPY ./ ./

RUN cargo build --release
CMD target/release/sidecar
