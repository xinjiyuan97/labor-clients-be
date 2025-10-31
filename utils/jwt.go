package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// JWT配置
var (
	JWTSecret     = []byte("your-secret-key") // 应该从配置文件读取
	TokenExpire   = 24 * time.Hour            // token过期时间
	RefreshExpire = 7 * 24 * time.Hour        // refresh token过期时间
)

// Claims JWT声明
type Claims struct {
	UserID  int64  `json:"user_id"`
	Role    string `json:"role"`
	BrandID *int64 `json:"brand_id,omitempty"` // 品牌ID（用于品牌管理员和门店管理员）
	StoreID *int64 `json:"store_id,omitempty"` // 门店ID（用于门店管理员）
	jwt.RegisteredClaims
}

// GenerateToken 生成JWT token
func GenerateToken(userID int64, role string) (string, error) {
	return GenerateTokenWithExtra(userID, role, nil, nil)
}

// GenerateTokenWithExtra 生成JWT token（包含额外的品牌ID和门店ID）
func GenerateTokenWithExtra(userID int64, role string, brandID, storeID *int64) (string, error) {
	claims := Claims{
		UserID:  userID,
		Role:    role,
		BrandID: brandID,
		StoreID: storeID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExpire)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "labor-clients",
			Subject:   "user-auth",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JWTSecret)
}

// GenerateRefreshToken 生成refresh token
func GenerateRefreshToken(userID int64, role string) (string, error) {
	return GenerateRefreshTokenWithExtra(userID, role, nil, nil)
}

// GenerateRefreshTokenWithExtra 生成refresh token（包含额外的品牌ID和门店ID）
func GenerateRefreshTokenWithExtra(userID int64, role string, brandID, storeID *int64) (string, error) {
	claims := Claims{
		UserID:  userID,
		Role:    role,
		BrandID: brandID,
		StoreID: storeID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(RefreshExpire)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "labor-clients",
			Subject:   "refresh-auth",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JWTSecret)
}

// ParseToken 解析JWT token
func ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return JWTSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

// ValidateToken 验证token是否有效
func ValidateToken(tokenString string) (*Claims, error) {
	return ParseToken(tokenString)
}
