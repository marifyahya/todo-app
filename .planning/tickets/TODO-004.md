# [TODO-004] Database Connection & GORM Models (Technical)

**Description:**
Initialize the Go project and connect it to the PostgreSQL database using GORM.

**Technical Steps:**
1. Initialize Go module: `cd backend && go mod init github.com/user/todos-api`.
2. Install dependencies: `go get gorm.io/gorm gorm.io/driver/postgres`.
3. Create file `database/db.go`:
   - Use `os.Getenv` to read DB credentials from environment variables.
   - Use `gorm.Open` to create a connection pool.
4. Create `models/user.go`, `models/task.go`:
   - `User` struct: `ID`, `Email` (unique), `Password`, `Name`, `CreatedAt`.
   - `Task` struct: `ID`, `Title`, `Description`, `Status` (default: pending), `UserID`, `CreatedAt`.
5. Run `db.AutoMigrate(&models.User{}, &models.Task{})` in the `InitDB` function.

**Acceptance Criteria:**
- Backend starts and logs "Connected to Database".
- Tables `users` and `tasks` are automatically created in PostgreSQL.
