<script setup>
const props = defineProps({
	name: {
		type: String,
		default: ''
	},
	size: {
		type: String,
		default: 'xs'
	},
});
</script>

<template>
	<button
		class="d-avatar d-avatar-placeholder d-btn d-btn-accent d-btn-circle"
		:class="`d-btn-${props.size}`"
	>
		<span class="text-base-100 text-[130%]">{{ props.name.charAt(0)?.toUpperCase() || '?' }}</span>
	</button>
</template>
