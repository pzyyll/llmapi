<script setup lang="ts">
const props = defineProps({
	title: { type: String, default: '' }
});
</script>
<template>
	<header class="border-b-1 border-base-300 box-content flex flex-col px-5 pt-4">
		<div class="flex h-8 items-center justify-between">
			<div>
				<slot name="left">
					<h1 class="whitespace-nowrap text-xl font-bold">{{ props.title }}</h1>
				</slot>
			</div>
			<div>
				<slot name="right" />
			</div>
		</div>
		<div>
			<slot name="footer">
				<div class="pt-4"></div>
			</slot>
		</div>
	</header>
</template>
