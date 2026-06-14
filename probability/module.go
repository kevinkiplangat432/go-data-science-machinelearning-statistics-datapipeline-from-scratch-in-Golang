package probability

import "fmt"

// Module implements the Runner interface.
// TODO: add your algorithm structs and functions below.
type Module struct{}

func NewModule() *Module { return &Module{} }

func (m *Module) Name() string { return "probability" }

func (m *Module) Run() error {
	// TODO: implement probability demo logic here
	fmt.Println("probability: not yet implemented")
	return nil
}
