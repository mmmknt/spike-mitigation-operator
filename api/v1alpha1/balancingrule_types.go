/*


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

// BalancingRuleSpec defines the desired state of BalancingRule
type BalancingRuleSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of BalancingRule. Edit BalancingRule_types.go to remove/update
	GatewayName              string                `json:"gatewayName,omitempty"`
	ExternalHost             string                `json:"externalHost,omitempty"`
	ExternalAuthorizationRef SecretRef             `json:"externalAuthorizationRef,omitempty"`
	InternalHost             string                `json:"internalHost,omitempty"`
	HostInfoHeaderKeyRef     SecretRef             `json:"hostInfoHeaderKeyRef,omitempty"`
	HPATriggerRate           int32                 `json:"hpaTriggerRate,omitempty"`
	MonitoredHPANamespace    string                `json:"monitoredHpaNamespace,omitempty"`
	MonitoredHPANames        []string              `json:"monitoredHpaNames,omitempty"`
	MitigationTriggerRate    int32                 `json:"mitigationTriggerRate,omitempty"`
	MetricsStoreSecretRef    MetricsStoreSecretRef `json:"metricsStoreSecretRef,omitempty"`
	MetricsCondition         MetricsCondition      `json:"metricsCondition,omitempty"`
	SecretNamespace          string                `json:"secretNamespace,omitempty"`
	OptionalAuthorization    OptionalAuthorization `json:"optionalAuthorization,omitempty"`
}

type OptionalAuthorization struct {
	KeyRef   SecretRef `json:"keyRef,omitempty"`
	ValueRef SecretRef `json:"valueRef,omitempty"`
}

type MetricsCondition struct {
	MetricsName string `json:"metricsName,omitempty"`
	ClusterName string `json:"clusterName,omitempty"`
}

type MetricsStoreSecretRef struct {
	DDApiKeyRef SecretRef `json:"ddApiKeyRef,omitempty"`
	DDAppKeyRef SecretRef `json:"ddAppKeyRef,omitempty"`
}

type SecretRef struct {
	Name string `json:"name,omitempty"`
	Key  string `json:"key,omitempty"`
}

// BalancingRuleStatus defines the observed state of BalancingRule
type BalancingRuleStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// BalancingRule is the Schema for the balancingrules API
type BalancingRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BalancingRuleSpec   `json:"spec,omitempty"`
	Status BalancingRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BalancingRuleList contains a list of BalancingRule
type BalancingRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BalancingRule `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BalancingRule{}, &BalancingRuleList{})
}
