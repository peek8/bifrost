package components

import (
	"fmt"

	corev1 "k8s.io/api/core/v1"
	apiequality "k8s.io/apimachinery/pkg/api/equality"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type PersistentVolumeClaim struct {
	corev1.PersistentVolumeClaim
}

func (p *PersistentVolumeClaim) DeepCopySpecInto(other Component) {
	p.Spec.DeepCopyInto(&other.(*PersistentVolumeClaim).Spec)
}

func (p *PersistentVolumeClaim) DiffersSemanticallyFrom(other Component) bool {
	return !apiequality.Semantic.DeepDerivative(p.Spec, other.(*PersistentVolumeClaim).Spec)
}

func (p *PersistentVolumeClaim) GetClientObject() client.Object {
	return &p.PersistentVolumeClaim
}

func (p *PersistentVolumeClaim) IsReady() (string, string, bool) {
	if p.Spec.VolumeName == "" {
		reason := "ClaimNotBound"
		message := fmt.Sprintf("PersistentVolumeClaim %s is not yet bound.", p.Name)

		return reason, message, false
	}

	return "", "", true
}

func (p *PersistentVolumeClaim) PropagateLabels(_ map[string]string) {}
