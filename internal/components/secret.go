/*
 * Copyright (c) 2025 peek8.io
 *
 * Created Date: Wednesday, November 26th 2025, 3:39:34 pm
 * Author: Md. Asraful Haque
 *
 */

package components

import (
	"bytes"
	"slices"

	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type Secret struct {
	corev1.Secret
	// The listed fields are ignored in DiffersSemanticallyFrom and
	// DeepCopySpecInto.
	IgnoredFields []string
}

func (s *Secret) DeepCopySpecInto(other Component) {
	otherSecret := other.(*Secret)

	otherSecret.Type = s.Secret.Type
	newData := make(map[string][]byte)

	for k, v := range s.Data {
		if slices.Contains(s.IgnoredFields, k) {
			newData[k] = otherSecret.Data[k]
		} else {
			newData[k] = v
		}
	}

	otherSecret.Data = newData
}

func (s Secret) DiffersSemanticallyFrom(other Component) bool {
	otherSecret := other.(*Secret)

	if s.Type == "" {
		s.Type = corev1.SecretTypeOpaque
	}

	if otherSecret.Type != s.Type {
		return true
	}

	for k, v := range s.Data {
		if slices.Contains(s.IgnoredFields, k) {
			continue
		}

		if !bytes.Equal(v, otherSecret.Data[k]) {
			return true
		}
	}

	return false
}

func mergeStringData(s *Secret) {
	if s.StringData != nil {
		if s.Data == nil {
			s.Data = make(map[string][]byte)
		}

		for k, v := range s.StringData {
			s.Data[k] = []byte(v)
		}
	}
}

func (s *Secret) GetClientObject() client.Object {
	return &s.Secret
}

func (s Secret) IsReady() (reason string, message string, ready bool) {
	return "", "", true
}

func (s *Secret) PropagateLabels(_ map[string]string) {}
