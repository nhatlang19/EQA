<!-- components/Header.vue -->
<template>
    <header class="bg-white shadow">
      <nav class="container mx-auto flex justify-between items-center p-4">
        <div class="flex items-center">
          <nuxt-link to="/" class="text-xl font-semibold text-gray-800">EQA PROGRAM</nuxt-link>
        </div>
        <div class="hidden md:flex md:space-x-8">
          <nuxt-link to="/" class="text-gray-800 hover:text-blue-600">Trang chủ</nuxt-link>
          <nuxt-link to="/programs" class="text-gray-800 hover:text-blue-600">Chương trình</nuxt-link>
        </div>
        <div class="md:hidden">
          <button @click="toggleMenu" class="text-gray-800 focus:outline-none">
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16m-7 6h7" />
            </svg>
          </button>
        </div>
      </nav>
      <div v-if="menuOpen" class="md:hidden bg-white shadow-md">
        <nuxt-link to="/" class="block px-4 py-2 text-gray-800 hover:bg-gray-200">Trang chủ</nuxt-link>
        <nuxt-link to="/programs" class="block px-4 py-2 text-gray-800 hover:bg-gray-200">Chương trình</nuxt-link>
      </div>
    </header>
  </template>
  
  <script setup lang="ts">
  import { ref } from 'vue';
  
  const menuOpen = ref(false);
  
  const toggleMenu = () => {
    menuOpen.value = !menuOpen.value;
  };
  </script>
  
  <style scoped>
  /* Additional styles if necessary */
  </style>
  