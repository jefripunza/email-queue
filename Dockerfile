FROM golang:tip-trixie AS backend
FROM debian:bookworm-slim AS runner
LABEL org.opencontainers.image.authors="Jefri Herdi Triyanto <jefriherditriyanto@gmail.com>"
LABEL description="Ketring.ID: Catering Management System"

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

# copy compiled files
COPY --from=be-builder /app/email-queue /app/email-queue

# run
CMD ["./email-queue"]
