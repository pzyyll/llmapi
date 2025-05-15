export {};

declare global {
	namespace Dto {
		interface User {
			user_id: number;
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
	}
}
