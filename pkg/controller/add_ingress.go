package controller

import (
	"github.com/openshift-knative/knative-openshift-ingress/pkg/controller/ingress"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, ingress.Add)
}
