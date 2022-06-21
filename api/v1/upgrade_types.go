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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// UpgradeSpec defines the desired state of Upgrade
type UpgradeSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of Upgrade. Edit upgrade_types.go to remove/update
	Release string `json:"release,omitempty"`
}

// UpgradeStatus defines the observed state of Upgrade
type UpgradeStatus struct {
	// Define observed state of cluster
	// The idea here is to have the current Release version of each node present
	// and It mast adhere to the UpgradeSpec Release string mentioned in the
	// UpgradeSpec
	// Important: Run "make" to regenerate code after modifying this file
	Nodes []string `json:"nodes"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Upgrade is the Schema for the upgrades API
type Upgrade struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   UpgradeSpec   `json:"spec,omitempty"`
	Status UpgradeStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// UpgradeList contains a list of Upgrade
type UpgradeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Upgrade `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Upgrade{}, &UpgradeList{})
}
