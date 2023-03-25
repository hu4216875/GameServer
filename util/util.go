package util

import "time"

func GetCurTime() uint32 {
	return uint32(time.Now().Unix())
}
