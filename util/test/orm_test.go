package test

import (
	"blog-backend/util"
	"fmt"
	"os"
	"testing"
	"time"
)

type User struct {
	Id         int    `gorm:"column:id;primary_key"`
	Name       string `gorm:"column:name"`
	Nickname   string `gorm:"column:nickname"`
	CreateTime time.Time
	Password   string `gorm:"column:password"`
	int
	gender int
}

func TestGetGormFields(t *testing.T) {
	var p *User
	columns := util.GetGormFields(p)
	fmt.Println(columns)
}

func TestMain(m *testing.M) {
	before()
	code := m.Run()
	after()
	os.Exit(code)
}

func before() {

}

func after() {

}
