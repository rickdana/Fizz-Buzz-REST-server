package Model

import (
	"encoding/json"
	"strings"
)

type FizzBuzz struct {
	FirstMultiple  uint   `json:"firstMultiple" valid:"required"`
	SecondMultiple uint   `json:"secondMultiple" valid:"required"`
	Limit          uint   `json:"limit" valid:"required"`
	FizzWord       string `json:"fizzWord" valid:"required"`
	BuzzWord       string `json:"buzzWord" valid:"required"`
}

func (fz *FizzBuzz) GetFizzBuzz() string {
	return fz.FizzWord + fz.BuzzWord
}

func FizzBuzzDtoFactory(stringArray string) ([]FizzBuzz, error) {
	split := strings.Split(stringArray, "\n")
	if len(split) > 0 {
		split = split[:len(split)-1]
	}
	fizzBuzzDtoList := make([]FizzBuzz, 0, len(split))

	for _, fz := range split {
		fizzBuzz := FizzBuzz{}
		err := json.Unmarshal([]byte(fz), &fizzBuzz)
		if err != nil {
			return nil, err
		}
		fizzBuzzDtoList = append(fizzBuzzDtoList, fizzBuzz)
	}
	return fizzBuzzDtoList, nil
}
