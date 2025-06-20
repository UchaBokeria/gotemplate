package enums

type OrderStatusType struct {
	New       uint
	Pending   uint
	Inprocess uint
	Issued    uint
	Canceled  uint
	Done      uint
}

var OrderStatus OrderStatusType = OrderStatusType{
	New:       uint(1),
	Pending:   uint(2),
	Inprocess: uint(3),
	Issued:    uint(4),
	Canceled:  uint(5),
	Done:      uint(6),
}
