<script setup lang="ts">
const props = defineProps({
	showTitle: { type: Boolean, default: true },
	title: { type: String, default: 'dashboard' },
	highlighted: { type: Boolean, default: false }
});
</script>

<template>
	<button
		class="d-btn d-btn-ghost gap-0 p-0 d-btn-sm"
		:class="{
			'd-btn-active': props.highlighted
		}"
	>
		<span class="mx-2 flex items-center justify-center">
			<slot name="icon">
				<MaterialSymbolsListsRounded class="size-6" />
			</slot>
		</span>
		<Transition>
			<span v-if="props.showTitle" class="w-32 whitespace-nowrap text-left">{{ props.title }}</span>
		</Transition>
	</button>
</template>

<style scoped>
.v-enter-active,
.v-leave-active {
	transition: all 0.25s ease;
}

.v-enter-from,
.v-leave-to {
	width: 0;
	opacity: 0;
}
</style>
