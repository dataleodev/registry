package jwt

import (
	"github.com/dataleodev/registry"
	"github.com/dataleodev/registry/pkg/errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)


var (
	_ registry.Tokenizer = (*tokenizer)(nil)
	ErrInvalidToken = errors.New("invalid token")
	ErrTokenExpired = errors.New("token expired")
)

const (
	issuerName = "igrid-iam-jwt-tokenizer"
	audienceName = "igrid-iam-server"
)

type claims struct {
	jwt.StandardClaims
	Purpose string `json:"purpose"`
}

func (c claims) Valid() error {
	if c.Issuer != issuerName {
		return ErrInvalidToken
	}

	if c.Audience != audienceName{
		return ErrInvalidToken
	}
	return c.StandardClaims.Valid()
}

func (c claims) toKey() registry.Key {
	key := registry.Key{
		Issuer:    c.Issuer,
		Purpose:   c.Purpose,
		Subject:   c.Subject,
		Audience:  c.Audience,
		IssuedAt:  time.Unix(c.IssuedAt,0).UTC(),
		ExpiresAt: time.Unix(c.ExpiresAt,0).UTC(),
	}

	return key
}

type tokenizer struct {

	secret string
}

func NewTokenizer() registry.Tokenizer {
	secret := "a6feaf75-2ff7-45c2-a7a3-ab469f7bed37"
	return &tokenizer{
		secret: secret,
	}
}

func (t tokenizer) Issue(key registry.Key) (string, error) {
	cs := claims{
		StandardClaims: jwt.StandardClaims{
			Audience:  audienceName,
			ExpiresAt: key.ExpiresAt.UTC().Unix(),
			IssuedAt:  key.IssuedAt.UTC().Unix(),
			Issuer:    issuerName,
			Subject:   key.Subject,
		},
		Purpose: key.Purpose,
	}

	if !key.ExpiresAt.IsZero() {
		cs.ExpiresAt = key.ExpiresAt.UTC().Unix()
	}
	if key.Subject != "" {
		cs.Subject = key.Subject
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, cs)
	return jwtToken.SignedString([]byte(t.secret))
}

func (t tokenizer) Parse(token string) (registry.Key, error) {
	c := claims{}
	_, err := jwt.ParseWithClaims(token, &c, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token")
		}
		return []byte(t.secret), nil
	})
	if err != nil {
		if e, ok := err.(*jwt.ValidationError); ok && e.Errors == jwt.ValidationErrorExpired {
			return registry.Key{}, errors.Wrap(ErrTokenExpired, err)
		}
		return registry.Key{}, errors.Wrap(errors.New("invalid token"), err)
	}
	return c.toKey(), nil
}

