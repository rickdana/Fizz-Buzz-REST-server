package Controller

import (
	"encoding/json"
	"github.com/asaskevich/govalidator"
	"github.com/rickdana/fizzbuzzApi/Model"
	"github.com/rickdana/fizzbuzzApi/Service"
	"net/http"
)

type FizzBuzzController struct {
	fizzBuzzPrinter Service.FizzBuzzPrinter
	statistics      Service.StatisticsManager
}

func NewFizzBuzzController(fizzBuzzPrinter Service.FizzBuzzPrinter, statistics Service.StatisticsManager) *FizzBuzzController {
	return &FizzBuzzController{fizzBuzzPrinter: fizzBuzzPrinter, statistics: statistics}
}

func (fzc *FizzBuzzController) PrintFizzBuzzer(w http.ResponseWriter, r *http.Request) {
	var fizzBuzzDto Model.FizzBuzz

	err := json.NewDecoder(r.Body).Decode(&fizzBuzzDto)

	if err != nil {
		respond(w, http.StatusBadRequest, "application/json", err.Error())
		return
	}

	if _, errs := govalidator.ValidateStruct(fizzBuzzDto); errs != nil {
		respond(w, http.StatusBadRequest, "application/json", govalidator.ErrorsByField(errs))
		return
	}

	fzc.statistics.SaveEntry(fizzBuzzDto)

	fizzBuzzOutput := fzc.fizzBuzzPrinter.Print(fizzBuzzDto)

	respond(w, http.StatusOK, "text/plain", fizzBuzzOutput)
}
