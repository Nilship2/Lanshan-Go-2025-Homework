package utils

import (
	"fmt"
	"time"
)

func MakeToken(username string, expireTime time.Time) (string, error) {
	token := username + "|" + expireTime.Format(time.RFC3339)
	return token, nil
}

func ParseToken(token string) (string, error) {
	leng := len(token)
	if leng < 2 {
		return "", nil
	}
	parts := 0
	for i := leng - 1; i >= 0; i-- {
		if token[i] == '|' {
			parts = i
		}
	}
	username := token[:parts]
	expireStr := token[parts+1:]
	fmt.Println(username)
	fmt.Println(expireStr)
	expireTime, err := time.Parse(time.RFC3339, expireStr)
	if err != nil {
		fmt.Println("114514")
		return "", err
	}
	if time.Now().After(expireTime) {
		fmt.Println("191980")
		return "", nil
	}
	return username, nil
}
