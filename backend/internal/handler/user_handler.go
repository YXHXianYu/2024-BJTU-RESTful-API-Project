package handler

import (
	"net/http"
	"restfulapi/internal/auth"
	"restfulapi/internal/model"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterUserRoutes(r *gin.Engine, db *gorm.DB) {
    r.POST("/api/v1/users", CreateUser(r, db))
    r.GET("/api/v1/users/:id", auth.Authenticate, auth.RequireRole("user", "admin"), GetUser(r, db))
	r.GET("/api/v1/users", auth.Authenticate, auth.RequireRole("admin"), GetAllUsers(r, db))
}

func CreateUser(r *gin.Engine, db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
        var newUser model.User
		// 从请求体中解析 JSON
        if err := c.ShouldBindJSON(&newUser); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{
				"code": http.StatusBadRequest,
				"message": err.Error(),
			})
            return
        }
		// 检查是否有字段为空
		if newUser.Username == "" || newUser.Password == "" || newUser.Email == "" || newUser.Telephone == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": http.StatusBadRequest,
				"message": "All fields are required",
			})
			return
		}
		// 设置默认角色
		newUser.Role = "user"
		// 加密密码
		if err := newUser.SetPassword(newUser.Password); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": http.StatusBadRequest,
				"message": err.Error(),
			})
			return
		}
		// 保存到数据库
        if err := db.Create(&newUser).Error; err != nil {
			if strings.Contains(err.Error(), "Duplicate entry") {
				c.JSON(http.StatusBadRequest, gin.H{
					"code": http.StatusBadRequest,
					"message": "Username already exists",
				})
				return
			}

			c.JSON(http.StatusBadRequest, gin.H{
				"code": http.StatusBadRequest,
				"message": err.Error(),
			})
			return
		}
		// 返回成功消息
        c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"message": "User created successfully!",
			"data": newUser.Simplify(),
		})
    }
}

func GetUser(r *gin.Engine, db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
        var targetUser model.User
		{
			id := c.Param("id")
			result := db.First(&targetUser, id)
			if result.Error != nil {
				c.JSON(http.StatusNotFound, gin.H{
					"code": http.StatusNotFound,
					"message": "User not found",
				})
				return
			}
		}

		userInterface, _ := c.Get("user")
		user := userInterface.(model.UserSimplified)
		if user.Role != "admin" && user.ID != targetUser.ID {
			c.JSON(http.StatusForbidden, gin.H{
				"code": http.StatusForbidden,
				"message": "You are not allowed to access this user",
			})
			return
		}

        c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"message": "User found successfully!",
			"data": targetUser.Simplify(),
		})
    }
}

func GetAllUsers(r *gin.Engine, db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var users []model.User
		db.Find(&users)

		var usersResponse []model.UserSimplified
		for _, user := range users {
			usersResponse = append(usersResponse, user.Simplify())
		}

		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"message": "Users found successfully!",
			"data": usersResponse,
		})
	}
}