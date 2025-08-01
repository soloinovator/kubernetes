/*
Copyright The Kubernetes Authors.

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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// DeviceClaimConfigurationApplyConfiguration represents a declarative configuration of the DeviceClaimConfiguration type for use
// with apply.
type DeviceClaimConfigurationApplyConfiguration struct {
	Requests                              []string `json:"requests,omitempty"`
	DeviceConfigurationApplyConfiguration `json:",inline"`
}

// DeviceClaimConfigurationApplyConfiguration constructs a declarative configuration of the DeviceClaimConfiguration type for use with
// apply.
func DeviceClaimConfiguration() *DeviceClaimConfigurationApplyConfiguration {
	return &DeviceClaimConfigurationApplyConfiguration{}
}

// WithRequests adds the given value to the Requests field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Requests field.
func (b *DeviceClaimConfigurationApplyConfiguration) WithRequests(values ...string) *DeviceClaimConfigurationApplyConfiguration {
	for i := range values {
		b.Requests = append(b.Requests, values[i])
	}
	return b
}

// WithOpaque sets the Opaque field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Opaque field is set to the value of the last call.
func (b *DeviceClaimConfigurationApplyConfiguration) WithOpaque(value *OpaqueDeviceConfigurationApplyConfiguration) *DeviceClaimConfigurationApplyConfiguration {
	b.DeviceConfigurationApplyConfiguration.Opaque = value
	return b
}
