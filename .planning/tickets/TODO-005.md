# [TODO-005] OpenAPI/Swagger YAML Setup

**Description:**
Define the API contract using OpenAPI 3.0.

**Technical Steps:**
1. Create a folder `docs/`.
2. Create `docs/swagger.yaml`.
3. Define the `/auth/register` and `/auth/login` endpoints first.
4. Specify request bodies and response schemas (200, 400, 401).
5. (Optional) Install `github.com/swaggo/swag` and use annotations in your Go code to generate the YAML automatically.

**Acceptance Criteria:**
- The YAML file can be imported into Swagger Editor or Postman without errors.
