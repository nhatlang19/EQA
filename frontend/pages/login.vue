<template>
    <div class="flex items-center justify-center min-h-screen bg-gray-100">
        <div class="w-full max-w-md p-8 space-y-4 bg-white rounded-lg shadow-md">
            <h2 class="text-2xl font-bold text-center">Login</h2>
            <form @submit.prevent="handleLogin" class="space-y-4">
                <div>
                    <label for="email" class="block text-sm font-medium text-gray-700">Email</label>
                    <input
                        type="email"
                        id="email"
                        v-model="email"
                        required
                        class="block w-full mt-1 p-2 border rounded-md focus:outline-none focus:ring focus:ring-blue-300"
                        placeholder="you@example.com"
                    />
                </div>
                <div>
                    <label for="password" class="block text-sm font-medium text-gray-700">Password</label>
                    <input
                        type="password"
                        id="password"
                        v-model="password"
                        required
                        class="block w-full mt-1 p-2 border rounded-md focus:outline-none focus:ring focus:ring-blue-300"
                        placeholder="••••••••"
                    />
                </div>
                <button
                    type="submit"
                    class="w-full px-4 py-2 font-bold text-white bg-blue-600 rounded-md hover:bg-blue-700 focus:outline-none focus:ring focus:ring-blue-300"
                >
                    Login
                </button>
            </form>
            <p class="text-center">
                <a href="#" class="text-sm text-blue-600 hover:underline">Forgot Password?</a>
            </p>
        </div>
    </div>
</template>

<script setup>
import { ref } from 'vue';

const email = ref('');
const password = ref('');

const handleLogin = () => {
    // Handle login logic here (e.g., API call)
    console.log('Email:', email.value);
    console.log('Password:', password.value);
};

definePageMeta({
    layout: false
});
</script>

<style scoped>
/* Additional styling if necessary */
</style>
