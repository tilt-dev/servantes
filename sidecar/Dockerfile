FROM rust:1.37.0-alpine

COPY ./ ./

RUN cargo build --release
CMD target/release/sidecar
