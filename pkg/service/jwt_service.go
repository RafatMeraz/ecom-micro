package service

import (
	"github.com/RafatMeraz/ecom-micro/pkg/errors"
	"github.com/gofiber/fiber/v2/log"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type JwtCustomClaims struct {
	UserId uint `json:"user_id"`
	jwt.RegisteredClaims
}

type JwtServiceConfig struct {
	JwtSecret               string
	RefreshJwtSecret        string
	TokenExpireAtDay        int
	RefreshTokenExpireAtDay int
}

type JwtService struct {
	config *JwtServiceConfig
}

func NewJwtService(config *JwtServiceConfig) *JwtService {
	return &JwtService{
		config: config,
	}
}

func (j JwtService) generateToken(userId uint) (string, int64, error) {
	tokenDurationInDays := time.Now().Add(time.Hour * 24 * time.Duration(j.config.TokenExpireAtDay))

	claims := JwtCustomClaims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(tokenDurationInDays),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(j.config.JwtSecret))

	return tokenString, tokenDurationInDays.UnixMilli(), err
}

func (j JwtService) generateRefreshToken(userId uint) (string, int64, error) {
	tokenDurationInDays := time.Now().Add(time.Hour * 24 * time.Duration(j.config.RefreshTokenExpireAtDay))

	claims := JwtCustomClaims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(tokenDurationInDays),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(j.config.RefreshJwtSecret))

	return tokenString, tokenDurationInDays.UnixMilli(), err
}

func (j JwtService) getClaimsFromToken(tokenString string) (JwtCustomClaims, error) {
	var claims JwtCustomClaims
	_, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.config.JwtSecret), nil
	})

	if err != nil {
		return JwtCustomClaims{}, err
	}
	return claims, nil
}

func (j JwtService) PrepareJwtTokens(userId uint) (map[string]interface{}, error) {
	token, expireAt, err := j.generateToken(userId)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	refreshToken, refreshExpireAt, err := j.generateRefreshToken(userId)
	if err != nil {
		return nil, err
	}
	return map[string]interface{}{
		"access_token":            token,
		"access_token_expire_at":  expireAt,
		"refresh_token":           refreshToken,
		"refresh_token_expire_at": refreshExpireAt,
	}, nil
}

func (j JwtService) CheckJwtTokenValidity(token string) (JwtCustomClaims, error) {
	claims, err := j.getClaimsFromToken(token)
	if err != nil {
		return JwtCustomClaims{}, err
	}
	return claims, nil
}

func (j JwtService) RefreshToken(refreshToken string) (map[string]interface{}, error) {
	claims, err := j.getClaimsFromRefreshToken(refreshToken)
	if err != nil {
		return nil, errors.ErrInvalidToken
	}
	tokens, err := j.PrepareJwtTokens(claims.UserId)
	if err != nil {
		return nil, errors.ErrInvalidToken
	}
	return tokens, nil
}

func (j JwtService) getClaimsFromRefreshToken(tokenString string) (JwtCustomClaims, error) {
	var claims JwtCustomClaims
	_, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.config.RefreshJwtSecret), nil
	})

	if err != nil {
		return JwtCustomClaims{}, err
	}
	return claims, nil
}
