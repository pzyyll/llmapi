<style scoped>
.fade-enter-active,
.fade-leave-active {
	transition:
		opacity 0.1s,
		transform 0.1s;
}

.fade-enter-from,
.fade-leave-to {
	opacity: 0;
	transform: scale(0.8);
}
</style>

<template>
	<div>
		<h1>Welcome to the demo page 2</h1>
		<p>This is a demo page for testing purposes.</p>
		<button @click="show = !show">Toggle</button>
		<Transition name="fade">
			<div v-if="show" class="flex justify-center">
				<p class="text-2xl font-bold">API Keys</p>
			</div>
		</Transition>
	</div>
</template>

<script setup lang="ts">
const show = ref(true);
</script>
