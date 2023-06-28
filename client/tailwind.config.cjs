/** @type {import('tailwindcss').Config} */
module.exports = {
    content: ["./src/**/*.{html,js,svelte,ts}"],
    theme: {
        extend: {
            backgroundImage: {
                "hive-pattern": "url('./bg.jpeg')",
            },
        },
    },
    plugins: [],
};
