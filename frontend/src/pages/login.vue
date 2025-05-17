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

const turnstileRef = ref<{ reset: () => void } | null>(null);
const isLoading = ref(false);

const isTurnstileEnabled = computed(() => turnstile.siteKey.length > 0);

const isFormInvalid = computed(() => {
	if (!username.value || !password.value) return true;
	if (isTurnstileEnabled.value && !cft_token.value) return true;
	return false;
});

function resetTurnstile() {
	cft_token.value = '';
	turnstileRef.value?.reset();
}

async function handleLogin() {
	try {
		const headers: Record<string, string> = {};
		if (isTurnstileEnabled.value) {
			headers[Header.TurnstileToken] = cft_token.value;
		}

		isLoading.value = true;
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

		if (data.user && data.access_token) {
			console.log('Login successful:', data);
			const authStore = useAuthStore();
			authStore.setUser(data.user);
			authStore.setToken(data.access_token);
			navigateTo('/');
		} else {
			console.error('Login failed: Invalid response', data);
			throw new Error('Invalid response');
		}
	} catch (error) {
		console.error('Login failed:', error);
	} finally {
		isLoading.value = false;
		resetTurnstile();
	}
}

async function turnstileExpired() {
	resetTurnstile();
}
</script>

<template>
	<div class="flex h-screen flex-col items-center justify-center">
		<div
			class="bg-base-100 sm:w-90 relative flex w-full flex-col justify-center overflow-hidden sm:rounded-2xl"
		>
			<progress
				class="d-progress absolute left-0 top-0 w-full rounded-none"
				v-if="isLoading"
			></progress>
			<form @submit.prevent="handleLogin" class="flex flex-col items-center gap-4 p-4">
				<input
					type="text"
					placeholder="Username"
					required
					v-model="username"
					class="d-input d-input-ghost w-full"
					:disabled="isLoading"
				/>
				<input
					type="password"
					placeholder="Password"
					required
					v-model="password"
					class="d-input d-input-ghost w-full"
					:disabled="isLoading"
				/>
				<NuxtTurnstile
					v-model="cft_token"
					v-if="isTurnstileEnabled"
					:options="{
						expiredCallback: turnstileExpired
					}"
					ref="turnstileRef"
				/>
				<button type="submit" class="d-btn w-full" :disabled="isFormInvalid || isLoading">
					{{ $t('login') }}
				</button>
			</form>
		</div>
	</div>
</template>
