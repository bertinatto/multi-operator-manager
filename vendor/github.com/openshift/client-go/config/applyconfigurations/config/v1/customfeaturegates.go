// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/openshift/api/config/v1"
)

// CustomFeatureGatesApplyConfiguration represents a declarative configuration of the CustomFeatureGates type for use
// with apply.
type CustomFeatureGatesApplyConfiguration struct {
	Enabled  []v1.FeatureGateName `json:"enabled,omitempty"`
	Disabled []v1.FeatureGateName `json:"disabled,omitempty"`
}

// CustomFeatureGatesApplyConfiguration constructs a declarative configuration of the CustomFeatureGates type for use with
// apply.
func CustomFeatureGates() *CustomFeatureGatesApplyConfiguration {
	return &CustomFeatureGatesApplyConfiguration{}
}

// WithEnabled adds the given value to the Enabled field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Enabled field.
func (b *CustomFeatureGatesApplyConfiguration) WithEnabled(values ...v1.FeatureGateName) *CustomFeatureGatesApplyConfiguration {
	for i := range values {
		b.Enabled = append(b.Enabled, values[i])
	}
	return b
}

// WithDisabled adds the given value to the Disabled field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Disabled field.
func (b *CustomFeatureGatesApplyConfiguration) WithDisabled(values ...v1.FeatureGateName) *CustomFeatureGatesApplyConfiguration {
	for i := range values {
		b.Disabled = append(b.Disabled, values[i])
	}
	return b
}