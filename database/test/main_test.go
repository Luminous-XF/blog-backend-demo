package test

import (
	"blog-backend/initialize"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	before()
	code := m.Run()
	after()
	os.Exit(code)
}

func before() {
	if err := initialize.InitProject(); err != nil {
		panic(err)
	}
}

func after() {

}
