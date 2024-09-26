<script setup lang="ts">
const nav = [
  { label: 'Home', to: '/' },
  { label: 'Custom', to: '/custom' },
  { label: 'Dynamic', to: '/dynamic' },
  { label: 'Other', to: '/other' },
]

definePageMeta({
    layout: 'default', // Use the 'auth' layout
});
</script>


<template>
  <NuxtExample dir="features/layouts" :nav="nav">
    <NuxtLayout class="layouts">
      <NuxtPage />
    </NuxtLayout>
  </NuxtExample>
</template>

<style scoped>
.layouts {
  padding: 1rem;
}
</style>
