/** @type {import('tailwindcss').Config} */
export default {
  content: [
    './components/**/*.{vue,js}',
    './layouts/**/*.{vue,js}',
    './pages/**/*.{vue,js}',
    './app.vue',
  ],
  theme: {
    extend: {
      width: {
        '1/15': '6.6667%', // Approximately 6.67% for each of 15 columns
      },
    },
  },
  plugins: [],
}

