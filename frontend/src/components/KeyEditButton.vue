<script setup lang="ts">
import path from 'pathe';

const props = defineProps<{
	keyInfo: Dto.Key;
}>();

const emit = defineEmits<{
	(e: 'onKeyChanged', key: Dto.Key): void;
}>();
const keyName = ref(props.keyInfo.name);

const onConfirm = async (event: MouseEvent) => {
	console.log('[KeyEditDialog] onConfirm');
	try {
		if (keyName.value === '') {
			event.preventDefault();
			console.error('Key name is required');
			return;
		}
		const { data } = await useAPI().put(path.join(RequestPath.UpdateKey, props.keyInfo.lookup_key), {
			name: keyName.value
		});
		emit('onKeyChanged', { ...props.keyInfo, name: keyName.value } as Dto.Key);
		console.log('update key: ', data);
	} catch (error) {
		console.error('Update Key Error', error);
	}
};

onMounted(() => {
	console.log('[KeyEditDialog] mounted', props.keyInfo);
});
</script>

<template>
	<EditDialog :title="$t('edit_key')" :desc="props.keyInfo.secret_brief" @on-confirm="onConfirm">
		<template #trigger>
			<button class="d-btn d-btn-ghost d-btn-xs">
				<MaterialSymbolsEditSquareOutlineRounded class="size-4" />
			</button>
		</template>
		<template #content>
			<fieldset>
				<div class="flex w-full flex-col gap-2">
					<div class="flex flex-col gap-1">
						<label for="key-name" class="text-sm">{{ $t('name') }}</label>
						<input
							label="$t('key_name')"
							v-model="keyName"
							class="d-input d-input-primary d-input-sm w-full"
							:placeholder="props.keyInfo.name"
						/>
					</div>
				</div>
			</fieldset>
		</template>
	</EditDialog>
</template>
