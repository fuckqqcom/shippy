package code

import (
	"github.com/dgrijalva/jwt-go"
	pb "shippy/user-service/proto/user"
	"time"
)

type Authable interface {
	Decode(tokenStr string) (*CustomClaims, error)
	Encode(user *pb.User) (string, error)
}

var privateKey = []byte("xs#a_1")

type CustomClaims struct {
	User *pb.User
	jwt.StandardClaims
}

type TokenService struct {
	repo Repository
}

func (t *TokenService) Encode(user *pb.User) (string, error) {
	expireTime := time.Now().Add(time.Hour * 24 * 3).Unix()

	claims := CustomClaims{
		user,
		jwt.StandardClaims{
			Issuer:    "go.micro.srv.user",
			ExpiresAt: expireTime,
		},
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return jwtToken.SignedString(privateKey)
}
