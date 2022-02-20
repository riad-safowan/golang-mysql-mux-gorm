package helpers

import (

	"log"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	// "go.mongodb.org/mongo-driver/mongo/options"
)

type SignedDetails struct {
	Email      string
	First_name string
	Last_name  string
	Uid        int
	User_type  string
	Token_type string
	jwt.StandardClaims
}

func GenerateAllToken(email string, firstName string, lastName string, userType string) (signedAccessToken string, signedRefreshToken string, err error) {
	accessClaims := &SignedDetails{
		Email:      email,
		First_name: firstName,
		Last_name:  lastName,
		User_type:  userType,
		Token_type: "access_token",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(30)).Unix(),
		},
	}

	refreshClaims := &SignedDetails{
		Email:      email,
		Token_type: "refresh_token",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24*7)).Unix(),
		},
	}

	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims).SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		log.Panic(err)
		return
	}
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		log.Panic(err)
		return
	}

	return accessToken, refreshToken, err

}

func ValidateToken(signedToken string) (claims *SignedDetails, msg string) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&SignedDetails{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SECRET_KEY")), nil
		},
	)

	if err != nil {
		msg = err.Error()
		return
	}

	claims, ok := token.Claims.(*SignedDetails)
	if !ok {
		msg = err.Error()
		return
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		msg = err.Error()
		return
	}

	return claims, msg
}
