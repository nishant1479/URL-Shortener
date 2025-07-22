# Go URL Shortener

A fast, scalable URL Shortener written in Go, designed to convert long URLs into short, easy-to-share links. Built for reliability and speed, this system uses MongoDB Atlas as its database and Redis for high-performance caching. Expiring URLs, API key security, and modular clean architecture make it suitable for production, integrations, or as a starting point for more advanced link management tools.

## Features

- **Shorten URLs:** Convert long URLs to short, unique keys.
- **Redirects:** Instantly redirects requests from short keys to original URLs.
- **Expiration:** Supports setting a TTL (Time-To-Live) for links; expired URLs auto-delete.
- **Caching:** Uses Redis for low-latency retrieval of active links.
- **API Key Authentication:** Secure endpoint to protect the URL creation process.

## Tech Stack

| Component       | Description                                            |
|------------------|--------------------------------------------------------|
| Go (net/http)    | Efficient HTTP server and RESTful API in pure Go.     |
| MongoDB Atlas    | Managed, cloud-hosted NoSQL database.                 |
| Redis            | In-memory store for fast link lookups (cache layer).  |
| API Key Auth     | Simple, secure access control for URL shortening.     |
| Clean Architecture | Separation: config, db, handler, middleware, utils. |

## Getting Started

### Prerequisites

- Go 1.18+ installed
- MongoDB Atlas account & cluster
- Redis (local or cloud, optional but recommended)

### Setup Instructions

1. **Clone the repository:**

    ```
    git clone https://github.com/your-username/go-url-shortener.git
    cd go-url-shortener
    ```

2. **Configure Environment Variables:**

    - Copy `.env.example` to `.env` and update:
      ```
      MONGO_URI=mongodb+srv://<user>:<pass>@cluster.mongodb.net
      REDIS_ADDR=localhost:6379
      API_KEY=<your-secret-key>
      ```

3. **Start Redis (optional, but improves performance):**

    - **Local:**
      ```
      docker run -p 6379:6379 redis
      ```

4. **Run the Go server:**

    ```
    go run main.go
    ```

## API Usage & Testing (with Postman)

### 1. Create a Short URL

- **Endpoint:** `POST /shorten`
- **Headers:**
  - `X-API-KEY: <your-api-key>`
- **Body (JSON):**
    ```
    {
      "url": "https://example.com/your-long-link",
      "expireAt": "2025-08-01T00:00:00Z" // optional ISO-8601
    }
    ```
- **Response:**
    ```
    {
      "shortUrl": "http://localhost:8080/abc123",
      "expireAt": "2025-08-01T00:00:00Z"
    }
    ```

### 2. Redirect to Original URL

- **Endpoint:** `GET /{shortKey}`
- **Example:**  
  `GET /abc123`  
  → Redirects to `https://example.com/your-long-link` 

## Future Enhancements

- 🧭 Click analytics (visits, referrers, device info)
- 🖥️ Frontend dashboard for link management & analytics
- 🔐 JWT-based authentication & user access control
