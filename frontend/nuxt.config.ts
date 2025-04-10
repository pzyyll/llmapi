import path from 'path';

// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
	compatibilityDate: '2024-11-01',
	devtools: { enabled: true },
	ssr: false,
	app: {
		baseURL: '/dashboard'
	},
	nitro: {
		output: {
			publicDir: path.resolve(__dirname, '../backend/src/static/dist')
		}
	},
	modules: [
		'@nuxt/eslint'
		// '@nuxt/icon',
		// '@nuxt/fonts',
		// '@nuxt/image'
	]
});
