// Package utils contain utility functions and interfaces
package utils

import (
	"context"
	"fmt"
	"reflect"

	"github.com/samber/lo"
	corev1 "k8s.io/api/core/v1"
)

type Option[T any] interface {
	// Applies the option to the given object.
	Apply() func(*T) error
}

func ApplyAll[Result any](target *Result, opts ...Option[Result]) error {
	for _, opt := range opts {
		err := opt.Apply()(target)

		if err != nil {
			return err
		}
	}

	return nil
}

type AddConfigMapAsVolume struct {
	ConfigMapName string
	VolumeName    string
	MountPath     string
	SubPath       string
	ContainerName string
	Mode          *int32
}

// Apply adds the ConfigMap as a volume to the Deployment and mounts it to the container.
// If the volume is already present inside the deployment, it will not be added again.
func (a AddConfigMapAsVolume) Apply(ctx context.Context) func(*corev1.PodTemplateSpec) error {
	return func(d *corev1.PodTemplateSpec) error {
		newVolume := corev1.Volume{
			Name: a.VolumeName,
			VolumeSource: corev1.VolumeSource{
				ConfigMap: &corev1.ConfigMapVolumeSource{
					LocalObjectReference: corev1.LocalObjectReference{Name: a.ConfigMapName},
					DefaultMode:          a.Mode,
				},
			},
		}

		volume, found := lo.Find(d.Spec.Volumes, func(vol corev1.Volume) bool {
			return vol.Name == a.VolumeName
		})

		if found {
			if volume.VolumeSource.ConfigMap == nil {
				return fmt.Errorf("volume %s already exists with a different source", a.VolumeName)
			}

			if volume.VolumeSource.ConfigMap.Name != a.ConfigMapName {
				return fmt.Errorf("volume %s already exists with a different configMap", a.VolumeName)
			}
		} else {
			d.Spec.Volumes = append(d.Spec.Volumes, newVolume)
		}

		container, err := mustFindContainer(d.Spec.Containers, a.ContainerName)

		if err != nil {
			return err
		}

		newVolumeMount := corev1.VolumeMount{
			Name:      a.VolumeName,
			MountPath: a.MountPath,
			SubPath:   a.SubPath,
		}

		container.VolumeMounts, err = addVolumeMount(&container, newVolumeMount)

		return err
	}
}

type AddPVC struct {
	PVCName       string
	MountPath     string
	ContainerName string
}

func (a AddPVC) Apply(ctx context.Context) func(*corev1.PodTemplateSpec) error {
	return func(d *corev1.PodTemplateSpec) error {
		newVolume := corev1.Volume{
			Name: a.PVCName,
			VolumeSource: corev1.VolumeSource{
				PersistentVolumeClaim: &corev1.PersistentVolumeClaimVolumeSource{
					ClaimName: a.PVCName,
				},
			},
		}

		found := lo.ContainsBy(d.Spec.Volumes, func(vol corev1.Volume) bool {
			return vol.Name == a.PVCName
		})

		if found {
			return fmt.Errorf("volume %s already exists ", a.PVCName)
		} else {
			d.Spec.Volumes = append(d.Spec.Volumes, newVolume)
		}

		container, err := mustFindContainer(d.Spec.Containers, a.ContainerName)

		if err != nil {
			return err
		}

		newVolumeMount := corev1.VolumeMount{
			Name:      a.PVCName,
			MountPath: a.MountPath,
		}

		container.VolumeMounts, err = addVolumeMount(&container, newVolumeMount)

		return err
	}
}

func addVolumeMount(container *corev1.Container, newMount corev1.VolumeMount) ([]corev1.VolumeMount, error) {
	volumeMount, found := lo.Find(container.VolumeMounts, func(vm corev1.VolumeMount) bool {
		return vm.MountPath == newMount.MountPath
	})

	if found {
		if !reflect.DeepEqual(volumeMount, newMount) {
			return container.VolumeMounts, fmt.Errorf("volume mount %s already exists in container %s", newMount.MountPath, container.Name)
		}
	} else {
		return append(container.VolumeMounts, newMount), nil
	}

	return container.VolumeMounts, nil
}

func mustFindContainer(containers []corev1.Container, name string) (corev1.Container, error) {
	container, found := lo.Find(containers, func(c corev1.Container) bool {
		return c.Name == name
	})

	if !found {
		return container, fmt.Errorf("unable to find container with name %s", name)
	}

	return container, nil
}
