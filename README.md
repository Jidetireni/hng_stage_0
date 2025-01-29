# HNG Stage 0 API  
This is a simple public API developed for the HNG12 Stage 0 task. The api is written in golang and gin web framework.
The API returns basic information such as my registered email address, current datetime in ISO 8601 format, and the GitHub repository URL.
The pro
hng_stage_0/ ├── main.go # Entry point of the application ├── model/ │ └── api_response.go # Defines the API response struct and logic ├── route/ │ └── route.go # API route definitions and handlers └── go.mod # Go module configuration

## Setup Instructions

### Prerequisites
- Go installed on your system

### Running Locally
1. Clone this repository:
   ```bash
   git clone https://github.com/yourusername/hng_stage_0.git
   cd hng_stage_0
   ```

2. Run the server:
   ```bash
   go run main.go
   ```

3. Visit `http://localhost:8080/hngstage0` in your browser to see the response.

## API Documentation

### Endpoint
**GET** `/hngstage0`

### Request Example
```bash
curl http://localhost:8080/hngstage0
```

### Response Example
```json
{
  "email": "jtirenipraise@gmail.com",
  "current_datetime": "2025-01-30T09:30:00Z",
  "github_url": "https://github.com/Jidetireni/hng_stage_0.git"
}
```

## Backlink
[Hire Golang Developers](https://hng.tech/hire/golang-developers)
```

Let me know if you need help customizing it further! 😊
