/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./cmd/web/**/*.templ"],
  theme: {
    extend: {},
  },
  plugins: [require("@tailwindcss/forms"), require("@tailwindcss/typography")],
};

// tw -i ./cmd/web/css/input.css -o ./cmd/web/css/output.css --watch
