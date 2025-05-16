package role

import "llmapi/src/internal/constants"

func IsAdmin(role constants.RoleType) bool {
	return GetRoleLevel(role) > GetRoleLevel(constants.RoleTypeUser)
}

func GetRoleLevel(role constants.RoleType) constants.RoleLevel {
	return constants.RoleLevelMap[role]
}
