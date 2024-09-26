<!-- components/Card.vue -->
<template>
    <nuxt-link :to="`${link}`" class="text-gray-800 hover:text-blue-600">
    <div
        class="bg-white border border-gray-200 rounded-lg shadow-md hover:shadow-lg transition-shadow duration-200"
    >
        
        <div class="p-10">
        <h3 class="text-xl font-semibold mb-2 text-center">{{ title }}</h3>
        <p class="text-gray-600">{{ description }}</p>
        </div>
    </div>
    </nuxt-link>
</template>

<script setup lang="ts">
const props = defineProps<{
  title: string;
  description: string;
  image: string;
  link: string;
}>();

console.log(props.link, 'asdasd')
</script>

<style scoped>
/* Additional styles if needed */
h3:hover {
    color:black;
}
</style>
