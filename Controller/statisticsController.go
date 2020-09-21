package Controller

import (
	"github.com/rickdana/fizzbuzzApi/Service"
	"net/http"
)

type StatisticsController struct {
	statistics Service.StatisticsManager
}

func NewStatisticsController(statistics Service.StatisticsManager) *StatisticsController {
	return &StatisticsController{statistics: statistics}
}

func (sc StatisticsController) GetStatistics(w http.ResponseWriter, r *http.Request) {
	requestStatistics, err := sc.statistics.GetRequestStatistics()

	if err != nil {
		respond(w, http.StatusInternalServerError, "application/json", err.Error())
		return
	}

	respond(w, http.StatusOK, "application/json", requestStatistics)
}
