package criteria

import "fmt"

type Filters struct {
	List []Filter `json:"filters"`
}

func NewFilters() Filters {
	return Filters{
		List: []Filter{},
	}
}

func (f *Filters) Clean() {
	f.List = f.List[:0]
}

func (f *Filters) Size() int {
	return len(f.List)
}

func (f *Filters) HasFilters() bool {
	return len(f.List) > 0
}

func (f *Filters) Serialize() string {
	result := ""
	for i, v := range f.List {
		result = result + fmt.Sprintf("  Filter #%d: %s %s %s\n", i, v.Field, v.Operator, v.Value)
	}
	return result
}
