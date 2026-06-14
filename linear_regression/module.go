package linear_regression

import "fmt"

// Module implements the Runner interface.
// TODO: add your algorithm structs and functions below.
type Module struct{}

func NewModule() *Module { return &Module{} }

func (m *Module) Name() string { return "linear_regression" }

func (m *Module) Run() error {
	// TODO: implement linear_regression demo logic here
	fmt.Println("linear_regression: not yet implemented")
	return nil
}
