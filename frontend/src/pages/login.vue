<script setup lang="ts">
import type { LoginResponse } from './pages.d';
definePageMeta({
	layout: 'login'
});

const username = ref('');
const password = ref('');

async function handleLogin() {
	const { data, error, status } = await useAPI<LoginResponse>('/api/login', {
		method: 'POST',
		body: {
			username: username.value,
			password: password.value
		},
		headers: {
			'Content-Type': 'application/json'
		}
	});

	if (error.value) {
		console.error('Login error:', error.value);
		return;
	}

	// console.log('data', data);
	// console.log('error', error);
	// console.log('status', status);
	// console.log('token', data.value?.access_token);
	// console.log('user_id', data.value?.user_id);
	// console.log('username', data.value?.username);
	// console.log('role', data.value?.role);
	// console.log('email', data.value?.email);
}
</script>

<template>
	<div class="flex h-screen items-center justify-center">
		<div class="bg-base-100 sm:w-90 flex w-full flex-col justify-center p-4 sm:rounded-2xl">
			<!-- <h1>Sign In</h1> -->
			<form @submit.prevent="handleLogin" class="flex flex-col gap-4">
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
				<button type="submit" class="d-btn">{{ $t('login') }}</button>
			</form>
		</div>
	</div>
</template>
