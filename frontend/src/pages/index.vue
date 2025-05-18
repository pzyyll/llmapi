<script setup lang="ts">
import path from 'pathe';
const { isAdmin } = useAuthStore();

const navItems = [
	{
		title: 'logs',
		path: '/logs',
		icon: MaterialSymbolsChartDataOutlineRounded,
		visible: true
	},
	{
		title: 'models',
		path: '/models',
		icon: MaterialSymbolsRobot2OutlineRounded,
		visible: true
	},
	{
		title: 'keys',
		path: '/api_key',
		icon: MaterialSymbolsKeyVerticalOutlineRounded,
		visible: true
	},
	{
		title: 'user',
		path: '/user',
		icon: MaterialSymbolsPersonOutlineRounded,
		visible: isAdmin
	},
	{
		title: 'settings',
		path: '/settings',
		icon: MaterialSymbolsSettingsOutlineRounded,
		visible: true
	}
];

definePageMeta({
	layout: false,
	redirect: (to) => {
		console.log("Index page redirecting from: ", to.path, " to: ", path.join(to.path, navItems[0].path));
		return path.join(to.path, navItems[0].path);
	}
});

import MaterialSymbolsPersonOutlineRounded from '~icons/material-symbols/person-outline-rounded';
import MaterialSymbolsKeyVerticalOutlineRounded from '~icons/material-symbols/key-vertical-outline-rounded';
import MaterialSymbolsSettingsOutlineRounded from '~icons/material-symbols/settings-outline-rounded';
import MaterialSymbolsRobot2OutlineRounded from '~icons/material-symbols/robot-2-outline-rounded';
import MaterialSymbolsChartDataOutlineRounded from '~icons/material-symbols/chart-data-outline-rounded';
import AvatarInfo from '~/components/AvatarInfo.vue';

const isSidebarCollapsed = ref(false);
const toggleSidebar = () => {
	isSidebarCollapsed.value = !isSidebarCollapsed.value;
};




const router = useRouter();
const isActive = (path: string) => {
	return router.currentRoute.value.path === path;
};

onMounted(() => {
	const currentPath = router.currentRoute.value.path;
	if (currentPath === '/') {
		router.push(navItems[0].path);
	}
});
</script>

<template>
	<NuxtLayout name="dashboard">
		<template #header>
			<NavHeader class="px-3 py-2 h-13">
				<template #left>
					<h1 class="text-2xl font-bold">LLMTech</h1>
				</template>
				<template #right>
					<AvatarInfo />
				</template>
			</NavHeader>
		</template>
		<template #sidebar>
			<SideListItem
				:showTitle="!isSidebarCollapsed"
				:title="$t('dashboard')"
				@click="toggleSidebar"
				class="mx-2"
			>
				<template #icon>
					<MaterialSymbolsListsRounded class="size-6" />
				</template>
			</SideListItem>

			<nav class="py-3">
				<ul class="flex flex-col gap-1">
					<li v-for="item in navItems.filter((i) => i.visible)" :key="item.path">
						<NuxtLink :to="`${item.path}`">
							<SideListItem
								:showTitle="!isSidebarCollapsed"
								:title="$t(item.title)"
								:highlighted="isActive(item.path)"
								:class="{
									'opacity-50': !isActive(item.path),
									'opacity-100': isActive(item.path)
								}"
								class="mx-2"
							>
								<template #icon>
									<component :is="item.icon" class="size-6" />
								</template>
							</SideListItem>
						</NuxtLink>
					</li>
					<!-- <div class="d-divider gap-0 m-0" /> -->
				</ul>
			</nav>
		</template>

		<div class="flex h-full flex-col pb-3 pr-3">
			<NuxtPage />
		</div>
	</NuxtLayout>
</template>
