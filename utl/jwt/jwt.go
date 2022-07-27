package jwt

import (
	"errors"
	"go-alodokter/model"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var Key = "tS9ybXMcXKSNDYt7d63BA3F4Y9cJ5wTBJusCbc"

type BearerClaims struct {
	Email    string    `json:"email"`
	Access   []string  `json:"access"`
	Hostname string    `json:"hostname"`
	Exp      int64     `json:"exp"`
	Iat      time.Time `json:"iat"`
	jwt.StandardClaims
}

//  generate jwt
func Generate(email string, access []string) (token model.AuthAccess, err error) {
	now := time.Now().Add(-90 * time.Second)
	age := now.Add(time.Hour * 8)
	ageString := age.Format("2006-01-02 15:04:05")
	hostname, err := os.Hostname()
	if hostname == "" || err != nil {
		hostname = "localhost"
	}

	claim := make(jwt.MapClaims)
	claim["email"] = email
	claim["exp"] = age.Unix()
	claim["iat"] = now
	claim["hostname"] = hostname
	claim["access"] = access

	sign := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), claim)

	token.AccessToken, err = sign.SignedString([]byte(Key))
	token.ExpiredAt = ageString

	return token, err
}

//  verify
func Verify(token string) (clm *BearerClaims, r bool) {
	sign, err := jwt.ParseWithClaims(token, &BearerClaims{}, func(t *jwt.Token) (interface{}, error) {
		if t.Method != jwt.GetSigningMethod("HS256") {
			return t, errors.New("false")
		}
		return []byte(Key), nil

	})

	if err != nil {
		return clm, false
	}

	if clm, ok := sign.Claims.(*BearerClaims); ok {
		return clm, true
	}

	return clm, sign.Valid
}
