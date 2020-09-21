package Service

import (
	"github.com/rickdana/fizzbuzzApi/Logger"
	"github.com/rickdana/fizzbuzzApi/Model"
	"github.com/rickdana/fizzbuzzApi/Repository"
)

type StatisticService struct {
	logger *Logger.Logger
	em     Repository.EntityManager
}

func NewStatisticService(logger *Logger.Logger, em Repository.EntityManager) *StatisticService {
	return &StatisticService{logger: logger, em: em}
}

func (sts *StatisticService) SaveEntry(dto Model.FizzBuzz) bool {
	err := sts.em.Save(dto)
	if err != nil {
		sts.logger.Warning("unable to save request data", err)
		return false
	}
	return true
}

func (sts *StatisticService) GetRequestStatistics() ([]Model.RequestStatistics, error) {
	fizzBuzzDtos, err := sts.em.FindAll()
	if err != nil {
		sts.logger.Error("something went wrong", err)
		return nil, err
	}

	statsMap := make(map[Model.FizzBuzz]int)

	for _, fz := range fizzBuzzDtos {
		if _, ok := statsMap[fz]; ok {
			statsMap[fz] = statsMap[fz] + 1
		} else {
			statsMap[fz] = 1
		}
	}

	stats := make([]Model.RequestStatistics, 0, len(fizzBuzzDtos))

	for k, v := range statsMap {
		s := Model.RequestStatistics{
			FizzBuzzDto: k,
			Hits:        v,
		}
		stats = append(stats, s)
	}
	sort(stats)

	return stats, nil
}

func sort(requestStatistics []Model.RequestStatistics) {
	l := len(requestStatistics)
	for i := 0; i < l-1; i++ {
		minIdx := i
		for j := i + 1; j < l; j++ {
			if requestStatistics[j].Hits > requestStatistics[minIdx].Hits {
				minIdx = j
			}
		}
		temp := requestStatistics[minIdx]
		requestStatistics[minIdx] = requestStatistics[i]
		requestStatistics[i] = temp
	}
}
