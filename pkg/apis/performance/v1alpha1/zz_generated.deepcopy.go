// +build !ignore_autogenerated

// Code generated by operator-sdk-v0.13.0-x86_64-linux-gnu. DO NOT EDIT.

package v1alpha1

import (
	v1 "github.com/openshift/custom-resource-status/conditions/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CPU) DeepCopyInto(out *CPU) {
	*out = *in
	if in.Reserved != nil {
		in, out := &in.Reserved, &out.Reserved
		*out = new(CPUSet)
		**out = **in
	}
	if in.Isolated != nil {
		in, out := &in.Isolated, &out.Isolated
		*out = new(CPUSet)
		**out = **in
	}
	if in.NonIsolated != nil {
		in, out := &in.NonIsolated, &out.NonIsolated
		*out = new(CPUSet)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CPU.
func (in *CPU) DeepCopy() *CPU {
	if in == nil {
		return nil
	}
	out := new(CPU)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HugePage) DeepCopyInto(out *HugePage) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HugePage.
func (in *HugePage) DeepCopy() *HugePage {
	if in == nil {
		return nil
	}
	out := new(HugePage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HugePages) DeepCopyInto(out *HugePages) {
	*out = *in
	if in.DefaultHugePagesSize != nil {
		in, out := &in.DefaultHugePagesSize, &out.DefaultHugePagesSize
		*out = new(HugePageSize)
		**out = **in
	}
	if in.Pages != nil {
		in, out := &in.Pages, &out.Pages
		*out = make([]HugePage, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HugePages.
func (in *HugePages) DeepCopy() *HugePages {
	if in == nil {
		return nil
	}
	out := new(HugePages)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PerformanceProfile) DeepCopyInto(out *PerformanceProfile) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PerformanceProfile.
func (in *PerformanceProfile) DeepCopy() *PerformanceProfile {
	if in == nil {
		return nil
	}
	out := new(PerformanceProfile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PerformanceProfile) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PerformanceProfileList) DeepCopyInto(out *PerformanceProfileList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PerformanceProfile, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PerformanceProfileList.
func (in *PerformanceProfileList) DeepCopy() *PerformanceProfileList {
	if in == nil {
		return nil
	}
	out := new(PerformanceProfileList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PerformanceProfileList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PerformanceProfileSpec) DeepCopyInto(out *PerformanceProfileSpec) {
	*out = *in
	if in.CPU != nil {
		in, out := &in.CPU, &out.CPU
		*out = new(CPU)
		(*in).DeepCopyInto(*out)
	}
	if in.HugePages != nil {
		in, out := &in.HugePages, &out.HugePages
		*out = new(HugePages)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.RealTimeKernel != nil {
		in, out := &in.RealTimeKernel, &out.RealTimeKernel
		*out = new(RealTimeKernel)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PerformanceProfileSpec.
func (in *PerformanceProfileSpec) DeepCopy() *PerformanceProfileSpec {
	if in == nil {
		return nil
	}
	out := new(PerformanceProfileSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PerformanceProfileStatus) DeepCopyInto(out *PerformanceProfileStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PerformanceProfileStatus.
func (in *PerformanceProfileStatus) DeepCopy() *PerformanceProfileStatus {
	if in == nil {
		return nil
	}
	out := new(PerformanceProfileStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RealTimeKernel) DeepCopyInto(out *RealTimeKernel) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RealTimeKernel.
func (in *RealTimeKernel) DeepCopy() *RealTimeKernel {
	if in == nil {
		return nil
	}
	out := new(RealTimeKernel)
	in.DeepCopyInto(out)
	return out
}