/*
 * Copyright (c) 2025 peek8.io
 *
 * Created Date: Wednesday, November 26th 2025, 3:39:34 pm
 * Author: Md. Asraful Haque
 *
 */

package components

import (
	"fmt"
	"maps"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	apiequality "k8s.io/apimachinery/pkg/api/equality"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type Deployment struct {
	appsv1.Deployment
}

func (d *Deployment) DeepCopySpecInto(other Component) {
	d.Spec.DeepCopyInto(&other.(*Deployment).Spec)
}

func (d *Deployment) DiffersSemanticallyFrom(other Component) bool {
	return !apiequality.Semantic.DeepDerivative(d.Spec, other.(*Deployment).Spec)
}

func (d *Deployment) GetClientObject() client.Object {
	return &d.Deployment
}

func (d *Deployment) IsReady() (string, string, bool) {
	status := metav1.ConditionUnknown

	var reason, message string

	if d.Status.ObservedGeneration != d.Generation {
		message = fmt.Sprintf("Deployment %s has not yet observed generation %d", d.Name, d.Generation)

		return "NotObserved", message, false
	}

	for _, c := range d.Status.Conditions {
		if c.Type == appsv1.DeploymentAvailable {
			status = metav1.ConditionStatus(c.Status)
			reason = c.Reason
			message = c.Message

			break
		}
	}

	if status != "True" {
		message = fmt.Sprintf("Deployment %s is not reporting Available: %s", d.Name, message)

		return reason, message, false
	}

	if status == "True" {
		return "", "", true
	}

	return "NotReconciled", "Deployment has not reconciled yet", false
}

func (d *Deployment) PropagateLabels(labels map[string]string) {
	if d.Spec.Template.ObjectMeta.Labels == nil {
		d.Spec.Template.ObjectMeta.Labels = map[string]string{}
	}

	maps.Copy(d.Spec.Template.ObjectMeta.Labels, labels)
}

func (d Deployment) GetWorkload() corev1.PodTemplateSpec {
	return d.Spec.Template
}

func (d *Deployment) SetWorkload(workload corev1.PodTemplateSpec) {
	d.Spec.Template = workload
}
