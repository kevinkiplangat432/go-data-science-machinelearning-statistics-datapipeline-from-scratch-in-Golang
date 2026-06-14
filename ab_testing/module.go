package ab_testing

import "fmt"

// Module implements the Runner interface.
// TODO: add your algorithm structs and functions below.
type Module struct{}

func NewModule() *Module { return &Module{} }

func (m *Module) Name() string { return "ab_testing" }

func (m *Module) Run() error {
	// TODO: implement ab_testing demo logic here
	fmt.Println("ab_testing: not yet implemented")
	return nil
}
