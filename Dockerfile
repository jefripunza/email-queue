FROM golang:tip-trixie AS backend
FROM debian:bookworm-slim AS runner
LABEL org.opencontainers.image.authors="Jefri Herdi Triyanto <jefriherditriyanto@gmail.com>"
LABEL description="Email Queue: A lightweight, self-hosted email queue service"

# =======================================================================================
# Build Backend
# =======================================================================================

FROM backend AS be-builder
WORKDIR /app

COPY ./go.mod ./
RUN go mod download

COPY . .

RUN go build -o email-queue main.go

# =======================================================================================
# Run
# =======================================================================================

FROM runner
WORKDIR /app

# install ca-certificates
RUN apt-get update && apt-get install -y \
    sudo unzip ca-certificates curl wget nano net-tools iputils-ping \
    && rm -rf /var/lib/apt/lists/*

# copy compiled files
COPY --from=be-builder /app/email-queue /app/email-queue

# run
CMD ["./email-queue"]
