package statistics

import "fmt"

// Module implements the Runner interface.
// TODO: add your algorithm structs and functions below.
type Module struct{}

func NewModule() *Module { return &Module{} }

func (m *Module) Name() string { return "statistics" }

func (m *Module) Run() error {
	// TODO: implement statistics demo logic here
	fmt.Println("statistics: not yet implemented")
	return nil
}
