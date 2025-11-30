/*
 * Copyright (c) 2025 peek8.io
 *
 * Created Date: Friday, November 28th 2025, 8:17:35 pm
 * Author: Md. Asraful Haque
 *
 */

package loki

import (
	"bytes"
	_ "embed"
	"text/template"
	"time"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

//go:embed loki-config.yaml
var lokiConfig string

type ConfigData struct {
	FromDate        string
	RetentionPeriod string
}

func lokiConfigMap(data Data) (corev1.ConfigMap, error) {
	var fromDate string
	if data.LogSpaceSpec.LokiConfig.Schema != nil &&
		data.LogSpaceSpec.LokiConfig.Schema.FromDate != nil {
		fromDate = *data.LogSpaceSpec.LokiConfig.Schema.FromDate
	} else {
		currentTime := time.Now()
		fromDate = currentTime.Format(time.DateOnly)
	}

	cd := ConfigData{
		FromDate:        fromDate,
		RetentionPeriod: *data.LogSpaceSpec.LokiConfig.RetentionPeriod,
	}
	configStr, err := generateConfig(cd)

	if err != nil {
		return corev1.ConfigMap{}, err
	}

	return corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      data.Name,
			Namespace: data.Namespace,
		},
		BinaryData: map[string][]byte{
			configFile: []byte(configStr),
		},
	}, nil
}

func generateConfig(cd ConfigData) (string, error) {
	tpl, err := template.New("loki").Parse(string(lokiConfig))
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	err = tpl.Execute(&buf, cd)

	if err != nil {
		return "", err
	}

	return buf.String(), nil

}
