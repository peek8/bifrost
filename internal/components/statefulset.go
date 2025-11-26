/*
 * Copyright (c) 2025 peek8.io
 *
 * Created Date: Wednesday, November 26th 2025, 3:43:35 pm
 * Author: Mss. Asraful Haque
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

type StatefulSet struct {
	appsv1.StatefulSet
}

func (ss *StatefulSet) DeepCopySpecInto(other Component) {
	ss.Spec.DeepCopyInto(&other.(*StatefulSet).Spec)
}

func (ss *StatefulSet) DiffersSemanticallyFrom(other Component) bool {
	return !apiequality.Semantic.DeepDerivative(ss.Spec, other.(*StatefulSet).Spec)
}

func (ss *StatefulSet) GetClientObject() client.Object {
	return &ss.StatefulSet
}

func (ss *StatefulSet) IsReady() (string, string, bool) {
	if ss.Status.ObservedGeneration != ss.Generation {
		message := fmt.Sprintf("StatefulSet %s has not yet observed generation %d", ss.Name, ss.Generation)

		return "NotObserved", message, false
	}

	if ss.Status.UpdateRevision == ss.Status.CurrentRevision &&
		ss.Status.ReadyReplicas == *ss.Spec.Replicas {
		return "", "", true
	}

	return "NotReconciled", "StatefulSet has not reconciled yet", false
}

func (ss *StatefulSet) PropagateLabels(labels map[string]string) {
	if ss.Spec.Template.ObjectMeta.Labels == nil {
		ss.Spec.Template.ObjectMeta.Labels = map[string]string{}
	}

	maps.Copy(ss.Spec.Template.ObjectMeta.Labels, labels)
}

func (ss *StatefulSet) GetWorkload() corev1.PodTemplateSpec {
	return ss.Spec.Template
}

func (ss *StatefulSet) SetWorkload(workload corev1.PodTemplateSpec) {
	ss.Spec.Template = workload
}
