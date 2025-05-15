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
		<!-- 头部导航区域 -->
		<header class="flex w-full justify-center">
			<slot name="header" />
		</header>

		<div class="flex flex-1 overflow-hidden">
			<!-- 左侧导航菜单区域 -->
			<aside class="flex flex-col overflow-y-auto">
				<slot name="sidebar" />
			</aside>

			<!-- 右侧内容展示区域 -->
			<main class="flex-1 overflow-y-auto">
				<slot />
			</main>
		</div>
	</div>
</template>

<style scoped>
.sidebar-title-enter-active,
.sidebar-title-leave-active {
	transition: all 0.25s ease;
}

.sidebar-title-enter-from,
.sidebar-title-leave-to {
	width: 0;
	opacity: 0;
}
</style>
