<script setup>
import { ref } from "vue";
import { useRouter } from "vue-router";
import { useAuthStore } from "../stores/auth";
import api from "../api/axios";

const router = useRouter();
const authStore = useAuthStore();

const form = ref({
  email: "",
  password: "",
});

const error = ref("");
const loading = ref(false);

const handleLogin = async () => {
  error.value = "";
  loading.value = true;
  try {
    const response = await api.post("/auth/login", form.value);
    authStore.setAuth(response.data.data.user, response.data.data.token);
    router.push("/dashboard");
  } catch (err) {
    error.value =
      err.response?.data?.message ||
      "Login failed. Please check your credentials.";
  } finally {
    loading.value = false;
  }
};
</script>

<template>
  <div
    class="min-h-screen bg-white flex flex-col justify-center py-8 sm:px-6 lg:px-8 font-sans"
  >
    <div class="sm:mx-auto sm:w-full sm:max-w-md px-6">
      <div class="flex justify-center">
        <div
          class="w-12 h-12 bg-black rounded-xl flex items-center justify-center shadow-lg"
        >
          <svg
            xmlns="http://www.w3.org/2000/svg"
            class="h-6 w-6 text-white"
            fill="none"
            viewBox="0 0 24 24"
            stroke="currentColor"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2m-6 9l2 2 4-4"
            />
          </svg>
        </div>
      </div>
      <h2
        class="mt-6 text-center text-2xl font-black text-black tracking-tighter"
      >
        Welcome back
      </h2>
      <p class="mt-1 text-center text-xs text-gray-400 font-medium">
        Don't have an account?
        <router-link
          to="/register"
          class="text-black font-bold hover:underline underline-offset-4"
        >
          Register here
        </router-link>
      </p>
    </div>

    <div class="mt-8 sm:mx-auto sm:w-full sm:max-w-md px-6">
      <div
        class="bg-white py-8 px-6 rounded-2xl border border-gray-100 shadow-sm"
      >
        <form class="space-y-4" @submit.prevent="handleLogin">
          <div
            v-if="error"
            class="bg-red-50 border border-red-100 p-3 rounded-xl text-center"
          >
            <p
              class="text-[10px] font-bold text-red-600 uppercase tracking-widest"
            >
              {{ error }}
            </p>
          </div>

          <div>
            <label
              for="email"
              class="block text-[10px] font-black text-gray-400 uppercase tracking-widest mb-1.5 ml-1"
            >
              Email
            </label>
            <input
              v-model="form.email"
              id="email"
              type="email"
              required
              class="appearance-none block w-full px-4 py-3 border border-gray-100 rounded-xl shadow-sm placeholder-gray-300 focus:outline-none focus:ring-1 focus:ring-black focus:border-black text-sm bg-gray-50/30"
              placeholder="you@example.com"
            />
          </div>

          <div>
            <label
              for="password"
              class="block text-[10px] font-black text-gray-400 uppercase tracking-widest mb-1.5 ml-1"
            >
              Password
            </label>
            <input
              v-model="form.password"
              id="password"
              type="password"
              required
              class="appearance-none block w-full px-4 py-3 border border-gray-100 rounded-xl shadow-sm placeholder-gray-300 focus:outline-none focus:ring-1 focus:ring-black focus:border-black text-sm bg-gray-50/30"
              placeholder="••••••••"
            />
          </div>

          <div class="pt-2">
            <button
              type="submit"
              :disabled="loading"
              class="w-full flex justify-center py-3 px-4 border border-transparent rounded-xl text-xs font-black text-white bg-black hover:bg-gray-800 transition-all disabled:opacity-50 uppercase tracking-widest"
            >
              <span v-if="loading">Signing in...</span>
              <span v-else>Sign In</span>
            </button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>
