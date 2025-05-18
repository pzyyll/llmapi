import { defineStore } from 'pinia';

export const useSettingsStore = defineStore('settings', () => {
	const loadSettings = ref<Dto.LoadSettings | null>(null);

	const turnstileSiteKey = computed(() => {
		return loadSettings.value?.turnstile_site_key || '';
	});
	const isTurnstileEnabled = computed(() => {
		return !!turnstileSiteKey.value;
	});

	const setLoadSettings = (newSettings: Dto.LoadSettings) => {
		loadSettings.value = newSettings;
	};

	const clearSettings = () => {
		loadSettings.value = null;
	};

	return {
		loadSettings,

		turnstileSiteKey,
		isTurnstileEnabled,

		setLoadSettings,
		clearSettings
	};
});
