package error

import (
	"errors"
	"fmt"
)

var ErrUserNotExist = errors.New("user not found")

// 登录
func login() {
	userInfo, err := getUserInfo(1)
	if errors.Is(err, ErrUserNotExist) {
		fmt.Errorf("login failed, %v\n", err)
	} else {
		fmt.Printf("Login successful, %v", userInfo)
	}
}

// 根据id获取用户信息
func getUserInfo(id int) (*User, error) {
	var user User = User{
		UserID:   1,
		UserName: "stx",
	}
	return &user, nil
}

type User struct {
	UserID   int64  `json-type:"user_id"`
	UserName string `json-type:"user_name"`
}
