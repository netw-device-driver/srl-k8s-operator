/*
Copyright 2020 Wim Henderickx.

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

// TargetFoundStatus defines the status of the resource target
type TargetFoundStatus string

const (
	// TargetFoundStatusSuccess means the target was found
	TargetFoundStatusSuccess TargetFoundStatus = "Success"

	// TargetFoundStatusFailed means the target was not found
	TargetFoundStatusFailed TargetFoundStatus = "Failed"
)

func TargetFoundStatusPtr(s TargetFoundStatus) *TargetFoundStatus { return &s }

// ValidationStatus defines the validation status of the resource object
type ValidationStatus string

const (
	// ValidationStatusSuccess means the validation was successfull
	ValidationStatusSuccess ValidationStatus = "Success"

	// ValidationStatusSuccess means the validation was successfull
	ValidationStatusFailed ValidationStatus = "Failed"
)

func ValidationStatusPtr(s ValidationStatus) *ValidationStatus { return &s }

// ValidationDetails provides the status of the configuration applied on this particular device
type ValidationDetails struct {
	// LocalLeafRef identifies the leafref value that should match the remote leafref
	// if empty it means the object does not exist.
	LocalLeafRefs *[]string `json:"localLeafRef,omitempty"`

	// RemoteLeafRef points to the remote leafref object
	RemoteLeafRefs *[]string `json:"remoteLeafRef,omitempty"`
}

// ValidationDetails2 provides the status of the configuration applied on this particular device
type ValidationDetails2 struct {
	// LocalResolvedLeafRefInfo provides the status of the remote leafref information
	LocalResolvedLeafRefInfo map[string]*RemoteLeafRefInfo `json:"localResolvedLeafRefInfo,omitempty"`
}

type RemoteLeafRefInfo struct {
	// RemoteLeafRef provides the reference path to the remote leafref
	RemoteLeafRef *string `json:"remoteLeafRef"`
	// DependencyCheck validates if the remote leafref is present or not
	// +kubebuilder:validation:Enum=Success;Failed
	DependencyCheck *DependencyCheck `json:"dependencyCheck"`
}

// DependencyCheck defines the status of the leafref dependency
type DependencyCheck string

const (
	// DependencyCheckSuccess means the leafref depdendency was not found
	DependencyCheckSuccess DependencyCheck = "Success"

	// DependencyCheckFailed means the leafref depdendency was not found
	DependencyCheckFailed DependencyCheck = "Failed"
)

func DependencyCheckPtr(c DependencyCheck) *DependencyCheck { return &c }

// ConfigStatus defines the states the resource object is reporting
type ConfigStatus string

const (
	// ConfigStatusNone means the object state is unknown
	ConfigStatusNone ConfigStatus = ""

	// ConfigStatusDeleting means the object gets deleted
	ConfigStatusDeleting ConfigStatus = "Deleting"

	// ConfigStatusDeleteFailed means the object delete failed
	ConfigStatusDeleteFailed ConfigStatus = "DeleteFailed"

	// ConfigStatusDeleteSuccess means the object delete succeeded
	ConfigStatusDeleteSuccess ConfigStatus = "DeleteSuccess"

	// ConfigStatusConfiguring means the object gets configured
	ConfigStatusConfiguring ConfigStatus = "Configuring"

	// ConfigStatusConfigureSuccess means the object configuration succeeded
	ConfigStatusConfigureSuccess ConfigStatus = "ConfiguredSuccess"

	// ConfigStatusConfigureFailed means the object configuration failed
	ConfigStatusConfigureFailed ConfigStatus = "ConfigStatusConfigureFailed"
)

func (c *ConfigStatus) String() string {
	switch *c {
	case ConfigStatusNone:
		return ""
	case ConfigStatusDeleting:
		return "Deleting"
	case ConfigStatusDeleteFailed:
		return "DeleteFailed"
	case ConfigStatusDeleteSuccess:
		return "DeleteSuccess"
	case ConfigStatusConfiguring:
		return "Configuring"
	case ConfigStatusConfigureSuccess:
		return "ConfiguredSuccess"
	case ConfigStatusConfigureFailed:
		return "ConfigStatusConfigureFailed"
	}
	return ""
}

func ConfigStatusPtr(c ConfigStatus) *ConfigStatus { return &c }

// TargetStatus provides the status of the configuration applied on this particular device
type TargetStatus struct {
	// +kubebuilder:default:=""
	// +kubebuilder:validation:Enum="";Deleting;DeleteFailed;DeleteSuccess;Configuring;ConfiguredSuccess;ConfigStatusConfigureFailed
	ConfigStatus *ConfigStatus `json:"configStatus"`
	// +kubebuilder:default:=""
	ConfigStatusDetails *string `json:"configStatusDetails,omitempty"`
	// ErrorCount records how many times the host has encoutered an error since the last successful operation
	// +kubebuilder:default:=0
	ErrorCount *int `json:"errorCount"`
}
