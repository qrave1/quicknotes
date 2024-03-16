package auth

import (
	"context"
	"github.com/golang-jwt/jwt/v5"
	"github.com/qrave1/logger-wrapper/logrus"
	"strconv"
	"time"
)

const CurrentUserKey = "userId"

func UserIdFromCtx(ctx context.Context) int {
	return ctx.Value(CurrentUserKey).(int)
}

func SetUserId(ctx context.Context, id int) context.Context {
	return context.WithValue(ctx, CurrentUserKey, id)
}

type Auth interface {
	Generate(userId int) (string, error)
	//Inspect(ctx context.Context, token string) (bool, error)
}

type AuthService struct {
	secret string
	log    logrus.Logger
}

func NewAuthService(secret string, log logrus.Logger) *AuthService {
	return &AuthService{secret: secret, log: log}
}

func (a *AuthService) Generate(userId int) (string, error) {
	claims := jwt.RegisteredClaims{
		Issuer:    "quicknotes",
		Subject:   strconv.Itoa(userId),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(a.secret))
	if err != nil {
		return "", err
	}

	return t, nil
}

//func (a *AuthService) Inspect(ctx context.Context, token string) (bool, error) {
//	t, err := jwt.ParseWithClaims(token, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
//		return a.secret, nil
//	})
//	if err != nil {
//		return false, err
//	}
//
//	claims, ok := t.Claims.(*jwt.RegisteredClaims)
//	if !ok || !t.Valid {
//		a.log.Warn("error cast to claims")
//		return false, fmt.Errorf("error cast to claims")
//	}
//
//	currentId := UserIdFromCtx(ctx)
//	id, err := strconv.Atoi(claims.Subject)
//	if err != nil {
//		a.log.Warn("error convert subject to int")
//		return false, fmt.Errorf("error convert claims to int")
//	}
//
//	if currentId == 0 || currentId != id {
//		return false, fmt.Errorf("wrong subjest")
//	}
//
//	return true, nil
//}
