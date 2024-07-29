package util

// import (
// 	"reflect"
// )

// func GetGormFields(stc any) []string {
// 	typ := reflect.TypeOf(stc)
// 	if typ.Kind() == reflect.Ptr { //若传的stc是指针类型，先解析指针
// 		typ = typ.Elem()
// 	}
// 	if typ.Kind() == reflect.Struct {
// 		count := typ.NumField()
// 		columns := make([]string, 0, count)
// 		for i := 0; i < count; i++ {
// 			typField := typ.Field(i)
// 			if typField.IsExported() { //只关注可导出成员
// 				if typField.Tag.Get("gorm") == "-" { //不做ORM映射的字段跳过
// 					continue
// 				}
// 				name := Camel2Snake(typField.Name) //若无gorm Tag,按默认情况处理，即name的蛇形形式对应数据库该字段
// 			}
// 		}
// 	}
// }



