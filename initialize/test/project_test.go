package test

import (
	"blog-backend/global"
	"blog-backend/initialize"
	"fmt"
	"testing"
)

// go test -v .\initialize\test\ -run=^TestInitProject -count=1
func TestInitProject(t *testing.T) {
	if err := initialize.InitProject(); err != nil {
		t.Error(err)
		t.Fail()
	}

	fmt.Printf("%#v\n", global.CONFIG.ServerConfig)
}
