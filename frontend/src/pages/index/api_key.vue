<script setup lang="ts">
import { gsap } from 'gsap';

const api = useAPI();
const keys = ref<Dto.Key[]>([]);
const onDataFetching = ref(false);

const onKeyCreated = (key: Dto.Key) => {
	keys.value.unshift(key);
};

const deleteKey = async (lookupKey: string) => {
	try {
		await api.delete(RequestPath.DeleteKey, {
			params: {
				lookup_key: lookupKey
			}
		});
		keys.value = keys.value.filter((key) => key.lookup_key !== lookupKey);
	} catch (error) {
		console.error('Delete Key Error', error);
	}
};

const onEnter = (el: Element, done: () => void) => {
	const duration = 0.1;
	const cells = Array.from(el.children);
	gsap.from(el, {
		opacity: 0,
		duration: duration
	});
	gsap.from(cells, {
		opacity: 0,
		paddingTop: 0,
		paddingBottom: 0,
		fontSize: 0,
		duration: duration,
		onComplete: done
	});
};

const onLeave = (el: Element, done: () => void) => {
	const duration = 0.1;
	const cells = Array.from(el.children);
	gsap.to(el, {
		opacity: 0,
		duration: duration
	});
	gsap.to(cells, {
		opacity: 0,
		paddingTop: 0,
		paddingBottom: 0,
		fontSize: 0,
		duration: duration,
		onComplete: done
	});
};

onMounted(async () => {
	// get keys from API
	try {
		onDataFetching.value = true;
		const { data } = await api.get<Dto.Keys>(RequestPath.GetKeys);
		console.log('data: ', data);
		keys.value = data.api_keys;
		onDataFetching.value = false;
	} catch (error) {
		console.error('Get Keys Error', error);
	}
});
</script>

<template>
	<DashboardMainLayout :title="$t('keys')" :on-loading="onDataFetching">
		<template #header-right>
			<NewKeyButton @on-key-created="onKeyCreated" />
		</template>
		<div class="flex flex-col gap-4">
			<article class="prose prose-sm text-base-content">
				<p>
					{{ keys.length > 0 ? $t('key_description') : $t('create_key_desc_on_empty') }}
				</p>
			</article>
			<div class="bg-base-100 overflow-x-auto" v-if="keys.length > 0">
				<table class="d-table w-full">
					<colgroup>
						<col class="w-2/6" />
						<col class="w-1/6" />
						<col class="w-1/6" />
						<col class="w-1/6" />
						<col class="w-1/8" />
					</colgroup>
					<thead>
						<tr class="text-base-content/80 text-xs">
							<th
								v-for="item in ['name', 'keys', 'created_at', 'last_used']"
								:key="item"
								class="first:pl-0"
							>
								{{ $t(item).toLocaleUpperCase() }}
							</th>
						</tr>
					</thead>
					<TransitionGroup name="list" tag="tbody" :css="false" @enter="onEnter" @leave="onLeave">
						<tr v-for="key in keys" :key="key.name" class="border-base-content/10 border-b">
							<td class="first:pl-0 last:pr-0">{{ key.name }}</td>
							<td class="first:pl-0 last:pr-0">{{ key.secret_brief }}</td>
							<td class="first:pl-0 last:pr-0">{{ formatDateTimeString(key.created_at) }}</td>
							<td class="first:pl-0 last:pr-0">
								{{ key.last_used_at ? formatDateTimeString(key.last_used_at) : $t('never') }}
							</td>
							<td class="flex items-center justify-end first:pl-0 last:pr-0">
								<button class="d-btn d-btn-ghost d-btn-xs">
									<MaterialSymbolsEditSquareOutlineRounded class="size-4" />
								</button>

								<ConfirmDialog @confirm="deleteKey(key.lookup_key)" :title="$t('delete_key')">
									<button
										class="d-btn d-btn-ghost d-btn-xs hover:bg-error/20 active:bg-error/20 focus:outline-error hover:border-transparent focus:border-transparent focus:bg-transparent active:border-transparent"
									>
										<MaterialSymbolsDeleteOutlineRounded class="text-error size-4" />
									</button>
									<template #message>
										<div class="flex flex-col gap-2">
											<p class="text-base-content/80 my-1 text-sm">
												{{ $t('delete_key_message') }}
											</p>
											<label class="d-input w-full">
												{{ key.secret_brief }}
											</label>
										</div>
									</template>
								</ConfirmDialog>
							</td>
						</tr>
					</TransitionGroup>
				</table>
			</div>
		</div>
	</DashboardMainLayout>
</template>
