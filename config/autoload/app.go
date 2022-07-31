package autoload

import (
	"github.com/liangdong/my-gin/pkg/convert"
	"github.com/liangdong/my-gin/pkg/utils"
)

type AppConfig struct {
	AppEnv         string `ini:"app_env" yaml:"app_env"`
	Debug          bool   `ini:"debug" yaml:"debug"`
	Language       string `ini:"language" yaml:"language"`
	StaticBasePath string `ini:"base_path" yaml:"base_path"`
}

var App = AppConfig{
	AppEnv:         "local",
	Debug:          true,
	Language:       "zh_CN",
	StaticBasePath: getDefaultPath(),
}

func getDefaultPath() (path string) {
	path, _ = utils.GetCurrentPath()
	path, _ = convert.GetString(utils.If(path != "", path, "/tmp"))
	return
}
