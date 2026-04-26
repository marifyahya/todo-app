# [TODO-013] UI: Login & Register Pages (Design)

**Description:**
Create the entry point for users with a premium look.

**Technical Steps:**
1. Create `src/views/Login.vue` and `src/views/Register.vue`.
2. Use a centered card layout.
3. Use `v-model` to bind form inputs to local state.
4. On submit: call `authStore.login(form)` and redirect to `/dashboard` on success.
5. Display error messages in a red alert box if the API returns 401.

**Acceptance Criteria:**
- Forms are responsive (look good on mobile).
- Input fields have focus states and clear labels.
