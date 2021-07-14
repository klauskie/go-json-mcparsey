package Optional

import (
	"fmt"
	"testing"
)

func TestArrayString(t *testing.T) {
	tests := []struct {
		name  string
		input struct {
			Origin       map[string]interface{}
			Key          string
			DefaultValue []string
		}
		output []string
	}{

		{
			name: "NonExistentKeys",
			input: struct {
				Origin       map[string]interface{}
				Key          string
				DefaultValue []string
			}{
				Origin: map[string]interface{}{
					"proxies": "1,2,3,TOR,luminati",
				}, Key: "NonExisting", DefaultValue: []string{}},
			output: []string{},
		},
		{
			name: "CaseString",
			input: struct {
				Origin       map[string]interface{}
				Key          string
				DefaultValue []string
			}{
				Origin: map[string]interface{}{
					"proxies": "1,2,3,TOR,luminati",
				}, Key: "proxies", DefaultValue: []string{}},
			output: []string{"1", "2", "3", "TOR", "luminati"},
		},
		{
			name: "UnsupportedType",
			input: struct {
				Origin       map[string]interface{}
				Key          string
				DefaultValue []string
			}{
				Origin: map[string]interface{}{
					"proxies": 3443,
				}, Key: "proxies", DefaultValue: []string{}},
			output: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ArrayString(tt.input.Origin, tt.input.Key, tt.input.DefaultValue)
			if fmt.Sprint(got) != fmt.Sprint(tt.output) {
				t.Errorf(
					"expected ArrayString(%v,%v,%v) = %v; got %v",
					tt.input.Origin,
					tt.input.Key,
					tt.input.DefaultValue,
					tt.output,
					got,
				)
			}
		})
	}

}
