# Todos App

A full-stack Todo application built with Go (Backend) and Vue 3 (Frontend).

## Prerequisites

- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)

## Getting Started

1. Clone the repository.
2. Copy the example environment file:
   ```bash
   cp .env.example .env
   ```
3. Start the application using Docker Compose:
   ```bash
   docker-compose up -d
   ```

## Services & Ports

- **Frontend**: [http://localhost:5174](http://localhost:5174)
- **Backend**: [http://localhost:8081](http://localhost:8081)
- **Database**: `localhost:5432` (PostgreSQL)

> **Note**: Default ports (5173/8080) were adjusted to 5174/8081 to avoid local environment conflicts.

## Project Structure

- `/backend`: Go backend source code.
- `/frontend`: Vue 3 frontend source code.
- `/docker-compose.yml`: Infrastructure definition.
- `/.planning`: Technical tickets and implementation plans.
