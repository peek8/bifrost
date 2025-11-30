// Package alloy contains all the workloads and cofig for alloy
package alloy

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"text/template"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

//go:embed config.alloy
var alloyConfig string

type ConfigData struct {
	LokiService string
	LokiPort    int
	Namespaces  []string
	ClusterName string
}

func AlloyConfigMap(data Data, cd ConfigData) (corev1.ConfigMap, error) {
	configStr, err := generateConfig(cd)

	if err != nil {
		return corev1.ConfigMap{}, err
	}

	return corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      data.Name,
			Namespace: data.Namespace,
		},
		Data: map[string]string{
			"config.alloy": configStr,
		},
	}, nil
}

func generateConfig(cd ConfigData) (string, error) {
	tpl, err := template.New("alloy").Funcs(funcMap).Parse(string(alloyConfig))
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

var funcMap = template.FuncMap{
	"jsonArray": jsonArray,
}

func jsonArray(items []string) string {
	out, _ := json.Marshal(items)

	return string(out)
}
