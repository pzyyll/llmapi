<!-- <script setup lang="ts">
definePageMeta({
	middleware: [
		async (to, from) => {
			const { data } = await useAPI().get<Dto.LoadSettings>(RequestPath.LoadSettings)
			if (data) {
				const { setLoadSettings } = useSettingsStore()
				setLoadSettings(data)
			}
		}
	],
})

</script> -->

<template>
	<div>
		<NuxtLayout>
			<NuxtPage />
		</NuxtLayout>
	</div>
</template>
