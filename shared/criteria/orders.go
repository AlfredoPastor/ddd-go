package criteria

import (
	"fmt"
)

type Orders struct {
	List []Order `json:"orders"`
}

func NewOrders() Orders {
	return Orders{
		List: []Order{},
	}
}

func (o *Orders) Clean() {
	o.List = o.List[:0]
}

func (o *Orders) Size() int {
	return len(o.List)
}

func (o *Orders) HasOrders() bool {
	return len(o.List) > 0
}

func (o *Orders) Serialize() string {
	result := ""
	for i, v := range o.List {
		result = result + fmt.Sprintf("  Order #%d: %s (%s) \n", i, v.OrderBy, v.OrderType)
	}
	return result
}
