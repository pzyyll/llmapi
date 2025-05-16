<script setup lang="ts">
const props = defineProps({
	title: { type: String, default: '' },
	onLoading: { type: Boolean, default: false }
});
</script>

<template>
	<div class="bg-base-100 flex flex-1 flex-col justify-start items-center rounded-md">
		<div class="flex flex-col w-full" v-show="!props.onLoading">
			<header class="border-b-1 border-base-300 p-2">
				<NavHeader>
					<template #left>
						<div>
							<slot name="header-left">
								<h1 class="text-xl font-bold">{{ props.title }}</h1>
							</slot>
						</div>
					</template>
					<template #right>
						<div>
							<slot name="header-right" />
						</div>
					</template>
				</NavHeader>
			</header>
			<main class="p-4">
				<slot />
				<!-- This is where the page content will be injected -->
			</main>
		</div>
	</div>
</template>
