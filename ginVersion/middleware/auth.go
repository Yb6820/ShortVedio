package middleware

import (
	"DouYin/models"
	"DouYin/utils/errmsg"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
	"net/http"
	"strings"
	"time"
)

var jwtKey = viper.GetString("jwt.key")

type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// 发放token
func SetToken(username string) (token string, err error) {
	expireTime := time.Now().AddDate(0, 0, 1)
	setClaims := MyClaims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			//过期时间
			ExpiresAt: expireTime.Unix(),
			//发放时间
			IssuedAt: time.Now().Unix(),
			//发放者
			Issuer: "shortVideo",
		},
	}
	reqClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, setClaims)
	token, err = reqClaim.SignedString(jwtKey)
	if err != nil {
		return "", jwt.ErrSignatureInvalid
	}
	return token, nil
}

// 解析token
func ParseToken(token string) (claims *MyClaims, err error) {
	setToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}
	//断言解析数据
	if key, _ := setToken.Claims.(*MyClaims); setToken.Valid {
		return key, nil
	} else {
		return nil, errors.New("token has invalid claims")
	}
}

// token中间件
func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHandler := c.Request.Header.Get("Authorization")
		code := errmsg.SUCCESS
		if tokenHandler == "" {
			code = errmsg.TOKEN_NOT_EXISTS
			c.JSON(http.StatusOK, gin.H{
				"status_code": code,
				"status_msg":  errmsg.GetErrMsg(code),
			})
			c.Abort()
		}
		//限制拆分的个数
		checkToken := strings.SplitN(tokenHandler, " ", 2)
		if len(checkToken) != 2 || checkToken[0] != "Bearer" {
			code = errmsg.TOKEN_TYPE_ERROR
			c.JSON(http.StatusOK, gin.H{
				"status_code": code,
				"status_msg":  errmsg.GetErrMsg(code),
			})
			c.Abort()
		}
		key, err := ParseToken(checkToken[1])
		if err != nil {
			code = errmsg.PARSE_TOKRN_ERROR
			c.JSON(http.StatusOK, gin.H{
				"status_code": code,
				"status_msg":  errmsg.GetErrMsg(code),
			})
			c.Abort()
		}
		if time.Now().Unix() > key.ExpiresAt {
			code = errmsg.TOKEN_RUNTIME_ERROR
			c.JSON(http.StatusOK, gin.H{
				"status_code": code,
				"status_msg":  errmsg.GetErrMsg(code),
			})
			c.Abort()
		}
		//获取用户信息并将有效的用户信息保存到ctx中
		c.Set("userInfo", models.UserInfo{
			Username: key.Username,
		})
		c.Next()
	}
}
