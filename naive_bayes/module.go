package naive_bayes

import "fmt"

// Module implements the Runner interface.
// TODO: add your algorithm structs and functions below.
type Module struct{}

func NewModule() *Module { return &Module{} }

func (m *Module) Name() string { return "naive_bayes" }

func (m *Module) Run() error {
	// TODO: implement naive_bayes demo logic here
	fmt.Println("naive_bayes: not yet implemented")
	return nil
}
