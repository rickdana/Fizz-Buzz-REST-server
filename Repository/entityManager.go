package Repository

import "github.com/rickdana/fizzbuzzApi/Model"

type EntityManager interface {
	Save(entity interface{}) (err error)
	FindAll() ([]Model.FizzBuzz, error)
}
