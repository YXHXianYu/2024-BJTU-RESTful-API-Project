package handler

import (
	"net/http"
	"restfulapi/internal/auth"
	"restfulapi/internal/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterAuthRoutes(r *gin.Engine, db *gorm.DB) {
	r.POST("/api/v1/sessions", Login(r, db))
	r.GET("/api/v1/secure", auth.Authenticate, auth.RequireRole("user", "admin"), CheckAuthenticate(r, db))
}

func Login(r *gin.Engine, db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginInfo struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}

		// 从请求体中解析 JSON
		if err := c.ShouldBindJSON(&loginInfo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
			return
		}

		// 从数据库中查找用户
		var user model.User
		if result := db.Where("username = ?", loginInfo.Username).First(&user); result.Error != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
			return
		}
		
		// 检查密码是否正确
		if !user.CheckPassword(loginInfo.Password) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
			return
		}

		// 生成token
		token, err := auth.GenerateToken(user.Simplify())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
			return
		}

		// 返回token
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"message": "Success.",
			"data": gin.H{
				"user": user.Simplify(),
				"token": token,
			},
		})
	}
}

func CheckAuthenticate(r *gin.Engine, db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"message": "Success",
		})
	}
}