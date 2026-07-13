package user

import (
	"backend/pkg/database"

	"gorm.io/gorm"
)

// ============================================================
//  用户数据库操作
// ============================================================

// db 快捷引用
func db() *gorm.DB {
	return database.DB
}

// AutoMigrate 自动迁移用户表
func AutoMigrate() error {
	return db().AutoMigrate(&User{})
}

// ---------- 增 ----------

func InsertUser(user *User) error {
	return db().Create(user).Error
}

// ---------- 删 ----------

func DeleteUser(id uint64) error {
	return db().Delete(&User{}, id).Error
}

// ---------- 改 ----------

func UpdateUser(user *User) error {
	return db().Save(user).Error
}

// ---------- 查 ----------

func GetUserByID(id uint64) (*User, error) {
	var user User
	err := db().First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserByUsername(username string) (*User, error) {
	var user User
	err := db().Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserByEmail(email string) (*User, error) {
	var user User
	err := db().Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func IsUsernameExist(username string) bool {
	var count int64
	db().Model(&User{}).Where("username = ?", username).Count(&count)
	return count > 0
}

func IsEmailExist(email string) bool {
	var count int64
	db().Model(&User{}).Where("email = ?", email).Count(&count)
	return count > 0
}
