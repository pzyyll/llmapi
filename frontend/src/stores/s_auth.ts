import { defineStore } from 'pinia';

export const useAuthStore = defineStore(
	'auth',
	() => {
		const accessToken = ref<string | null>(null);
		const user = ref<Dto.User | null>(null);

		const isLoggedIn = computed(() => !!accessToken.value);
		const getToken = computed(() => accessToken.value);
		const getUser = computed(() => user.value);

		const setUser = (newUser: Dto.User) => {
			user.value = newUser;
		};
		const setToken = (newToken: string) => {
			accessToken.value = newToken;
		};

		const clearUser = () => {
			user.value = null;
		};
		const clearToken = () => {
			accessToken.value = null;
		};
		const clearAuth = () => {
			clearUser();
			clearToken();
		};

		return {
			// State
			accessToken,
			user,

			// Getters
			isLoggedIn,
			getToken,
			getUser,

			setUser,
			setToken,

			// Actions
			clearUser,
			clearToken,
			clearAuth
		};
	},
	{
		persist: [
			{
				pick: ['user']
			},
			{
				pick: ['accessToken'],
				storage: sessionStorage
			}
		]
	}
);
