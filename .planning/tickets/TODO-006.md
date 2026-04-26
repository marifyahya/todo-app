# [TODO-006] API: POST /auth/register (Implementation Detail)

**Description:**
Create the endpoint to allow new users to sign up.

**Technical Steps:**
1. Create `handlers/auth_handler.go`.
2. Create a `RegisterRequest` struct with `Email`, `Password`, and `Name` tags for JSON.
3. In the handler:
   - Bind JSON input to the struct.
   - Validate that Email is not empty and is valid.
   - Hash the password using `golang.org/x/crypto/bcrypt` with `DefaultCost`.
   - Save the user to the database using `db.Create(&user)`.
   - Handle errors (e.g., if email already exists, return 400).
4. Register the route in `main.go`: `r.POST("/api/auth/register", handlers.Register)`.

**Acceptance Criteria:**
- Sending a POST request to `/api/auth/register` with valid JSON returns 201 Created.
- Database shows a new record with a hashed password (starts with `$2a$`).
