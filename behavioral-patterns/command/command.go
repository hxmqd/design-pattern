// 命令模式是一种数据驱动的设计模式，它属于行为型模式。 请求以命令的形式包裹在对象中，并传给调用对象。
// 调用对象寻找可以处理该命令的合适的对象，并把该命令传给相应的对象，该对象执行命令。
// 意图：将一个请求封装成一个对象，从而使您可以用不同的请求对客户进行参数化。 主要解决：在软件系统中，
// 行为请求者与行为实现者通常是一种紧耦合的关系，但某些场合，比如需要对行为进行记录、撤销或重做、事务等处理时，
// 这种无法抵御变化的紧耦合的设计就不太合适

package command

import "fmt"

type Order interface {
	Execute()
}

type Stock struct {
	Name string
	Quantity int
}

func (s *Stock)Buy(){
	fmt.Printf("stock [Name:%s, Quantity:%d] bought\n", s.Name, s.Quantity)
}

func (s *Stock) Sell(){
	fmt.Printf("Stock [Name:%s, Quantity:%d] sold\n", s.Name, s.Quantity)
}

type SellStock struct {
	stock *Stock
}

func  NewSellStock(stock *Stock) *SellStock{
	return &SellStock{
		stock:stock,
	}
}

func (b *SellStock) Execute(){
	b.stock.Sell()
}

type BuyStock struct {
	stock *Stock
}

func  NewBuyStock(stock *Stock) *BuyStock{
	return &BuyStock{
		stock:stock,
	}
}

func (b *BuyStock) Execute(){
	b.stock.Buy()
}

type Broker struct {
	orderList []Order
}

func (b *Broker) TakeOrder(order Order){
	b.orderList = append(b.orderList, order)
}

func (b *Broker) placeOrders(){
	for _, order :=range b.orderList {
		order.Execute()
	}
	b.orderList = make([]Order, 0 )
}
