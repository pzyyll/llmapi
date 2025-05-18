import { AxiosError, HttpStatusCode } from 'axios';

export default defineNuxtRouteMiddleware(async (to, from) => {
	// Check if the user is authenticated
	const authStore = useAuthStore();
	const isLoggedIn = authStore.isLoggedIn;

	if (isLoggedIn) {
		return;
	}

	if (authStore.getUser !== null) {
		try {
			const { data } = await useAPI().post<Dto.RenewTokenResponse>(RequestPath.RenewToken);
			authStore.setToken(data.access_token);
			authStore.setUser(data.user);
			return;
		} catch (error: any) {
			// TODO: if bad gateway, show error page
			console.error('Error while checking user authentication:', error);
			return;
		}
	}

	return navigateTo('/login');
});
