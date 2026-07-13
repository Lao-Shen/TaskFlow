package auth

import (
	"backend/internal/user"
	"backend/pkg/jwt"
	"backend/pkg/password"
	"errors"

	"gorm.io/gorm"
)

// ============================================================
//  业务错误定义
// ============================================================

var (
	ErrUserNotFound   = errors.New("用户不存在")
	ErrWrongPassword  = errors.New("密码错误")
	ErrUsernameTaken  = errors.New("用户名已被占用")
	ErrEmailTaken     = errors.New("邮箱已被注册")
	ErrRegisterFailed = errors.New("注册失败")
)

// ============================================================
//  注册
// ============================================================

func registerUser(req RegisterRequest) (*LoginResponse, error) {
	// 检查用户名是否重复
	if user.IsUsernameExist(req.Username) {
		return nil, ErrUsernameTaken
	}

	// 检查邮箱是否重复
	if user.IsEmailExist(req.Email) {
		return nil, ErrEmailTaken
	}

	// 密码哈希
	hashed, err := password.HashPassword(req.Password)
	if err != nil {
		return nil, ErrRegisterFailed
	}

	userInfo := &user.User{
		Username: req.Username,
		Password: hashed,
		Email:    req.Email,
	}

	if err := user.InsertUser(userInfo); err != nil {
		return nil, ErrRegisterFailed
	}

	// 生成 token
	token, err := authjwt.GenerateToken(userInfo.ID, userInfo.Username)
	if err != nil {
		return nil, ErrRegisterFailed
	}

	return &LoginResponse{Token: token, User: *userInfo}, nil
}

// ============================================================
//  登录
// ============================================================

func loginUser(req LoginRequest) (*LoginResponse, error) {
	userInfo, err := user.GetUserByUsername(req.Username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	if !password.CheckPassword(userInfo.Password, req.Password) {
		return nil, ErrWrongPassword
	}

	token, err := authjwt.GenerateToken(userInfo.ID, userInfo.Username)
	if err != nil {
		return nil, err
	}

	return &LoginResponse{Token: token, User: *userInfo}, nil
}
