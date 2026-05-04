# Step 2: Create Main Entry Point

## What we did:

1. ✅ Created `cmd/api/main.go` - The main application entry point
2. ✅ Built a basic HTTP server using Go's standard library (`net/http`)
3. ✅ Added essential middleware:
   - **CORS**: Enables Cross-Origin Resource Sharing for web clients
   - **Logging**: Logs all HTTP requests to console
4. ✅ Created basic routes:
   - `GET /health` - Health check endpoint
   - `GET /api/v1/` - API root endpoint
5. ✅ Used proper JSON responses with struct types

## Code Explanation:

### Main Components:

**Response Struct:**
```go
type Response struct {
    Status  string `json:"status,omitempty"`
    Message string `json:"message"`
    Version string `json:"version,omitempty"`
}
```

**Middleware Functions:**
- `corsMiddleware`: Handles CORS headers and preflight requests
- `loggingMiddleware`: Logs each request method and path

**Route Handlers:**
- `healthHandler`: Returns server status
- `apiRootHandler`: Returns API welcome message

**Server Setup:**
```go
mux := http.NewServeMux()
handler := corsMiddleware(loggingMiddleware(mux))
http.ListenAndServe(":8080", handler)
```

## Testing the Server:

To test this step, run:
```bash
export PATH="/usr/local/go/bin:$PATH"
go run cmd/api/main.go
```

Then visit in browser or use curl:
- `http://localhost:8080/health` - Should return `{"status":"ok","message":"Simple Go CRUD API is running"}`
- `http://localhost:8080/api/v1/` - Should return `{"message":"Welcome to Simple Go CRUD API","version":"v1"}`

## Why Standard Library Instead of Fiber?

Due to network connectivity issues, we're using Go's standard `net/http` package for now. This is actually a great learning opportunity! In future steps, we can easily replace this with Fiber or Gin when dependencies are available.

## Next Steps:

Once you review and test this, we'll move to **Step 3**: Setup database connection with GORM and SQLite.

Ready to proceed? 🚀
