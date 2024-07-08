package test

import (
	"blog-backend/database"
	"blog-backend/util"
	"os"
	"sync"
	"testing"
)

// go test -v .\database\test\ -run=^TestGetBlogDBConnection$ -count=1
func TestGetBlogDBConnection(t *testing.T) {
	const testCount = 100
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(testCount)
	for i := 0; i < testCount; i++ {
		go func() {
			defer waitGroup.Done()
			database.GetBlogDBConnection()
		}()
	}
	waitGroup.Wait()
}

func TestMain(m *testing.M) {
	before()
	code := m.Run()
	after()
	os.Exit(code)
}

func before() {
	util.InitLogger("log")
}

func after() {

}
