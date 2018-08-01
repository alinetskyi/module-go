package orderbook

import (
	"sort"
)

type Orderbook struct {
	Ask      []*Order
	Bid      []*Order
	Trades   []*Trade
	Rejected *Order
}

func New() *Orderbook {
	s := make([]*Order, 0)
	b := make([]*Order, 0)
	t := make([]*Trade, 0)
	return &Orderbook{s, b, t, nil}
}

type DescOrderArray []*Order

func (orderbook *Orderbook) Match(order *Order) ([]*Trade, *Order) {
	orderbook.Trades = nil
	orderbook.Rejected = nil
	OrderExists(order)
	switch order.Side.String() {
	case "BID":
		PlaceBid(order, orderbook)
	case "ASK":
		PlaceAsk(order, orderbook)
	}
	return orderbook.Trades, orderbook.Rejected
}

func PlaceBid(order *Order, orderbook *Orderbook) {
	var ar []*Order = make([]*Order, 0)
	if order.Price > 0 {
		orderbook.Bid = append(orderbook.Bid, order)
	}
	SortBid(orderbook)
	for i := 0; i < len(orderbook.Ask); i++ {
		r := orderbook.Ask[i]
		if r.Price <= order.Price && OrderExists(r) {
			StartTrade(i, &ar, order, r, orderbook.Bid, orderbook)
		}
		if order.Price == 0 && order.Volume > 0 {
			StartTrade(i, &ar, order, r, orderbook.Bid, orderbook)
		}
	}
	orderbook.Ask = ar
	if order.Kind == 1 && order.Volume > 0 {
		orderbook.Rejected = order
	}
}

func PlaceAsk(order *Order, orderbook *Orderbook) {
	var ar []*Order = make([]*Order, 0)
	if order.Price > 0 {
		orderbook.Ask = append(orderbook.Ask, order)
	}
	SortBid(orderbook)
	for i := 0; i < len(orderbook.Bid); i++ {
		r := orderbook.Bid[i]
		if r.Price >= order.Price && OrderExists(r) && order.Price > 0 {
			StartTrade(i, &ar, order, r, orderbook.Ask, orderbook)
		}
		if order.Price == 0 && order.Volume > 0 {
			StartTrade(i, &ar, order, r, orderbook.Ask, orderbook)
		}
	}
	orderbook.Bid = ar
	if order.Kind == 1 && order.Volume > 0 {
		orderbook.Rejected = order
	}
}

func StartTrade(i int, ar *[]*Order, curOrder *Order, tradeOrder *Order, orders []*Order, orderbook *Orderbook) {
	if tradeOrder.Volume < curOrder.Volume {
		orderbook.Trades = MakeTrade(false, orderbook.Trades, curOrder, tradeOrder)
		curOrder.Volume -= tradeOrder.Volume
	} else if tradeOrder.Volume > curOrder.Volume {
		orderbook.Trades = MakeTrade(true, orderbook.Trades, curOrder, tradeOrder)
		tradeOrder.Volume -= curOrder.Volume
		curOrder.Volume = 0
		//orders = orders[:len(orders)-1]
		*ar = append(*ar, tradeOrder)
	} else {
		orderbook.Trades = MakeTrade(false, orderbook.Trades, curOrder, tradeOrder)
		curOrder.Volume -= curOrder.Volume
		//orders = orders[:len(orders)-1]
	}
}

func MakeTrade(isCurOrder bool, curTrades []*Trade, curOrder *Order, tradeOrder *Order) []*Trade {
	if !isCurOrder {
		curTrades = append(curTrades, &Trade{curOrder, tradeOrder, tradeOrder.Volume, tradeOrder.Price})
	} else {
		curTrades = append(curTrades, &Trade{curOrder, tradeOrder, curOrder.Volume, tradeOrder.Price})
	}
	return curTrades
}

func OrderExists(order *Order) bool {
	if order == nil {
		panic("[ERROR]: Your order doesn't exist! ")
		return false
	} else {
		return true
	}
}

func SortBid(orderbook *Orderbook) {
	sort.Sort(sort.Reverse(DescOrderArray(orderbook.Bid)))
	sort.Sort(DescOrderArray(orderbook.Ask))
}

func (o DescOrderArray) Len() int {
	return len(o)
}
func (o DescOrderArray) Swap(i, j int) {
	o[i], o[j] = o[j], o[i]
}
func (o DescOrderArray) Less(i, j int) bool {
	var less bool
	if o[i] != nil && o[j] != nil && o[i].Price < o[j].Price {
		less = true
	}
	return less
}
