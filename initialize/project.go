package initialize

import (
	"blog-backend/global"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
	"path"
	"runtime"
	"time"
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

	// gin router 路由配置
	engine := initGin()
	addr := fmt.Sprintf(":%d", global.CONFIG.ServerConfig.Addr)
	ReadTimeout := global.CONFIG.ServerConfig.ReadTimeout
	WriteTimeout := global.CONFIG.ServerConfig.WriteTimeout

	s := &http.Server{
		Addr:           addr,
		Handler:        engine,
		ReadTimeout:    ReadTimeout * time.Second,
		WriteTimeout:   WriteTimeout * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	// 初始化参数校验工具
	if validate := initValidator(); validate == nil {
		return errors.New("初始化校验器失败!")
	}

	if err := s.ListenAndServe(); err != nil {
		return errors.New("System Server Start Error!")
	}

	return nil
}

func InitProjectForTest() error {
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
