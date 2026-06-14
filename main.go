package main

import (
	"fmt"
	"os"

	"go-data-science/ab_testing"
	"go-data-science/agent_monitoring"
	"go-data-science/analytics"
	"go-data-science/data_engineering"
	"go-data-science/decision_trees"
	"go-data-science/hallucination_detection"
	"go-data-science/kmeans"
	"go-data-science/knn"
	"go-data-science/linear_regression"
	"go-data-science/logistic_regression"
	"go-data-science/naive_bayes"
	"go-data-science/pii_detection"
	"go-data-science/probability"
	"go-data-science/risk_scoring"
	"go-data-science/statistics"
)

// Runner is the contract every module must satisfy.
type Runner interface {
	Name() string
	Run() error
}

// Registry holds all registered modules.
type Registry struct {
	modules []Runner
}

func (r *Registry) Register(m Runner) {
	r.modules = append(r.modules, m)
}

func (r *Registry) RunAll() {
	for _, m := range r.modules {
		fmt.Printf("\n=== %s ===\n", m.Name())
		if err := m.Run(); err != nil {
			fmt.Printf("ERROR: %v\n", err)
		}
	}
}

func (r *Registry) RunOne(name string) error {
	for _, m := range r.modules {
		if m.Name() == name {
			fmt.Printf("\n=== %s ===\n", m.Name())
			return m.Run()
		}
	}
	return fmt.Errorf("module %q not found", name)
}

func main() {
	reg := &Registry{}

	reg.Register(statistics.NewModule())
	reg.Register(probability.NewModule())
	reg.Register(ab_testing.NewModule())
	reg.Register(linear_regression.NewModule())
	reg.Register(logistic_regression.NewModule())
	reg.Register(decision_trees.NewModule())
	reg.Register(knn.NewModule())
	reg.Register(naive_bayes.NewModule())
	reg.Register(kmeans.NewModule())
	reg.Register(data_engineering.NewModule())
	reg.Register(analytics.NewModule())
	reg.Register(pii_detection.NewModule())
	reg.Register(risk_scoring.NewModule())
	reg.Register(agent_monitoring.NewModule())
	reg.Register(hallucination_detection.NewModule())

	if len(os.Args) > 1 {
		if err := reg.RunOne(os.Args[1]); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	} else {
		reg.RunAll()
	}
}
