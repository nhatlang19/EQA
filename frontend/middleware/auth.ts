import { defineNuxtRouteMiddleware, navigateTo } from 'nuxt/app';

export default defineNuxtRouteMiddleware((to, from) => {
    const isAuthenticated = !!localStorage.getItem('token'); // Replace this with your actual authentication logic
console.log('aaaaaaaaaaaa', isAuthenticated);
    // If the user is not authenticated and is trying to access a protected route
    if (!isAuthenticated && to.path !== '/login') {
        return navigateTo('/login');
    }
});