# [TODO-011] Frontend: Vue 3 + Tailwind Setup

**Description:**
Initialize the Vue 3 project and configure the styling engine.

**Technical Steps:**
1. Run `npm create vite@latest frontend -- --template vue`.
2. Install Tailwind: `npm install -D tailwindcss postcss autoprefixer`.
3. Run `npx tailwindcss init -p`.
4. Configure `tailwind.config.js` to include all vue files.
5. Add tailwind directives to `src/style.css`.

**Acceptance Criteria:**
- `App.vue` can use Tailwind classes like `<div class="bg-blue-500 text-white">`.
