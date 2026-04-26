# [TODO-008] API: GET /tasks (Filtering & Auth)

**Description:**
Fetch tasks belonging only to the authenticated user.

**Technical Steps:**
1. Create a Middleware `middleware/auth.go` to verify the JWT token in the `Authorization` header.
2. In the handler:
   - Get `user_id` from the JWT claims (set by middleware).
   - Query DB: `db.Where("user_id = ?", userID).Find(&tasks)`.
   - Add query params for status: `?status=completed`. Use `if status != "" { query = query.Where("status = ?", status) }`.

**Acceptance Criteria:**
- Request without `Authorization` header returns 401.
- Request returns only the tasks created by the logged-in user.
