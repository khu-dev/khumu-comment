package rest

import (
	"bytes"
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/khu-dev/khumu-comment/ent"
	"github.com/khu-dev/khumu-comment/ent/khumuuser"
	"github.com/khu-dev/khumu-comment/errorz"
	"github.com/khu-dev/khumu-comment/util"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type Authenticator struct {
	Repo *ent.Client
}

func (a *Authenticator) Authenticate(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
	// middleware는 handlerFunc를 받아서 handlerFunc를 리턴함
	// handlerFunc는 그 속의 context가 http 응답을 할 수 있는녀석인듯함. 리턴 자체는 error만
	// 응답을하는 handlerFunc를 직접 리턴할 수도 있고,
	// handlerFunc를 호출해서 걔의 리턴값을 리턴할 수도 있음.
	// middleware는 자기가 다음에 수행해야할 handlerFunc를 인자로 받아서
	// 괜찮으면 받았던 handlerFunc를 수행
	// 안괜찮으면 error로 응답하는 handlerFunc를 수행하는 방식
	return func(ctx echo.Context) error {
		logger := log.WithField("middleware", "Authenticator.Authenticate")
		if strings.HasPrefix(ctx.Request().Header.Get("Authorization"), "Bearer") {
			logger.Debug("Try JWT Authentication")
			return middleware.JWTWithConfig(KhumuJWTConfig)(
				// 토큰 속의 유저가 존재하는 유저인지 확인해서 분기하는 http Handler 끼워넣기
				func(ctx echo.Context) error {
					background := context.Background()
					if token, ok := ctx.Get(KhumuJWTConfig.ContextKey).(*jwt.Token); ok {
						if mapClaim, ok := token.Claims.(jwt.MapClaims); ok {
							username := mapClaim["user_id"].(string)
							user, err := a.Repo.KhumuUser.Query().
								Select("username").
								Where(khumuuser.ID(username)).
								Only(background)
							if err != nil {
								log.Error(err, "JWT 인증 도중 에러 발생")
								return ctx.JSON(401, map[string]interface{}{
									"statusCode": 401,
									"body":       "Request with a non-existing user.",
								})
							}

							ctx.Set("user_id", username)
							//여기까지 왔으면 존재하는 유저의 토큰
							logger.WithField("user_id", user.ID).Println("Pass JWT Authentication.")
							return handlerFunc(ctx)
						}
					}
					logger.Error("JWT Authentication failed")
					return ctx.JSON(401, map[string]interface{}{
						"statusCode": 401,
						"body":       "Request with a non-existing user.",
					})
				})(ctx)
		} else if strings.HasPrefix(ctx.Request().Header.Get("Authorization"), "Basic") {
			logger.Debug("Try Basic Authentication")
			return ctx.JSON(401, map[string]interface{}{
				"statusCode": 401,
				"body":       "Basic authentication is not supported",
			})
		}

		return handlerFunc(ctx)
	}
}

var KhumuJWTConfig = middleware.JWTConfig{
	Skipper: func(c echo.Context) bool {
		return false
	},
	SigningKey:    []byte(os.Getenv("KHUMU_SECRET")),
	SigningMethod: "HS256",
	ContextKey:    "user",
	TokenLookup:   "header:" + echo.HeaderAuthorization,
	AuthScheme:    "Bearer",
	Claims:        jwt.MapClaims{},
}

// 이제 basic auth를 지원할 수 없음. 댓글 서버가 자체적으로 인증을 해줄 수 없기 때문
// 댓글 서버는 인증을 할 수 있는 credential이 존재하지 않음.
//func (a *Authenticator) KhumuBasicAuth(username, password string, c echo.Context) (bool, error) {
//	background := context.Background()
//	user, err := a.Repo.KhumuUser.Query().
//		Select("username").
//		Where(khumuuser.ID(username)).
//		Only(background)
//	if err != nil {
//		log.Error(err, "Basic 인증 도중 에러 발생")
//		return false, err
//	}
//	if user == nil {
//		return false, nil
//	} else {
//		found, err := hashers.CheckPassword(password, user.Password)
//		c.Set("user_id", username)
//		return found, err
//	}
//}

// application/json 요청인 경우 바디를 출력.
func KhumuRequestLog(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		logger := log.WithField("middleware", "KhumuRequestLog")
		req := context.Request()
		if req.Header.Get("Content-Type") != "" {
			logger.Println("Content-Type:", req.Header.Get("Content-Type"))
		}

		if (req.Method == http.MethodPost || req.Method == http.MethodPut || req.Method == http.MethodPatch) &&
			strings.HasPrefix(req.Header.Get("Content-Type"), "application/json") {
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
			if err != nil {
				logger.Error("Body bind error:", err)
				return err
			}

			logger.Println("body: ", body)
		}
		return handlerFunc(context)
	}
}

// HTTPErrorHandler 참고
// https://echo.labstack.com/guide/error-handling/
func CustomHTTPErrorHandler(err error, c echo.Context) {
	log.Errorf("%+v", err)
	if errors.Is(err, errorz.ErrBadRequest) {
		if err := c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"data":    nil,
			"message": util.GetSimpleErrorMessage(err),
		}); err != nil {
			log.Error(err)
		}
	} else if errors.Is(err, errorz.ErrUnauthorized) {
		if respErr := c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"data":    nil,
			"message": util.GetSimpleErrorMessage(err),
		}); respErr != nil {
			log.Error(respErr)
		}
	} else if errors.Is(err, errorz.ErrResourceNotFound) {
		if respErr := c.JSON(http.StatusNotFound, map[string]interface{}{
			"data":    nil,
			"message": util.GetSimpleErrorMessage(err),
		}); respErr != nil {
			log.Error(respErr)
		}
	} else {
		if respErr := c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"data":    nil,
			"message": "알 수 없는 오류로 작업을 수행하지 못했습니다. 쿠뮤에 문의해주세요.",
		}); respErr != nil {
			log.Error(respErr)
		}
	}
}
