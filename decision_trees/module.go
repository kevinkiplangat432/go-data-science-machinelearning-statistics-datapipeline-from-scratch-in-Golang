package decision_trees

import "fmt"

// Module implements the Runner interface.
// TODO: add your algorithm structs and functions below.
type Module struct{}

func NewModule() *Module { return &Module{} }

func (m *Module) Name() string { return "decision_trees" }

func (m *Module) Run() error {
	// TODO: implement decision_trees demo logic here
	fmt.Println("decision_trees: not yet implemented")
	return nil
}
