package alloy

import (
	"context"
	"fmt"
	"path"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	"k8s.io/utils/ptr"

	"github.com/peek8/bifrost/api/v1alpha1"
	"github.com/peek8/bifrost/internal/components"
	"github.com/peek8/bifrost/internal/components/factory"
	"github.com/peek8/bifrost/internal/utils"
)

const (
	primaryContainer = "main"
	mainPort         = 12345
	mainPortName     = "main"
	containerImage   = "grafana/alloy:v1.11.3"

	configPath  = "/etc/alloy"
	configFile  = "config.alloy"
	storagePath = "/var/lib/alloy/data"
	storageSize = "5Gi"
)

type Data struct {
	Name            string
	LogSpaceSpec    v1alpha1.LogSpaceSpec
	Namespace       string
	LokiServiceName string
}

type Alloy struct {
	config    corev1.ConfigMap
	pvc       corev1.PersistentVolumeClaim
	daemonSet appsv1.DaemonSet
	//rbac
	serviceAccount corev1.ServiceAccount
	role           rbacv1.ClusterRole
	roleBinding    rbacv1.ClusterRoleBinding
}

func (al Alloy) ToComponents() components.Components {
	items := components.Components{
		&components.ConfigMap{ConfigMap: al.config},
		&components.PersistentVolumeClaim{PersistentVolumeClaim: al.pvc},
		&components.ServiceAccount{ServiceAccount: al.serviceAccount},
		&components.ClusterRole{ClusterRole: al.role},
		&components.ClusterRoleBinding{ClusterRoleBinding: al.roleBinding},
		&components.DaemonSet{DaemonSet: al.daemonSet},
	}

	return items.NonEmptyComponents()
}

type Builder struct{}

func (b Builder) New(ctx context.Context, data Data) (Alloy, error) {
	role := *clusterRole(data)
	sa := serviceAccount(data)

	cd := ConfigData{
		LokiService: data.LokiServiceName,
		LokiPort:    3100,
		Namespaces:  data.LogSpaceSpec.TargetNamespaces,
		ClusterName: "default",
	}

	cm, err := AlloyConfigMap(data, cd)
	if err != nil {
		return Alloy{}, err
	}

	alloy := Alloy{
		daemonSet:      daemonSet(data),
		config:         cm,
		pvc:            factory.NewPVC(data.Name, data.Namespace, data.LogSpaceSpec.Collector.Storage.Size, data.LogSpaceSpec.PVCStorage.StorageClass),
		serviceAccount: sa,
		role:           role,
		roleBinding:    clusterRoleBinding(data, role, sa),
	}

	options := []utils.Option[corev1.PodTemplateSpec]{
		utils.AddConfigMapAsVolume{
			ConfigMapName: alloy.config.Name,
			VolumeName:    "alloy-config",
			MountPath:     configPath,
			ContainerName: primaryContainer,
			Mode:          ptr.To(int32(0755)),
		},
		utils.AddPVC{
			PVCName:       alloy.pvc.Name,
			MountPath:     storagePath,
			ContainerName: primaryContainer,
		},
	}

	err = utils.ApplyAll(&alloy.daemonSet.Spec.Template, options...)

	if err != nil {
		return Alloy{}, err
	}

	return alloy, nil
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

	return *factory.NewDaemonSet(data.Name, data.Namespace, container).
		WithLabels(factory.K8sLabels(data.Name, "collector")).
		WithServiceAccount(data.Name).
		Get()
}

func serviceAccount(d Data) corev1.ServiceAccount {
	return factory.NewServiceAccount(d.Name, d.Namespace)
}

func clusterRole(d Data) *rbacv1.ClusterRole {
	return factory.NewClusterRole(d.Name, d.Namespace).
		WithRules("", []string{"pods", "namespaces"}, []string{"get", "list", "watch"}).
		WithRules("", []string{"pods/log"}, []string{"get", "list", "watch"}).
		Get()
}

func clusterRoleBinding(d Data, role rbacv1.ClusterRole, sa corev1.ServiceAccount) rbacv1.ClusterRoleBinding {
	return factory.ClusterRoleBinding(d.Name, role, sa)
}

func listenAddr() string {
	return fmt.Sprintf("0.0.0.0:%d", mainPort)
}

func fullConfigPath() string {
	return path.Join(configPath, configFile)
}
