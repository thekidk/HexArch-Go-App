package api

import (
	"hex/internal/ports"
)

type Adapter struct {
	db ports.DBPort
	arith ports.ArithmeticPort
}

func NewAdapter(db ports.DBPort, arith ports.ArithmeticPort) *Adapter {
	return &Adapter{db: db, arith: arith}
}

func (apiAdapter Adapter) GetAddition(a, b int32) (int32, error) {

	answer, err := apiAdapter.arith.Addition(a, b)
	if err != nil {
		return 0, err
	}

	err = apiAdapter.db.AddToHistory(answer, "addition")
	if err != nil {
		return 0, err
	}

	return answer, nil
}

func (apiAdapter Adapter) GetSubtraction(a, b int32) (int32, error) {

	answer, err := apiAdapter.arith.Subtraction(a, b)
	if err != nil {
		return 0, err
	}

	return answer, nil
}

func (apiAdapter Adapter) GetMultiplication(a, b int32) (int32, error) {

	answer, err := apiAdapter.arith.Multiplication(a, b)
	if err != nil {
		return 0, err
	}

	return answer, nil
}

func (apiAdapter Adapter) GetDivision(a, b int32) (int32, error) {

	answer, err := apiAdapter.arith.Division(a, b)
	if err != nil {
		return 0, err
	}

	return answer, nil
}
