<script setup lang="ts">
import { gsap } from 'gsap';

const emit = defineEmits<{
	(e: 'onKeyCreated', key: Dto.Key): void;
}>();

const open = ref(false);
const createSuccess = ref(false);
const secretKey = ref('');
const showCheck = ref(false);
const secretKeyName = ref('');
const keyCreating = ref(false);

async function handleSubmit() {
	try {
		name = secretKeyName.value ? secretKeyName.value : 'Secret Key';
		console.log('name: ', name, secretKeyName.value);
		if (keyCreating.value) return;

		keyCreating.value = true;
		const { data } = await useAPI().post<Dto.CreateKeyResponse>(RequestPath.CreateKey, {
			name
		});

		console.log('new key: ', data);
		emit('onKeyCreated', data.key);
		// keys.value.unshift(data.key);
		keyCreating.value = false;
		createSuccess.value = true;
		secretKey.value = data.secret;
	} catch (error) {
		console.error('Create Key Error', error);
	}
}

async function handleCopy() {
	if (secretKey.value) {
		try {
			if (showCheck.value) return;
			await navigator.clipboard.writeText(secretKey.value);
			showCheck.value = true;
			setTimeout(() => {
				showCheck.value = false;
			}, 2000);
		} catch (err) {
			console.error('Failed to copy: ', err);
		}
	}
}

function handleInteractOutside(event) {
	if (createSuccess.value || keyCreating.value) {
		event.preventDefault();
	}
}

function overlayFadeIn(el, done) {
	gsap.from(el, {
		duration: 0.2,
		opacity: 0,
		onComplete: done
	});
}
function overlayFadeOut(el, done) {
	gsap.to(el, {
		duration: 0.2,
		opacity: 0,
		onComplete: done
	});
}
function contentFadeIn(el, done) {
	gsap.from(el, {
		duration: 0.2,
		// scale: 0.8,
		y: 20,
		opacity: 0,
		onComplete: done
	});
}
function contentFadeOut(el, done) {
	gsap.to(el, {
		duration: 0.2,
		// scale: 0.8,
		y: 20,
		opacity: 0,
		onComplete: done
	});
}

watch(open, (val) => {
	if (val) {
		// Reset it the next time it's opened, leaving it for the exit animation to display.
		createSuccess.value = false;
		secretKey.value = '';
		showCheck.value = false;
		secretKeyName.value = '';
		keyCreating.value = false;
	} else {
		secretKey.value = '';
	}
});
</script>

<template>
	<div>
		<DialogRoot v-model:open="open">
			<DialogTrigger asChild>
				<button class="d-btn d-btn-xs"><MaterialSymbolsAdd2Rounded />{{ $t('new_key') }}</button>
			</DialogTrigger>
			<DialogPortal>
				<Transition name="overlay-fade" @enter="overlayFadeIn" @leave="overlayFadeOut">
					<DialogOverlay class="fixed inset-0 bg-black/20" />
				</Transition>
				<Transition name="content-fade" @enter="contentFadeIn" @leave="contentFadeOut">
					<DialogContent
						class="bg-base-100 fixed left-[50%] top-[50%] z-[100] flex max-h-[85vh] w-[40vw] min-w-64 max-w-sm translate-x-[-50%] translate-y-[-50%] flex-col gap-2 rounded-xl px-5 py-4 shadow-xl"
						@interact-outside="handleInteractOutside"
					>
						<DialogTitle>
							<h2 class="text-base-content text-lg">
								{{ createSuccess ? $t('save_key') : $t('new_key') }}
							</h2>
						</DialogTitle>
						<form @submit.prevent="handleSubmit" v-show="!createSuccess">
							<div class="mt-2 flex flex-col gap-4">
								<div class="flex flex-col gap-2">
									<div class="flex justify-start gap-2">
										<label for="key-name" class="text-sm">{{ $t('name') }}</label>
										<label for="key-name" class="text-base-content/50 text-sm">
											{{ $t('option') }}
										</label>
									</div>
									<input
										id="key-name"
										type="text"
										class="d-input d-input-sm d-input-primary w-full"
										:placeholder="$t('key_name_placeholder')"
										v-model="secretKeyName"
										:disabled="keyCreating"
									/>
								</div>
								<div class="mt-2 flex justify-end gap-4">
									<DialogClose asChild>
										<button class="d-btn d-btn-sm min-w-20" @click="open = false" :disabled="keyCreating">
											{{ $t('cancel') }}
										</button>
									</DialogClose>
									<button class="d-btn d-btn-sm d-btn-primary min-w-20" type="submit" :disabled="keyCreating">
										<span v-if="keyCreating" class="d-loading d-loading-spinner"></span>
										<span v-else>{{ $t('submit') }}</span>
									</button>
								</div>
							</div>
						</form>
						<fieldset v-if="createSuccess" class="flex flex-col gap-4">
							<DialogDescription>
								<i18n-t
									keypath="create_key_description"
									tag="p"
									class="text-base-content/80 my-1 text-sm"
								>
									<template #create_key_description_bold>
										<span class="text-base-content font-bold">{{
											$t('create_key_description_bold')
										}}</span>
									</template>
								</i18n-t>
							</DialogDescription>
							<label class="d-input d-input-primary w-full">
								<input id="key-name" type="text" class="grow" :value="secretKey" readonly />
								<button class="d-btn d-btn-primary d-btn-xs" @click="handleCopy">
									<Transition name="check-fade" mode="out-in">
										<MaterialSymbolsCheckRounded v-if="showCheck" />
										<MaterialSymbolsContentCopyOutlineRounded v-else />
									</Transition>
									{{ $t('copy') }}
								</button>
							</label>
							<div class="mt-2 flex justify-end gap-4">
								<DialogClose asChild>
									<button class="d-btn d-btn-sm min-w-20" @click="open = false">
										{{ $t('done') }}
									</button>
								</DialogClose>
							</div>
						</fieldset>
					</DialogContent>
				</Transition>
			</DialogPortal>
		</DialogRoot>
	</div>
</template>

<style scoped>
.check-fade-enter-active,
.check-fade-leave-active {
	transition: all 0.1s ease-in-out;
}

.check-fade-enter-from,
.check-fade-leave-to {
	opacity: 0;
	transform: scale(0.6);
}
</style>
