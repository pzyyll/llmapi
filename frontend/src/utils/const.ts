import type { constants } from 'node:sqlite';

// 使用 TypeScript 的枚举类型
export enum Status {
	Pending = 'pending',
	Success = 'success',
	Error = 'error'
}

export enum RequestPath {
	LoadSettings = '/_dashboard/load_settings',
	Login = '/_dashboard/login',
	Register = '/_dashboard/register',
	Logout = '/_dashboard/logout',
	UserInfo = '/_dashboard/user_info',
	RenewToken = '/_dashboard/renew_token',
	Users = '/_dashboard/users',
	DeleteUser = '/_dashboard/delete_user',
	CreateKey = '/_dashboard/create_api_key',
	GetKeys = '/_dashboard/api_keys',
	DeleteKey = '/_dashboard/delete_api_key',
	UpdateKey = '/_dashboard/update_api_key',
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
