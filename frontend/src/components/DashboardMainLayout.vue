<script setup lang="ts">
const props = defineProps({
	title: { type: String, default: '' }
});

</script>

<template>
	<div class="flex-1 flex flex-col bg-base-100 rounded-md">
		<header class="p-2 border-b-1 border-base-300">
			<NavHeader>
				<template #left>
					<slot name="header-left">
						<h1 class="text-xl font-bold">{{ props.title }}</h1>
					</slot>
				</template>
			</NavHeader>
		</header>
		<main class="p-4">
			<slot />
			<!-- This is where the page content will be injected -->
		</main>
	</div>
</template>
