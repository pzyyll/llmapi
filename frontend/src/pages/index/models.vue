<script setup lang="ts">
import path from 'pathe';

const items = [
	{ name: 'Endpoint', path: '/models/endpoint' },
	{ name: 'Channel', path: '/models/channel' }
];

definePageMeta({
	middleware: [
		(to, from) => {
			if (path.basename(to.path) === 'models') {
				return navigateTo(items[0].path);
			}
		}
	]
});

const defaultValue = ref(items[0].path);

const router = useRouter();
onMounted(() => {
	const currentPath = router.currentRoute.value.path;
	defaultValue.value = currentPath.endsWith('/') ? currentPath.slice(0, -1) : currentPath;
	console.log('Current path: ', currentPath, ' defaultValue: ', defaultValue.value);
});
</script>

<template>
	<MainLayout :title="$t('models')">
		<template #header-footer>
			<TabsRoot as-child class="mt-2" v-model="defaultValue">
				<nav class="flex items-center justify-start">
					<TabsList aria-label="Model Tabs" class="d-tabs d-tabs-xs relative flex shrink-0 gap-2">
						<TabsIndicator
							class="translate-x-(--reka-tabs-indicator-position) w-(--reka-tabs-indicator-size) translate-y-[1px] transition-x absolute bottom-0 left-0 h-[1px] rounded-full duration-100"
						>
							<div class="bg-accent h-full w-full"></div>
						</TabsIndicator>
						<TabsTrigger
							v-for="(item, index) in items"
							:key="index"
							:value="item.path"
							class="shrink-0 leading-none"
							as-child
						>
							<NuxtLink :to="item.path" v-slot="{ isActive }">
								<div class="d-tab p-0" :class="{ 'd-tab-active': isActive, 'font-bold': isActive }">
									<span>{{ item.name }}</span>
								</div>
							</NuxtLink>
						</TabsTrigger>
					</TabsList>
				</nav>
			</TabsRoot>
		</template>
		<NuxtPage />
	</MainLayout>
</template>
