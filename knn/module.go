package knn

import "fmt"

// Module implements the Runner interface.
// TODO: add your algorithm structs and functions below.
type Module struct{}

func NewModule() *Module { return &Module{} }

func (m *Module) Name() string { return "knn" }

func (m *Module) Run() error {
	// TODO: implement knn demo logic here
	fmt.Println("knn: not yet implemented")
	return nil
}
