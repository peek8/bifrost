/*
Copyright 2025.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// LogSpaceSpec defines the desired state of LogSpace.
type LogSpaceSpec struct {
	// Logs will be collected from the pods which are in these Namespaces.
	TargetNamespaces []string `json:"targetNamespaces,omitempty"`
	// General PVC storage settings of the cluster
	PVCStorage *PVCStorageConfig `json:"pvc,omitempty"`

	// Collector controls how Grafana Alloy (or other collector) is deployed.
	Collector *CollectorSpec `json:"collector,omitempty"`

	// Configuration for Loki
	LokiConfig *LokiConfig `json:"loki,omitempty"`
}

type PVCStorageConfig struct {
	// Storage Cluster of the cluster. if you are not sure the default storage class of the cluster,
	// you can get it by the command : $ kubectl get storageclass
	StorageClass string `json:"storageClass,omitempty"`
}

type LokiConfig struct {
	// Loki is a distributed system consisting of many microservices.
	// It also has a unique build model where all of those microservices exist within the same binary.
	// For now only valid value is monolithic

	// +optional
	// +kubebuilder:validation:Enum=monolithic;simpleScalable;microservice
	// +kubebuilder:default:=monolithic
	DeploymentMode *string `json:"deploymentMode,omitempty"`

	// Retention policy for logs, e.g. "7d", "30d"
	// +optional
	// +kubebuilder:default:=7d
	RetentionPeriod *string `json:"retentionPeriod,omitempty"`
	// Storage Configuration for loki
	Storage *StorageConfig `json:"storage,omitempty"`

	// Schema Configuration for loki
	Schema *LokiSchemaConfig `json:"schema,omitempty"`
}

type CollectorSpec struct {
	// Type of collector. e.g. "grafana-alloy", "fluentbit", "fluentd", Now only grafana-alloy is supported
	// +kubebuilder:validation:Enum=grafana-alloy;fluentbit;fluentd
	// +kubebuilder:default:=grafana-alloy
	// +optional
	Type string `json:"type,omitempty"`

	// Replicas for collector controller (if applicable).
	// +kubebuilder:validation:Minimum=1
	// +optional
	Replicas *int32 `json:"replicas,omitempty"`

	// How much storage is needed for collector eg aloy default is 5Gi
	Storage *StorageConfig `json:"storage,omitempty"`

	// PerNamespaceDeployment if true deploys collector per-target-namespace.
	// If false, operator may choose a cluster-level deployment.
	// +optional
	// PerNamespaceDeployment bool `json:"perNamespaceDeployment,omitempty"`
}

type StorageConfig struct {
	// Size of storage to be used in PVC Storage, eg 5Gi
	Size string `json:"size,omitempty"`
}

type LokiSchemaConfig struct {
	// Loki will use this storage schema for all logs from this date.
	// For example, if it is set to 2025-01-01, The schema will be used from 1st Jan, 2025
	// If not set, the current date will be used
	// +optional
	FromDate *string `json:"fromDate,omitempty"`

	// The supported values are `filesystem, s3, gcs, azure` where `filesystem`	store TSDB blocks locally on disk
	// and others store blocks in cloud object storage.
	// For now only filesystem is supported
	// +kubebuilder:validation:Enum=filesystem;s3;gcs;azure
	// +kubebuilder:default:=filesystem
	// +optional
	ObjectStore string `json:"objectStore,omitempty"`
}

// LogSpaceStatus defines the observed state of LogSpace.
type LogSpaceStatus struct {
	// State is a high-level state like "Ready", "Reconciling", "Error"
	// +optional
	State string `json:"state,omitempty"`

	// ObservedNamespaces lists namespaces the operator selected/observed
	// +optional
	ObservedNamespaces []string `json:"observedNamespaces,omitempty"`

	// Dashboard URLs created in Grafana
	// +optional
	// DashboardURLs []string `json:"dashboardUrls,omitempty"`

	GrafanaEndpoint string `json:"grafanaEndpoint,omitempty"`
	LokiEndpoint string `json:"lokiEndpoint,omitempty"`

	// LastError captures last reconciliation error (short message)
	// +optional
	LastError string `json:"lastError,omitempty"`

	// LastUpdated timestamp for status
	// +optional
	LastUpdated *metav1.Time `json:"lastUpdated,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:shortName=ls,scope=Namespaced
// +kubebuilder:printcolumn:name="State",type=string,JSONPath=`.status.state`
// +kubebuilder:printcolumn:name="Namespaces",type=string,JSONPath=`.status.observedNamespaces`
// +kubebuilder:printcolumn:name="Age",type=date,JSONPath=`.metadata.creationTimestamp`

// LogSpace is the Schema for the logspaces API.
type LogSpace struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   LogSpaceSpec   `json:"spec,omitempty"`
	Status LogSpaceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LogSpaceList contains a list of LogSpace.
type LogSpaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LogSpace `json:"items"`
}

func init() {
	SchemeBuilder.Register(&LogSpace{}, &LogSpaceList{})
}
