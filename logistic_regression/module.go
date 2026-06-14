package logistic_regression

import "fmt"

// Module implements the Runner interface.
// TODO: add your algorithm structs and functions below.
type Module struct{}

func NewModule() *Module { return &Module{} }

func (m *Module) Name() string { return "logistic_regression" }

func (m *Module) Run() error {
	// TODO: implement logistic_regression demo logic here
	fmt.Println("logistic_regression: not yet implemented")
	return nil
}
