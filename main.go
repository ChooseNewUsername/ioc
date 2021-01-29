package main

import (
	"ioc/Injector"
	"ioc/services"
)

func main()  {
	userService := services.NewUserService(services.NewOrderService())
	userService.GetOrderInfo()
	userService.GetUserInfo()

	Injector.BeanFactory.Set()
}
