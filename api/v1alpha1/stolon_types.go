/*
Copyright 2022.

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

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// StolonSpec defines the desired state of Stolon
type StolonStore struct {
	// stolon store options for specs

	Backend   string `json:"backend"`
	Endpoints string `json:"endpoints"`
}
type StolonProxy struct {
	// stolon proxy options for specs

	Replicas    int    `json:"replicas"`
	ServiceType string `json:"serviceType"`
}
type StolonKeeper struct {
	// stolon keeper options for specs

	Replicas    int    `json:"replicas"`
	ServiceType string `json:"serviceType"`
	UidPrefix   string `json:"uid_prefix"`
}
type StolonClientSSL struct {
	// stolon client ssl options for specs

	Enabled         bool   `json:"enabled"`
	CertsSecretName string `json:"certs_secret_name"`
}
type StolonComputeResource struct {
	CPU    string `json:"cpu"`
	Memory string `json:"memory"`
}
type StolonResources struct {
	// stolon resources options for specs

	Requests StolonComputeResource `json:"requests"`
	Limits   StolonComputeResource `json:"limits"`
}
type StolonPersistence struct {
	// stolon resources options for specs

	Enabled          bool   `json:"enabled"`
	AccessMode       string `json:"accessMode"`
	StorageClassName string `json:"storageClassName"`
	Size             string `json:"size"`
}
type StolonRBAC struct {
	// stolon RBAC options for specs

	Create bool `json:"create"`
}
type StolonServiceAccount struct {
	// stolon RBAC options for specs

	Create bool   `json:"create"`
	Name   string `json:"name"`
}
type StolonPorts struct {
	// stolon RBAC options for specs

	InternalPort int `json:"internalPort"`
	ExternalPort int `json:"externalPort"`
}
type StolonSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	Image                       string               `json:"image"`
	ImageTag                    string               `json:"imageTag"`
	ImagePullPolicy             string               `json:"imagePullPolicy"`
	ClusterName                 string               `json:"clusterName"`
	Debug                       bool                 `json:"debug"`
	Store                       StolonStore          `json:"store"`
	PostgresReplicationUsername string               `json:"pgReplUsername"`
	PostgresReplicationPassword string               `json:"pgReplPassword"`
	PostgresSuperUserUsername   string               `json:"pgSuperuserName"`
	PostgresSuperUserPassword   string               `json:"pgSuperuserPassword"`
	Proxy                       StolonProxy          `json:"proxy"`
	ClientSSL                   StolonClientSSL      `json:"client_ssl"`
	Resources                   StolonResources      `json:"resources"`
	Persistence                 StolonPersistence    `json:"persistence"`
	RBAC                        StolonRBAC           `json:"rbac"`
	ServiceAccount              StolonServiceAccount `json:"serviceAccount"`
	Ports                       StolonPorts          `json:"ports"`
}

// StolonStatus defines the observed state of Stolon
type StolonStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Stolon is the Schema for the stolons API
type Stolon struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   StolonSpec   `json:"spec,omitempty"`
	Status StolonStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// StolonList contains a list of Stolon
type StolonList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Stolon `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Stolon{}, &StolonList{})
}
