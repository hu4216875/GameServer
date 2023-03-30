package model

// Account 账户表
type Account struct {
	UserId     string
	Pwd        string
	AccountId  int64
	Nick       string
	ServerId   uint32
	CreateTime uint32
	UpdateTime uint32
	Forbidden  bool // 是否禁止登录
}
