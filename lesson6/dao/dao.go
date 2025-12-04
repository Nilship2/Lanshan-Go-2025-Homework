package dao

// 模拟数据库
var database = map[string]string{}

func AddUser(username string, password string) {
	database[username] = password
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
