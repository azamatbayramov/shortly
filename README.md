# shortly

[![CI](https://github.com/azamatbayramov/shortly/actions/workflows/ci.yml/badge.svg)](https://github.com/azamatbayramov/shortly/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/azamatbayramov/shortly)](https://goreportcard.com/report/github.com/azamatbayramov/shortly)
[![Go Version](https://img.shields.io/github/go-mod/go-version/azamatbayramov/shortly)](https://github.com/azamatbayramov/shortly)

Shortly is a **simple API** for a link shortening service. It provides endpoints to shorten URLs, redirect users from a short URL back to the original URL, and even features a minimal HTML frontend for quick link shortening.

---

## Table of Contents

- [Features](#features)
- [Endpoints](#endpoints)
- [Installation](#installation)
  - [Run Locally](#run-locally)
  - [Run with Docker](#run-with-docker)
  - [Docker Compose (with PostgreSQL)](#using-docker-compose-with-postgresql)
- [Environment Variables](#environment-variables)
- [Testing](#testing)
- [CI/CD](#cicd)
- [License](#license)

---

## Features

- **URL Shortening:** Generate a short alias for any given URL.
- **Redirection:** Redirect a short URL back to its original URL.
- **Frontend Interface:** Built-in HTML page for quick link shortening.
- **Flexible Storage:** Choose between in-memory storage or PostgreSQL.
- **Custom Encoder:** Configure your own alphabet and fixed-length encoding for short links.
- **Docker Support:** Easily build and run the application with Docker and Docker Compose.
- **CI:** Automated testing, linting, and building with GitHub Actions.

---

## Endpoints

### Shorten a URL

`POST /shorten`

**Request:**

```json
{
  "full_link": "https://www.example.com"
}
```

**Response:**

```json
{
  "short_link": "abCD_1234_"
}
```

### Redirect to Original URL

`GET /{short_link}`

Automatically redirects to the original URL.

### Get Frontend HTML Page

`GET /`

Returns a simple HTML page featuring a form to shorten URLs.

---

## Installation

### Prerequisites

- [Go 1.23.6](https://golang.org/dl/) or higher (for local development)
- [Docker](https://www.docker.com/) (for containerized setup)
- (Optional) [Docker Compose](https://docs.docker.com/compose/) for PostgreSQL support

### Run Locally

1. **Clone the Repository:**

   ```bash
   git clone https://github.com/azamatbayramov/shortly.git
   cd shortly
   ```

2. **Set Up Environment Variables:**  
   Configure the application using environment variables (see [Environment Variables](#environment-variables)).

3. **Download Dependencies and Start the Server:**

   ```bash
   go mod download
   go run cmd/server/main.go
   ```

   The server will run at the host and port defined by `APP_HOST` and `APP_PORT` (default: `0.0.0.0:8000`).

### Run with Docker

1. **Build the Docker Image:**

   ```bash
   docker build -t shortly .
   ```

2. **Run the Docker Container:**

   ```bash
   docker run -d -p 8000:8000 shortly
   ```

### Using Docker Compose (with PostgreSQL)

The project includes a `docker-compose.yml` file that sets up both the application and a PostgreSQL database.

1. **Configure Environment Variables:**  
   Create a `.env` file with the necessary variables (refer to [Environment Variables](#environment-variables)).

2. **Start with Docker Compose:**

   ```bash
   docker-compose up
   ```

   This command launches both the PostgreSQL database and the shortly application.

---

## Environment Variables

The application supports the following environment variables:

| Variable                  | Description                                                                              | Default                                                         |
| ------------------------- | ---------------------------------------------------------------------------------------- | --------------------------------------------------------------- |
| `APP_HOST`                | Address to bind the server.                                                              | `0.0.0.0`                                                       |
| `APP_PORT`                | Port to run the server on.                                                               | `8000`                                                          |
| `STORAGE_TYPE`            | Storage backend. Acceptable values: `"in_memory"` or `"postgresql"`.                     | `in_memory`                                                     |
| `POSTGRES_HOST`           | Host address for PostgreSQL (required if `STORAGE_TYPE` is set to `postgresql`).           | —                                                               |
| `POSTGRES_PORT`           | Port for PostgreSQL.                                                                     | `5432`                                                          |
| `POSTGRES_DB`             | PostgreSQL database name.                                                                | —                                                               |
| `POSTGRES_USER`           | Username for PostgreSQL.                                                                 | —                                                               |
| `POSTGRES_PASSWORD`       | Password for PostgreSQL.                                                                 | —                                                               |
| `CODER_ALPHABET`          | Alphabet used for encoding IDs into short links.                                         | `abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_` |
| `CODER_LENGTH`            | Fixed length for the encoded short link.                                                 | `10`                                                            |
| `ORIGINAL_LINK_MAX_LENGTH`| Maximum allowed length for the original URL.                                             | `2048`                                                          |

---

## Testing

Run the tests with the following command:

```bash
go test ./tests
```

---

## CI

The project utilizes GitHub Actions for continuous integration. The workflow configuration is available at [`.github/workflows/ci.yml`](.github/workflows/ci.yml).

---

## License

This project is licensed under the [MIT License](LICENSE).

---

Happy shortening!
