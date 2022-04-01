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

package v1alpha1

// ImageReviewStatusApplyConfiguration represents an declarative configuration of the ImageReviewStatus type for use
// with apply.
type ImageReviewStatusApplyConfiguration struct {
	Allowed          *bool             `json:"allowed,omitempty"`
	Reason           *string           `json:"reason,omitempty"`
	AuditAnnotations map[string]string `json:"auditAnnotations,omitempty"`
}

// ImageReviewStatusApplyConfiguration constructs an declarative configuration of the ImageReviewStatus type for use with
// apply.
func ImageReviewStatus() *ImageReviewStatusApplyConfiguration {
	return &ImageReviewStatusApplyConfiguration{}
}

// WithAllowed sets the Allowed field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Allowed field is set to the value of the last call.
func (b *ImageReviewStatusApplyConfiguration) WithAllowed(value bool) *ImageReviewStatusApplyConfiguration {
	b.Allowed = &value
	return b
}

// WithReason sets the Reason field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Reason field is set to the value of the last call.
func (b *ImageReviewStatusApplyConfiguration) WithReason(value string) *ImageReviewStatusApplyConfiguration {
	b.Reason = &value
	return b
}

// WithAuditAnnotations puts the entries into the AuditAnnotations field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the AuditAnnotations field,
// overwriting an existing map entries in AuditAnnotations field with the same key.
func (b *ImageReviewStatusApplyConfiguration) WithAuditAnnotations(entries map[string]string) *ImageReviewStatusApplyConfiguration {
	if b.AuditAnnotations == nil && len(entries) > 0 {
		b.AuditAnnotations = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.AuditAnnotations[k] = v
	}
	return b
}
