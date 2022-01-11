package criteria

import (
	"errors"
	"fmt"
	"strings"
)

type operator string
type group string

const (
	EQUAL        operator = "="
	NOT_EQUAL    operator = "!="
	GT           operator = ">"
	GTE          operator = ">="
	LT           operator = "<"
	LTE          operator = "<="
	CONTAINS     operator = "CONTAINS"
	NOT_CONTAINS operator = "NOT_CONTAINS"
	UNKNOWN      operator = ""
	AND          group    = "AND"
	OR           group    = "OR"
	MATCH        operator = "MATCH"
)

type Filter struct {
	Field    string   `json:"field"`
	Operator operator `json:"operator"`
	Value    string   `json:"value"`
	Group    group    `json:"group"`
}

func (o *operator) UnmarshalText(data []byte) error {
	switch strings.Trim(string(data), `"`) {
	case string(EQUAL):
		*o = EQUAL
	case string(NOT_EQUAL):
		*o = NOT_EQUAL
	case string(GT):
		*o = GT
	case string(GTE):
		*o = GTE
	case string(LT):
		*o = LT
	case string(LTE):
		*o = LTE
	case string(CONTAINS):
		*o = CONTAINS
	case string(NOT_CONTAINS):
		*o = NOT_CONTAINS
	case string(MATCH):
		*o = MATCH
	case string(UNKNOWN):
		return errors.New("El operador se encuentra vacio")
	default:
		return errors.New(fmt.Sprintf("El operador '%s' no existe", string(data)))
	}
	return nil
}

func (g *group) UnmarshalText(data []byte) error {
	switch strings.Trim(string(data), `"`) {
	case string(AND):
		*g = AND
	case string(OR):
		*g = OR
	case string(UNKNOWN):
		*g = AND
	default:
		return errors.New(fmt.Sprintf("El grupo '%s' no existe", string(data)))
	}
	return nil
}

func (o *operator) String() string {
	return string(*o)
}

func (g *group) String() string {
	return string(*g)
}
