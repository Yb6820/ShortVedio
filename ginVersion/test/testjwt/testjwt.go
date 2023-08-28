//引用文章
//https://blog.csdn.net/qq_50737715/article/details/124335666

package main

import (
	"DouYin/utils"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Telephone string `json:"telephone"`
	Password  string `json:"password"`
}

// jwt加密密钥
var jwtKey = []byte("a_secret_crect")

// token的claim
type Claims struct {
	UserId uint
	Name   string
	jwt.StandardClaims
}

// 发放token
func ReleaseToken(user User) (string, error) {

	//token的有效期
	expirationTime := time.Now().Add(7 * 24 * time.Hour)

	claims := &Claims{

		//自定义字段
		UserId: user.ID,
		Name:   user.Name,
		//标准字段
		StandardClaims: jwt.StandardClaims{

			//过期时间
			ExpiresAt: expirationTime.Unix(),
			//发放的时间
			IssuedAt: time.Now().Unix(),
			//发放者
			Issuer: "127.0.0.1",
			//主题
			Subject: "user token",
		},
	}

	//使用jwt密钥生成token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}

	//返回token
	return tokenString, nil
}

// 从tokenString中解析出claims并返回
func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtKey, nil
	})
	return token, claims, err
}
func AuthMiddleware() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		// 获取authorization header
		tokenString := ctx.GetHeader("Authorization")

		// validate token formate
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "message": "权限不足"})
			ctx.Abort()
			return
		}

		//提取token的有效部分（"Bearer "共占7位)
		tokenString = tokenString[7:]

		token, claims, err := ParseToken(tokenString)
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "message": "权限不足"})
			ctx.Abort()
			return
		}

		// 验证通过后获取claim 中的userId
		userId := claims.UserId
		DB := utils.DB
		var user User
		DB.First(&user, userId)

		// 用户不存在
		if user.ID == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "message": "权限不足"})
			ctx.Abort()
			return
		}

		// 用户存在将user的信息写入上下文，方便读取
		ctx.Set("user", user)

		ctx.Next()
	}
}

// 登录
func Login(ctx *gin.Context) {

	db := utils.DB

	//获取参数
	//此处使用Bind()函数，可以处理不同格式的前端数据
	var requestUser User
	ctx.Bind(&requestUser)
	telephone := requestUser.Telephone
	password := requestUser.Password

	//数据验证
	if len(telephone) != 11 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "手机号必须为11位",
		})
		return
	}
	if len(password) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "密码不能少于6位",
		})
		return
	}

	//判断手机号是否存在
	var user User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID == 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "用户不存在",
		})
		return
	}

	//判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "密码错误",
		})
	}

	//发放token
	token, err := ReleaseToken(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "系统异常",
		})
		//记录下错误
		fmt.Printf("token generate error: %v", err)
		return
	}

	//返回结果
	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"data":    gin.H{"token": token},
		"message": "登录成功",
	})
}

func Info(ctx *gin.Context) {

	user, _ := ctx.Get("user")
	//将用户信息返回
	ctx.JSON(http.StatusOK, gin.H{
		"data": gin.H{"user": user},
	})

}
func main() {
	r := gin.Default()
	r.GET("/login", func(ctx *gin.Context) {
		user_id, _ := strconv.Atoi(ctx.Query("user_id"))
		name := ctx.Query("name")
		user := User{
			ID:   uint(user_id),
			Name: name,
		}
		token, err := ReleaseToken(user)
		if err != nil {
			ctx.JSON(200, gin.H{
				"code": 200,
				"msg":  err,
			})
			fmt.Printf("something err:%v\n", err)
			return
		}
		ctx.JSON(200, gin.H{
			"code": 200,
			"msg":  token,
		})
	})
	r.GET("/get", func(ctx *gin.Context) {
		tokenString := ctx.Query("token")
		// 获取authorization header

		// validate token formate
		if tokenString == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "message": "权限不足1"})
			return
		}

		//提取token的有效部分（"Bearer "共占7位)
		//tokenString = tokenString[7:]

		token, claims, err := ParseToken(tokenString)
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "message": "权限不足2"})
			return
		}

		// 验证通过后获取claim 中的userId
		user := User{
			ID:   claims.UserId,
			Name: claims.Name,
		}
		ctx.JSON(200, gin.H{
			"code":    200,
			"message": user,
		})
	})
	r.Run()
}
