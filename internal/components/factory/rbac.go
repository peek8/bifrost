/*
 * Copyright (c) 2025 peek8.io
 *
 * Created Date: Thursday, November 27th 2025, 12:12:59 pm
 * Author: Md. Asraful Haque
 *
 */

package factory

import (
	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ClusterRoleBuilder struct {
	cr *rbacv1.ClusterRole
}

func (crb ClusterRoleBuilder) WithRules(apiGroup string, resources, verbs []string) ClusterRoleBuilder {
	crb.cr.Rules = append(crb.cr.Rules, rbacv1.PolicyRule{
		APIGroups: []string{apiGroup},
		Resources: resources,
		Verbs:     verbs,
	})

	return crb
}

func (crb ClusterRoleBuilder) Get() *rbacv1.ClusterRole {
	return crb.cr
}

func NewClusterRole(name, namespace string) ClusterRoleBuilder {
	return ClusterRoleBuilder{
		cr: &rbacv1.ClusterRole{
			ObjectMeta: metav1.ObjectMeta{
				Name:      name,
				Namespace: namespace,
			},
		},
	}
}

func ClusterRoleBinding(name string, role rbacv1.ClusterRole, sa corev1.ServiceAccount) rbacv1.ClusterRoleBinding {
	return rbacv1.ClusterRoleBinding{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: role.Namespace,
		},
		Subjects: []rbacv1.Subject{{
			Kind:      "ServiceAccount",
			Name:      sa.Name,
			Namespace: sa.Namespace,
		}},
		RoleRef: rbacv1.RoleRef{
			APIGroup: "rbac.authorization.k8s.io",
			Kind:     "ClusterRole",
			Name:     role.Name,
		},
	}
}

func NewServiceAccount(name, namespace string) corev1.ServiceAccount {
	return corev1.ServiceAccount{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		}}
}
