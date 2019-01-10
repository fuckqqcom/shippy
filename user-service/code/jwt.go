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

var privateKey = []byte("123wrgdgxs#ahfthfjy_119gdxfsdf9145sdnl")

type CustomClaims struct {
	User *pb.User
	jwt.StandardClaims
}

type TokenService struct {
	repo Repository
}

func (t *TokenService) Decode(token string) (*CustomClaims, error) {
	to, err := jwt.ParseWithClaims(token, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return privateKey, nil
	})

	if claims, ok := to.Claims.(*CustomClaims); ok && to.Valid {
		return claims, nil
	} else {
		return nil, err
	}
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
