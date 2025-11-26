/*
 * Copyright (c) 2025 peek8.io
 *
 * Created Date: Wednesday, November 26th 2025, 3:39:34 pm
 * Author: Md. Asraful Haque
 *
 */

package components

import (
	"reflect"
	"slices"

	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type ConfigMap struct {
	corev1.ConfigMap
	// The listed fields are ignored in DiffersSemanticallyFrom and
	// DeepCopySpecInto.
	IgnoredFields []string
}

func (c *ConfigMap) DeepCopySpecInto(other Component) {
	otherConfigMap := other.(*ConfigMap)

	newData := make(map[string]string)

	for k, v := range c.Data {
		if slices.Contains(c.IgnoredFields, k) {
			newData[k] = otherConfigMap.Data[k]
		} else {
			newData[k] = v
		}
	}

	otherConfigMap.Data = newData

	newBynaryData := make(map[string][]byte)

	for k, v := range c.BinaryData {
		if slices.Contains(c.IgnoredFields, k) {
			newBynaryData[k] = otherConfigMap.BinaryData[k]
		} else {
			newBynaryData[k] = v
		}
	}

	otherConfigMap.BinaryData = newBynaryData
}

func (c *ConfigMap) DiffersSemanticallyFrom(other Component) bool {
	otherConfigMap := other.(*ConfigMap)

	for k, v := range c.Data {
		if slices.Contains(c.IgnoredFields, k) {
			continue
		}

		if v != otherConfigMap.Data[k] {
			return true
		}
	}

	for k, v := range c.BinaryData {
		if slices.Contains(c.IgnoredFields, k) {
			continue
		}

		if !reflect.DeepEqual(v, otherConfigMap.BinaryData[k]) {
			return true
		}
	}

	return false
}

func (c *ConfigMap) GetClientObject() client.Object {
	return &c.ConfigMap
}

func (c *ConfigMap) IsReady() (string, string, bool) {
	return "", "", true
}

func (c *ConfigMap) PropagateLabels(_ map[string]string) {}
