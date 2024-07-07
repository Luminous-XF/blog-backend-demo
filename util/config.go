package util

import (
	"errors"
	"fmt"
	"path"
	"runtime"

	"github.com/spf13/viper"
)

// ProjectRootPath 项目根路径
var (
	ProjectRootPath = getCurrentPath() + "/../"
)

// 用于获取当前 config.go 文件所在路径
func getCurrentPath() string {
	_, fileName, _, _ := runtime.Caller(0)
	return path.Dir(fileName)
}

// ParseConfig 解析配置文件
func ParseConfig(fileName string) *viper.Viper {
	config := viper.New()
	// 配置文件所在目录
	configPath := ProjectRootPath + "config/"
	config.AddConfigPath(configPath)
	config.SetConfigName(fileName)
	config.SetConfigType("yaml")
	configFile := configPath + fileName + ".yaml"

	// 读取配置文件
	if err := config.ReadInConfig(); err != nil {
		// 若系统初始化阶段发生错误，则直接结束进程
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			panic(fmt.Errorf("找不到配置文件: %s", configFile))
		} else {
			panic(fmt.Errorf("解析配置文件 %s 出错: %s", configFile, err))
		}
	}

	return config
}
