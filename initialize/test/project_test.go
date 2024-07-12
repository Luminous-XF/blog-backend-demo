package test

import (
	"blog-backend/initialize"
	"testing"
)

// go test -v .\initialize\test\ -run=^TestInitProject -count=1
func TestInitProject(t *testing.T) {
	if err := initialize.InitProject(); err != nil {
		t.Error(err)
		t.Fail()
	}
}
