package main

import (
    "github.com/gin-gonic/gin"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

type Product struct {
    gorm.Model
    Code  string
    Price uint
}

func main() {
    r := gin.Default()
    db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

    // 自动迁移schema
    db.AutoMigrate(&Product{})

    // 设置路由
    r.GET("/products", func(c *gin.Context) {
        var products []Product
        db.Find(&products)
        c.JSON(200, products)
    })

    r.POST("/products", func(c *gin.Context) {
        var product Product
        if err := c.ShouldBindJSON(&product); err != nil {
            c.JSON(400, gin.H{"error": err.Error()})
            return
        }
        db.Create(&product)
        c.JSON(200, product)
    })

    // 启动服务
    r.Run() // 默认监听并在 0.0.0.0:8080 上运行
}
