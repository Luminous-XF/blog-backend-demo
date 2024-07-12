package initialize

import (
	"blog-backend/global"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

func initViper() (config *viper.Viper) {
	config = viper.New()
	configPath := global.ProjectRootPath + "/config"
	config.AddConfigPath(configPath)
	config.SetConfigName("config")
	config.SetConfigType("yaml")

	// 读取配置文件
	if err := config.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if !errors.As(err, &configFileNotFoundError) {
			panic(fmt.Errorf("找不到配置文件: %s", err))
		} else {
			panic(fmt.Errorf("读取配置文件失败: %s", err))
		}
	}

	// 解析配置文件
	if err := config.Unmarshal(&global.CONFIG); err != nil {
		panic(fmt.Errorf("解析配置文件失败: %s", err))
	}

	config.WatchConfig()
	config.OnConfigChange(func(e fsnotify.Event) {
		if err := config.Unmarshal(&global.CONFIG); err != nil {
			panic(fmt.Errorf("解析配置文件失败: %s", err))
		}
	})

	return config
}
