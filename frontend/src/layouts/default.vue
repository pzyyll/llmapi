<script setup lang="ts">
useHead({
	htmlAttrs: {
		class: 'bg-base-200 h-screen'
	},
	bodyAttrs: {
		class: 'bg-base-200 h-screen'
	}
});
</script>

<template>
	<div class="flex h-screen w-full flex-col">
		<slot />
	</div>
</template>
