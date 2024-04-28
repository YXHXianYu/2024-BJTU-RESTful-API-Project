// internal/auth/jwt.go
package auth

import (
	"net/http"
	"restfulapi/internal/model"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtKey = []byte("your_secret_key")

type Claims struct {
    User model.UserSimplified `json:"userSimplified"`
    jwt.StandardClaims
}

// GenerateToken 生成 JWT 令牌
func GenerateToken(user model.UserSimplified) (string, error) {
    expirationTime := time.Now().Add(1 * time.Hour)
    claims := &Claims{
        User: user,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expirationTime.Unix(),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString(jwtKey)

    return tokenString, err
}

// Authenticate 验证中间件
func Authenticate(c *gin.Context) {
    tokenString := c.GetHeader("Authorization")

    claims := &Claims{}

    token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
        return jwtKey, nil
    })

    if err != nil || !token.Valid {
        c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code": http.StatusUnauthorized,
			"message": "Unauthorized",
		})
        return
    }

    c.Set("user", claims.User)
    c.Next()
}

