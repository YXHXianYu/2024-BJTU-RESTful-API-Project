package utils

import (
	"restfulapi/internal/auth"
	"restfulapi/internal/config"
	"restfulapi/internal/handler"
	"restfulapi/internal/model"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Setup() *gin.Engine {
	r := gin.Default()
	r.Use(cors.New(config.GetCorsConfig()))

	dsn := config.GetDSN()
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

    db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Blog{})

	auth.InitializeRBAC()

	handler.RegisterUserRoutes(r, db)
	handler.RegisterAuthRoutes(r, db)
	handler.RegisterBlogRoutes(r, db)

	return r
}