# Todo Application Planning

## 1. Project Overview
A robust Todo application to help users manage their daily tasks efficiently. The system will consist of a high-performance RESTful API built with Golang and a modern web frontend built with Vue.js, all fully containerized using Docker.

## 2. Core Features
- **User Authentication**: Register, Login, and JWT-based session management.
- **Task Management (CRUD)**:
  - Create tasks with title, description, priority, and due date.
  - List tasks with filtering (by status, priority, category) and sorting.
  - Update task details or toggle completion status.
  - Delete tasks.
- **Categorization**: Group tasks by categories.
- **Priority Levels**: Low, Medium, High.
- **API Documentation**: Interactive OpenAPI/Swagger documentation.
- **Reliability**: Comprehensive Unit Testing for backend logic.
- **Infrastructure**: Fully Dockerized development and production environments.

## 3. Tech Stack
- **Backend**: **Golang** (Gin or Echo).
- **Database**: **PostgreSQL**.
- **ORM**: GORM.
- **API Specification**: OpenAPI 3.0 (YAML).
- **Testing**: Go `testing` + `testify`.
- **Frontend**: **Vue.js 3** (Vite) + **Tailwind CSS**.
- **Containerization**: **Docker & Docker Compose**.
- **State Management**: Pinia.

## 4. Database Schema
### User Table
- `id`: uint (Primary Key)
- `email`: string (Unique)
- `password`: string (Hashed)
- `name`: string
- `created_at`: time.Time

### Task Table
- `id`: uint (Primary Key)
- `title`: string
- `description`: text
- `status`: string (pending, in_progress, completed)
- `priority`: string (low, medium, high)
- `due_date`: time.Time
- `user_id`: uint (Foreign Key)
- `category_id`: uint (Foreign Key)

### Category Table
- `id`: uint (Primary Key)
- `name`: string
- `user_id`: uint (Foreign Key)

## 5. API Endpoints & Specification
- **OpenAPI**: A `swagger.yaml` file will define all endpoints.
- **Endpoints**:
  - `POST /auth/register`
  - `POST /auth/login`
  - `GET /tasks`
  - `POST /tasks`
  - `PUT /tasks/:id`
  - `DELETE /tasks/:id`

## 6. Implementation Steps
1.  **Phase 0: Environment & Dockerization (MANDATORY START)**:
    - Set up `docker-compose.yml` to orchestrate Backend, Frontend, and PostgreSQL.
    - Create `Dockerfile` for the Golang Backend (multi-stage build).
    - Create `Dockerfile` for the Vue.js Frontend.
    - Configure `.env` files for environment variables.
2.  **Phase 1: Backend Initialization**:
    - Initialize Go module and project structure.
    - Set up GORM with PostgreSQL.
    - Create initial **OpenAPI YAML** spec.
3.  **Phase 2: Backend Core & Testing**:
    - Implement JWT Auth and Middlewares.
    - Implement CRUD logic for Tasks and Categories.
    - Write and run **Unit Tests**.
4.  **Phase 3: Frontend Initialization**:
    - Initialize Vue.js 3 with Vite and Tailwind CSS.
    - Set up Pinia and Axios.
5.  **Phase 4: UI/UX & Integration**:
    - Build authentication and dashboard views.
    - Connect to Backend via Docker network.
    - Add responsive styling and transitions.
6.  **Phase 5: Finalization**:
    - Optimize Docker images.
    - Final end-to-end testing and documentation.
