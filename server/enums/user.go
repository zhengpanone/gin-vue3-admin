package enums

type OrderState int 

const (
	// 待确认
	OrderConfirmed OrderState = iota + 1
	// 待支付
	OrderPaying
	// 待发货
	OrderShipping
	// 待收货
	OrderReceiving
	// 已收货
	OrderReceived
	// 已完成
	OrderCompleted
)

var (
	OrderStateItems = map[OrderState] string {
		OrderConfirmed:"待确认",
		OrderPaying:    "待支付",
		OrderShipping:  "待发货",
		OrderReceiving: "待收货",
		OrderReceived:  "已收货",
		OrderCompleted: "已完成",
	}
)
// 枚举类
type OrderStateEnum struct{
	v OrderState
}

// Name 获取订单状态名称
func (o OrderStateEnum) Name() string{
	return OrderStateItems[o.v]
}
// 检查枚举值是否合法
func (o OrderStateEnum) IsValid() bool{
	_,ok:=OrderStateItems[o.v]
	return ok
}

