import path from 'path';
import Icons from 'unplugin-icons/vite';
import IconsResolver from 'unplugin-icons/resolver';
import UnpluginViteComponents from 'unplugin-vue-components/vite';
import tailwindcss from '@tailwindcss/vite';
import type { NuxtPage } from 'nuxt/schema';

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
		'pinia-plugin-persistedstate/nuxt',
		'@nuxt/scripts',
		'@nuxtjs/turnstile'
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
			},
			{
				code: 'jp',
				name: '日本語',
				file: 'jp.json'
			},
			{
				code: 'de',
				name: 'Deutsch',
				file: 'de.json'
			}
		],
		defaultLocale: 'zh',
		restructureDir: 'src/i18n',
		bundle: {
			optimizeTranslationDirective: false
		},
		strategy: 'no_prefix'
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
				apiBase: '',
				turnstile: {
					// This can be overridden at runtime via the NUXT_PUBLIC_TURNSTILE_SITE_KEY environment variable.
					siteKey: ''
				}
			},
		},
		pages: {
			pattern: ['!**/demo/**']
		}
	},
	$development: {
		runtimeConfig: {
			public: {
				apiBase: '/dashboard/proxy/',
				goBase: 'http://localhost:13140/',
				turnstile: {
					// This can be overridden at runtime via the NUXT_PUBLIC_TURNSTILE_SITE_KEY environment variable.
					siteKey: '0x4AAAAAABdJ7yz22densFFG'
				}
			}
		}
	},
	hooks: {
		'pages:extend': (pages) => {
			function setMiddleware(pages: NuxtPage[]) {
				for (const page of pages) {
					// console.log(page);
					if (page.name === 'index') {
						page.meta ||= {};
						// middleware appened `mid_auth`
						// Ensure page.meta.middleware is an array
						// Nuxt middleware can be a string, an array, or undefined.
						if (typeof page.meta.middleware === 'string') {
							// If it's a single middleware string, convert it to an array
							page.meta.middleware = [page.meta.middleware];
						} else if (!Array.isArray(page.meta.middleware)) {
							// If it's undefined or not an array (and not a string), initialize as an empty array
							page.meta.middleware = [];
						}

						// Now, page.meta.middleware is guaranteed to be an array.
						// Add 'mid-auth' if it's not already present.
						if (!page.meta.middleware.includes('mid-auth')) {
							page.meta.middleware.push('mid-auth');
						}
					}
					if (page.children) {
						setMiddleware(page.children);
					}
				}
			}
			setMiddleware(pages);
		}
	}
});
