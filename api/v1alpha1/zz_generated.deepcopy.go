// +build !ignore_autogenerated

/*

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/intstr"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AwsUpgradeStrategy) DeepCopyInto(out *AwsUpgradeStrategy) {
	*out = *in
	if in.CRDType != nil {
		in, out := &in.CRDType, &out.CRDType
		*out = new(CRDUpgradeStrategy)
		**out = **in
	}
	if in.RollingUpdateType != nil {
		in, out := &in.RollingUpdateType, &out.RollingUpdateType
		*out = new(RollingUpdateStrategy)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AwsUpgradeStrategy.
func (in *AwsUpgradeStrategy) DeepCopy() *AwsUpgradeStrategy {
	if in == nil {
		return nil
	}
	out := new(AwsUpgradeStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CRDUpgradeStrategy) DeepCopyInto(out *CRDUpgradeStrategy) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CRDUpgradeStrategy.
func (in *CRDUpgradeStrategy) DeepCopy() *CRDUpgradeStrategy {
	if in == nil {
		return nil
	}
	out := new(CRDUpgradeStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EKSCFConfiguration) DeepCopyInto(out *EKSCFConfiguration) {
	*out = *in
	if in.NodeSecurityGroups != nil {
		in, out := &in.NodeSecurityGroups, &out.NodeSecurityGroups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Subnets != nil {
		in, out := &in.Subnets, &out.Subnets
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]map[string]string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = make(map[string]string, len(*in))
				for key, val := range *in {
					(*out)[key] = val
				}
			}
		}
	}
	if in.ManagedPolicies != nil {
		in, out := &in.ManagedPolicies, &out.ManagedPolicies
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.MetricsCollection != nil {
		in, out := &in.MetricsCollection, &out.MetricsCollection
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EKSCFConfiguration.
func (in *EKSCFConfiguration) DeepCopy() *EKSCFConfiguration {
	if in == nil {
		return nil
	}
	out := new(EKSCFConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EKSCFSpec) DeepCopyInto(out *EKSCFSpec) {
	*out = *in
	if in.EKSCFConfiguration != nil {
		in, out := &in.EKSCFConfiguration, &out.EKSCFConfiguration
		*out = new(EKSCFConfiguration)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EKSCFSpec.
func (in *EKSCFSpec) DeepCopy() *EKSCFSpec {
	if in == nil {
		return nil
	}
	out := new(EKSCFSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EKSConfiguration) DeepCopyInto(out *EKSConfiguration) {
	*out = *in
	if in.NodeSecurityGroups != nil {
		in, out := &in.NodeSecurityGroups, &out.NodeSecurityGroups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Volumes != nil {
		in, out := &in.Volumes, &out.Volumes
		*out = make([]NodeVolume, len(*in))
		copy(*out, *in)
	}
	if in.Subnets != nil {
		in, out := &in.Subnets, &out.Subnets
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]map[string]string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = make(map[string]string, len(*in))
				for key, val := range *in {
					(*out)[key] = val
				}
			}
		}
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Taints != nil {
		in, out := &in.Taints, &out.Taints
		*out = make([]v1.Taint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ManagedPolicies != nil {
		in, out := &in.ManagedPolicies, &out.ManagedPolicies
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EKSConfiguration.
func (in *EKSConfiguration) DeepCopy() *EKSConfiguration {
	if in == nil {
		return nil
	}
	out := new(EKSConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EKSManagedConfiguration) DeepCopyInto(out *EKSManagedConfiguration) {
	*out = *in
	if in.NodeLabels != nil {
		in, out := &in.NodeLabels, &out.NodeLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.NodeSecurityGroups != nil {
		in, out := &in.NodeSecurityGroups, &out.NodeSecurityGroups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]map[string]string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = make(map[string]string, len(*in))
				for key, val := range *in {
					(*out)[key] = val
				}
			}
		}
	}
	if in.Subnets != nil {
		in, out := &in.Subnets, &out.Subnets
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EKSManagedConfiguration.
func (in *EKSManagedConfiguration) DeepCopy() *EKSManagedConfiguration {
	if in == nil {
		return nil
	}
	out := new(EKSManagedConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EKSManagedSpec) DeepCopyInto(out *EKSManagedSpec) {
	*out = *in
	if in.EKSManagedConfiguration != nil {
		in, out := &in.EKSManagedConfiguration, &out.EKSManagedConfiguration
		*out = new(EKSManagedConfiguration)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EKSManagedSpec.
func (in *EKSManagedSpec) DeepCopy() *EKSManagedSpec {
	if in == nil {
		return nil
	}
	out := new(EKSManagedSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EKSSpec) DeepCopyInto(out *EKSSpec) {
	*out = *in
	if in.EKSConfiguration != nil {
		in, out := &in.EKSConfiguration, &out.EKSConfiguration
		*out = new(EKSConfiguration)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EKSSpec.
func (in *EKSSpec) DeepCopy() *EKSSpec {
	if in == nil {
		return nil
	}
	out := new(EKSSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceGroup) DeepCopyInto(out *InstanceGroup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceGroup.
func (in *InstanceGroup) DeepCopy() *InstanceGroup {
	if in == nil {
		return nil
	}
	out := new(InstanceGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InstanceGroup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceGroupCondition) DeepCopyInto(out *InstanceGroupCondition) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceGroupCondition.
func (in *InstanceGroupCondition) DeepCopy() *InstanceGroupCondition {
	if in == nil {
		return nil
	}
	out := new(InstanceGroupCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceGroupList) DeepCopyInto(out *InstanceGroupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]InstanceGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceGroupList.
func (in *InstanceGroupList) DeepCopy() *InstanceGroupList {
	if in == nil {
		return nil
	}
	out := new(InstanceGroupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InstanceGroupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceGroupSpec) DeepCopyInto(out *InstanceGroupSpec) {
	*out = *in
	if in.EKSCFSpec != nil {
		in, out := &in.EKSCFSpec, &out.EKSCFSpec
		*out = new(EKSCFSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.EKSManagedSpec != nil {
		in, out := &in.EKSManagedSpec, &out.EKSManagedSpec
		*out = new(EKSManagedSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.EKSSpec != nil {
		in, out := &in.EKSSpec, &out.EKSSpec
		*out = new(EKSSpec)
		(*in).DeepCopyInto(*out)
	}
	in.AwsUpgradeStrategy.DeepCopyInto(&out.AwsUpgradeStrategy)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceGroupSpec.
func (in *InstanceGroupSpec) DeepCopy() *InstanceGroupSpec {
	if in == nil {
		return nil
	}
	out := new(InstanceGroupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceGroupStatus) DeepCopyInto(out *InstanceGroupStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]InstanceGroupCondition, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceGroupStatus.
func (in *InstanceGroupStatus) DeepCopy() *InstanceGroupStatus {
	if in == nil {
		return nil
	}
	out := new(InstanceGroupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeVolume) DeepCopyInto(out *NodeVolume) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeVolume.
func (in *NodeVolume) DeepCopy() *NodeVolume {
	if in == nil {
		return nil
	}
	out := new(NodeVolume)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RollingUpdateStrategy) DeepCopyInto(out *RollingUpdateStrategy) {
	*out = *in
	if in.MaxUnavailable != nil {
		in, out := &in.MaxUnavailable, &out.MaxUnavailable
		*out = new(intstr.IntOrString)
		**out = **in
	}
	if in.SuspendProcesses != nil {
		in, out := &in.SuspendProcesses, &out.SuspendProcesses
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RollingUpdateStrategy.
func (in *RollingUpdateStrategy) DeepCopy() *RollingUpdateStrategy {
	if in == nil {
		return nil
	}
	out := new(RollingUpdateStrategy)
	in.DeepCopyInto(out)
	return out
}
