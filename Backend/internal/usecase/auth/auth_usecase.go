package auth

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/qrave1/logger-wrapper/logrus"
	"strconv"
	"time"
)

type defaultUserKey struct{}

func SetUserId(ctx context.Context, id int) context.Context {
	return context.WithValue(ctx, defaultUserKey{}, id)
}

func UserId(ctx context.Context) int {
	return ctx.Value(defaultUserKey{}).(int)
}

type Auth interface {
	Generate(userId int) (string, error)
	Inspect(ctx context.Context, token string) (bool, error)
}

type AuthImpl struct {
	token string
	log   logrus.Logger
}

func (a *AuthImpl) Generate(userId int) (string, error) {
	claims := jwt.RegisteredClaims{
		Issuer:    "quicknotes",
		Subject:   strconv.Itoa(userId),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(a.token))
	if err != nil {
		return "", err
	}

	return t, nil
}

func (a *AuthImpl) Inspect(ctx context.Context, token string) (bool, error) {
	t, err := jwt.ParseWithClaims(token, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return a.token, nil
	})
	if err != nil {
		return false, err
	}

	claims, ok := t.Claims.(*jwt.RegisteredClaims)
	if !ok || !t.Valid {
		a.log.Warn("error cast to claims")
		return false, fmt.Errorf("error cast to claims")
	}

	id := UserId(ctx)
	i, err := strconv.Atoi(claims.Subject)
	if err != nil {
		a.log.Warn("error convert claims to int")
		return false, fmt.Errorf("error convert claims to int")
	}

	if id == 0 || id != i {
		return false, fmt.Errorf("wrong subjest")
	}

	return true, nil
}
