package dao

import (
	"bufio"
	"fmt"
	"os"
)

// 模拟数据库
var database = map[string]string{}

func Fileread() {
	logfile, _ := os.OpenFile("userdata.txt", os.O_RDONLY, 0666)
	reader := bufio.NewReader(logfile)
	for line, _, _ := reader.ReadLine(); line != nil; line, _, _ = reader.ReadLine() {
		var i int
		var username string
		var password string
		username = ""
		password = ""
		for i = 0; line[i] != ','; i++ {
			username += string(line[i])
		}
		i++
		for ; i < len(line); i++ {
			password += string(line[i])
		}
		database[username] = password
	}
	logfile.Close()
}

func Fileadd(key string, value string) {
	logfile, _ := os.OpenFile("userdata.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	//fmt.Println("嗨嗨嗨")
	fmt.Fprintf(logfile, "%s,%s\n", key, value)
	defer logfile.Close()

}

func Fileope(key string, value string) {
	logfile2, _ := os.OpenFile("userdata.txt", os.O_RDONLY, 0666)
	reader := bufio.NewReader(logfile2)
	logfile, _ := os.OpenFile("userdata.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)

	for line, _, _ := reader.ReadLine(); line != nil; line, _, _ = reader.ReadLine() {
		var i int
		var username string
		var password string
		username = ""
		password = ""
		for i = 0; line[i] != ','; i++ {
			username += string(line[i])
		}
		i++
		for ; i < len(line); i++ {
			password += string(line[i])
		}
		if username == key {
			password = value
			//fmt.Println("嗨嗨嗨！！！")
		}
		fmt.Fprintf(logfile, "%s,%s\n", username, password)
	}
}
func AddUser(username string, password string) {

	if UserExists(username) {
		//fmt.Println("嗨嗨嗨！！！")
		Fileope(username, password)
		database[username] = password
		return
	}
	database[username] = password
	//fmt.Println("嚯嚯嚯！！！")
	Fileadd(username, password)
}

func UserExists(username string) bool {
	_, ok := database[username]
	return ok
}

func FindUser(username string, password string) bool {
	if pwd, ok := database[username]; ok {
		if pwd == password {
			return true
		}
	}
	return false
}
func SelectPasswordFromUsername(username string) string {
	return database[username]
}
