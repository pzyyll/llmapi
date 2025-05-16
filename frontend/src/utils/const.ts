import type { constants } from 'node:sqlite';

// 使用 TypeScript 的枚举类型
export enum Status {
	Pending = 'pending',
	Success = 'success',
	Error = 'error'
}

export enum RequestPath {
	Login = '/api/login',
	Register = '/api/register',
	Logout = '/api/logout',
	UserInfo = '/api/user_info',
	RenewToken = '/api/renew_token',
	Users = '/api/users',
	DeleteUser = '/api/delete_user',
	CreateKey = '/api/create_api_key',
	GetKeys = '/api/api_keys',
	DeleteKey = '/api/delete_api_key'
}

export enum Role {
	Super = 'super',
	Admin = 'admin',
	User = 'user'
}

export const RoleNames: Record<Role, string> = {
	[Role.Super]: 'super_admin',
	[Role.Admin]: 'admin',
	[Role.User]: 'normal_user'
};

export enum Header {
	TurnstileToken = 'X-Turnstile-Token'
}
