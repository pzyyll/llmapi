<script setup lang="ts">
useHead({
  htmlAttrs: {
    class: "bg-base-200 h-screen",
  },
  bodyAttrs: {
    class: "bg-base-200 h-screen",
  },
});

</script>

<template>
	<div class="w-full py-2 h-screen md:px-4">
		<AppHeader />
		<slot />
	</div>
</template>
