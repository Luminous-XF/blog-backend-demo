package test

import (
	"blog-backend/util"
	"fmt"
	"testing"
	"time"
)

func TestGetCurrentPath(t *testing.T) {
	fmt.Println(util.ProjectRootPath)
}

func TestParseConfig_1(t *testing.T) {
	testViper := util.ParseConfig("test")

	// 判断配置文件种是否配置该值
	if !testViper.IsSet("app-name") {
		t.Fail()
	}
	appName := testViper.GetString("app-name")

	if !testViper.IsSet("user-info.name") {
		t.Fail()
	}
	name := testViper.GetString("user-info.name")

	if !testViper.IsSet("user-info.age") {
		t.Fail()
	}
	age := testViper.GetInt("user-info.age")

	if !testViper.IsSet("user-info.password") {
		t.Fail()
	}
	password := testViper.GetString("user-info.password")

	fmt.Println(appName, name, age, password)
}

func TestParseConfig_2(t *testing.T) {
	type UserInfo struct {
		Name     string
		Age      int
		Password string
	}

	type Config struct {
		AppName  string   `json:"app-name" mapstructure:"app-name"`
		UserInfo UserInfo `json:"user-info" mapstructure:"user-info"`
	}

	var config Config
	testViper := util.ParseConfig("test")
	testViper.WatchConfig()
	if err := testViper.Unmarshal(&config); err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("Config: %+v\n", config)
}

func TestParseConfig_3(t *testing.T) {
	testViper := util.ParseConfig("test")
	// 实时监听配置文件变化
	testViper.WatchConfig()

	// 判断配置文件种是否配置该值
	if !testViper.IsSet("app-name") {
		t.Fail()
	}
	appName := testViper.GetString("app-name")
	fmt.Println(appName)

	time.Sleep(10 * time.Second)
	appName = testViper.GetString("app-name")
	fmt.Println(appName)
}
