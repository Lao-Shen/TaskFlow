package user

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// ============================================================
//  业务错误定义
// ============================================================

var (
	ErrUserNotFound      = errors.New("用户不存在")
	ErrWrongPassword     = errors.New("密码错误")
	ErrUsernameTaken     = errors.New("用户名已被占用")
	ErrEmailTaken        = errors.New("邮箱已被注册")
	ErrRegisterFailed    = errors.New("注册失败")
)

// ============================================================
//  JWT
// ============================================================

func generateToken(user *User) (string, error) {
	claims := &Claims{
		UserID:   user.ID,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "TaskFlow",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(JWTSecret))
}

// ============================================================
//  密码
// ============================================================

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func checkPassword(hashed, plain string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain)) == nil
}

// ============================================================
//  注册
// ============================================================

func registerUser(req RegisterRequest) (*LoginResponse, error) {
	// 检查用户名是否重复
	if IsUsernameExist(req.Username) {
		return nil, ErrUsernameTaken
	}

	// 检查邮箱是否重复
	if IsEmailExist(req.Email) {
		return nil, ErrEmailTaken
	}

	// 密码哈希
	hashed, err := hashPassword(req.Password)
	if err != nil {
		return nil, ErrRegisterFailed
	}

	user := &User{
		Username: req.Username,
		Password: hashed,
		Email:    req.Email,
	}

	if err := InsertUser(user); err != nil {
		return nil, ErrRegisterFailed
	}

	// 生成 token
	token, err := generateToken(user)
	if err != nil {
		return nil, ErrRegisterFailed
	}

	return &LoginResponse{Token: token, User: *user}, nil
}

// ============================================================
//  登录
// ============================================================

func loginUser(req LoginRequest) (*LoginResponse, error) {
	user, err := GetUserByUsername(req.Username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	if !checkPassword(user.Password, req.Password) {
		return nil, ErrWrongPassword
	}

	token, err := generateToken(user)
	if err != nil {
		return nil, err
	}

	return &LoginResponse{Token: token, User: *user}, nil
}