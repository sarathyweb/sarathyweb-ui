const tailwindcss = require("tailwindcss");
const autoprefixer = require("autoprefixer");
const purgecss = require('@fullhuman/postcss-purgecss');

module.exports = {
	  plugins: [
		      tailwindcss,
		      autoprefixer,
		      purgecss({
		      	content: ['./**/*.html'],
		      	defaultExtractor: content => content.match(/[\w-/.:]+(?<!:)/g) || []
		      })]
};
