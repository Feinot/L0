package cashe

import (
	"nats/models"
)

var cashe map[string]models.Order

func Init() {
	cashe = make(map[string]models.Order)
}

func Set(order models.Order) {
	Init()
	cashe[order.OrderUID] = order
}

func GetOrder(id string) (models.Order, bool) {
	order, ok := cashe[id]

	return order, ok
}
