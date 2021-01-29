package services

import "fmt"

type OrderService struct {
	Version string
}

func NewOrderService() *OrderService {
	return &OrderService{Version:"1.0"}
}
func  (this *OrderService) GetOrderInfo(){
	fmt.Println("获取订单信息")
}