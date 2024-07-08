package util

// func GetGormFields(stc any) []string {
// 	typ := reflect.TypeOf(stc)
// 	// 如果入参是指针则先解析指针
// 	if typ.Kind() == reflect.Ptr {
// 		typ = typ.Elem()
// 	}
//
// 	if typ.Kind() == reflect.Struct {
// 		columns := make([]string, 0, typ.NumField())
// 		for i := 0; i < typ.NumField(); i++ {
// 			field := typ.Field(i)
//
// 			// 不是可导出成员或不做ORM映射的字段则不处理
// 			if !field.IsExported() || field.Tag.Get("gorm") == "-" {
// 				continue
// 			}
//
//             name :=
// 		}
// 	}
// }
