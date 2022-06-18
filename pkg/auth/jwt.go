package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type Token struct {
	Key      string
	Expireat time.Duration
}

func NewToken(key string, time time.Duration) *Token {
	return &Token{
		Key:      key,
		Expireat: time,
	}
}
func (t *Token) RfTokenWithSession(id uuid.UUID) (error, string) {
	exipre := time.Now().Add(time.Hour * 24 * 30 * 12)
	rftoken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(exipre),
		NotBefore: jwt.NewNumericDate(time.Now()),
		Subject:   id.String(),
	})

	rftokenstring, err := rftoken.SignedString([]byte(t.Key))
	if err != nil {
		return err, ""
	}
	return nil, rftokenstring
}

func (t *Token) TokenWithUser(id uuid.UUID) (error, string) {
	expire := time.Now().Add(t.Expireat)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(expire),
		NotBefore: jwt.NewNumericDate(time.Now()),
		Subject:   id.String(),
	})
	tokenstring, err := token.SignedString([]byte(t.Key))
	if err != nil {
		return err, ""
	}
	return nil, tokenstring
}

func (t *Token) VerifyToken(jwtkey string) (error, uuid.UUID) {
	claimes := jwt.RegisteredClaims{}

	decoded, err := jwt.ParseWithClaims(jwtkey, &claimes, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("False")
		}
		return []byte(t.Key), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Name}))

	if err != nil {
		return err, uuid.UUID{}
	}
	if !decoded.Valid {
		return errors.New("Not Valid Token"), uuid.UUID{}
	}
	u, _ := uuid.Parse(claimes.Subject)
	return nil, u
}
