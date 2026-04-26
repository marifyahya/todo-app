# [TODO-001] Setup Docker Compose & Environment (Low-Level)

**Description:**
Set up the base infrastructure using Docker Compose. This ensures everyone on the team has the same Database, Backend, and Frontend environment.

**Technical Steps:**
1. Create a `.env` file in the root directory.
2. Define the following variables in `.env`:
   - `DB_HOST=db`
   - `DB_USER=user`
   - `DB_PASSWORD=password`
   - `DB_NAME=todos`
   - `DB_PORT=5432`
3. Create `docker-compose.yml` in the root.
4. Define 3 services:
   - `db`: Use image `postgres:15-alpine`. Map port `5432:5432`. Use volumes for persistence.
   - `backend`: Build from `./backend` directory. Map port `8080:8080`. Link to `db`.
   - `frontend`: Build from `./frontend` directory. Map port `5173:5173`.
5. Create empty directories `/backend` and `/frontend` so Docker doesn't fail.

**Acceptance Criteria:**
- Running `docker-compose up -d` starts all 3 services without errors.
- `docker ps` shows `todos-db`, `todos-backend`, and `todos-frontend` running.
- You can connect to the database using a tool (like TablePlus or DBeaver) on `localhost:5432`.
