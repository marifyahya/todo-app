# [TODO-007] API: POST /auth/login (JWT Implementation)

**Description:**
Handle user login and issue a JWT token.

**Technical Steps:**
1. Install JWT library: `go get github.com/golang-jwt/jwt/v5`.
2. In `handlers/auth_handler.go`:
   - Find user by Email: `db.Where("email = ?", input.Email).First(&user)`.
   - Compare password: `bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))`.
   - If match, create JWT claim with `user_id` and expiration (e.g., 24 hours).
   - Sign token with a `JWT_SECRET` from `.env`.
3. Return the token in JSON: `{"token": "..."}`.

**Acceptance Criteria:**
- Login with correct credentials returns a JWT token.
- Login with wrong password returns 401 Unauthorized.
