package registry

import (
	"fmt"
	"time"
)

//iss (issuer): Issuer of the JWT
//sub (subject): Subject of the JWT (the user)
//aud (audience): Recipient for which the JWT is intended
//exp (expiration time): Time after which the JWT expires
//nbf (not before time): Time before which the JWT must
//not be accepted for processing
//iat (issued at time): Time at which the JWT was issued;
//can be used to determine age of the JWT
type Key struct {
	Issuer string `json:"iss"`
	Purpose   string    `json:"purpose"` //api for things, access and refresh for users,
	Subject   string    `json:"sub"` //userid
	Audience  string    `json:"aud"` //igrid-message-bus, igrid-user-services
	IssuedAt  time.Time `json:"iat"`
	ExpiresAt time.Time `json:"exp"`
}

func NewKey(id, purpose string) Key {
	return Key{
		Purpose:   purpose,
		Subject:   id,
		IssuedAt:  time.Now(),
		ExpiresAt: time.Now().Add(5 * time.Minute),
	}
}

func (key Key)String() string {
	return fmt.Sprintf("iss: %v, purpose: %v, sub: %v aud: %v, iat: %v, exp: %v\n",
		key.Issuer,key.Purpose,key.Subject,key.Audience,key.IssuedAt, key.ExpiresAt)
}
