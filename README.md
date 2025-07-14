# 🚦 Go Rate Limit

A lightweight, IP-based rate limiting middleware for Go applications using the Gin framework and Redis.

---

## ✨ Features

* 🔐 **IP-based** rate limiting
* ⏱️ **Configurable** request limits and time windows
* ⚡ **Redis-backed** for fast and scalable tracking
* 🧩 **Easy integration** with existing Gin apps

---

## 📦 Prerequisites

* Go `v1.16+`
* [Redis](https://redis.io/) server
* [Gin](https://github.com/gin-gonic/gin) web framework
* [godotenv](https://github.com/joho/godotenv) for loading `.env` files

---

## 🚀 Installation

### 1. Clone the Repository

```bash
git clone https://github.com/yourusername/go-rate-limit.git
cd go-rate-limit
```

### 2. Install Dependencies

```bash
go mod download
```

### 3. Configure Environment Variables

Create a `.env` file in the root directory:

```env
REDIS_HOST=localhost
REDIS_PORT=6379
REDIS_PASSWORD=""
REDIS_DB=0
```

---

## 🧪 Usage

### 1. Run Redis via Docker

```bash
docker pull redis
docker run --name redis -p 6379:6379 -d redis
```

Ensure it’s running:

```bash
docker ps
```

### 2. Start the Application

```bash
go run main.go
```

By default, the server runs at:
👉 `http://localhost:8080`

### 3. Test

Send requests to the root endpoint (`/`) to test the rate limiter in action.

---

## ⚙️ Configuration

Modify rate limit settings in `main.go`:

```go
const (
    limitPeriod = 30 * time.Second  // Time window for rate limiting
    limitCount  = 5                 // Max number of requests per IP during the window
)
```

---

## 📈 Rate Limiting Behavior

* Each **IP address** can make up to `limitCount` requests per `limitPeriod`
* Exceeding the limit results in **HTTP 429 - Too Many Requests**
* Redis keys **auto-expire** after the time window ends

---

## 🛠 Dependencies

* [Gin](https://github.com/gin-gonic/gin) – HTTP web framework
* [go-redis](https://github.com/redis/go-redis) – Redis client for Go
* [godotenv](https://github.com/joho/godotenv) – Loads `.env` into `os.Environ()`

---

## 📜 License

This project is licensed under the **MIT License**.

---

## 🤝 Contributing

Pull requests are welcome! For major changes, please [open an issue](https://github.com/yourusername/go-rate-limit/issues) first to discuss your ideas.

---

## 💬 Support

Found a bug or need help? Open an issue on the [GitHub repo](https://github.com/yourusername/go-rate-limit/issues).
