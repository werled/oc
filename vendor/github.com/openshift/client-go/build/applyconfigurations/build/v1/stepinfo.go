// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/openshift/api/build/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// StepInfoApplyConfiguration represents an declarative configuration of the StepInfo type for use
// with apply.
type StepInfoApplyConfiguration struct {
	Name                 *v1.StepName `json:"name,omitempty"`
	StartTime            *metav1.Time `json:"startTime,omitempty"`
	DurationMilliseconds *int64       `json:"durationMilliseconds,omitempty"`
}

// StepInfoApplyConfiguration constructs an declarative configuration of the StepInfo type for use with
// apply.
func StepInfo() *StepInfoApplyConfiguration {
	return &StepInfoApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *StepInfoApplyConfiguration) WithName(value v1.StepName) *StepInfoApplyConfiguration {
	b.Name = &value
	return b
}

// WithStartTime sets the StartTime field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the StartTime field is set to the value of the last call.
func (b *StepInfoApplyConfiguration) WithStartTime(value metav1.Time) *StepInfoApplyConfiguration {
	b.StartTime = &value
	return b
}

// WithDurationMilliseconds sets the DurationMilliseconds field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DurationMilliseconds field is set to the value of the last call.
func (b *StepInfoApplyConfiguration) WithDurationMilliseconds(value int64) *StepInfoApplyConfiguration {
	b.DurationMilliseconds = &value
	return b
}