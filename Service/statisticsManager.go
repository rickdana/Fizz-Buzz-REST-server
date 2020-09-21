package Service

import "github.com/rickdana/fizzbuzzApi/Model"

type StatisticsManager interface {
	SaveEntry(dto Model.FizzBuzz) bool
	GetRequestStatistics() ([]Model.RequestStatistics, error)
}
