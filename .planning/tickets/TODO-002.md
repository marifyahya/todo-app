# [TODO-002] Backend Dockerfile (Multi-stage)

**Description:**
Create a production-ready Dockerfile for the Go backend.

**Technical Steps:**
1. Create `backend/Dockerfile`.
2. Stage 1 (Build):
   - Use `golang:1.21-alpine`.
   - Copy `go.mod` and `go.sum`, run `go mod download`.
   - Copy source code and build using `go build -o main .`.
3. Stage 2 (Run):
   - Use `alpine:latest`.
   - Copy the binary from Stage 1.
   - Expose port 8080 and set `CMD ["./main"]`.

**Acceptance Criteria:**
- `docker build -t todos-backend ./backend` completes successfully.
- Image size is small (under 50MB).
