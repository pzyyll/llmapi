<script setup lang="ts">
const props = defineProps({
	onConfirm: {
		type: Function
	},
	onCancel: {
		type: Function
	},
	title: {
		type: String,
		default: 'Confirm'
	},
	message: {
		type: String,
		default: 'Are you sure you want to proceed?'
	}
});

const confirm = () => {
	props.onConfirm?.();
};

const cancel = () => {
	props.onCancel?.();
};
</script>

<template>
	<DialogRoot>
		<DialogTrigger asChild>
			<slot>
				<button class="d-btn d-btn-primary"></button>
			</slot>
		</DialogTrigger>
		<DialogPortal>
			<Transition name="overlay-fade">
				<DialogOverlay class="fixed inset-0 bg-black/30" />
			</Transition>
			<Transition name="content-fade">
				<DialogContent
					class="bg-base-100 fixed left-[50%] top-[50%] z-[100] flex max-h-[85vh] w-[40vw] min-w-64 max-w-sm translate-x-[-50%] translate-y-[-50%] flex-col gap-2 rounded-xl px-5 py-4 shadow-xl"
				>
					<DialogTitle class="text-lg text-base-content">
						{{ props.title }}
					</DialogTitle>
					<div>
						<DialogDescription asChild>
							<slot name="message">
								<p class="text-sm text-base-content/80">
									{{ props.message }}
								</p>
							</slot>
						</DialogDescription>
					</div>
					<div class="mt-2 flex justify-end gap-4">
						<DialogClose asChild>
							<button class="d-btn d-btn-sm min-w-20" @click="cancel">{{ $t('cancel') }}</button>
						</DialogClose>
						<DialogClose asChild>
							<button class="d-btn d-btn-sm d-btn-error min-w-20" @click="confirm">
								{{ $t('confirm') }}
							</button>
						</DialogClose>
					</div>
				</DialogContent>
			</Transition>
		</DialogPortal>
	</DialogRoot>
</template>

<style scoped>
.content-fade-enter-active,
.content-fade-leave-active {
	transition:
		transform 0.1s ease-in-out,
		opacity 0.1s ease-in-out;
}

.content-fade-enter-from,
.content-fade-leave-to {
	transform: scale(0.9);
	opacity: 0;
}

.overlay-fade-enter-active,
.overlay-fade-leave-active {
	transition: opacity 0.1s ease-in-out;
}
.overlay-fade-enter-from,
.overlay-fade-leave-to {
	opacity: 0;
}
</style>
