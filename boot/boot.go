/*
 * @Author: liangdong09
 * @Date: 2022-07-19 00:31:13
 * @LastEditTime: 2022-10-04 21:59:08
 * @LastEditors: liangdong09
 * @Description:
 * @FilePath: /my_gin/boot/boot.go
 */
package boot

import (
	"flag"
	"fmt"

	"github.com/liangdong/my-gin/config"
	"github.com/liangdong/my-gin/data"
	"github.com/liangdong/my-gin/internal/pkg/logger"
	"github.com/liangdong/my-gin/internal/routers"
	"github.com/liangdong/my-gin/internal/task"
	"github.com/liangdong/my-gin/internal/validator"
)

func init() {
	var configPath string

	flag.StringVar(&configPath, "c", "", "请输入配置文件绝对路径")
	flag.Parse()

	// 1、初始化配置
	config.InitConfig(configPath)

	// 2、初始化zap日志
	logger.InitLogger()

	// 3、初始化数据库
	data.InitData()

	// 4、初始化验证器
	validator.InitValidatorTrans("zh")

	// 5、task
	task.InitTask()
	// task.SendDayilyMsg()
}

func Run() {
	r := routers.SetRouters()
	err := r.Run(fmt.Sprintf("%s:%d", config.Config.Server.Host, config.Config.Server.Port))
	if err != nil {
		panic(err)
	}
}
