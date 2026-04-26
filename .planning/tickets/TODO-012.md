# [TODO-012] Frontend: Pinia & Axios Setup

**Description:**
Set up global state management and HTTP client.

**Technical Steps:**
1. Install: `npm install pinia axios`.
2. Create `src/stores/auth.js`:
   - State: `user`, `token`.
   - Action: `login(credentials)`, `logout()`.
3. Create `src/api/axios.js`:
   - Set `baseURL` to `http://localhost:8080/api`.
   - Add Interceptor: If `token` exists in store, add `Authorization: Bearer <token>` to headers.

**Acceptance Criteria:**
- Once logged in, every Axios request automatically sends the JWT token.
