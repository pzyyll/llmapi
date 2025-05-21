<script setup lang="ts">
const props = defineProps({
	title: { type: String, default: '' },
	onLoading: { type: Boolean, default: false }
});
</script>

<template>
	<div class="bg-base-100 flex flex-1 flex-col justify-start items-center rounded-md">
		<div class="flex flex-col w-full" v-show="!props.onLoading">
			<header class="box-content border-b-1 border-base-300 px-4 flex flex-col">
				<NavHeader class="pt-3">
					<template #left>
						<div>
							<slot name="header-left">
								<h1 class="text-xl font-bold whitespace-nowrap">{{ props.title }}</h1>
							</slot>
						</div>
					</template>
					<template #right>
						<div>
							<slot name="header-right" />
						</div>
					</template>
				</NavHeader>
				<slot name="header-footer">
					<div class="pt-3"></div>
				</slot>
			</header>
			<main class="p-4">
				<slot />
				<!-- This is where the page content will be injected -->
			</main>
		</div>
	</div>
</template>
