package Service

import (
	"github.com/rickdana/fizzbuzzApi/Logger"
	"github.com/rickdana/fizzbuzzApi/Model"
	"strconv"
)

type FizzBuzzService struct {
	logger *Logger.Logger
}

func NewFizzBuzzService(logger *Logger.Logger) *FizzBuzzService {
	return &FizzBuzzService{logger: logger}
}

func (fz *FizzBuzzService) Print(dto Model.FizzBuzz) []string {
	output := make([]string, dto.Limit)

	for i := uint(1); i <= dto.Limit; i++ {
		m := dto.FirstMultiple * dto.SecondMultiple
		if i%m == 0 {
			output[i-1] = dto.GetFizzBuzz()
		} else if i%dto.FirstMultiple == 0 {
			output[i-1] = dto.FizzWord
		} else if i%dto.SecondMultiple == 0 {
			output[i-1] = dto.BuzzWord
		} else {
			output[i-1] = strconv.Itoa(int(i))
		}
	}
	return output
}
