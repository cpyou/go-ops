package notice

import "go-ops/models/users"

type Notice interface {
	Notify(users []users.User) error // 发通知
	Record()  // 记录
 }

