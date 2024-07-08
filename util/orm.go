package util

import (
	"reflect"
	"strings"
)

func GetGormFields(stc any) []string {
	typ := reflect.TypeOf(stc)
	// 如果入参是指针则先解析指针
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	// 如果 stc 不是结构体则返回空切片
	if typ.Kind() != reflect.Struct {
		return nil
	}

	columns := make([]string, 0, typ.NumField())
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)

		// 不是可导出成员或不做ORM映射的字段则不处理
		if !field.IsExported() || field.Tag.Get("gorm") == "-" {
			continue
		}

		name := CamelToSnake(field.Name)
		if len(field.Tag.Get("gorm")) > 0 {
			content := field.Tag.Get("gorm")
			if strings.HasPrefix(content, "column:") {
				content = content[7:]
				pos := strings.Index(content, ";")
				if pos > 0 {
					name = content[0:pos]
				} else if pos < 0 {
					name = content
				}
			}
		}
		columns = append(columns, name)
	}

	return columns
}
