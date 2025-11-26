/*
 * Copyright (c) 2025 peek8.io
 *
 * Created Date: Wednesday, November 26th 2025, 3:55:51 pm
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
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type DaemonSet struct {
	appsv1.DaemonSet
}

func (ds *DaemonSet) DeepCopySpecInto(other Component) {
	ds.Spec.DeepCopyInto(&other.(*DaemonSet).Spec)
}

func (ds *DaemonSet) DiffersSemanticallyFrom(other Component) bool {
	return !apiequality.Semantic.DeepDerivative(ds.Spec, other.(*DaemonSet).Spec)
}

func (ds *DaemonSet) GetClientObject() client.Object {
	return &ds.DaemonSet
}

func (ds *DaemonSet) IsReady() (string, string, bool) {
	if ds.Status.ObservedGeneration != ds.Generation {
		message := fmt.Sprintf("DaemonSet %s has not yet observed generation %d", ds.Name, ds.Generation)

		return "NotObserved", message, false
	}

	if ds.Status.NumberReady == ds.Status.DesiredNumberScheduled {
		return "", "", true
	}

	return "NotReconciled", "DaemonSet has not reconciled yet", false
}

func (ds *DaemonSet) PropagateLabels(labels map[string]string) {
	if ds.Spec.Template.ObjectMeta.Labels == nil {
		ds.Spec.Template.ObjectMeta.Labels = map[string]string{}
	}

	maps.Copy(ds.Spec.Template.ObjectMeta.Labels, labels)
}

func (ds *DaemonSet) GetWorkload() corev1.PodTemplateSpec {
	return ds.Spec.Template
}

func (ds *DaemonSet) SetWorkload(workload corev1.PodTemplateSpec) {
	ds.Spec.Template = workload
}
