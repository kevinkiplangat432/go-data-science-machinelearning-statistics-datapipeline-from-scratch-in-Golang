package analytics

import "fmt"

// Module implements the Runner interface.
// TODO: add your algorithm structs and functions below.
type Module struct{}

func NewModule() *Module { return &Module{} }

func (m *Module) Name() string { return "analytics" }

func (m *Module) Run() error {
	// TODO: implement analytics demo logic here
	fmt.Println("analytics: not yet implemented")
	return nil
}
