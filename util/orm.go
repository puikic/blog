package util

import (
	"reflect"
	"strings"
)

func GetGormFields(stc any) []string {
	typ := reflect.TypeOf(stc)
	columns := make([]string, 0, typ.NumField())
	if typ.Kind() == reflect.Ptr { //若传的stc是指针类型，先解析指针
		typ = typ.Elem()
	}
	if typ.Kind() == reflect.Struct {
		count := typ.NumField()
		for i := 0; i < count; i++ {
			typField := typ.Field(i)
			if typField.IsExported() { //只关注可导出成员
				if typField.Tag.Get("gorm") == "-" { //不做ORM映射的字段跳过
					continue
				}
				name := Camel2Snake(typField.Name) //若无gorm Tag,按默认情况处理，即name的蛇形形式对应数据库该字段
				if len(typField.Tag.Get("gorm")) > 0 {
					content := typField.Tag.Get("gorm")
					if strings.HasPrefix(content, "column:") {
						content = content[7:]
						pos := strings.Index(content, ";")
						if pos > 0 {
							content = content[:pos]
							name = content
						} else {
							name = content
						}
					}
				}
				columns = append(columns, name)
			}
		}
		return columns
	} else { //如果stc不是结构体则返回空切片
		return nil
	}
}
