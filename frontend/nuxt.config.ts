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
	css: ['~/assets/css/main.css'],
	nitro: {
		output: {
			publicDir: path.resolve(__dirname, '../backend/src/internal/router/dashboard/static/dist')
		}
	},
	modules: [
		'@nuxt/fonts',
		'@nuxt/eslint',
		'@pinia/nuxt',
		'unplugin-icons/nuxt',
		'reka-ui/nuxt',
		'@nuxtjs/i18n',
		'pinia-plugin-persistedstate/nuxt',
		'@nuxt/scripts',
		'@nuxtjs/turnstile',
		'@vueuse/nuxt'
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
			}
		},
		pages: {
			pattern: ['**', '!demo/**']
		}
	},
	$development: {
		runtimeConfig: {
			public: {
				apiBase: '/proxy/',
				goBase: 'http://localhost:13140/'
			}
		}
	},
	hooks: {
		'prerender:routes'({ routes }) {
			routes.clear(); // Do not generate any routes (except the defaults)
		} // This will clear all routes, including the default ones, (SPA mode, only the root index.html needed.)
	}
});
