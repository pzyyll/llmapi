import { defineStore } from 'pinia';

export const useAuthStore = defineStore(
	'auth',
	() => {
		const accessToken = ref<string | null>(null);
		const user = ref<Dto.User | null>(null);

		const isLoggedIn = computed(() => !!accessToken.value);
		const getToken = computed(() => accessToken.value);
		const getUser = computed(() => user.value);
		const isAdmin = computed(() => {
			if (!user.value) return false;
			return user.value.role === Role.Admin || user.value.role === Role.Super;
		});

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

		const isSelf = (userId: string) => {
			return user.value?.user_id === userId;
		};

		return {
			// State
			accessToken,
			user,

			// Getters
			isLoggedIn,
			getToken,
			getUser,
			isAdmin,

			setUser,
			setToken,

			// Actions
			clearUser,
			clearToken,
			clearAuth,

			// Methods
			isSelf
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
