/*
 * Copyright (c) 2025 peek8.io
 *
 * Created Date: Sunday, November 30th 2025, 5:44:32 pm
 * Author: Md. Asraful Haque
 *
 */

package grafana

import (
	_ "embed"

	"github.com/peek8/bifrost/internal/utils"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	grafanaProvisioningDir           = "/etc/grafana/provisioning"
	grafanaDataSourceProvisioningDir = grafanaProvisioningDir + "/datasources"
	dataSourceFileName               = "datasource.yaml"
	// will be used at grafana dashboard
	lokiDataSourceName = "bifrost-loki"

	grafanaDashboardProvisioningDir = grafanaProvisioningDir + "/dashboards"
	dashboardProviderFileName       = "dashboard-providers.yaml"
	dashboardFileName               = "bf-dashboard.json"
)

//go:embed configs/datasource.yaml
var grafanaDatasource string

//go:embed configs/dashboard-providers.yaml
var grafanaDashboardProvider string

//go:embed configs/bf-dashboard.json
var grafanaDashboard string

func grafanaDSConfigMap(data Data) (corev1.ConfigMap, error) {
	cd := struct {
		Name        string
		LokiService string
	}{
		Name:        lokiDataSourceName,
		LokiService: data.LokiServiceName,
	}
	configStr, err := utils.RenderConfTemplate("grafana-ds", grafanaDatasource, cd)

	if err != nil {
		return corev1.ConfigMap{}, err
	}

	return corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      data.Name + "-ds",
			Namespace: data.Namespace,
		},

		Data: map[string]string{
			dataSourceFileName: configStr,
		},
	}, nil
}

func grafanaDashboardConfigMap(data Data) (corev1.ConfigMap, error) {
	cd := struct {
		Name       string
		Datasource string
	}{
		Name:       data.Name + "-dashboard",
		Datasource: lokiDataSourceName,
	}
	dashboard, err := utils.RenderConfTemplate("grafana-dash", grafanaDashboard, cd)

	if err != nil {
		return corev1.ConfigMap{}, err
	}

	return corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      data.Name,
			Namespace: data.Namespace,
		},

		Data: map[string]string{
			dashboardProviderFileName: grafanaDashboardProvider,
			dashboardFileName:         dashboard,
		},
	}, nil
}
