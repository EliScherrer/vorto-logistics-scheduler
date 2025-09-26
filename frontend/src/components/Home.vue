<template>
  <div class="flex items-center justify-center min-h-screen">
    <div class="bg-white p-8 rounded-lg shadow-md">
      <h1 class="text-3xl font-bold text-center mb-8">Logistics Scheduler</h1>
      <div class="space-y-4">
        <router-link
          to="/admin/login"
          class="block w-full bg-blue-500 hover:bg-blue-600 text-white font-bold py-3 px-6 rounded text-center"
        >
          Admin Portal
        </router-link>
        <router-link
          to="/driver/login"
          class="block w-full bg-green-500 hover:bg-green-600 text-white font-bold py-3 px-6 rounded text-center"
        >
          Driver Portal
        </router-link>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
// Home component
</script>