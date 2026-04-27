<script setup>
import { ref } from "vue";
import { useRouter } from "vue-router";
import api from "../api/axios";

const router = useRouter();

const form = ref({
  name: "",
  email: "",
  password: "",
  confirm_password: "",
});

const error = ref("");
const loading = ref(false);

const handleRegister = async () => {
  if (form.value.password !== form.value.confirm_password) {
    error.value = "Passwords do not match.";
    return;
  }

  error.value = "";
  loading.value = true;
  try {
    await api.post("/auth/register", {
      name: form.value.name,
      email: form.value.email,
      password: form.value.password,
    });
    router.push("/login");
  } catch (err) {
    error.value =
      err.response?.data?.message || "Registration failed. Please try again.";
  } finally {
    loading.value = false;
  }
};
</script>

<template>
  <div class="min-h-screen bg-white flex flex-col justify-center py-8 sm:px-6 lg:px-8 font-sans">
    <div class="sm:mx-auto sm:w-full sm:max-w-md px-6">
      <div class="flex justify-center">
        <div class="w-12 h-12 bg-black rounded-xl flex items-center justify-center shadow-lg">
           <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6 text-white" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 9v3m0 0v3m0-3h3m-3 0h-3m-2-5a4 4 0 11-8 0 4 4 0 018 0zM3 20a6 6 0 0112 0v1H3v-1z" />
          </svg>
        </div>
      </div>
      <h2 class="mt-6 text-center text-2xl font-black text-black tracking-tighter">
        Get Started
      </h2>
      <p class="mt-1 text-center text-xs text-gray-400 font-medium">
        Already have an account?
        <router-link to="/login" class="text-black font-bold hover:underline underline-offset-4">
          Login here
        </router-link>
      </p>
    </div>

    <div class="mt-8 sm:mx-auto sm:w-full sm:max-w-md px-6">
      <div class="bg-white py-8 px-6 rounded-2xl border border-gray-100 shadow-sm">
        <form class="space-y-4" @submit.prevent="handleRegister">
          <div v-if="error" class="bg-red-50 border border-red-100 p-3 rounded-xl text-center">
            <p class="text-[10px] font-bold text-red-600 uppercase tracking-widest">{{ error }}</p>
          </div>

          <div>
            <label for="name" class="block text-[10px] font-black text-gray-400 uppercase tracking-widest mb-1.5 ml-1"> Full Name </label>
            <input v-model="form.name" id="name" type="text" required class="appearance-none block w-full px-4 py-3 border border-gray-100 rounded-xl shadow-sm placeholder-gray-300 focus:outline-none focus:ring-1 focus:ring-black focus:border-black text-sm bg-gray-50/30" placeholder="John Doe" />
          </div>

          <div>
            <label for="email" class="block text-[10px] font-black text-gray-400 uppercase tracking-widest mb-1.5 ml-1"> Email </label>
            <input v-model="form.email" id="email" type="email" required class="appearance-none block w-full px-4 py-3 border border-gray-100 rounded-xl shadow-sm placeholder-gray-300 focus:outline-none focus:ring-1 focus:ring-black focus:border-black text-sm bg-gray-50/30" placeholder="you@example.com" />
          </div>

          <div class="grid grid-cols-2 gap-3">
            <div>
              <label for="password" class="block text-[10px] font-black text-gray-400 uppercase tracking-widest mb-1.5 ml-1"> Password </label>
              <input v-model="form.password" id="password" type="password" required class="appearance-none block w-full px-3 py-3 border border-gray-100 rounded-xl shadow-sm placeholder-gray-300 focus:outline-none focus:ring-1 focus:ring-black focus:border-black text-sm bg-gray-50/30" placeholder="••••••••" />
            </div>
            <div>
              <label for="confirm_password" class="block text-[10px] font-black text-gray-400 uppercase tracking-widest mb-1.5 ml-1"> Confirm </label>
              <input v-model="form.confirm_password" id="confirm_password" type="password" required class="appearance-none block w-full px-3 py-3 border border-gray-100 rounded-xl shadow-sm placeholder-gray-300 focus:outline-none focus:ring-1 focus:ring-black focus:border-black text-sm bg-gray-50/30" placeholder="••••••••" />
            </div>
          </div>

          <div class="pt-2">
            <button type="submit" :disabled="loading" class="w-full flex justify-center py-3 px-4 border border-transparent rounded-xl text-xs font-black text-white bg-black hover:bg-gray-800 transition-all disabled:opacity-50 uppercase tracking-widest">
              <span v-if="loading">Creating...</span>
              <span v-else>Register</span>
            </button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>
