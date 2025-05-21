<script setup lang="ts">
import { gsap } from 'gsap';

const emit = defineEmits<{
	(e: 'interactOutside', event: Event): void;
	(e: 'onConfirm', event: MouseEvent): void;
}>();

const props = defineProps<{
	title?: string;
	desc?: string;
}>();

const open = defineModel('open', { default: false, required: false });

function handleInteractOutside(event: Event) {
	emit('interactOutside', event);
}

function onConfirm(event: MouseEvent) {
	emit('onConfirm', event);
}

function overlayFadeIn(el: any, done: () => void) {
	gsap.from(el, {
		duration: 0.2,
		opacity: 0,
		onComplete: done
	});
}
function overlayFadeOut(el: any, done: () => void) {
	gsap.to(el, {
		duration: 0.2,
		opacity: 0,
		onComplete: done
	});
}
function contentFadeIn(el: any, done: () => void) {
	gsap.from(el, {
		duration: 0.2,
		// scale: 0.8,
		y: 20,
		opacity: 0,
		onComplete: done
	});
}
function contentFadeOut(el: any, done: () => void) {
	gsap.to(el, {
		duration: 0.2,
		// scale: 0.8,
		y: 20,
		opacity: 0,
		onComplete: done
	});
}

watch(open, (val) => {});
</script>

<template>
	<div>
		<DialogRoot v-model:open="open">
			<DialogTrigger asChild>
				<slot name="trigger">
					<button class="d-btn d-btn-xs"></button>
				</slot>
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
						<DialogTitle v-if="props.title">
							<h2 class="text-base-content text-lg">
								{{ props.title }}
							</h2>
						</DialogTitle>
						<slot name="title"> </slot>
						<DialogDescription v-if="props.desc">
							<p class="text-base-content/80 text-sm">
								{{ props.desc }}
							</p>
						</DialogDescription>
						<slot name="content"> </slot>
						<slot name="bottom">
							<div class="mt-2 flex justify-end gap-4">
								<DialogClose asChild>
									<button class="d-btn d-btn-sm min-w-20">
										{{ $t('cancel') }}
									</button>
								</DialogClose>
								<DialogClose asChild>
									<button class="d-btn d-btn-sm d-btn-primary min-w-20" @click="onConfirm">
										<span>{{ $t('confirm') }}</span>
									</button>
								</DialogClose>
							</div>
						</slot>
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
