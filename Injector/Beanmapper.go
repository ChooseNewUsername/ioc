package Injector

import "reflect"

type BeanMapper map[reflect.Type]reflect.Value

func (this BeanMapper) add(bean interface{})  {
	t:=reflect.TypeOf(bean)
	if t.Kind() != reflect.Ptr{
		panic("require ptr object ")
	}
	this[t] = reflect.ValueOf(bean)
}