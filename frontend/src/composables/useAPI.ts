import type { UseFetchOptions, AsyncData } from 'nuxt/app';
// import type { DefaultAsyncDataErrorValue, DefaultAsyncDataValue } from 'nuxt/app/defaults';
// import type { PickFrom, KeysOf } from '#app/composables/asyncData';

export function useAPI<T>(
	url: string | (() => string),
	options?: UseFetchOptions<T>
) {
	// return AsyncData<PickFrom<DataT, PickKeys> | DefaultT, ErrorT | DefaultAsyncDataErrorValue>
	return useFetch(url, {
		...options,
		$fetch: useNuxtApp().$api as typeof $fetch
	});
}
