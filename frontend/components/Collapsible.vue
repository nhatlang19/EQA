<template>
    <div class="border rounded-md shadow-md">
      <button
        @click="toggle"
        class="flex justify-between items-center w-full p-4 text-left bg-gray-200 hover:bg-gray-300 focus:outline-none"
      >
        <span>{{ title }}</span>
        <svg
          xmlns="http://www.w3.org/2000/svg"
          class="w-6 h-6 transition-transform"
          :class="isOpen ? 'rotate-180' : 'rotate-0'"
          fill="none"
          viewBox="0 0 24 24"
          stroke="currentColor"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            d="M19 9l-7 7-7-7"
          />
        </svg>
      </button>
      <transition name="fade">
        <div v-if="isOpen" class="p-4 bg-white">
          <slot />
        </div>
      </transition>
    </div>
  </template>
  
  <script setup>
  import { ref } from 'vue';
  
  defineProps({
    title: {
      type: String,
      required: true,
    },
  });
  
  const isOpen = ref(false);
  const toggle = () => {
    isOpen.value = !isOpen.value;
  };
  </script>
  
  <style>
  .fade-enter-active, .fade-leave-active {
    transition: opacity 0.3s;
  }
  .fade-enter, .fade-leave-to /* .fade-leave-active in <2.1.8 */ {
    opacity: 0;
  }
  </style>
  