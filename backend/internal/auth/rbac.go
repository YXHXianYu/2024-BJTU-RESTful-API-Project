package auth

import (
	"net/http"
	"restfulapi/internal/model"

	"github.com/gin-gonic/gin"
	"github.com/mikespook/gorbac/v3"
)

var rbac *gorbac.RBAC[string]

func InitializeRBAC() {
    rbac = gorbac.New[string]()

    // 创建权限
    readPermission := gorbac.NewPermission("read")
    writePermission := gorbac.NewPermission("write")

    // 创建角色并分配权限
    adminRole := gorbac.NewRole("admin")
    adminRole.Assign(readPermission)
    adminRole.Assign(writePermission)

    userRole := gorbac.NewRole("user")
    userRole.Assign(readPermission)

    // 初始化RBAC结构并添加角色
    rbac.Add(adminRole)
    rbac.Add(userRole)
}

func RequireRole(allowedRoles ...string) gin.HandlerFunc {
    return func(c *gin.Context) {
        user, _ := c.Get("user")
        userRole := user.(model.UserSimplified).Role

        for _, role := range allowedRoles {
            if userRole == role {
                c.Next()
                return
            }
        }

        c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
            "code": http.StatusForbidden,
            "message": "Forbidden",
        })
    }
}