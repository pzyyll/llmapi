export {};

declare global {
	namespace Dto {
		interface User {
			user_id: string;
			username: string;
			role: string;
			email: string;
			created_at: string;
		}

		interface Users {
			users: User[];
		}

		interface LoginResponse {
			access_token: string;
			user: User;
		}

		interface RegisterResponse {
			access_token: string;
			user: User;
		}

		interface RenewTokenResponse {
			access_token: string;
			user: User;
		}
		interface Key {
			user_id: string;
			name: string;
			scopes: number;
			lookup_key: string;
			created_at: string;
			updated_at: string;
			last_used_at: string | null;
			expire_at: string | null;
			secret_brief: string;
		}

		interface Keys {
			api_keys: Key[];
		}

		interface CreateKeyResponse {
			key: Key;
			secret: string;
		}
	}
}
