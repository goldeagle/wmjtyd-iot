package utils

import (
	"errors"
	"reflect"
)

/**
 * 拷贝不同结构体的属性值，通常用在 DTO 与 Model之间
 * 从 from 拷贝到 to
 */
func StructCopy(from, to interface{}) error {

	fromValue := reflect.ValueOf(from)
	toValue := reflect.ValueOf(to)

	// 必须是指针类型
	if fromValue.Kind() != reflect.Ptr || toValue.Kind() != reflect.Ptr {
		return errors.New("必须指针")
	}

	if fromValue.IsNil() || toValue.IsNil() {
		return errors.New("空值")
	}

	// 获取到来源数据
	fromElem := fromValue.Elem()
	// 需要的数据
	toElem := toValue.Elem()

	for i := 0; i < toElem.NumField(); i++ {

		toField := toElem.Type().Field(i)
		// 看看来源的结构体中是否有这个属性
		fromFieldName, ok := fromElem.Type().FieldByName(toField.Name)

		// 存在相同的属性名称并且类型一致
		// todo 可以根据需要判断是否是空值
		if ok && fromFieldName.Type == toField.Type {
			toElem.Field(i).Set(fromElem.FieldByName(toField.Name))
		}
	}

	return nil
}
