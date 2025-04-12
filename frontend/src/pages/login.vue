<script setup lang="ts">
definePageMeta({
	layout: 'login'
});

const username = ref('');
const password = ref('');

async function handleLogin() {
	const { data } = await useAPI('/api/login', {
		method: 'POST',
		body: {
			username: username.value,
			password: password.value
		},
		headers: {
			'Content-Type': 'application/json'
		}
	});
}
</script>

<template>
	<div class="flex h-screen items-center justify-center">
		<div class="bg-base-100 rounded-2xl sm:w-90 w-full flex flex-col justify-center p-4 gap-4">
			<!-- <h1>Sign In</h1> -->
			<form @submit.prevent="handleLogin" class="flex flex-col gap-4">
				<input type="text" placeholder="Username" required v-model="username" class="d-input d-input-ghost"/>
				<input type="password" placeholder="Password" required v-model="password" class="d-input d-input-ghost"/>
				<button type="submit" class="d-btn">{{ $t('login') }}</button>
			</form>
		</div>
	</div>
</template>
