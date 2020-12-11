package http

import (
	"bytes"
	"github.com/dgrijalva/jwt-go"
	"github.com/khu-dev/khumu-comment/repository"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/meehow/go-django-hashers"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type Authenticator struct {
	UserRepository repository.UserRepositoryInterface
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
		logger := logrus.WithField("middleware", "Authenticator.Authenticate")
		if strings.HasPrefix(context.Request().Header.Get("Authorization"), "Bearer") {
			logger.Debug("Try JWT Authentication")
			return middleware.JWTWithConfig(KhumuJWTConfig)(
				// 토큰 속의 유저가 존재하는 유저인지 확인해서 분기하는 http Handler 끼워넣기
				func(context echo.Context) error {
					if token, ok := context.Get(KhumuJWTConfig.ContextKey).(*jwt.Token); ok {
						if mapClaim, ok := token.Claims.(jwt.MapClaims); ok{
							username := mapClaim["user_id"].(string)
							user := a.UserRepository.GetUserForAuth(username)
							if user != nil {
								context.Set("user_id", username)
								//여기까지 왔으면 존재하는 유저의 토큰
								logger.WithField("user_id", mapClaim["user_id"]).Println("Pass JWT Authentication.")
								return handlerFunc(context)
							}
						}
					}
					logger.Error("JWT Authentication failed")
					return context.JSON(401, map[string]interface{}{
						"statusCode": 401,
						"body":       "Request with a non-existing user.",
					})

				})(context)
		} else if strings.HasPrefix(context.Request().Header.Get("Authorization"), "Basic") {
			logger.Debug("Try Basic Authentication")
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
	if user == nil {
		return false, nil
	} else {
		found, err := hashers.CheckPassword(password, user.Password)
		c.Set("user_id", username)
		return found, err
	}
}

// application/json 요청인 경우 바디를 출력.
func KhumuRequestLog(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		logger := logrus.WithField("middleware", "KhumuRequestLog")
		req := context.Request()
		if req.Header.Get("Content-Type") != ""{
			logger.Println("Content-Type:", req.Header.Get("Content-Type"))
		}

		if (req.Method == http.MethodPost || req.Method == http.MethodPut || req.Method == http.MethodPatch) &&
			strings.HasPrefix(req.Header.Get("Content-Type"), "application/json"){
			rawBody, err := ioutil.ReadAll(req.Body)
			if err != nil {
				logger.Error(err)
			}
			// body는 stream 형태이므로 한 번 읽으면 다시 원상복구 시켜줘야함.
			// Restore the io.ReadCloser to it's original state
			req.Body = ioutil.NopCloser(bytes.NewBuffer(rawBody))
			body := &echo.Map{}
			err = context.Bind(body)
			// Restore the io.ReadCloser to it's original state
			// Bind에서 한 번 또 읽었으니 원상복구
			req.Body = ioutil.NopCloser(bytes.NewBuffer(rawBody))
			if err != nil{
				logger.Error("Body bind error:", err)
				return err
			}

			logger.Println("body: ", body)
		}
		return handlerFunc(context)
	}
}

func isAdmin(username string) bool{
	return username == "admin"
}
