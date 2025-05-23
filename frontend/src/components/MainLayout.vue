<script setup lang="ts">
const props = defineProps({
	title: { type: String, default: '' },
	onLoading: { type: Boolean, default: false }
});
</script>

<template>
	<div class="bg-base-100 flex h-full flex-1 flex-col items-center justify-start rounded-md">
		<div class="flex h-full w-full flex-col" v-show="!props.onLoading">
			<slot />
			<!-- <MainHeader :title="props.title"/> -->
			<!-- <section class="overflow-auto h-full">
				<slot />
			</section> -->
		</div>
	</div>
</template>
