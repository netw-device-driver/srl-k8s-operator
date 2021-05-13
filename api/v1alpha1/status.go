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
	// Values identifies the object value that should match the leafref value
	// if empty it means the object does not exist.
	Values *[]string `json:"values,omitempty"`

	// LeafRefPath points to the leafref object value
	LeafRefPath *string `json:"leafRefParh,omitempty"`

	// LeafRefValues identifies the values assigned in the leaf ref
	LeafRefValues *[]string `json:"leafRefValues,omitempty"`
}

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
