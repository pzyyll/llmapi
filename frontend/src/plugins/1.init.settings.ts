export default defineNuxtPlugin((nuxtApp) => {
	const settingsStore = useSettingsStore();

	useAPI()
		.get<Dto.LoadSettings>(RequestPath.LoadSettings)
		.then((res) => {
      console.log('Settings loaded successfully:', res.data);
			settingsStore.setLoadSettings(res.data);
		})
		.catch((err) => {
			console.error('Error loading settings:', err);
			settingsStore.clearSettings();
		})
		.finally(() => {
			console.log('Settings loaded');
		});
});
