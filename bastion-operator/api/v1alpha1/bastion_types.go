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

// BastionSpec defines the desired state of Bastion
type BastionSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// +optional
	// +kubebuilder:validation:Minimum=0
	ServingPort int32 `json:"servingPort,omitempty"`

	// +optional
	// +kubebuilder:validation:Minimum=0
	Replicas *int32 `json:"replicas,omitempty"`
}

// BastionStatus defines the observed state of Bastion
type BastionStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Information when was the last time the keys were successfully updated.
	// +optional
	LastUpdateTime *metav1.Time `json:"lastUpdateTime,omitempty"`
}

// +kubebuilder:object:root=true

// Bastion is the Schema for the bastions API
type Bastion struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BastionSpec   `json:"spec,omitempty"`
	Status BastionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BastionList contains a list of Bastion
type BastionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Bastion `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Bastion{}, &BastionList{})
}
