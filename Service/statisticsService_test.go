package Service

import (
	"errors"
	"github.com/rickdana/fizzbuzzApi/Logger"
	"github.com/rickdana/fizzbuzzApi/Model"
	"github.com/rickdana/fizzbuzzApi/Repository"
	"log"
	"os"
	"reflect"
	"testing"
)

var requests1 = Model.FizzBuzz{
	FirstMultiple:  3,
	SecondMultiple: 5,
	Limit:          90,
	FizzWord:       "Fizz",
	BuzzWord:       "Buzz",
}

var requests2 = Model.FizzBuzz{
	FirstMultiple:  7,
	SecondMultiple: 9,
	Limit:          100,
	FizzWord:       "Hello",
	BuzzWord:       "World",
}

var expectedStats = []Model.RequestStatistics{
	{requests1, 2},
	{requests2, 1},
}

type testRepository struct {
}

type errRepository struct {
}

func (tr testRepository) Save(entry interface{}) (err error) {
	return nil
}

func (tr testRepository) FindAll() ([]Model.FizzBuzz, error) {
	return []Model.FizzBuzz{requests1, requests1, requests2}, nil
}

func (er errRepository) Save(entry interface{}) (err error) {
	return errors.New("oups something went wrong")
}

func (er errRepository) FindAll() ([]Model.FizzBuzz, error) {
	return nil, errors.New("oups something went wrong")
}

var l = log.New(os.Stderr, "", log.Ldate|log.Ltime|log.Lshortfile)
var logger = Logger.NewFZLogger(l, l, l)

func TestStatisticService_GetRequestStatistics(t *testing.T) {
	type fields struct {
		logger *Logger.Logger
		em     Repository.EntityManager
	}

	tests := []struct {
		name    string
		fields  fields
		want    []Model.RequestStatistics
		wantErr bool
	}{
		{"Should return RequestStatistics", fields{logger: logger, em: testRepository{}}, expectedStats, false},
		{"Should return error when fail to fetch request history", fields{logger: logger, em: errRepository{}}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sts := &StatisticService{
				logger: tt.fields.logger,
				em:     tt.fields.em,
			}
			got, err := sts.GetRequestStatistics()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRequestStatistics() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRequestStatistics() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStatisticService_SaveEntry(t *testing.T) {
	type fields struct {
		logger *Logger.Logger
		em     Repository.EntityManager
	}
	type args struct {
		dto Model.FizzBuzz
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{"Should save fizzBuzz request body", fields{logger: logger, em: testRepository{}}, args{dto: requests1}, true},
		{"Should fail to save fizzBuzz request body", fields{logger: logger, em: errRepository{}}, args{dto: requests1}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sts := &StatisticService{
				logger: tt.fields.logger,
				em:     tt.fields.em,
			}
			if got := sts.SaveEntry(tt.args.dto); got != tt.want {
				t.Errorf("SaveEntry() = %v, want %v", got, tt.want)
			}
		})
	}
}
