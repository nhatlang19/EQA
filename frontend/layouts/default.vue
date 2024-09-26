<template>
    <div>
        <Header />
        <main class="container mx-auto my-4">
            <slot />
        </main>
    </div>
</template>

<script setup>
import Header from '~/components/Header.vue';
</script>