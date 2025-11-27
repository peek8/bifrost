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
	rbacv1 "k8s.io/api/rbac/v1"
	apiequality "k8s.io/apimachinery/pkg/api/equality"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type Role struct {
	rbacv1.Role
}

func (r *Role) DeepCopySpecInto(other Component) {
	other.(*Role).Rules = r.DeepCopy().Rules
}

func (r *Role) DiffersSemanticallyFrom(other Component) bool {
	return !apiequality.Semantic.DeepDerivative(r.Rules, other.(*Role).Rules)
}

func (r *Role) GetClientObject() client.Object {
	return &r.Role
}

func (r *Role) IsReady() (string, string, bool) {
	return "", "", true
}

func (s *Role) PropagateLabels(_ map[string]string) {}

type ClusterRole struct {
	rbacv1.ClusterRole
}

func (cr *ClusterRole) DeepCopySpecInto(other Component) {
	other.(*ClusterRole).Rules = cr.DeepCopy().Rules
}

func (cr *ClusterRole) DiffersSemanticallyFrom(other Component) bool {
	return !apiequality.Semantic.DeepDerivative(cr.Rules, other.(*ClusterRole).Rules)
}

func (cr *ClusterRole) GetClientObject() client.Object {
	return &cr.ClusterRole
}

func (cr *ClusterRole) IsReady() (string, string, bool) {
	return "", "", true
}

func (s *ClusterRole) PropagateLabels(_ map[string]string) {}

// RoleBinding

type RoleBinding struct {
	rbacv1.RoleBinding
}

func (rb *RoleBinding) DeepCopySpecInto(other Component) {
	other.(*RoleBinding).Subjects = rb.DeepCopy().Subjects
	other.(*RoleBinding).RoleRef = rb.RoleRef
}

func (rb *RoleBinding) DiffersSemanticallyFrom(other Component) bool {
	return !(apiequality.Semantic.DeepDerivative(rb.Subjects, other.(*RoleBinding).Subjects) &&
		apiequality.Semantic.DeepDerivative(rb.RoleRef, other.(*RoleBinding).RoleRef))
}

func (rb *RoleBinding) GetClientObject() client.Object {
	return &rb.RoleBinding
}

func (rb *RoleBinding) IsReady() (string, string, bool) {
	return "", "", true
}

func (rb *RoleBinding) PropagateLabels(_ map[string]string) {}

type ClusterRoleBinding struct {
	rbacv1.ClusterRoleBinding
}

func (rb *ClusterRoleBinding) DeepCopySpecInto(other Component) {
	other.(*ClusterRoleBinding).Subjects = rb.DeepCopy().Subjects
	other.(*ClusterRoleBinding).RoleRef = rb.RoleRef
}

func (rb *ClusterRoleBinding) DiffersSemanticallyFrom(other Component) bool {
	return !(apiequality.Semantic.DeepDerivative(rb.Subjects, other.(*ClusterRoleBinding).Subjects) &&
		apiequality.Semantic.DeepDerivative(rb.RoleRef, other.(*ClusterRoleBinding).RoleRef))
}

func (rb *ClusterRoleBinding) GetClientObject() client.Object {
	return &rb.ClusterRoleBinding
}

func (rb *ClusterRoleBinding) IsReady() (string, string, bool) {
	return "", "", true
}

func (rb *ClusterRoleBinding) PropagateLabels(_ map[string]string) {}

// Service account

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
