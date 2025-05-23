<script setup lang="ts">
import path from 'pathe';

const items = [
	{ name: 'endpoint', path: '/models/endpoint' },
	{ name: 'channel', path: '/models/channel' }
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

	const pathMatch = currentPath.match(/^\/models\/[^/]+/);
	if (pathMatch) {
		defaultValue.value = pathMatch[0];
	} else {
		defaultValue.value = items[0].path;
	}
});
</script>

<template>
	<MainLayout :title="$t('models')">
		<MainHeader :title="$t('models')">
			<template #footer>
				<TabsRoot as-child class="mt-2" v-model="defaultValue" activation-mode="manual">
					<div class="flex items-center justify-start">
						<TabsList
							aria-label="Model Tabs"
							class="d-tabs d-tabs-sm relative flex shrink-0 flex-nowrap gap-3"
						>
							<TabsIndicator
								class="translate-x-(--reka-tabs-indicator-position) w-(--reka-tabs-indicator-size) transition-x absolute bottom-0 left-0 h-[1px] translate-y-[1px] rounded-full duration-100"
							>
								<div class="bg-accent h-full w-full"></div>
							</TabsIndicator>
							<TabsTrigger
								v-for="(item, index) in items"
								:key="index"
								:value="item.path"
								class="shrink-0 leading-none"
							>
								<NuxtLink :to="item.path" v-slot="{ isActive, navigate }" custom>
									<button
										@mousedown="navigate"
										class="d-tab p-0"
										:class="{ 'font-bold': isActive, 'd-tab-active': isActive }"
									>
										<span>{{ $t(item.name) }}</span>
									</button>
									<!-- <div class="p-0" :class="{'font-bold': isActive }">
									<span>{{ item.name }}</span>
								</div> -->
								</NuxtLink>
							</TabsTrigger>
						</TabsList>
					</div>
				</TabsRoot>
			</template>
		</MainHeader>
		<NuxtPage />
	</MainLayout>
</template>
