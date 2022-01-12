package domain

type ProductQuantity struct {
	value int
}

func NewProductQuantity(value int) ProductQuantity {
	return ProductQuantity{
		value: value,
	}
}

func (p *ProductQuantity) Primitive() int {
	return p.value
}
