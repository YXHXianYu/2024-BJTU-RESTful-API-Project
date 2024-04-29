package handler

import (
	"net/http"
	"restfulapi/internal/auth"
	"restfulapi/internal/model"
<<<<<<< HEAD
	"strings"
=======
>>>>>>> 21de4b8415dfe06201b4e0052e7c443b0dae38a8

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterBlogRoutes(r *gin.Engine, db *gorm.DB) {
    r.POST("/api/v1/blogs",       auth.Authenticate, auth.RequireRole("user", "admin"), CreateBlog(r, db))
    r.GET("/api/v1/blogs/:id",    auth.Authenticate, auth.RequireRole("user", "admin"), GetBlog(r, db))
	r.GET("/api/v1/blogs",        auth.Authenticate, auth.RequireRole("user", "admin"), GetAllBlogs(r, db))
	r.PATCH("./api/v1/blogs/:id", auth.Authenticate, auth.RequireRole("user", "admin"), UpdateBlog(r, db))
}

func CreateBlog(r *gin.Engine, db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var blog model.Blog
		// 1. get data
		if err := c.ShouldBindJSON(&blog); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": http.StatusBadRequest,
				"message": "Invalid JSON format",
			})
			return
		}
		// 2. check data
		if blog.Title == "" || blog.Content == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": http.StatusBadRequest,
				"message": "All fields are required",
			})
			return
		}
		user, _ := c.Get("user")
		blog.CreatedUserID = user.(model.UserSimplified).ID
		// 3. save data
		if err := db.Create(&blog).Error; err != nil {
<<<<<<< HEAD
			if strings.Contains(err.Error(), "Duplicate entry") {
				c.JSON(http.StatusBadRequest, gin.H{
					"code": http.StatusBadRequest,
					"message": "Blog title already exists",
				})
				return
			}

=======
>>>>>>> 21de4b8415dfe06201b4e0052e7c443b0dae38a8
			c.JSON(http.StatusBadRequest, gin.H{
				"code": http.StatusBadRequest,
				"message": err.Error(),
			})
			return
		}
		// 4. return data
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"message": "Blog created successfully!",
			"data": blog,
		})
	}
}

func GetBlog(r *gin.Engine, db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. get data
		id := c.Param("id")
		// 2. check data
		var blog model.Blog
		if err := db.First(&blog, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"code": http.StatusNotFound,
				"message": "Blog not found",
			})
			return
		}
		// 3. return data
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"message": "Blog found successfully!",
			"data": blog,
		})
	}
}

func GetAllBlogs(r *gin.Engine, db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. get data
		var blogs []model.Blog
		// 2. check data
		db.Find(&blogs)
		// 3. return data
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"message": "Blogs found successfully!",
			"data": blogs,
		})
	}
}

func UpdateBlog(r *gin.Engine, db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. get data
		id := c.Param("id")
		var blog model.Blog
		// 2. check data
		if err := db.First(&blog, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"code": http.StatusNotFound,
				"message": "Blog not found",
			})
			return
		}
		// 3. authenticate
		userInterface, _ := c.Get("user")
		user := userInterface.(model.UserSimplified)
		if user.Role != "admin" && user.ID != blog.CreatedUserID {
			c.JSON(http.StatusForbidden, gin.H{
				"code": http.StatusForbidden,
				"message": "You are not update to access this blog",
			})
			return
		}
		// 3. update data
		var newBlog model.Blog
		if err := c.ShouldBindJSON(&newBlog); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": http.StatusBadRequest,
				"message": "Invalid JSON format",
			})
			return
		}
		if newBlog.Title != "" { blog.Title = newBlog.Title }
		if newBlog.Content != "" { blog.Content = newBlog.Content }
		if err := db.Save(&blog).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": http.StatusBadRequest,
				"message": err.Error(),
			})
			return
		}
		// 4. return data
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"message": "Blog updated successfully!",
			"data": blog,
		})
	}
}