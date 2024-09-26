<template>
    <div class="container mx-auto p-6">
    <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
      <Card
        v-for="(card, index) in cards"
        :key="index"
        :title="card.title"
        :description="card.description"
        :image="card.image"
        :link="card.link"
      />
    </div>
  </div>
</template>

<script setup lang="ts">
import Card from '~/components/Card.vue';

const cards = [
  {
    title: 'Reminder',
    description: '',
    image: 'https://via.placeholder.com/300x200?text=Card+1',
    link: '/reminders/plan-result'
  },
];

// definePageMeta({
//     middleware: 'auth'
// });
</script>