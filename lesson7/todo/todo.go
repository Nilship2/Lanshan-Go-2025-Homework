package todo

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Member struct {
	gorm.Model
	//index int    `gorm:"primaryKey;autoIncrement;comment:任务ID"`
	Name     string `gorm:"size:32;not null;comment:任务名称"`
	Deadline string `gorm:"default:'宇宙毁灭之时';comment:任务截止时间"`
	Kind     string `gorm:"size:64;index;comment:任务类型"`
}

func connectDB() (*gorm.DB, error) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/todolist?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("连接数据库失败: ", err)
		return nil, err
	}
	return db, nil
}

func Create(c *gin.Context) {
	var req Member
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
		})
		return
	}
	db, err := connectDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "database connection error",
		})
		return
	}
	db.Create(&req)
}

func Read(c *gin.Context) {
	var req int
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
		})
		return
	}
	db, err := connectDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "database connection error",
		})
		return
	}
	var m2 Member
	db.First(&m2, req)
	c.JSON(http.StatusOK, gin.H{
		"member": m2,
	})
}

func Update(c *gin.Context) {
	var req Member
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
		})
		return
	}
	db, err := connectDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "database connection error",
		})
		return
	}
	/*db.Model(&Member{}).Where("id = ?", req.ID).Updates(map[string]interface{}{
	    "name":     req.Name,
	    "deadline": req.Deadline,
	    "kind":     req.Kind,
	})*/
	//db.Model(&req).Where("id=?", req.ID).Updates(Member{Name: req.Name, Deadline: req.Deadline, Kind: req.Kind})
	result := db.Model(&Member{}).Where("id = ?", req.ID).Updates(req)

	// 或者使用 Select 指定要更新的字段
	// result := db.Model(&Member{}).Where("id = ?", req.ID).
	//     Select("Name", "Deadline", "Kind").
	//     Updates(req)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "更新失败: " + result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("更新成功，影响行数: %d", result.RowsAffected),
	})
}

func Delete(c *gin.Context) {
	var req int
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
		})
		return
	}
	db, err := connectDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "database connection error",
		})
		return
	}
	db.Delete(&Member{}, req)
}

func Justdoit() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/todolist?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("连接数据库失败: ", err)
	} else {
		err = db.AutoMigrate(&Member{})
		if err != nil {
			log.Fatal("自动迁移失败: ", err)
		}
		log.Println("自动迁移成功")

		log.Println("所有操作完成")
	}

}
