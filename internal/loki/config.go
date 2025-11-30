/*
 * Copyright (c) 2025 peek8.io
 *
 * Created Date: Friday, November 28th 2025, 8:17:35 pm
 * Author: Md. Asraful Haque
 *
 */

package loki

import (
	_ "embed"
	"time"

	"github.com/peek8/bifrost/internal/utils"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

//go:embed loki-config.yaml
var lokiConfig string

func lokiConfigMap(data Data) (corev1.ConfigMap, error) {
	cd := struct {
		StoragePath     string
		FromDate        string
		RetentionPeriod string
	}{
		StoragePath:     storagePath,
		FromDate:        extractFromDate(data),
		RetentionPeriod: *data.LogSpaceSpec.LokiConfig.RetentionPeriod,
	}

	configStr, err := utils.RenderConfTemplate("loki-conf", lokiConfig, cd)

	if err != nil {
		return corev1.ConfigMap{}, err
	}

	return corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      data.Name,
			Namespace: data.Namespace,
		},

		Data: map[string]string{
			configFile: configStr,
		},
	}, nil
}

func extractFromDate(data Data) string {
	if data.LogSpaceSpec.LokiConfig.Schema != nil &&
		data.LogSpaceSpec.LokiConfig.Schema.FromDate != nil {
		return *data.LogSpaceSpec.LokiConfig.Schema.FromDate
	}

	currentTime := time.Now()
	return currentTime.Format(time.DateOnly)
}
