import path from 'path';
import Icons from 'unplugin-icons/vite';
import IconsResolver from 'unplugin-icons/resolver';
import UnpluginViteComponents from 'unplugin-vue-components/vite';
import tailwindcss from '@tailwindcss/vite';

// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
	compatibilityDate: '2024-11-01',
	devtools: { enabled: true },
	ssr: false,
	srcDir: 'src',
	app: {
		baseURL: '/dashboard'
	},
	css: ['~/assets/css/main.css'],
	nitro: {
		output: {
			publicDir: path.resolve(__dirname, '../backend/src/internal/router/dashboard/static/dist')
		}
	},
	modules: [
		// '@nuxt/image'
		'@nuxt/fonts',
		'@nuxt/eslint',
		'@pinia/nuxt',
		'unplugin-icons/nuxt',
		'reka-ui/nuxt',
		'@nuxtjs/i18n',
		'nuxt-auth-utils',
		'pinia-plugin-persistedstate/nuxt'
	],
	vite: {
		plugins: [
			tailwindcss(),
			UnpluginViteComponents({
				resolvers: [
					IconsResolver({
						prefix: '',
						strict: true
					})
				]
			}),
			Icons({
				autoInstall: true
			})
		]
	},
	i18n: {
		locales: [
			{
				code: 'en',
				name: 'English',
				file: 'en.json'
			},
			{
				code: 'zh',
				name: '简体中文',
				file: 'zh.json'
			}
		],
		defaultLocale: 'zh',
		restructureDir: 'src/i18n',
		bundle: {
			optimizeTranslationDirective: false
		},
		strategy: 'no_prefix',
		// detectBrowserLanguage: {
		// 	useCookie: true,
		// 	cookieCrossOrigin: false,
		// 	fallbackLocale: 'en'
		// }
	},
	devServer: {
		port: 13001,
		host: 'localhost'
	},
	$production: {
		runtimeConfig: {
			public: {
				apiBase: ''
			}
		},
		pages: {
			pattern: ["!**/demo/**"]
		}
	},
	$development: {
		runtimeConfig: {
			public: {
				apiBase: 'http://localhost:13001'
			}
		}
	}
});
