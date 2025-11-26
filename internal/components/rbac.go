/*
 * Copyright (c) 2025 peek8.io
 *
 * Created Date: Wednesday, November 26th 2025, 3:39:34 pm
 * Author: Md. Asraful Haque
 *
 */

package components

import (
	rbacv1 "k8s.io/api/rbac/v1"
	apiequality "k8s.io/apimachinery/pkg/api/equality"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type Role struct {
	rbacv1.Role
}

func (s *Role) DeepCopySpecInto(other Component) {
	other.(*Role).Rules = s.DeepCopy().Rules
}

func (s *Role) DiffersSemanticallyFrom(other Component) bool {
	return !apiequality.Semantic.DeepDerivative(s.Rules, other.(*Role).Rules)
}

func (s *Role) GetClientObject() client.Object {
	return &s.Role
}

func (s *Role) IsReady() (string, string, bool) {
	return "", "", true
}

func (s *Role) PropagateLabels(_ map[string]string) {}

type RoleBinding struct {
	rbacv1.RoleBinding
}

func (s *RoleBinding) DeepCopySpecInto(other Component) {
	other.(*RoleBinding).Subjects = s.DeepCopy().Subjects
	other.(*RoleBinding).RoleRef = s.RoleRef
}

func (s *RoleBinding) DiffersSemanticallyFrom(other Component) bool {
	return !(apiequality.Semantic.DeepDerivative(s.Subjects, other.(*RoleBinding).Subjects) &&
		apiequality.Semantic.DeepDerivative(s.RoleRef, other.(*RoleBinding).RoleRef))
}

func (s *RoleBinding) GetClientObject() client.Object {
	return &s.RoleBinding
}

func (s *RoleBinding) IsReady() (string, string, bool) {
	return "", "", true
}

func (s *RoleBinding) PropagateLabels(_ map[string]string) {}
