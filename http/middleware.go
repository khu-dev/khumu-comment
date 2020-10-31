package http

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/khu-dev/khumu-comment/repository"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/meehow/go-django-hashers"
	"log"
	"os"
	"strings"
)

type Authenticator struct {
	UserRepository repository.UserRepository
}

func (a *Authenticator) Authenticate(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
	// middleware는 handlerFunc를 받아서 handlerFunc를 리턴함
	// handlerFunc는 그 속의 context가 http 응답을 할 수 있는녀석인듯함. 리턴 자체는 error만
	// 응답을하는 handlerFunc를 직접 리턴할 수도 있고,
	// handlerFunc를 호출해서 걔의 리턴값을 리턴할 수도 있음.
	// middleware는 자기가 다음에 수행해야할 handlerFunc를 인자로 받아서
	// 괜찮으면 받았던 handlerFunc를 수행
	// 안괜찮으면 error로 응답하는 handlerFunc를 수행하는 방식
	return func(context echo.Context) error {
		if strings.HasPrefix(context.Request().Header.Get("Authorization"), "Bearer") {
			fmt.Println("JWT 인증")
			return middleware.JWTWithConfig(KhumuJWTConfig)(
				// 토큰 속의 유저가 존재하는 유저인지 확인해서 분기하는 http Handler 끼워넣기
				func(context echo.Context) error {
					if token, ok := context.Get(KhumuJWTConfig.ContextKey).(*jwt.Token); ok {
						if mapClaim, ok := token.Claims.(jwt.MapClaims); ok && mockCheckUserExists(mapClaim["user_id"].(string)) {
							context.Set("user_id", mapClaim["user_id"])
							//여기까지 왔으면 존재하는 유저의 토큰
							return handlerFunc(context)
						}
					}
					return context.JSON(401, map[string]interface{}{
						"statusCode": 401,
						"body":       "Request with a non-existing user.",
					})

				})(context)
		} else if strings.HasPrefix(context.Request().Header.Get("Authorization"), "Basic") {
			fmt.Println("Basic 인증")
			return middleware.BasicAuth(a.KhumuBasicAuth)(handlerFunc)(context)
		} else {
			return context.JSON(401, map[string]interface{}{
				"statusCode": 401,
				"response":   "Unauthorized error. Please pass a valid JWT token or Basic Auth information.",
			})
		}
	}
}

var KhumuJWTConfig middleware.JWTConfig = middleware.JWTConfig{
	Skipper: func(c echo.Context) bool {
		// 이 미들웨어를 pass 시키지 않음.
		return false
	},
	SigningKey:    []byte(os.Getenv("KHUMU_SECRET")),
	SigningMethod: "HS256",
	ContextKey:    "user",
	TokenLookup:   "header:" + echo.HeaderAuthorization,
	AuthScheme:    "Bearer",
	Claims:        jwt.MapClaims{},
}

func (a *Authenticator) KhumuBasicAuth(username, password string, c echo.Context) (bool, error) {
	user := a.UserRepository.GetUserForAuth(username)
	log.Println("Try Authenticating ", username)
	if user == nil {
		return false, nil
	} else {
		found, err := hashers.CheckPassword(password, user.Password)
		log.Println("Authentication result: ", found, username)
		return found, err
	}
}

func mockCheckUserExists(username string) bool {
	if username == "jinsu" {
		return true
	} else {
		return false
	}

}
