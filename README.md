# 📧 Email Queue

A lightweight, self-hosted email queue service built with **Go**. Queue emails via REST API and let the background worker handle SMTP delivery with automatic retry tracking.

## Features

- **REST API** — Simple JSON endpoints to queue emails and check delivery status
- **Background Worker** — Automatically picks up unsent emails and delivers via SMTP
- **Error Tracking** — Failed emails are flagged with error details for easy debugging
- **Multi-Database** — Supports SQLite (default), MySQL, and PostgreSQL
- **API Key Auth** — Protect endpoints with `X-API-Key` header
- **Configurable Delay** — Control polling interval via environment variable
- **Lightweight Docker Image** — Multi-stage build on `debian:bookworm-slim`

## Quick Start

### Docker Run

```bash
docker run -d \
  --name email-queue \
  -p 3000:3000 \
  -e SERVER_PORT=3000 \
  -e API_KEY=your-secret-key \
  -e DATABASE_PROVIDER=sqlite \
  -e DATABASE_NAME=data \
  -e EMAIL_FROM_EMAIL=no-reply@example.com \
  -e EMAIL_FROM_NAME="My App" \
  -e EMAIL_PASSWORD=your-email-password \
  -e EMAIL_SMTP_HOST=smtp.example.com \
  -e EMAIL_SMTP_PORT=587 \
  -e DELAY_SECOND=5 \
  jefriherditriyanto/email-queue:latest
```

### Docker Compose

```yaml
version: "3.8"
services:
  email-queue:
    image: jefriherditriyanto/email-queue:latest
    ports:
      - "3000:3000"
    environment:
      SERVER_PORT: 3000
      API_KEY: your-secret-key
      DATABASE_PROVIDER: sqlite
      DATABASE_NAME: data
      EMAIL_FROM_EMAIL: no-reply@example.com
      EMAIL_FROM_NAME: My App
      EMAIL_PASSWORD: your-email-password
      EMAIL_SMTP_HOST: smtp.example.com
      EMAIL_SMTP_PORT: 587
      DELAY_SECOND: 5
    volumes:
      - email-data:/app

volumes:
  email-data:
```

## Environment Variables

| Variable            | Description                                     | Default         |
| ------------------- | ----------------------------------------------- | --------------- |
| `SERVER_PORT`       | Server port                                     | `3000`          |
| `API_KEY`           | API key for authentication                      | `test-key`      |
| `DATABASE_PROVIDER` | Database engine (`sqlite`, `mysql`, `postgres`) | `sqlite`        |
| `DATABASE_HOST`     | Database host                                   | `localhost`     |
| `DATABASE_PORT`     | Database port                                   | auto            |
| `DATABASE_USER`     | Database username                               | —               |
| `DATABASE_PASS`     | Database password                               | —               |
| `DATABASE_NAME`     | Database name / file path                       | `./data.sqlite` |
| `EMAIL_FROM_EMAIL`  | Sender email address                            | —               |
| `EMAIL_FROM_NAME`   | Sender display name                             | —               |
| `EMAIL_PASSWORD`    | SMTP password                                   | —               |
| `EMAIL_SMTP_HOST`   | SMTP server host                                | —               |
| `EMAIL_SMTP_PORT`   | SMTP server port                                | —               |
| `DELAY_SECOND`      | Worker polling interval (seconds)               | `5`             |

## API Reference

All endpoints require the `X-API-Key` header.

### Queue an Email

```
POST /api/message/send
```

**Request Body:**

```json
{
  "to": "recipient@example.com",
  "subject": "Hello",
  "body": "<h1>Hello World</h1>"
}
```

**Response** `201`:

```json
{
  "status": 201,
  "message": "Email queued successfully",
  "data": {
    "key": "dfc9e68f-9faf-4d21-9ebc-84b92d7183fd"
  }
}
```

### Check Email Status

```
GET /api/message/status/:key
```

**Response** `200`:

```json
{
  "status": 200,
  "message": "Email status retrieved",
  "data": {
    "key": "dfc9e68f-9faf-4d21-9ebc-84b92d7183fd",
    "is_sended": true,
    "is_error": false,
    "error_message": "",
    "sended_at": "2026-03-31T15:00:05Z"
  }
}
```

## How It Works

1. Client sends a `POST /api/message/send` request with email details
2. The email is saved to the database with `is_sended: false`
3. A unique `key` is returned for status tracking
4. The background worker polls for unsent emails every `DELAY_SECOND` seconds
5. On success → `is_sended` is set to `true`
6. On failure → `is_error` is set to `true` with `error_message` details
7. Client can check delivery status anytime via `GET /api/message/status/:key`

## Tech Stack

- **[Go](https://go.dev/)** — Language
- **[Fiber](https://gofiber.io/)** — HTTP framework
- **[GORM](https://gorm.io/)** — ORM (SQLite / MySQL / PostgreSQL)
- **net/smtp** — SMTP delivery

## License

MIT
