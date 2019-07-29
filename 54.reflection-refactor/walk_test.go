package main

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {

	t.Run("with sortable data", func(t *testing.T) {
		cases := []struct{
			Name string
			Input interface{}
			ExpectCalls []string
		} {
			{
				"struct with two string fields",
				struct {
					Name string
				}{"Carrie"},
				[]string {"Carrie"},
			},
			{
				"struct with one string field",
				struct {
					Name string
					City string
				}{"Carrie", "Paris"},
				[]string {"Carrie", "Paris"},
			},
			{
				"struct with not string field",
				struct {
					Name string
					Age int
				}{"Carrie", 18},
				[]string {"Carrie"},
			},
			{
				"Nested fields",
				Person{
					"Carrie",
					Profile{18, "Paris"},
				},
				[]string{"Carrie", "Paris"},
			},
			{
				"Pointer to things",
				&Person{
					"Carrie",
					Profile{18, "Paris"},
				},
				[]string {"Carrie", "Paris"},
			},
			{
				"Slices",
				[]Profile {
					{18, "Paris"},
					{19, "London"},
				},
				[]string {"Paris","London"},
			},
			{
				"Array",
				[2]Profile {
					{18, "Paris"},
					{19, "London"},
				},
				[]string {"Paris","London"},
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
	})


	t.Run("with maps", func(t *testing.T) {
		// // 记住，Go 中的 map 不能保证顺序一致。因此，你的测试有时会失败，因为我们断言对 fn 的调用是以特定的顺序完成的。
		aMap := map[string]string {
			"Foo": "Bar",
			"Baz": "Boz",
		}

		var actual []string
		walk(aMap, func(input string) {
			actual = append(actual, input)
		})

		assert.Contains(t, actual, "Bar")
		assert.Contains(t, actual, "Boz")
	})


}

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}
