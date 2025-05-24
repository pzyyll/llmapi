<script setup lang="ts">
interface ModelItem {
	id: string;
	name: string;
	desc: string;
}

const testItems = ref<ModelItem[]>([
	{
		id: '1',
		name: 'gpt-3.5-turbo',
		desc: 'gpt-3.5-turbo is a family of state-of-the-art open-access large language models developed by Meta AI.'
	},
	{
		id: '2',
		name: 'gpt-4',
		desc: 'gpt-4 is a family of state-of-the-art open-access large language models developed by Meta AI.'
	},
	{
		id: '3',
		name: 'deepseek-r1',
		desc: 'deepseek-r1 is a family of state-of-the-art open-access large language models developed by Meta AI.'
	},
	{
		id: '4',
		name: 'deepseek-v3',
		desc: 'deepseek-v3 is a family of state-of-the-art open-access large language models developed by Meta AI.'
	},
	{
		id: '5',
		name: 'gemini-1.5-turbo-32k',
		desc: 'gemini-1.5-turbo-32k is a family of state-of-the-art open-access large language models developed by Meta AI.'
	},
	{
		id: '6',
		name: 'gemini-1.5-turbo-64k',
		desc: 'gemini-1.5-turbo-64k is a family of state-of-the-art open-access large language models developed by Meta AI.'
	},
	{
		id: '7',
		name: 'claude-sonnet-4',
		desc: 'claude-sonnet-4 is a family of state-of-the-art open-access large language models developed by Meta AI.'
	}
]);
</script>

<template>
	<MainContent>
		<div class="grid grid-cols-[repeat(auto-fit,minmax(180px,1fr))] gap-4">
			<template v-for="item in testItems" :key="item.id">
				<ModelCard :model="item"/>
			</template>
		</div>
	</MainContent>
</template>
