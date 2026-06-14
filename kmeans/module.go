package kmeans

import "fmt"

// Module implements the Runner interface.
// TODO: add your algorithm structs and functions below.
type Module struct{}

func NewModule() *Module { return &Module{} }

func (m *Module) Name() string { return "kmeans" }

func (m *Module) Run() error {
	// TODO: implement kmeans demo logic here
	fmt.Println("kmeans: not yet implemented")
	return nil
}
