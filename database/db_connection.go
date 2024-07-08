package database

import (
	"blog-backend/util"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	ormlog "gorm.io/gorm/logger"
)

var (
	blogMySQL     *gorm.DB
	blogMySQLOnce sync.Once

	dbLog ormlog.Interface
)

type DBConfig struct {
	Host     string `yaml:"host" mapstructure:"host"`
	Port     int    `yaml:"port" mapstructure:"port"`
	User     string `yaml:"user" mapstructure:"user"`
	Password string `yaml:"password" mapstructure:"password"`
}

func init() {
	dbLog = ormlog.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		ormlog.Config{
			// 慢 SQL 阈值
			SlowThreshold: time.Millisecond * 100,
			// Log level, Silent 表示不输出日志
			LogLevel: ormlog.Info,
			// 是否使用彩色打印
			Colorful: true,
		},
	)
}

func createMySQLDB(dbName string, dbConfig DBConfig) *gorm.DB {
	dataSourceName := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbName,
	)

	// 启用 PrepareStmt, SQL预编译, 提高查询效率
	var err error
	db, err := gorm.Open(mysql.Open(dataSourceName), &gorm.Config{Logger: dbLog, PrepareStmt: true})
	if err != nil {
		util.Logger.Panicf("connect to mysql use dataSourceName %s failed: %s", dataSourceName, err)
	}

	// 设置数据库连接池参数, 提高并发性能
	sqlDB, _ := db.DB()
	// 设置数据库连接池最大连接数
	sqlDB.SetMaxOpenConns(100)
	// 设置连接池最大允许的空闲连接shu, 若没有 sql 任务需要执行的连接数大于 20, 超过的连接会被连接池关闭
	sqlDB.SetMaxIdleConns(20)

	util.Logger.Infof("connect to mysql database %s success", dbName)

	return db
}

// GetBlogDBConnection 获取数据库连接
func GetBlogDBConnection() *gorm.DB {
	// 单例, 只创建一次
	blogMySQLOnce.Do(func() {
		dbName := "blog"
		dbConfig := *readDBConfig(dbName)
		blogMySQL = createMySQLDB(dbName, dbConfig)
	})

	return blogMySQL
}

// 读取数据库配置信息
func readDBConfig(dbName string) *DBConfig {
	var dbConfig DBConfig
	viper := util.ParseConfig("mysql")
	if err := viper.UnmarshalKey(dbName, &dbConfig); err != nil {
		util.Logger.Panicf("read db %s config file failed: %s", dbName, err)
	}
	return &dbConfig
}
