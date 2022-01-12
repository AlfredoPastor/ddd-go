package domain

type ProductName struct {
	value string
}

func NewProductName(value string) ProductName {
	return ProductName{
		value: value,
	}
}
