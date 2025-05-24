<script setup lang="ts">
const props = defineProps({
	model: {
		type: Object,
		required: true
	}
});

const tips = ref(false);
let { start, stop } = useTimeoutFn(() => {
	tips.value = false;
}, 500);

const onClick = () => {
	stop();
	// Logic to handle click event
	console.log('Model card clicked:', props.model);
	tips.value = true;
	start();
};

const onClickEdit = () => {
	// Logic to handle edit button click
	console.log('Edit button clicked for model:', props.model);
};
</script>
<template>
	<div class="d-card relative shadow-sm" @click="onClick">
		<div class="d-card-body">
			<h2 class="d-card-title">{{ props.model.name }}</h2>
			<p class="truncate">{{ props.model.desc }}</p>
			<div class="d-card-actions justify-end">
				<button class="d-btn d-btn-ghost d-btn-xs" @click.stop="onClickEdit">
					<MaterialSymbolsEditSquareOutlineRounded class="size-4" />
				</button>
			</div>
		</div>
		<Transition name="tips">
			<div
				class="d-badge d-badge-success d-badge-xs absolute left-1/2 top-3 z-10 -translate-x-1/2 transform p-1 text-xs"
				v-if="tips"
			>
				Copied
			</div>
		</Transition>
	</div>
</template>

<style scoped>
.tips-enter-active,
.tips-leave-active {
	transition: opacity 0.1s ease, transform 0.1s ease;
}
.tips-enter-from,
.tips-leave-to {
	opacity: 0;
	transform: translateY(-10px);
}
</style>
