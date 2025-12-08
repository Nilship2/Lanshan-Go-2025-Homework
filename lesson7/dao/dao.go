// dao/user_dao.go
package dao

import (
	"lesson7/model"
	"time"

	"gorm.io/gorm"
)

var db *gorm.DB

// 初始化数据库连接
func InitDB(database *gorm.DB) {
	db = database
	// 自动迁移用户表
	db.AutoMigrate(&model.User{})
}

// AddUser 添加用户
func AddUser(username, password string) error {
	user := model.User{
		Username: username,
		Password: password, // 注意：实际应用中密码应该加密存储
	}
	result := db.Create(&user)
	return result.Error
}

// FindUser 查找用户（登录验证）
func FindUser(username, password string) bool {
	var user model.User
	result := db.Where("username = ? AND password = ?", username, password).First(&user)
	return result.Error == nil && result.RowsAffected > 0
}

// CheckUserExists 检查用户是否存在（注册时用）
func CheckUserExists(username string) bool {
	var user model.User
	result := db.Where("username = ?", username).First(&user)
	return result.Error == nil && result.RowsAffected > 0
}

// SaveRefreshToken 保存刷新令牌
func SaveRefreshToken(username, refreshToken string) error {
	// 设置令牌过期时间（比如7天后）
	expiry := time.Now().Add(7 * 24 * time.Hour)

	result := db.Model(&model.User{}).
		Where("username = ?", username).
		Updates(map[string]interface{}{
			"refresh_token": refreshToken,
			"token_expiry":  expiry,
		})
	return result.Error
}

// GetUserByRefreshToken 通过刷新令牌获取用户
func GetUserByRefreshToken(refreshToken string) (*model.User, error) {
	var user model.User
	result := db.Where("refresh_token = ? AND token_expiry > ?",
		refreshToken, time.Now()).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// UpdateUserToken 更新用户令牌
func UpdateUserToken(username, token string, expiry time.Time) error {
	result := db.Model(&model.User{}).
		Where("username = ?", username).
		Updates(map[string]interface{}{
			"refresh_token": token,
			"token_expiry":  expiry,
		})
	return result.Error
}
