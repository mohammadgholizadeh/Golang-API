package token

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("my_secret_key")

type Claims struct {
	Id       int    `json:"id"`
	Role     string `json:"role"`
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func Jwt_maker(username string, id int, role string, password string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		Id:       id,
		Role:     role,
		Username: username,
		Password: password,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	} else {
		return tokenString, err
	}
}

func Jwt_Decoder(tknStr string) (int, string, string, error) {
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return 0, "", "", err
		}
	} else if !tkn.Valid {
		err = jwt.ValidationError{}
		return 0, "", "", err
	}

	// username given in the token
	//w.Write([]byte(fmt.Sprintf("Welcome %s!", claims.Username)))
	return claims.Id, claims.Username, claims.Role, err
}
