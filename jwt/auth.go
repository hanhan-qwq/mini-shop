package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var jwtSecret = []byte("user_jwt_secret")

const (
	AccessTokenExpireTime  = time.Hour * 2
	RefreshTokenExpireTime = time.Hour * 24 * 7
)

type Claims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

// GenerateTokens 生成 Access 和 Refresh Token
func GenerateTokens(userID uint, username, role string) (accessToken, refreshToken string, err error) {
	// Access Token
	accessClaims := Claims{
		UserID:   userID,
		Username: username,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(AccessTokenExpireTime)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "bookstore",
			Subject:   "access_token",
		},
	}
	accessToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims).SignedString(jwtSecret)
	if err != nil {
		return
	}

	// Refresh Token
	refreshClaims := Claims{
		UserID:   userID,
		Username: username,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(RefreshTokenExpireTime)),
			Issuer:    "bookstore",
			Subject:   "refresh_token",
		},
	}
	refreshToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString(jwtSecret)
	return
}

// ParseToken 解析token
func ParseToken(tokenStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, jwt.ErrTokenInvalidClaims
}
