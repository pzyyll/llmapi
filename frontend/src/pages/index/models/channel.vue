<script setup lang="ts">
type ChannelType = 'Open AI' | 'Deepseek' | 'Gemini AI';

interface ChannelItem {
	name: string;
	id: number;
	type: ChannelType;
	status: 'active' | 'inactive';
}

const typeColors: Record<ChannelType, string> = {
	'Open AI': 'neutral',
	Deepseek: 'info',
	'Gemini AI': 'success'
};

const route = useRoute();
const channels = ref<ChannelItem[]>([
	{ name: 'Channel 1', id: 1, type: 'Open AI', status: 'active' },
	{ name: 'Channel 2', id: 2, type: 'Deepseek', status: 'inactive' },
	{ name: 'Channel 3', id: 3, type: 'Gemini AI', status: 'inactive' }
]);
const navChannel = ref<ChannelItem | null>(null);
async function openChannel(channel: ChannelItem) {
	// Logic to open the channel
	// navagato to channel details in ./channel/[id].vue
	navChannel.value = channel;
	await navigateTo(`/models/channel/${channel.id}`);
}

watch(
	() => route.params.id,
	(newId) => {
		if (newId) {
			const channel = channels.value.find((c) => c.id === Number(newId));
			if (channel) {
				navChannel.value = channel;
			} else {
				navChannel.value = null;
			}
		} else {
			navChannel.value = null;
		}
	},
	{ immediate: true }
);
</script>

<template>
	<NuxtPage :transition="{ name: 'fade', mode: 'out-in', appear: true }" v-if="navChannel" />
	<MainContent class="flex flex-col gap-4" v-else>
		<div class="border-base-content/5 rounded-box w-full overflow-x-auto border">
			<table class="d-table">
				<thead>
					<tr class="text-base-content/80">
						<th class="text-left">ID</th>
						<th class="text-left">Name</th>
						<th class="text-left">Type</th>
						<th class="text-left">Status</th>
						<th></th>
					</tr>
				</thead>
				<tbody v-if="channels.length > 0">
					<tr
						v-for="channel in channels"
						:key="channel.id"
						class="hover:bg-base-200 active:bg-base-300 group cursor-pointer"
						@click="openChannel(channel)"
					>
						<td>{{ channel.id }}</td>
						<td>{{ channel.name }}</td>
						<td>
							<div
								class="d-badge d-badge-xs text-base-content/80"
								:class="['d-badge-' + typeColors[channel.type]]"
							>
								{{ channel.type }}
							</div>
						</td>
						<td>
							<div
								class="d-badge d-badge-xs text-base-content/80"
								:class="['d-badge-' + (channel.status === 'active' ? 'success' : 'error')]"
							>
								{{ channel.status }}
							</div>
						</td>
						<td class="invisible flex justify-end group-hover:visible">
							<MaterialSymbolsArrowForwardIosRounded />
						</td>
					</tr>
				</tbody>
			</table>
		</div>
	</MainContent>
</template>
