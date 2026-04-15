# AI Customer Support API (Golang)

A production-ready REST API built with Go and Gin, integrated with Google Gemini AI to provide smart customer support responses.

## 🚀 Features
- **Clean Architecture**: Separation of concerns (Handlers, Services, Models, Config).
- **AI Integration**: Uses Gemini 1.5 Flash for intent classification and response generation.
- **Robust Error Handling**: Proper HTTP status codes and JSON error responses.
- **Logging**: Request and response logging for monitoring.
- **Environment Driven**: Configuration via `.env` or environment variables.

## 🛠 Tech Stack
- **Language**: Go 1.21+
- **Framework**: Gin Gonic
- **AI SDK**: `github.com/google/generative-ai-go`
- **Config**: `github.com/joho/godotenv`

## 📋 Project Structure
```
.
├── main.go             # Entry point
├── config/             # Configuration loading
├── handlers/           # HTTP request handlers
├── services/           # Business logic & AI integration
├── models/             # Data structures
├── routes/             # API route definitions
├── .env.example        # Environment variable template
└── Dockerfile          # Containerization
```

## 🚦 Getting Started

### 1. Prerequisites
- Go 1.21 or higher installed.
- A Google AI Studio API Key (Gemini API).

### 2. Installation
1. Clone the repository.
2. Install dependencies:
   ```bash
   go mod tidy
   ```
3. Create a `.env` file from the example:
   ```bash
   cp .env.example .env
   ```
4. Add your `GEMINI_API_KEY` to the `.env` file.

### 3. Running the API
```bash
go run main.go
```
The server will start on `http://localhost:3000`.

## 🧪 API Usage

### Send a Message
**Endpoint**: `POST /message`

**Request Body**:
```json
{
  "user_id": "user_123",
  "message": "I am having trouble with my payment."
}
```

**Example Curl**:
```bash
curl -X POST http://localhost:3000/message \
     -H "Content-Type: application/json" \
     -d '{"user_id": "user_123", "message": "I am having trouble with my payment."}'
```

**Response**:
```json
{
  "response": "I'm sorry to hear you're having trouble with your payment. Could you please provide more details about the error you're seeing?",
  "category": "payment"
}
```

## 🐳 Docker
Build and run with Docker:
```bash
docker build -t ai-support-api .
docker run -p 3000:3000 --env-file .env ai-support-api
```
