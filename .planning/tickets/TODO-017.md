# [TODO-017] Unit Testing: Backend Services

**Description:**
Ensure business logic is reliable using automated tests.

**Technical Steps:**
1. Install `testify`: `go get github.com/stretchr/testify`.
2. Create `services/task_service_test.go`.
3. Use a mock database or SQLite in-memory for testing.
4. Write tests for:
   - Task creation validation.
   - Ownership verification logic.
   - Filtering logic for `GET /tasks`.
5. Run tests using `go test ./...`.

**Acceptance Criteria:**
- At least 80% code coverage for core business logic.
- All tests pass in the CI/CD or local environment.
