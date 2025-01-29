# HNG Stage 0 API

This is a simple public API developed for the HNG12 Stage 0 task. The API is written in Go and uses the Gin web framework. It returns basic information such as:

- My registered email address
- Current datetime in ISO 8601 format
- The GitHub repository URL

## Project Structure

```
hngstage0/
├── /model                # Data models (if applicable)
├── api.go                # API entry point
├── route/router.go       # Routing logic
├── README.md             # Project documentation
├── get.http              # HTTP request file (for testing or examples)
├── go.mod                # Go dependencies
├── go.sum                # Go dependencies checksum
└── main.go               # Main entry point to run the server
```

## Setup Instructions

### Prerequisites

- Go must be installed on your system.

### Running Locally

1. Clone this repository:

   ```bash
   git clone https://github.com/yourusername/hng_stage_0.git
   ```

2. Navigate to the project directory:

   ```bash
   cd hng_stage_0
   ```

3. Run the server:

   ```bash
   go run main.go
   ```

4. Visit [http://localhost:8080/hngstage0](http://localhost:8080/hngstage0) in your browser to see the response.

## API Documentation

### Endpoint

- **GET** `https://hngstage0-production.up.railway.app/hngstage0`

### Request Example

You can use `curl` to make the request:

```bash
curl https://hngstage0-production.up.railway.app/hngstage0
```

### Response Example

The API will return a JSON response with your email, current date-time in ISO 8601 format, and the GitHub repository URL:

```json
{
  "email": "jtirenipraise@gmail.com",
  "current_datetime": "2025-01-30T09:30:00Z",
  "github_url": "https://github.com/Jidetireni/hng_stage_0.git"
}
```

## Backlink

If you're interested in hiring Golang developers, check out this [link](https://hng.tech/hire/golang-developers) for more details on the skillset.
