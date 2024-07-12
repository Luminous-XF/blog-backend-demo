package global

import (
	"blog-backend/config"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	ProjectRootPath string
	GDB             *gorm.DB
	VP              *viper.Viper
	CONFIG          config.Config
	Logger          *logrus.Logger
)
