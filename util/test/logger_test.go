package test

import (
	"blog-backend/util"
	"testing"
)

func TestLogger(t *testing.T) {
	util.InitLogger("log")
	util.Logger.Debug("this is a debug log")
	util.Logger.Info("this is a info log")
	util.Logger.Warn("this is a warn log")
	util.Logger.Error("this is a error log")
	// util.Logger.Panic("this is a panic log")
}
