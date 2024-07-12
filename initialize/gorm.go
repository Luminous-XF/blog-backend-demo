package initialize

import (
	"blog-backend/global"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	ormlog "gorm.io/gorm/logger"
)

func initDB() *gorm.DB {
	dbConfig := global.CONFIG.DatabaseConfig
	if len(dbConfig.Name) == 0 {
		fmt.Printf("%#v\n", dbConfig)
		return nil
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Name,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: ormlog.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			ormlog.Config{
				// 慢 SQL 阈值
				SlowThreshold: time.Millisecond * 100,
				// Log level, Silent 表示不输出日志
				LogLevel: ormlog.Info,
				// 是否使用彩色打印
				Colorful: true,
			},
		),
		PrepareStmt: true,
	})

	if err != nil {
		global.Logger.Panicf("connect to mysql use dataSourceName %s failed: %s", dsn, err)
		return nil
	}

	// 设置数据库连接池参数, 提高并发性能
	mysqlConfig := global.CONFIG.MySQLConfig
	sqlDB, _ := db.DB()
	// 设置数据库连接池最大连接数
	sqlDB.SetMaxOpenConns(mysqlConfig.MaxOpenConnections)
	// 设置连接池最大允许的空闲连接shu, 若没有 sql 任务需要执行的连接数大于 20, 超过的连接会被连接池关闭
	sqlDB.SetMaxIdleConns(mysqlConfig.MaxIdleConnections)

	global.Logger.Infof("connect to mysql database %s success", dbConfig.Name)

	return db
}
