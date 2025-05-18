<script setup lang="ts">
import type { NuxtError } from '#app';
import { HttpStatusCode } from 'axios';
const props = defineProps({
	error: Object as () => NuxtError
});

const refreshPage = () => {
	clearError().then(() => {
		console.log('Refreshing page...');
		window.location.reload();
	});
};

const goHome = () => {
	clearError().then(() => {
		navigateTo('/');
	});
};
</script>

<template>
	<div class="bg-base-200 flex h-screen w-screen items-center justify-center">
		<div
			class="bg-base-100 w-full max-w-lg transform rounded-xl p-8 text-center shadow-2xl transition-all duration-500 hover:scale-105 md:p-12 lg:p-16"
		>
			<div v-if="error?.statusCode === HttpStatusCode.NotFound">
				<h1 class="text-info mb-4 text-6xl font-extrabold md:text-7xl">{{ error?.statusCode }}</h1>
				<h2 class="text-base-content mb-6 text-2xl font-semibold md:text-3xl">
					{{ $t('page_not_found') }}
				</h2>
				<p class="text-base-content mb-8 text-lg">
					{{ $t('page_not_found_desc') }}
				</p>
				<button class="d-btn d-btn-info d-btn-lg" @click="goHome">
					{{ $t('go_home') }}
				</button>
			</div>
			<div v-else-if="error?.statusCode === HttpStatusCode.BadGateway">
				<h1 class="text-info mb-4 text-6xl font-extrabold md:text-7xl">{{ error.statusCode }}</h1>
				<h2 class="text-base-content mb-6 text-2xl font-semibold md:text-3xl">
					{{ $t('bad_gateway') }}
				</h2>
				<p class="text-base-content mb-8 text-lg">
					{{ $t('bad_gateway_desc') }}
				</p>
				<button class="d-btn d-btn-info d-btn-lg" @click="refreshPage">
					{{ $t('refresh_page') }}
				</button>
			</div>
		</div>
	</div>
</template>
