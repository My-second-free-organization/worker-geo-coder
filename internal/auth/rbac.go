package auth

type Permission string
const (
	PermWorkflowRead   Permission = "workflow:read"
	PermWorkflowWrite  Permission = "workflow:write"
	PermTaskRead       Permission = "task:read"
	PermTaskWrite      Permission = "task:write"
	PermAdmin          Permission = "admin"
)

var rolePermissions = map[string][]Permission{
	"admin":   {PermWorkflowRead, PermWorkflowWrite, PermTaskRead, PermTaskWrite, PermAdmin},
	"manager": {PermWorkflowRead, PermWorkflowWrite, PermTaskRead, PermTaskWrite},
	"user":    {PermWorkflowRead, PermTaskRead, PermTaskWrite},
	"viewer":  {PermWorkflowRead, PermTaskRead},
}

func HasPermission(roles []string, perm Permission) bool {
	for _, role := range roles {
		for _, p := range rolePermissions[role] {
			if p == perm { return true }
		}
	}
	return false
}
