package service

import (
	"restfulapi/internal/model"

	"gorm.io/gorm"
)

func CreateBlog(db *gorm.DB, blog model.Blog) (error) {
	err := db.Create(&blog);
	return err.Error
}

func GetBlog(db *gorm.DB, id string) (model.Blog, error) {
	var blog model.Blog
	err := db.First(&blog, id);
	return blog, err.Error
}

func GetAllBlogs(db *gorm.DB) ([]model.Blog, error) {
	var blogs []model.Blog
	err := db.Find(&blogs)
	return blogs, err.Error
}

func UpdateBlog(db *gorm.DB, id string, newTitle string, newContent string) (error) {
	var blog model.Blog
	
	err := db.First(&blog, id);
	if err != nil { return err.Error }

	if newTitle != "" { blog.Title = newTitle }
	if newContent != "" { blog.Content = newContent }
	err = db.Save(&blog);

	if err != nil { return err.Error }

	return nil
}