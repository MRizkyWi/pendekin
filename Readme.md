# Pendekin URL Shortener

Pendekin is a simple URL shortener service built with Go. It allows you to shorten long URLs and redirect to the original URL using the shortened version.

## Features

- Shorten long URLs to a concise format
- Redirect to the original URL using the shortened URL
- Update the status of a URL (activate/deactivate)
- RESTful API endpoints

## Getting Started

### Prerequisites

- Go 1.16+
- Docker

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/mrizkywi/pendekin.git
   cd pendekin
   ```

2. Run the docker compose build
   ```bash
   docker compose up --build
   ```

3. Copy the example environment file and update it with your database credentials:

   ```bash
   cp .env.example .env
   ```

4. Run the database migrations to set up the MySQL database:

   ```bash
   make migrate-up
   ```

5. Build and run the application:

   ```bash
   go build -o pendekin cmd/main.go
   ./pendekin
   ```

### Usage

- Shorten URL: `POST /urls` with JSON body `{"actual_url": "http://example.com"}`
- Redirect to original URL: `GET /urls/{shortUrl}`
- Update URL status: `PUT /urls` with JSON body `{"id": 1, "status": false}`

### Project Structure

- `cmd/http/main.go`: Entry point for the application.
- `cache/`: Cache folder
- `config/`: Configuration files and database connection setup.
- `controller/`: HTTP handlers for API endpoints.
- `model/`: Data models and request/response structures.
- `repository/`: Database interaction and queries.
- `service/`: Business logic and service layer.
- `routes/`: API route definitions.

### Contributing

Contributions are welcome! Please fork the repository and submit a pull request for any improvements or bug fixes.

### License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

