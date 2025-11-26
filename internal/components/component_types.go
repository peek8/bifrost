/*
 * Copyright (c) 2025 peek8.io
 *
 * Created Date: Wednesday, November 26th 2025, 3:39:34 pm
 * Author: Md. Asraful Haque
 *
 */

// Package components contains all the components that will be applied to k8s cluster.
package components

import (
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// +kubebuilder:object:generate=false
type Component interface {
	client.Object
	GetClientObject() client.Object
	DeepCopySpecInto(other Component)
	DiffersSemanticallyFrom(other Component) bool
	IsReady() (reason string, message string, ready bool)
	PropagateLabels(labels map[string]string)
}

// Named is a filtering function, that can be used in e.g. DeleteFunc or
// FindFunc, to find a Component with a Componenticular name.
func Named(name string) func(Component) bool {
	return func(p Component) bool {
		return p.GetName() == name
	}
}

// +kubebuilder:object:generate=false
type Components []Component

func (p Components) FindComponentWithKindAndName(kind string, name string) Component {
	for _, Component := range p {
		if Component.GetObjectKind().GroupVersionKind().Kind == kind && Component.GetName() == name {
			return Component
		}
	}

	return nil
}

// +kubebuilder:object:generate=true
type ComponentStatus struct {
	APIVersion string  `json:"apiVersion,omitempty"`
	Kind       string  `json:"kind,omitempty"`
	Name       string  `json:"name,omitempty"`
	UID        string  `json:"uid,omitempty"`
	Ready      bool    `json:"ready"`
	Reason     *string `json:"reason,omitempty"`
	Message    *string `json:"message,omitempty"`
}
