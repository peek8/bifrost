/*
 * Copyright (c) 2025 peek8.io
 *
 * Created Date: Thursday, November 27th 2025, 9:38:01 am
 * Author: Md. Asraful Haque
 *
 */
// Package factory contains function to build different components
package factory

import (
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/utils/ptr"
)

func NewDeployment(name string, namespace string, labels map[string]string, container corev1.Container) appsv1.Deployment {
	return appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
			Labels:    labels,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: ptr.To(int32(1)),
			Selector: appLabelSeclector(name),
			Template: podTemplateSpec(name, container),
		},
	}
}

// DaemonSetBuilder makes daemonSet
type DaemonSetBuilder struct {
	daemonSet *appsv1.DaemonSet
}

func (dsb DaemonSetBuilder) WithLabels(labels map[string]string) DaemonSetBuilder {
	dsb.daemonSet.Labels = labels

	return dsb
}

func (dsb DaemonSetBuilder) WithServiceAccount(sa string) DaemonSetBuilder {
	dsb.daemonSet.Spec.Template.Spec.ServiceAccountName = sa

	return dsb
}

func (dsb DaemonSetBuilder) Get() *appsv1.DaemonSet {
	return dsb.daemonSet
}

func NewDaemonSet(name string, namespace string, mainContainer corev1.Container) DaemonSetBuilder {
	return DaemonSetBuilder{
		daemonSet: &appsv1.DaemonSet{
			ObjectMeta: metav1.ObjectMeta{
				Name:      name,
				Namespace: namespace,
			},
			Spec: appsv1.DaemonSetSpec{
				Selector: appLabelSeclector(name),
				Template: podTemplateSpec(name, mainContainer),
			},
		},
	}
}

// StatefulSetBuilder makes statefulset
type StatefulSetBuilder struct {
	statefulSet *appsv1.StatefulSet
}

func (dsb StatefulSetBuilder) WithLabels(labels map[string]string) StatefulSetBuilder {
	dsb.statefulSet.Labels = labels

	return dsb
}

func (dsb StatefulSetBuilder) WithServiceAccount(sa string) StatefulSetBuilder {
	dsb.statefulSet.Spec.Template.Spec.ServiceAccountName = sa

	return dsb
}

func (dsb StatefulSetBuilder) Get() *appsv1.StatefulSet {
	return dsb.statefulSet
}

func NewStatefulSet(name string, namespace string, mainContainer corev1.Container) StatefulSetBuilder {
	return StatefulSetBuilder{
		statefulSet: &appsv1.StatefulSet{
			ObjectMeta: metav1.ObjectMeta{
				Name:      name,
				Namespace: namespace,
			},
			Spec: appsv1.StatefulSetSpec{
				Replicas:    ptr.To(int32(1)),
				ServiceName: name,
				Selector:    appLabelSeclector(name),
				Template:    podTemplateSpec(name, mainContainer),
			},
		},
	}
}

func podTemplateSpec(name string, containers ...corev1.Container) corev1.PodTemplateSpec {
	return corev1.PodTemplateSpec{
		ObjectMeta: metav1.ObjectMeta{
			Labels: map[string]string{
				"app": name,
			},
		},
		Spec: corev1.PodSpec{
			Containers: containers,
		},
	}
}

func appLabelSeclector(app string) *metav1.LabelSelector {
	return &metav1.LabelSelector{
		MatchLabels: map[string]string{
			"app": app,
		},
	}

}

func K8sLabels(name, component string) map[string]string {
	return map[string]string{
		"app": name,
		//"app.kubernetes.io/name":      name,
		"app.kubernetes.io/component": component,
		"app.kubernetes.io/instance":  name,
		"app.kubernetes.io/part-of":   "bifrost",
	}
}
