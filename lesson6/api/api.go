package api

import (
	"net/http"
	"time"

	"nilmod/dao"
	"nilmod/middleware"
	"nilmod/model"
	"nilmod/utils"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var req model.User
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
		})
		return
	}
	// 如果用户存在，这里这种是用户名可以一致的，即只要密码不一致就视为不同用户
	if dao.FindUser(req.Username, req.Password) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "user already exists",
		})
		return
	}
	dao.AddUser(req.Username, req.Password)
	c.JSON(http.StatusOK, gin.H{
		"message": "register success",
	})
}

func Login(c *gin.Context) {
	var req model.User
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
		})
		return
	}
	// 检查用户是否存在且密码是否正确
	if !dao.FindUser(req.Username, req.Password) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "user not found",
		})
		return
	}
	// 生成jwt token
	token, err := utils.MakeToken(req.Username, time.Now().Add(10*time.Minute))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
		return
	}
	// 返回token
	c.JSON(http.StatusOK, gin.H{
		"message": "login",
		"token":   token,
	})
}

func Ping1(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func Changepwd(c *gin.Context) {
	var req model.User
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
		})
		return
	}
	//c.Next()
	// 从中间件注入的 username（token 里的用户名）
	//fmt.Println(">>>", c.GetString("username"))
	tokenUser := c.GetString("username")
	if tokenUser == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}

	// 如果前端传了 username，确保与 token 中一致；否则使用 token 的用户名
	if req.Username != "" && req.Username != tokenUser {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "username mismatch"})
		return
	}
	req.Username = tokenUser

	if !dao.UserExists(req.Username) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "有这个用户吗你就瞎几把改改改",
		})
		return
	}

	dao.AddUser(req.Username, req.Password)
	c.JSON(http.StatusOK, gin.H{
		"message": "password changed successfully",
	})
}

func InitRouter_gin() {
	r := gin.Default()
	dao.Fileread()
	// 中间件使用示例
	r.GET("/ping", middleware.Example1(), middleware.Example2(), Ping1)

	r.POST("register", Register)
	r.POST("login", Login)
	r.POST("changepwd", middleware.Auth(), Changepwd)
	r.Run(":8080")
}
