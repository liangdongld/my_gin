package utils

import (
	"errors"

	"github.com/liangdong/my-gin/config/autoload"
)

func FindRobotByName(config autoload.WeChatConfig, name string) (autoload.Robot, error) {
	for i := 0; i < len(config.Rotbot); i++ {
		if config.Rotbot[i].Name == name {
			return config.Rotbot[i], nil
		}
	}
	return autoload.Robot{}, errors.New("未找到机器人")
}
