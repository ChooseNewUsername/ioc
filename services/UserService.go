package services

import "fmt"

type UserService struct {
	order *OrderService
}

func NewUserService(order *OrderService) *UserService {
	return &UserService{order: order}
}

func  (this *UserService) GetUserInfo(){
	fmt.Println("获取用户信息")
}
func  (this *UserService) GetOrderInfo(){

	this.order.GetOrderInfo()
}