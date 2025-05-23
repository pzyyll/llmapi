<script setup lang="ts">
const router = useRoute();

const channelId = router.params.id;
</script>

<template>
	<MainContent class="relative flex h-full flex-col gap-4">
		<template #header>
			<div name="back" class="bg-base-100/60 sticky top-0 z-10 flex px-5 backdrop-blur-sm">
				<div
					@click="$router.back()"
					class="text-base-content/50 hover:text-base-content/80 relative my-4 flex flex-none cursor-pointer items-center"
				>
					<MaterialSymbolsArrowBackIosNewRounded class="absolute -left-3 size-3" />
					<span class="text-sm">{{ $t('back') }}</span>
				</div>
			</div>
		</template>

		<div class="flex h-[1100px] flex-col gap-4">
			<h2 class="text-md font-bold"><span>Channel: </span>{{ channelId }}</h2>
			<p class="text-sm">
				This is a channel page. You can add your content here.
				<br />
			</p>
		</div>
	</MainContent>
</template>
