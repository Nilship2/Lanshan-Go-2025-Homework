package api

import (
	"log"
	"net/http"
	"time"

	"lesson7/dao"
	"lesson7/middleware"
	"lesson7/model"
	"lesson7/todo"
	"lesson7/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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
		c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		return
	}
	if !dao.FindUser(req.Username, req.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "user not found"})
		return
	}

	token, err := utils.MakeToken(req.Username, time.Now().Add(100*time.Minute))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return
	}

	refreshToken := utils.MakeRefreshToken()
	dao.SaveRefreshToken(req.Username, refreshToken)

	c.JSON(http.StatusOK, gin.H{
		"message":       "login",
		"token":         token,
		"refresh_token": refreshToken,
	})
}

func Ping1(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func InitRouter_gin() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/todolist?charset=utf8mb4&parseTime=True&loc=Local"
	//dsn := "root:123456@tcp(127.0.0.1:3306)/todolist?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("连接数据库失败: ", err)
		return
	}
	dao.InitDB(db)
	r := gin.Default()
	// 中间件使用示例
	//r.GET("/ping", middleware.Example1(), middleware.Example2(), Ping1)

	r.POST("register", Register)
	r.POST("login", Login)
	r.POST("create", middleware.Auth(), todo.Create)
	r.POST("read", middleware.Auth(), todo.Read)
	r.POST("update", middleware.Auth(), todo.Update)
	r.POST("delete", middleware.Auth(), todo.Delete)
	r.Run(":8080")
}
