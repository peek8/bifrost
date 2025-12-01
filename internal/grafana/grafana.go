/*
 * Copyright (c) 2025 peek8.io
 *
 * Created Date: Thursday, November 27th 2025, 4:08:10 pm
 * Author: Md. Asraful Haque
 *
 */

// Package grafana contains grafana related components
package grafana

import (
	"context"

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
	mainPort         = 3000
	mainPortName     = "main"
	containerImage   = "grafana/grafana:12.1"

	storagePath = "/var/lib/grafana"
	storageSize = "5Gi"

	// grafana
	grafanaAdminUserKey = "GF_SECURITY_ADMIN_USER"
	grafanaAdminPassKey = "GF_SECURITY_ADMIN_PASSWORD"
)

type Data struct {
	Name            string
	LogSpaceSpec    v1alpha1.LogSpaceSpec
	Namespace       string
	LokiServiceName string
}

type Grafana struct {
	datasourceConfig corev1.ConfigMap
	dashboardConfig  corev1.ConfigMap
	pvc              corev1.PersistentVolumeClaim
	deployment       appsv1.Deployment
	service          corev1.Service
}

func (g Grafana) ToComponents() components.Components {
	items := components.Components{
		&components.ConfigMap{ConfigMap: g.datasourceConfig},
		&components.ConfigMap{ConfigMap: g.dashboardConfig},
		&components.PersistentVolumeClaim{PersistentVolumeClaim: g.pvc},
		&components.Deployment{Deployment: g.deployment},
		&components.Service{Service: g.service},
	}

	return items.NonEmptyComponents()
}

type Builder struct{}

func (b Builder) New(ctx context.Context, data Data) (Grafana, error) {
	dsCM, err := grafanaDSConfigMap(data)
	if err != nil {
		return Grafana{}, err
	}

	dbCM, err := grafanaDashboardConfigMap(data)
	if err != nil {
		return Grafana{}, err
	}

	grafana := Grafana{
		datasourceConfig: dsCM,
		dashboardConfig:  dbCM,
		pvc:              factory.NewPVC(data.Name+"-data", data.Namespace, data.LogSpaceSpec.GrafanaConfig.Storage.Size, data.LogSpaceSpec.PVCStorage.StorageClass),
		deployment:       deployment(data),
		service:          service(data),
	}

	options := []utils.Option[corev1.PodTemplateSpec]{
		utils.AddConfigMapAsVolume{
			ConfigMapName: dsCM.Name,
			VolumeName:    dsCM.Name,
			MountPath:     grafanaDataSourceProvisioningDir,
			ContainerName: primaryContainer,
			Mode:          ptr.To(int32(0755)),
		},
		utils.AddConfigMapAsVolume{
			ConfigMapName: dbCM.Name,
			VolumeName:    dbCM.Name,
			MountPath:     grafanaDashboardProvisioningDir,
			ContainerName: primaryContainer,
			Mode:          ptr.To(int32(0755)),
		},
		utils.AddPVC{
			PVCName:       grafana.pvc.Name,
			MountPath:     storagePath,
			ContainerName: primaryContainer,
		},
	}

	err = utils.ApplyAll(&grafana.deployment.Spec.Template, options...)

	if err != nil {
		return Grafana{}, err
	}

	return grafana, nil

}

func deployment(data Data) appsv1.Deployment {
	container := corev1.Container{
		Name:  primaryContainer,
		Image: containerImage,
		Ports: []corev1.ContainerPort{
			{
				Name:          mainPortName,
				ContainerPort: mainPort,
				Protocol:      corev1.ProtocolTCP,
			},
		},
		Env: []corev1.EnvVar{
			{
				Name:  grafanaAdminUserKey,
				Value: "admin",
			},
			{
				Name:  grafanaAdminPassKey,
				Value: "admin",
			},
		},
	}

	return factory.NewDeployment(data.Name, data.Namespace, factory.K8sLabels(data.Name, "grafana"), container)
}

func service(d Data) corev1.Service {
	return corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			// TODO: for now hardcoded to grafana
			Name:      "grafana",
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
