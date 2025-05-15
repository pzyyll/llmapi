<script setup lang="ts">
const props = defineProps({
	submit_text: {
		type: String,
		default: 'login'
	}
});

const username = ref('');
const password = ref('');

const emit = defineEmits(['submit']);
</script>

<template>
	<div class="flex h-screen items-center justify-center">
		<div class="bg-base-100 sm:w-90 flex w-full flex-col justify-center p-4 sm:rounded-2xl">
			<!-- <h1>Sign In</h1> -->
			<form @submit.prevent="emit('submit', username, password)" class="flex flex-col gap-4">
				<input
					type="text"
					placeholder="Username"
					required
					v-model="username"
					class="d-input d-input-ghost w-full"
				/>
				<input
					type="password"
					placeholder="Password"
					required
					v-model="password"
					class="d-input d-input-ghost w-full"
				/>
				<button type="submit" class="d-btn">{{ props.submit_text }}</button>
			</form>
		</div>
	</div>
</template>
