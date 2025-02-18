# shortly

[![CI](https://github.com/azamatbayramov/shortly/actions/workflows/ci.yml/badge.svg)](https://github.com/azamatbayramov/shortly/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/azamatbayramov/shortly)](https://goreportcard.com/report/github.com/azamatbayramov/shortly)
[![Go Version](https://img.shields.io/github/go-mod/go-version/azamatbayramov/shortly)](https://github.com/azamatbayramov/shortly)

shortly is a simple API for a link shortening service. It provides endpoints to shorten URLs, redirect users from a
short URL to the original URL, and even a minimal HTML frontend for quick link shortening.

## Features

- **URL Shortening:** Generate a short alias for a given URL.
- **Redirection:** Redirect a short URL to the original URL.
- **Frontend:** Built-in HTML page for quick usage.
- **Storage Options:** Choose between in-memory storage or PostgreSQL.
- **Custom Encoder:** Configurable alphabet and length for generating short links.
- **Docker Support:** Easily build and run with Docker and Docker Compose.
- **CI/CD:** GitHub Actions are configured for continuous integration and automated Docker image builds.

## Endpoints

### Shorten a URL

`POST /shorten`

#### Request

```json
{
  "full_link": "https://www.example.com"
}
```

#### Response

```json
{
  "short_link": "abCD_1234_"
}
```

### Redirect to Original URL

`GET /{short_link}`

#### Response

Redirects to the original URL

### Get Frontend HTML Page

`GET /`

#### Response

Returns the HTML page with the form to shorten a URL

## Installation

### Prerequisites

- [Go 1.23.6](https://golang.org/dl/) or higher (if running locally)
- [Docker](https://www.docker.com/) (if running via Docker)
- (Optional) [Docker Compose](https://docs.docker.com/compose/) for PostgreSQL support

### Run Locally

1. **Clone the repository:**

   ```bash
   git clone https://github.com/azamatbayramov/shortly.git
   cd shortly
   ```

2. **Set Up Environment Variables:**

   You can configure the application via environment variables (see [Environment Variables](#environment-variables)
   below).

3. **Download Dependencies and Run the Server:**

   ```bash
   go mod download
   go run cmd/server/main.go
   ```

   The server will start at the host and port defined by the `APP_HOST` and `APP_PORT` environment variables (default:
   `0.0.0.0:8000`).

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
   Create a `.env` file (refer to the [Environment Variables](#environment-variables) section) if needed.

2. **Run Docker Compose:**

   ```bash
   docker-compose up
   ```

   This will start the PostgreSQL database and the shortly application.

## Environment Variables

The application supports the following environment variables:

- **APP_HOST**  
  Address to bind the server.  
  _Default:_ `0.0.0.0`

- **APP_PORT**  
  Port to run the server on.  
  _Default:_ `8000`

- **STORAGE_TYPE**  
  Storage backend to use. Acceptable values are `"in_memory"` or `"postgresql"`.  
  _Default:_ `in_memory`

- **POSTGRES_HOST**  
  Host address for PostgreSQL (required if `STORAGE_TYPE` is set to `postgresql`).

- **POSTGRES_PORT**  
  Port for PostgreSQL.  
  _Default:_ `5432`

- **POSTGRES_DB**  
  PostgreSQL database name.

- **POSTGRES_USER**  
  Username for PostgreSQL.

- **POSTGRES_PASSWORD**  
  Password for PostgreSQL.

- **CODER_ALPHABET**  
  Alphabet used for encoding IDs into short links.  
  _Default:_ `abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_`

- **CODER_LENGTH**  
  The fixed length of the encoded short link.  
  _Default:_ `10`

- **ORIGINAL_LINK_MAX_LENGTH**  
  Maximum allowed length for the original URL.  
  _Default:_ `2048`

## Testing

Run the tests with the following command:

```bash
go test ./tests
```

## CI/CD

The project uses GitHub Actions for continuous integration. The workflow definition is located in [
`.github/workflows/ci.yml`](.github/workflows/ci.yml).

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

Happy shortening!
