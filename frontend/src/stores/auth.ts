import { defineStore } from 'pinia';

interface User {
	userid: string;
	username: string;
	role: string;
}

export const useAuthStore = defineStore(
	'auth',
	() => {
		const accessToken = ref<string | null>(null);
		const user = ref<User | null>(null);

		const isLoggedIn = computed(() => !!accessToken.value);
		const getToken = computed(() => accessToken.value);
		const getUser = computed(() => user.value);

		return {
			// State
			accessToken,
			user,

			// Getters
			isLoggedIn,
			getToken,
			getUser
		};
	},
	{
		persist: {
			pick: ['accessToken', 'user', 'initialized']
		}
	}
);
