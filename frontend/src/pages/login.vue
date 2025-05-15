<script setup lang="ts">
definePageMeta({
	layout: 'login',
	middleware: [
		function (to, from) {
			const authStore = useAuthStore();
			if (authStore.user) {
				return navigateTo('/');
			}
		}
	]
});

const {
	public: { turnstile }
} = useRuntimeConfig();

const username = ref('');
const password = ref('');
const cft_token = ref('');

const isTurnstileEnabled = computed(() => turnstile.siteKey.length > 0);

const isFormInvalid = computed(() => {
	if (!username.value || !password.value) return true;
	if (isTurnstileEnabled.value && !cft_token.value) return true;
	return false;
});

async function handleLogin() {
	try {
		const headers: Record<string, string> = {};
		if (isTurnstileEnabled.value) {
			headers[Header.TurnstileToken] = cft_token.value;
		}

		const { data } = await useAPI().post<Dto.LoginResponse>(
			RequestPath.Login,
			{
				username: username.value,
				password: password.value
			},
			{
				headers
			}
		);

		const { setUser, setToken } = useAuthStore();

		// console.log('Login successful:', data);
		setUser(data.user);
		setToken(data.access_token);
		navigateTo('/');
	} catch (error) {
		console.error('Login failed:', error);
	}
}

async function turnstileExpired() {
	cft_token.value = '';
}
</script>

<template>
	<div class="flex h-screen items-center justify-center">
		<div class="bg-base-100 sm:w-90 flex w-full flex-col justify-center p-4 sm:rounded-2xl">
			<form @submit.prevent="handleLogin" class="flex flex-col items-center gap-4">
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
				<NuxtTurnstile
					v-model="cft_token"
					v-if="isTurnstileEnabled"
					:options="{
						expiredCallback: turnstileExpired
					}"
				/>
				<button type="submit" class="d-btn w-full" :disabled="isFormInvalid">
					{{ $t('login') }}
				</button>
			</form>
		</div>
	</div>
</template>
