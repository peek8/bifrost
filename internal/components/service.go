/*
 * Copyright (c) 2025 peek8.io
 *
 * Created Date: Wednesday, November 26th 2025, 3:39:34 pm
 * Author: Md. Asraful Haque
 *
 */

package components

import (
	corev1 "k8s.io/api/core/v1"
	apiequality "k8s.io/apimachinery/pkg/api/equality"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type Service struct {
	corev1.Service
}

func (s *Service) DeepCopySpecInto(other Component) {
	s.Spec.DeepCopyInto(&other.(*Service).Spec)
}

func (s *Service) DiffersSemanticallyFrom(other Component) bool {
	return !apiequality.Semantic.DeepDerivative(s.Spec, other.(*Service).Spec)
}

func (s *Service) GetClientObject() client.Object {
	return &s.Service
}

func (s *Service) IsReady() (string, string, bool) {
	return "", "", true
}

func (s *Service) PropagateLabels(_ map[string]string) {}
