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
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type ServiceAccount struct {
	corev1.ServiceAccount
}

func (s *ServiceAccount) DeepCopySpecInto(other Component) {
}

func (s *ServiceAccount) DiffersSemanticallyFrom(other Component) bool {
	return false
}

func (s *ServiceAccount) GetClientObject() client.Object {
	return &s.ServiceAccount
}

func (s *ServiceAccount) IsReady() (string, string, bool) {
	return "", "", true
}

func (s *ServiceAccount) PropagateLabels(_ map[string]string) {}
