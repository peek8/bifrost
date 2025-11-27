package alloy

import (
	"context"
	"fmt"
	"path"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"

	"github.com/peek8/bifrost/api/v1alpha1"
	"github.com/peek8/bifrost/internal/components"
	"github.com/peek8/bifrost/internal/components/factory"
)

const (
	primaryContainer = "main"
	mainPort         = 12345
	mainPortName     = "main"
	containerImage   = "grafana/alloy:v1.11.3"

	configPath  = "/etc/alloy"
	configFile  = "config.alloy"
	storagePath = "/var/lib/alloy/data"
)

type Data struct {
	Name         string
	LogSpaceSpec v1alpha1.LogSpaceSpec
	Namespace    string
}

type Alloy struct {
	config    corev1.ConfigMap
	pvc       corev1.PersistentVolumeClaim
	daemonSet appsv1.DaemonSet
}

func (al Alloy) ToComponents() components.Components {
	items := components.Components{
		&components.ConfigMap{ConfigMap: al.config},
		&components.DaemonSet{DaemonSet: al.daemonSet},
		&components.PersistentVolumeClaim{PersistentVolumeClaim: al.pvc},
	}

	return items.NonEmptyComponents()
}

type Builder struct{}

func (b Builder) New(ctx context.Context, data Data) (Alloy, error) {
	return Alloy{
		daemonSet: daemonSet(data),
	}, nil
}

func daemonSet(data Data) appsv1.DaemonSet {
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
		Args: []string{
			"run",
			fullConfigPath(),
			fmt.Sprintf("--storage.path=%s", storagePath),
			fmt.Sprintf("--server.http.listen-addr=%s", listenAddr()),
		},
	}

	return factory.NewDaemonSet(data.Name, data.Namespace, factory.K8sLabels(data.Name, "collector"), container)
}

func listenAddr() string {
	return fmt.Sprintf("0.0.0.0:%d", mainPort)
}

func fullConfigPath() string {
	return path.Join(configPath, configFile)
}
