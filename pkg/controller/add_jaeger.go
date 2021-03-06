package controller

import (
	"github.com/jaegertracing/jaeger-operator/pkg/controller/jaeger"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, jaeger.Add)
}
