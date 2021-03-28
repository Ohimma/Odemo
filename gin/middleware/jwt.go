package middleware

import (
	"errors"
	"time"

	"github.com/ohimma/odemo/gin/config"
	"github.com/ohimma/odemo/gin/handler"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type JWT struct {
	SigningKey []byte
}

type UserClaims struct {
	Name string
	jwt.StandardClaims
}

// 工厂函数，定义一系列 jwt 方法
func NewJWT() *JWT {
	return &JWT{
		[]byte(config.Conf.Jwt.Signkey),
	}
}

func (j *JWT) GenerateToken(Name string) (string, error) {
	claims := UserClaims{
		Name: Name,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,    // 签名生效时间
			ExpiresAt: time.Now().Unix() + 60*60*2, // 过期时间 2h
			Issuer:    "odemo",                     // 签名的发行者
		},
	}
	Logger.Info("生成token claims = ", claims)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

func (j *JWT) ParseToken(t string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(t, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		v, ok := err.(*jwt.ValidationError)
		if ok {
			if v.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errors.New("That's not even a token")
			} else if v.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, errors.New("Token is expired")
			} else if v.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, errors.New("Token not active yet")
			} else {
				return nil, errors.New("Couldn't handle this token:")
			}
		}
	}
	if token != nil {
		claims, ok := token.Claims.(*UserClaims)
		if ok && token.Valid {
			return claims, nil
		}
		return nil, errors.New("Couldn't handle this token:")

	} else {
		return nil, errors.New("Couldn't handle this token:")

	}
}

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		Logger.Info("enter middleware auth ..... token = ", token)
		if token == "" {
			handler.ResponseUnauthorized(c, 401, token, "token为空")
			c.Abort()
			return
		}

		j := *NewJWT()

		claims, err := j.ParseToken(token)

		if err != nil {
			handler.ResponseUnauthorized(c, 401, err, "解析token失败")
			c.Abort()
			return
		}
		c.Set("claims", claims)
		c.Next()
	}

}
