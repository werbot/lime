/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./src/**/*.{vue,js,ts}",
    "./index.html",
  ],
  safelist: [
    {
      pattern: /(bg|border)-(gray|green|red|blue|yellow)-(100|500|600)$/,
      variants: ['after', 'hover', 'focus', 'peer-checked', 'peer-checked:after'],
    },
  ],
  theme: {
    extend: {},
  },
  plugins: [require("@tailwindcss/forms")],
};