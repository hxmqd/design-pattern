package command

import "testing"

func TestCommand(t *testing.T) {

	 stock := &Stock{
	 	Name: "abc",
	 	Quantity:10,
	 }

	 buyStockOrder := NewBuyStock(stock)
	 sellStockOrder := NewSellStock(stock)

	 broker := new(Broker)

	 broker.TakeOrder(buyStockOrder)
	 broker.TakeOrder(sellStockOrder)

	 broker.placeOrders()

}
