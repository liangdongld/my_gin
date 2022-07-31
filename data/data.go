package data

import (
	"sync"

	c "github.com/liangdong/my-gin/config"
)

var once sync.Once

func InitData() {
	once.Do(func() {
		if c.Config.Mysql.Enable {
			// 初始化 mysql
			initMysql()
		}

		if c.Config.Redis.Enable {
			// 初始化 redis
			initRedis()
		}
	})
}
