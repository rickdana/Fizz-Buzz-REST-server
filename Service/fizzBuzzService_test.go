package Service

import (
	"github.com/rickdana/fizzbuzzApi/Model"
	"reflect"
	"testing"
)

func TestFizzBuzzService_Print(t *testing.T) {

	tests := []struct {
		name string
		arg  Model.FizzBuzz
		want []string
	}{
		{"Print should return empty string when limit is less than 1", Model.FizzBuzz{
			FirstMultiple:  3,
			SecondMultiple: 2,
			Limit:          0,
			FizzWord:       "Fizz",
			BuzzWord:       "Buzz",
		}, []string{}},
		{"Print should return empty string when limit is 1", Model.FizzBuzz{
			FirstMultiple:  3,
			SecondMultiple: 2,
			Limit:          1,
			FizzWord:       "Fizz",
			BuzzWord:       "Buzz",
		}, []string{"1"}},
		{"Print should contain only FizzWord", Model.FizzBuzz{
			FirstMultiple:  3,
			SecondMultiple: 5,
			Limit:          4,
			FizzWord:       "Fizz",
			BuzzWord:       "Buzz",
		}, []string{"1", "2", "Fizz", "4"}},
		{"Print should contain only FizzWord and BuzzWord", Model.FizzBuzz{
			FirstMultiple:  3,
			SecondMultiple: 5,
			Limit:          6,
			FizzWord:       "Fizz",
			BuzzWord:       "Buzz",
		}, []string{"1", "2", "Fizz", "4", "Buzz", "Fizz"}},
		{"Print should contain only FizzBuzzWord", Model.FizzBuzz{
			FirstMultiple:  5,
			SecondMultiple: 2,
			Limit:          15,
			FizzWord:       "Fizz",
			BuzzWord:       "Buzz",
		}, []string{"1", "Buzz", "3", "Buzz", "Fizz", "Buzz", "7", "Buzz", "9", "FizzBuzz", "11", "Buzz", "13", "Buzz", "Fizz"}},
		{"Print should match the expected Output", Model.FizzBuzz{
			FirstMultiple:  3,
			SecondMultiple: 5,
			Limit:          15,
			FizzWord:       "Hello",
			BuzzWord:       "World",
		}, []string{"1", "2", "Hello", "4", "World", "Hello", "7", "8", "Hello", "World", "11", "Hello", "13", "14", "HelloWorld"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fz := &FizzBuzzService{}
			if got := fz.Print(tt.arg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Print() = %v, want %v", got, tt.want)
			}
		})
	}
}
