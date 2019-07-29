package main

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {

	cases := []struct{
		Name string
		Input interface{}
		ExpectCalls []string
	} {
		{
		"struct with one string field",
		struct {
			Name string
		}{"Carrie"},
		[]string {"Carrie"},
		},
	}


	for _, test:= range cases {
		t. Run(test.Name, func(t *testing.T) {
			var actual []string
			walk(test.Input, func(input string) {
				actual = append(actual, input)
			})

			if !reflect.DeepEqual(actual, test.ExpectCalls) {
				t.Errorf("expect %v, actual %v", test.ExpectCalls, actual)
			}
		})
	}


}
