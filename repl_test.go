package main

import (
	"reflect"
	"testing"
)

func TestCleanInput(t *testing.T) {
	tests := map[string]struct {
		got    string
		expect []string
	}{
		"Start & End Whitespace": {" hello world ", []string{"hello", "world"}},
		"double Whitespace":      {"hello  world", []string{"hello", "world"}},
		"Upper to Lower":         {"Hello World", []string{"hello", "world"}},
		"Single Word":            {"Hello ", []string{"hello"}},
		"Empty":                  {" ", []string{}},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			response := cleanInput(tc.got)
			if !reflect.DeepEqual(response, tc.expect) {
				t.Errorf("cleanInput() = %v, expect %v", response, tc.expect)
			}
		})
	}
}
