package model

import (
	"errors"
	"time"
)

type BaseUser struct {
	Id string
	Username string
	Password string
	Status bool
	CreatedAt string
	UpdatedAt string
}

// 检测帐号状态
func (user *BaseUser) CheckStatus() error {
	if !user.Status {
		return errors.New("帐号已禁用.")
	}
	return nil
}

// 自动更改更新时间
func (user *BaseUser) AutoUpdateAt()  {
	user.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
}
