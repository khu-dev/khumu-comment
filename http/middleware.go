package http

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/khu-dev/khumu-comment/repository"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/meehow/go-django-hashers"
	"log"
	"os"
)

type Authenticator struct{
	UserRepository repository.UserRepository
}

var KhumuJWTConfig middleware.JWTConfig = middleware.JWTConfig{
	Skipper: func(c echo.Context) bool{
		// or 조건의 auth를 하고싶은 경우
		// authHeader := c.Request().Header.Get("Authorization")
		// return !strings.HasPrefix(authHeader, "Bearer")
		return false
	},
  SigningKey: []byte(os.Getenv("KHUMU_SECRET")),
  SigningMethod: "HS256",
  ContextKey:    "user",
  TokenLookup:   "header:" + echo.HeaderAuthorization,
  AuthScheme:    "Bearer",
  Claims:        jwt.MapClaims{},
}

//func (a *Authenticator) KhumuJWTAuth(token string, c echo.Context) (bool, error){
//
//}

func (a *Authenticator) KhumuBasicAuth(username, password string, c echo.Context)(bool, error) {
	user := a.UserRepository.GetUserForAuth(username)
	log.Println("Try Authenticating ", username)
	if user == nil{
		return false, nil
	}else{
		found, err := hashers.CheckPassword(password, user.Password)
		log.Println("Authentication result: ", found, username)
		return found, err
	}
}