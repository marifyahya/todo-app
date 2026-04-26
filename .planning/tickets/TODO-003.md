# [TODO-003] Frontend Dockerfile (Dev Mode)

**Description:**
Create a Dockerfile for the Vue.js frontend optimized for development with hot-reload.

**Technical Steps:**
1. Create `frontend/Dockerfile`.
2. Use `node:20-alpine`.
3. Set working directory to `/app`.
4. Copy `package.json`, run `npm install`.
5. Copy project files.
6. Set `CMD ["npm", "run", "dev", "--", "--host"]`.
7. Ensure port 5173 is exposed.

**Acceptance Criteria:**
- Changes made in `frontend/src` are immediately reflected in the browser (Hot Module Replacement).
