package utils

import (
	"nilmod/model"
	"time"
)

func MakeToken(username string, expireTime time.Time) (string, error) {
	token := username + "|" + expireTime.Format(time.RFC3339)
	return token, nil
}

func ParseToken(token string) (model.User, error) {
	parts := len(token)
	if parts < 2 {
		return model.User{}, nil
	}
	username := token[:parts-1]
	expireStr := token[parts-1:]
	expireTime, err := time.Parse(time.RFC3339, expireStr)
	if err != nil {
		return model.User{}, err
	}
	if time.Now().After(expireTime) {
		return model.User{}, nil
	}
	return model.User{Username: username}, nil
}
