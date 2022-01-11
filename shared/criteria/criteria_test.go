package criteria

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_UnmarshalCriteria(t *testing.T) {
	tests := []struct {
		name     string
		peticion string
		want     int
	}{
		{
			name: "Prueba de creacion de criteria",
			peticion: `{
				"filters": [
					{
						"field": "municipality",
						"operator": "=",
						"value": "Bucaramanga"
					},
					{
						"field": "neighborhood",
						"operator": "=",
						"value": "San Francisco"
					}
				],
				"orders": [
					{
						"order_by": "municipality",
						"order_type": "desc"
					},
					{
						"order_by": "neighborhood",
						"order_type": "asc"
					}
				],
				"limit": 10,
				"offset": 1
			}`,
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewCriteria()
			err := json.Unmarshal([]byte(tt.peticion), &got)
			if err != nil {
				t.Errorf("Hubo un error en el Unmarshal: %s", err.Error())
			}
			if got.Filters.Size() != tt.want {
				t.Errorf("Tamaño filtros = %v, want %v", got.Filters.Size(), tt.want)
			}
			assert.Equal(t, got.Orders.List[0].OrderBy, "municipality", "Son diferentes")
			assert.Equal(t, got.Orders.List[0].OrderType.String(), "desc", "Son diferentes")
			assert.Equal(t, got.Limit, 10, "Son diferentes")
			assert.Equal(t, got.Offset, 1, "Son diferentes")
			assert.Equal(t, got.Filters.List[0].Field, "municipality", "Son diferentes")
			assert.Equal(t, got.Filters.List[0].Operator.String(), "=", "Son diferentesddd")
			assert.Equal(t, got.Filters.List[0].Value, "Bucaramanga", "Son diferentes")
		})
	}
}
