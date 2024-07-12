package initialize

import (
	"blog-backend/global"
	"github.com/pkg/errors"
	"path"
	"runtime"
)

func InitProject() error {
	// 设置项目根目录
	setProjectRootPath()

	// viper 初始化, 读取配置文件
	global.VP = initViper()
	if global.VP == nil {
		return errors.New("配置初始化失败!")
	}

	// logger 初始化
	global.Logger = initLogger()
	if global.Logger == nil {
		return errors.New("初始化日志工具失败!")
	}

	// gorm 初始化数据库
	global.GDB = initDB()
	if global.GDB == nil {
		return errors.New("数据库连接失败!")
	}

	return nil
}

func setProjectRootPath() {
	_, fileName, _, _ := runtime.Caller(0)
	global.ProjectRootPath = path.Dir(fileName) + "/../"
}
