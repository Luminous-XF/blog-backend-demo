package test

import (
	"os"
	"testing"
)

func beforeTest() {
	// initialize.InitProject()
	// initialize.InitViper()
}

func afterTest() {

}

func TestMain(m *testing.M) {
	beforeTest()
	code := m.Run()
	afterTest()
	os.Exit(code)
}
