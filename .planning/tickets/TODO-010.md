# [TODO-010] API: PUT /tasks/:id (Update & Ownership)

**Description:**
Update an existing task's status or details.

**Technical Steps:**
1. Retrieve `id` from URL params.
2. Find task in DB: `db.First(&task, id)`.
3. **CRITICAL**: Check if `task.UserID == currentUserID`. If not, return 403 Forbidden.
4. Updates: `db.Model(&task).Updates(map[string]interface{}{"status": input.Status, "title": input.Title})`.
5. Return the updated task.

**Acceptance Criteria:**
- User A cannot update User B's task.
- Changing status to "completed" works.
