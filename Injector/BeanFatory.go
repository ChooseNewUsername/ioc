package Injector

import "reflect"

var BeanFactory *BeanFactoryImpl

func init()  {
	BeanFactory = NewBeanFactory()
}

type BeanFactoryImpl struct {
	beanMapper BeanMapper
}

func (this *BeanFactoryImpl)Set(vlist ...interface{})  {
	if vlist==nil|| len(vlist)==0{
		return
	}
	for _,v:=range vlist{
		this.beanMapper.add(v)
	}
}
func (this *BeanFactoryImpl)Get(bean interface{})  {
	t:=reflect.TypeOf(bean)
	if bean==nil{
		return
	}
	for _,v:=range vlist{
		this.beanMapper.add(v)
	}
}



func NewBeanFactory() *BeanFactoryImpl {
	return &BeanFactoryImpl{beanMapper: make(BeanMapper)}
}