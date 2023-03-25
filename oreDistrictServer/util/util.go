package util

import (
	"time"
)

func GetLocalTimeStr() string {
	now := time.Now()
	// 格式化时间为标准字符串
	dateString := now.Format("2006-01-02 15:04:05")
	return dateString
}
