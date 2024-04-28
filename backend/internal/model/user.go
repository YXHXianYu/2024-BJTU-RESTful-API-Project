package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User 定义用户模型，对应数据库中的一个表
type User struct {
    gorm.Model
    Username    string `gorm:"column:username;unique;not null" json:"username"`
    Password    string `gorm:"column:password;not null" json:"password"`
    Email       string `gorm:"column:email;not null" json:"email"`
    Telephone   string `gorm:"column:telephone;not null" json:"telephone"`
	Role        string `gorm:"column:role;not null" json:"role"`
}

type UserSimplified struct {
	ID          uint   `json:"id"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	Telephone   string `json:"telephone"`
	Role        string `json:"role"`
}

// SetPassword 加密密码并存储
func (u *User) SetPassword(password string) error {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    if err != nil {
        return err
    }
    u.Password = string(bytes)
    return nil
}

// CheckPassword 验证密码
func (u *User) CheckPassword(password string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
    return err == nil
}

func (u *User) Simplify() UserSimplified {
	return UserSimplified{
		ID: u.ID,
		Username: u.Username,
		Email: u.Email,
		Telephone: u.Telephone,
		Role: u.Role,
	}
}