<script setup lang="ts">
import joinPath  from 'pathe';

const items = [
	{ name: 'All Models', path: 'default' },
	{ name: 'To Do', path: 'channel' }
];

definePageMeta({
	redirect: (to) => {
		console.log('Redirecting from: ', to.path, ' to: ', joinPath.join(to.path, items[0].path));
		return joinPath.join(to.path, items[0].path);
	}
});
</script>

<template>
	<DashboardMainLayout :title="$t('models')">
		<nav>
			<div class="d-tabs d-tabs-border">
				<NuxtLink
					v-for="(item, index) in items"
					:key="index"
					class="d-tab"
					:to="item.path"
					type="radio"
					:checked="index === 0 ? 'checked' : undefined"
				>
					{{ item.name }}
				</NuxtLink>
			</div>
		</nav>
		<NuxtPage />
	</DashboardMainLayout>
</template>
