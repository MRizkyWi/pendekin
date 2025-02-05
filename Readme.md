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
- MySQL

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/pendekin.git
   cd pendekin
   ```

2. Copy the example environment file and update it with your database credentials:

   ```bash
   cp .env.example .env
   ```

3. Run the database migrations to set up the MySQL database:

   ```bash
   make migrate-up
   ```

4. Build and run the application:

   ```bash
   go build -o pendekin cmd/main.go
   ./pendekin
   ```

### Usage

- Shorten URL: `POST /urls` with JSON body `{"actual_url": "http://example.com"}`
- Redirect to original URL: `GET /urls/{shortUrl}`
- Update URL status: `PUT /urls` with JSON body `{"id": 1, "status": false}`

### Project Structure

- `cmd/main.go`: Entry point for the application.
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
