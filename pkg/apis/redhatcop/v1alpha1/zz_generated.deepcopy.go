// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeepalivedGroup) DeepCopyInto(out *KeepalivedGroup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeepalivedGroup.
func (in *KeepalivedGroup) DeepCopy() *KeepalivedGroup {
	if in == nil {
		return nil
	}
	out := new(KeepalivedGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KeepalivedGroup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeepalivedGroupList) DeepCopyInto(out *KeepalivedGroupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KeepalivedGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeepalivedGroupList.
func (in *KeepalivedGroupList) DeepCopy() *KeepalivedGroupList {
	if in == nil {
		return nil
	}
	out := new(KeepalivedGroupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KeepalivedGroupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeepalivedGroupSpec) DeepCopyInto(out *KeepalivedGroupSpec) {
	*out = *in
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeepalivedGroupSpec.
func (in *KeepalivedGroupSpec) DeepCopy() *KeepalivedGroupSpec {
	if in == nil {
		return nil
	}
	out := new(KeepalivedGroupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeepalivedGroupStatus) DeepCopyInto(out *KeepalivedGroupStatus) {
	*out = *in
	in.ReconcileStatus.DeepCopyInto(&out.ReconcileStatus)
	if in.RouterIDs != nil {
		in, out := &in.RouterIDs, &out.RouterIDs
		*out = make(map[string]int, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeepalivedGroupStatus.
func (in *KeepalivedGroupStatus) DeepCopy() *KeepalivedGroupStatus {
	if in == nil {
		return nil
	}
	out := new(KeepalivedGroupStatus)
	in.DeepCopyInto(out)
	return out
}
