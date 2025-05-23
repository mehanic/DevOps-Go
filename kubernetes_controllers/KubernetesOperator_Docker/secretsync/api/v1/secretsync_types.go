package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type SecretSyncSpec struct {
	SourceNamespace       string   `json:"sourceNamespace"`
	DestinationNamespaces []string `json:"destinationNamespaces"`
	SecretName            string   `json:"secretName"`
}

// SecretSyncStatus defines the observed state of SecretSync
type SecretSyncStatus struct {
	LastSyncTime metav1.Time `json:"lastSyncTime"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// SecretSync is the Schema for the secretsyncs API
type SecretSync struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SecretSyncSpec   `json:"spec,omitempty"`
	Status SecretSyncStatus `json:"status,omitempty"`
}

//+kubebuilder:rbac:groups=core,resources=secrets,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:object:root=true

// SecretSyncList contains a list of SecretSync
type SecretSyncList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecretSync `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SecretSync{}, &SecretSyncList{})
}
