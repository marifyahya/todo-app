# [TODO-015] API: DELETE /tasks/:id (Removal)

**Description:**
Allow users to delete their own tasks.

**Technical Steps:**
1. Retrieve `id` from URL params.
2. Find task: `db.First(&task, id)`.
3. **CRITICAL**: Verify `task.UserID == currentUserID`.
4. Delete: `db.Delete(&task)`.
5. Return 204 No Content on success.

**Acceptance Criteria:**
- Task is removed from DB.
- User cannot delete another user's task.
