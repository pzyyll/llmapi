export default defineNuxtRouteMiddleware(async (to, from) => {
	if (['/login', '/register'].includes(to.path)) {
		console.log(
			'[mid-auth] No authentication required for login or register. Allowing navigation.'
		);
		return;
	}

	const authStore = useAuthStore();
	const isLoggedIn = authStore.isLoggedIn;
	console.log('[mid-auth] isLoggedIn:', isLoggedIn);

	if (isLoggedIn) {
		return;
	}

	if (authStore.getUser !== null) {
		console.log('[mid-auth] User object exists, attempting token renew.');
		try {
			const { data } = await useAPI().post<Dto.RenewTokenResponse>(RequestPath.RenewToken);
			authStore.setToken(data.access_token);
			authStore.setUser(data.user);
			console.log('[mid-auth] Token renewed successfully. Allowing navigation.');
			return;
		} catch (error: any) {
			// TODO: if bad gateway, show error page
			console.error('[mid-auth] Error during token renewal:', error);
			// 即使刷新失败，也会继续执行到最后的 navigateTo('/login')
		}
	} else {
		console.log('[mid-auth] User object is null.');
	}

	console.log('[mid-auth] User not authenticated and token not renewed. Redirecting to /login.');
	return navigateTo('/login');
});
