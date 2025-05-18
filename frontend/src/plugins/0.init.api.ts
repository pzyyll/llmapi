import { HttpStatusCode } from 'axios';
import axios from 'axios';

export default defineNuxtPlugin((nuxtApp) => {
	const runtimeConfig = useRuntimeConfig();
	const authStore = useAuthStore();

	const api = axios.create({
		baseURL: runtimeConfig.public.apiBase,
		withCredentials: true,
		headers: {
			'Content-Type': 'application/json'
		}
	});

	api.interceptors.request.use(
		(config) => {
			if (authStore.getToken) {
				config.headers.Authorization = `Bearer ${authStore.getToken}`;
			}
			return config;
		},
		(error) => {
			console.error('Error in request interceptor:', error);
			return Promise.reject(error);
		}
	);

	api.interceptors.response.use(
		(response) => {
			return response;
		},
		async (error) => {
			const originalRequest = error.config;
			if (error.response.status === HttpStatusCode.Unauthorized) {
				if (originalRequest._retry || originalRequest.url === RequestPath.RenewToken) {
					console.error('Token refresh failed, redirecting to login');
					authStore.clearAuth();
					await nuxtApp.runWithContext(() => navigateTo('/login'));
					return Promise.reject(error);
				}
				originalRequest._retry = true;
				try {
					const { data } = await api.post<Dto.RenewTokenResponse>(RequestPath.RenewToken);
					authStore.setToken(data.access_token);
					authStore.setUser(data.user);
					return api(originalRequest);
				} catch (err) {
					console.error('Error while refreshing token:', err);
					// await nuxtApp.runWithContext(() => navigateTo('/login'));
					return Promise.reject(err);
				}
			} else if (error.response.status === HttpStatusCode.BadGateway) {
				console.error('Bad Gateway error:', error);
				showError({
					statusCode: error.response.status
				});
			}
			return Promise.reject(error);
		}
	);

	return {
		provide: {
			api
		}
	};
});
