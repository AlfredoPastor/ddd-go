package criteria

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

type Criteria struct {
	Filters
	Orders
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

func NewCriterias(raw string) (Criteria, error) {
	raw = strings.ReplaceAll(raw, "%23", "#")
	criteria := NewCriteria()
	if len(raw) == 0 {
		return criteria, nil
	}
	rawIn := json.RawMessage(raw)
	rawBytes, err := rawIn.MarshalJSON()
	if err != nil {
		return Criteria{}, err
	}
	err = json.Unmarshal(rawBytes, &criteria)
	if err != nil {
		return Criteria{}, err
	}
	return criteria, nil
}

func NewCriteria() Criteria {
	return Criteria{
		Filters: NewFilters(),
		Orders:  NewOrders(),
		Limit:   10,
		Offset:  0,
	}
}

func (c *Criteria) ValidateFields() error {
	for _, v := range c.Filters.List {
		if v.Field == "" {
			return errors.New("el campo 'field' es requerido")
		}
		if v.Operator.String() == "" {
			return errors.New("el campo 'operador' es requerido")
		}
		if v.Value == "" {
			return errors.New("el campo 'value' es requerido")
		}
	}
	for _, v := range c.Orders.List {

		if v.OrderBy == "" {
			return errors.New("el campo 'order_by' es requerido")
		}
		if v.OrderType.String() == "" {
			return errors.New("el campo 'order_type' es requerido")
		}
	}
	return nil
}

func (c *Criteria) Serialize() string {
	return fmt.Sprintf("Filters\n%sOrders\n%sLimit %d\nOffset %d", c.Filters.Serialize(), c.Orders.Serialize(), c.Limit, c.Offset)
}

// field, operator, value and group
func (c *Criteria) AddFilter(field, op, value, gr string) {
	filter := Filter{
		Field:    field,
		Operator: operator(op),
		Value:    value,
		Group:    group(gr),
	}
	c.Filters.List = append(c.Filters.List, filter)
}

func (c *Criteria) AddLimit(limit int) {
	c.Limit = limit
}

func (c *Criteria) AddOffset(offset int) {
	c.Offset = offset
}

func (c *Criteria) AddOrder(order_by string, order_type kind) {
	order := Order{
		OrderBy:   order_by,
		OrderType: order_type,
	}

	c.Orders.List = append(c.Orders.List, order)
}
