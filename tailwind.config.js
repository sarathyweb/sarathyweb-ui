const tailwindui = require('@tailwindcss/ui')({
      layout: 'sidebar',
});

const defaultTheme = require('tailwindcss/defaultTheme');

module.exports = {
  theme: {
    extend: {
      fontFamily: {
        sans: ['Inter var', ...defaultTheme.fontFamily.sans],
      },
    },
  },
  variants: {},
  plugins: [
    tailwindui,
  ]
};
