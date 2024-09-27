<template>
    <div class="relative inline-block" @mouseenter="show" @mouseleave="hide">
      <slot></slot>
      <div
        v-if="visible && tooltipContent.length > 0"
        class="absolute min-w-64 bottom-full mb-2 left-1/2 transform -translate-x-1/2 bg-gray-700 text-white text-xs font-medium rounded-lg p-2 transition-opacity duration-200 z-10 whitespace-pre-wrap"
        :style="{ opacity: visible ? 1 : 0 }"
        v-html="tooltipContent"
      ></div>
    </div>
  </template>
  
  <script setup>
  import { ref, defineProps } from 'vue';
  
  const props = defineProps({
    text: {
      type: String,
      required: true,
    },
  });
  
  const visible = ref(false);
  const tooltipContent = props.text;
  
  const show = () => {
    visible.value = true;
  };
  
  const hide = () => {
    visible.value = false;
  };
  </script>
  
  <style scoped>
  /* Optional: Additional styles */
  </style>
  