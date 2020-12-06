package auth

import (
	"fmt"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// UploadClaims struct
type uploadClaims struct {
	UID string `json:"uid"`
	jwt.StandardClaims
}

// Data struct
type Data struct {
	Secret []byte
}

// GetIDFromToken get id Upload in token JWT
func (state *Data) GetIDFromToken(token string) string {
	claim := parseToken(token, state.Secret)
	// log.Println(claim.UID)
	if claim != nil {
		return claim.UID
	}
	return ""
}

func parseToken(tokenString string, secret []byte) *uploadClaims {


	token, err := jwt.ParseWithClaims(tokenString, &uploadClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return secret, nil
	})

	if err != nil {
		log.Println(err)
		// return nil
	}

	claims, ok := token.Claims.(*uploadClaims)
	if !ok || !token.Valid {
		return nil
	}
	return claims
}

// GenerateNewTokenWithID create token JWT with UID (Upload ID) claims
func (state *Data) GenerateNewTokenWithID(id string) (string, error) {
	claims := uploadClaims{
		id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * time.Duration(10)).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "test",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(state.Secret)
}
