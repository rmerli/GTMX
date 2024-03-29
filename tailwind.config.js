/** @type {import('tailwindcss').Config} */
module.exports = {
	content: [
		'src/view/**/*.templ',
		'src/view/**/*.go',
	],
	darkMode: 'class',
	theme: {
		container: {
			center: true,
		},
	},
	plugins: [
		require("@tailwindcss/typography"),
		require("daisyui"),
	],
	corePlugins: {
		preflight: true,
	}
}

