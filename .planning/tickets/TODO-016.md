# [TODO-016] API: Categories CRUD (Organization)

**Description:**
Implement endpoints to manage task categories.

**Technical Steps:**
1. Implement `POST /api/categories`: Create a new category for the user.
2. Implement `GET /api/categories`: List categories for the user.
3. Update `models.Task` to include an optional `CategoryID`.
4. Ensure `POST /api/tasks` can accept a `category_id`.

**Acceptance Criteria:**
- Categories are unique per user.
- Tasks can be successfully linked to a category.
