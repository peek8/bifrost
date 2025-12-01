/*
 * Copyright (c) 2025 peek8.io
 *
 * Created Date: Thursday, November 27th 2025, 3:23:00 pm
 * Author: Md. Asraful Haque
 *
 */

// Package loki contains loki related components
package loki

import (
	"context"
	"fmt"
	"path"

	"github.com/peek8/bifrost/api/v1alpha1"
	"github.com/peek8/bifrost/internal/components"
	"github.com/peek8/bifrost/internal/components/factory"
	"github.com/peek8/bifrost/internal/utils"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/utils/ptr"
)

const (
	primaryContainer = "main"
	mainPort         = 3100
	mainPortName     = "main"
	containerImage   = "grafana/loki:3.5.8"

	configPath  = "/etc/loki"
	configFile  = "loki.yaml"
	storagePath = "/loki"
	storageSize = "5Gi"
)

type Data struct {
	Name            string
	LogSpaceSpec    v1alpha1.LogSpaceSpec
	Namespace       string
	LokiServiceName string
}

type Loki struct {
	config      corev1.ConfigMap
	pvc         corev1.PersistentVolumeClaim
	statefulSet appsv1.StatefulSet
	service     corev1.Service
}

func (lk Loki) ToComponents() components.Components {
	items := components.Components{
		&components.ConfigMap{ConfigMap: lk.config},
		&components.PersistentVolumeClaim{PersistentVolumeClaim: lk.pvc},
		&components.StatefulSet{StatefulSet: lk.statefulSet},
		&components.Service{Service: lk.service},
	}

	return items.NonEmptyComponents()
}

type Builder struct{}

func (b Builder) New(ctx context.Context, data Data) (Loki, error) {
	cm, err := lokiConfigMap(data)

	if err != nil {
		return Loki{}, err
	}

	loki := Loki{
		config:      cm,
		pvc:         factory.NewPVC(data.Name+"-data", data.Namespace, data.LogSpaceSpec.LokiConfig.Storage.Size, data.LogSpaceSpec.PVCStorage.StorageClass),
		statefulSet: statefulSet(data),
		service:     service(data),
	}

	options := []utils.Option[corev1.PodTemplateSpec]{
		utils.AddConfigMapAsVolume{
			ConfigMapName: loki.config.Name,
			VolumeName:    "loki-config",
			MountPath:     configPath,
			ContainerName: primaryContainer,
			Mode:          ptr.To(int32(0755)),
		},
		utils.AddPVC{
			PVCName:       loki.pvc.Name,
			MountPath:     storagePath,
			ContainerName: primaryContainer,
		},
	}

	err = utils.ApplyAll(&loki.statefulSet.Spec.Template, options...)

	if err != nil {
		return Loki{}, err
	}

	return loki, nil

}

func statefulSet(data Data) appsv1.StatefulSet {
	container := corev1.Container{
		Name:  primaryContainer,
		Image: containerImage,
		Ports: []corev1.ContainerPort{
			{
				Name:          mainPortName,
				ContainerPort: mainPort,
				Protocol:      corev1.ProtocolTCP,
			},
			{
				Name:          "grpc",
				ContainerPort: 9095,
			},
		},
		Args: []string{
			fmt.Sprintf("-config.file=%s", path.Join(configPath, configFile)),
		},
	}

	return *factory.NewStatefulSet(data.Name, data.Namespace, container).
		WithLabels(factory.K8sLabels(data.Name, "loki")).
		Get()
}

func service(d Data) corev1.Service {
	return corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      d.LokiServiceName,
			Namespace: d.Namespace,
		},
		Spec: corev1.ServiceSpec{
			Selector: map[string]string{
				"app": d.Name,
			},
			Ports: []corev1.ServicePort{
				{
					Name:       "http",
					Protocol:   corev1.ProtocolTCP,
					Port:       mainPort,
					TargetPort: intstr.FromString(mainPortName),
				},
			},
			Type: corev1.ServiceTypeClusterIP,
		},
	}
}
