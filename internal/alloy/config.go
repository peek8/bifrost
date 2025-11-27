// Package alloy contains all the workloads and cofig for alloy
package alloy

import (
	_ "embed"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

//go:embed config.alloy
var alloyConfig string

func AlloyConfigMap(data Data) corev1.ConfigMap {
	return corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "alloy",
			Namespace: data.Namespace,
		},
		BinaryData: map[string][]byte{
			"config.alloy": []byte(alloyConfig),
		},
	}
}
