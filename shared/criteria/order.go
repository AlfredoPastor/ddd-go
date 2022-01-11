package criteria

import (
	"errors"
	"fmt"
	"strings"
)

const (
	NONE kind = "none"
	ASC  kind = "asc"
	DESC kind = "desc"
)

type kind string

type Order struct {
	OrderBy   string `json:"order_by"`
	OrderType kind   `json:"order_type"`
}

func (o *Order) IsNone() bool {
	return o.OrderType == NONE || o.OrderType.String() == ""
}

func (o *Order) IsAsc() bool {
	return o.OrderType == ASC
}

func (o *Order) HasOrder() bool {
	return !o.IsNone()
}

func (k *kind) UnmarshalText(data []byte) error {
	switch strings.Trim(string(data), `"`) {
	case string(NONE):
		*k = NONE
	case string(ASC):
		*k = ASC
	case string(DESC):
		*k = DESC
	default:
		return errors.New(fmt.Sprintf("El tipo de ordenamiento '%s' no existe", string(data)))
	}
	return nil
}

func (k *kind) String() string {
	return string(*k)
}
