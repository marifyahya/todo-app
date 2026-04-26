# [TODO-014] UI: Task Dashboard (Management)

**Description:**
The main interface where users see and manage their tasks.

**Technical Steps:**
1. Create `src/views/Dashboard.vue`.
2. On `onMounted`: call `taskStore.fetchTasks()`.
3. Use `v-for` to render a list of `TaskItem` components.
4. Implement a "New Task" modal.
5. Each task should have a checkbox (to toggle status) and a delete button (using `taskStore.deleteTask(id)`).

**Acceptance Criteria:**
- List updates automatically after a task is added or deleted.
- Loading states (spinners) are shown while data is fetching.
