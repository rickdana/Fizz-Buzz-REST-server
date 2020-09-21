package Service

import "github.com/rickdana/fizzbuzzApi/Model"

type FizzBuzzPrinter interface {
	Print(fizzBuzzDto Model.FizzBuzz) []string
}
