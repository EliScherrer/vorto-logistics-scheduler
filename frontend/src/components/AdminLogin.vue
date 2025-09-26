<template>
  <div class="flex items-center justify-center min-h-screen">
    <div class="bg-white p-8 rounded-lg shadow-md w-96">
      <h2 class="text-2xl font-bold text-center mb-6">Admin Login</h2>
      <form @submit.prevent="handleLogin">
        <div class="mb-4">
          <label for="username" class="block text-sm font-medium text-gray-700 mb-2">
            Username
          </label>
          <input
            id="username"
            v-model="username"
            type="text"
            required
            class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
          />
        </div>
        <button
          type="submit"
          class="w-full bg-blue-500 hover:bg-blue-600 text-white font-bold py-2 px-4 rounded"
        >
          Login
        </button>
      </form>
      <div class="mt-4 text-center">
        <router-link to="/" class="text-blue-500 hover:text-blue-600">
          Back to Home
        </router-link>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

const username = ref('')

const handleLogin = () => {
  console.log('Admin login:', username.value)
  // TODO: Implement admin login logic
}
</script>