package Controller

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/rickdana/fizzbuzzApi/Logger"
	"github.com/rickdana/fizzbuzzApi/Model"
	"github.com/rickdana/fizzbuzzApi/Service"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

var l = log.New(os.Stderr, "", log.Ldate|log.Ltime|log.Lshortfile)
var logger = Logger.NewFZLogger(l, l, l)

type testStatisticsService struct{}

func (tss testStatisticsService) SaveEntry(dto Model.FizzBuzz) bool {
	return true
}

func (tss testStatisticsService) GetRequestStatistics() ([]Model.RequestStatistics, error) {
	return nil, nil
}

func TestFizzBuzzController_PrintFizzBuzzer(t *testing.T) {

	fzc := FizzBuzzController{fizzBuzzPrinter: Service.NewFizzBuzzService(logger), statistics: &testStatisticsService{}}
	tt := []struct {
		name             string
		fizzBuzzDto      interface{}
		httpCode         int
		expectedResponse string
	}{
		{"PrintFizzBuzzer should return the expected array of string", Model.FizzBuzz{
			FirstMultiple:  3,
			SecondMultiple: 5,
			Limit:          15,
			FizzWord:       "Fizz",
			BuzzWord:       "Buzz",
		}, http.StatusOK, `["1","2","Fizz","4","Buzz","Fizz","7","8","Fizz","Buzz","11","Fizz","13","14","FizzBuzz"]`},
		{"PrintFizzBuzzer should return a HTTP 400 when firstMultiple field is nil in the request body", Model.FizzBuzz{
			SecondMultiple: 5,
			Limit:          15,
			FizzWord:       "Fizz",
			BuzzWord:       "Buzz",
		}, http.StatusBadRequest, `{"firstMultiple":"non zero value required"}`},
		{"PrintFizzBuzzer should return a HTTP 400 when SecondMultiple field is nil in the request body", Model.FizzBuzz{
			FirstMultiple: 3,
			Limit:         15,
			FizzWord:      "Fizz",
			BuzzWord:      "Buzz",
		}, http.StatusBadRequest, `{"secondMultiple":"non zero value required"}`},
		{"PrintFizzBuzzer should return a HTTP 400 when limit field is nil in the request body", Model.FizzBuzz{
			FirstMultiple:  3,
			SecondMultiple: 5,
			FizzWord:       "Fizz",
			BuzzWord:       "Buzz",
		}, http.StatusBadRequest, `{"limit":"non zero value required"}`},
		{"PrintFizzBuzzer should return a HTTP 400 when BuzzWord field is nil in the request body", Model.FizzBuzz{
			FirstMultiple:  3,
			SecondMultiple: 5,
			Limit:          15,
			FizzWord:       "Fizz",
		}, http.StatusBadRequest, `{"buzzWord":"non zero value required"}`},
		{"PrintFizzBuzzer should return a HTTP 400 when FizzWord field is nil in the request body", Model.FizzBuzz{
			FirstMultiple:  3,
			SecondMultiple: 5,
			Limit:          15,
			BuzzWord:       "Buzz",
		}, http.StatusBadRequest, `{"fizzWord":"non zero value required"}`},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			reqBody, err := json.Marshal(tc.fizzBuzzDto)

			if err != nil {
				t.Fatal(err)
			}
			req, err := http.NewRequest(http.MethodPost, "/fizz-buzz", bytes.NewBuffer(reqBody))

			if err != nil {
				t.Fatal(err)
			}
			rr := httptest.NewRecorder()

			router := mux.NewRouter()
			router.HandleFunc("/fizz-buzz", fzc.PrintFizzBuzzer)
			router.ServeHTTP(rr, req)

			resBody := strings.TrimSuffix(rr.Body.String(), "\n")
			resStatus := rr.Code

			if resStatus != tc.httpCode || tc.expectedResponse != resBody {
				t.Errorf("handler should have failed on routeVariable /fizz-buzz: got Http_code: %v / output: %v want Http_code: %v / output: %v", resStatus, resBody, tc.httpCode, tc.expectedResponse)
			}
		})

	}
}
