# [TODO-009] API: POST /tasks (Creation Logic)

**Description:**
Allow users to create new tasks with full details (title, description, priority, due date).

**Technical Steps:**
1. In `handlers/task_handler.go`:
   - Bind JSON input: `Title`, `Description`, `Priority` (low, medium, high), and `DueDate` (ISO 8601 string).
   - Retrieve `UserID` from context.
   - Create task object: `task := models.Task{Title: input.Title, Description: input.Description, Priority: input.Priority, DueDate: input.DueDate, UserID: userID, Status: "pending"}`.
   - Save to DB: `db.Create(&task)`.
2. Return the created task with 201 status.

**Acceptance Criteria:**
- Successfully creates a task with priority and due date.
- Default status is "pending".
- Fails if Title is empty.
