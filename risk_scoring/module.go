package risk_scoring

import "fmt"

// Module implements the Runner interface.
// TODO: add your algorithm structs and functions below.
type Module struct{}

func NewModule() *Module { return &Module{} }

func (m *Module) Name() string { return "risk_scoring" }

func (m *Module) Run() error {
	// TODO: implement risk_scoring demo logic here
	fmt.Println("risk_scoring: not yet implemented")
	return nil
}
