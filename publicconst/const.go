package publicconst

type PlayerState int

const (
	Logining PlayerState = iota // 登录中
	Online                      // 在线
	Offline                     // 离线
)
