package service

import (
	"crypto/sha1"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	pass := string(hash.Sum([]byte(salt)))
	return pass
}

func (s *AuthService) GenerateToken(userLogin, userPassword string) (string, error) {
	user, err := s.repo.GetUser(userLogin, generatePasswordHash(userPassword))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	return token.SignedString([]byte(singInKey))
}

func (s *AuthService) ParseToken(accessToken string) (int64, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signin method")
		}
		return []byte(singInKey), nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("invalid claimss")
	}

	return claims.UserId, nil
}
