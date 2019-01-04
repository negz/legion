// +build !ignore_autogenerated

/*
Copyright 2018 Planet Labs Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
implied. See the License for the specific language governing permissions
and limitations under the License.
*/

// Code generated by main. DO NOT EDIT.

package kubernetes

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodMutation) DeepCopyInto(out *PodMutation) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodMutation.
func (in *PodMutation) DeepCopy() *PodMutation {
	if in == nil {
		return nil
	}
	out := new(PodMutation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PodMutation) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodMutationSpec) DeepCopyInto(out *PodMutationSpec) {
	*out = *in
	out.Strategy = in.Strategy
	in.Template.DeepCopyInto(&out.Template)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodMutationSpec.
func (in *PodMutationSpec) DeepCopy() *PodMutationSpec {
	if in == nil {
		return nil
	}
	out := new(PodMutationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodMutationStrategy) DeepCopyInto(out *PodMutationStrategy) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodMutationStrategy.
func (in *PodMutationStrategy) DeepCopy() *PodMutationStrategy {
	if in == nil {
		return nil
	}
	out := new(PodMutationStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodMutationTemplate) DeepCopyInto(out *PodMutationTemplate) {
	*out = *in
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodMutationTemplate.
func (in *PodMutationTemplate) DeepCopy() *PodMutationTemplate {
	if in == nil {
		return nil
	}
	out := new(PodMutationTemplate)
	in.DeepCopyInto(out)
	return out
}
