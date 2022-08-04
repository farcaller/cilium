//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

// Code generated by main. DO NOT EDIT.

package v1

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *Bytes) DeepEqual(other *Bytes) bool {
	if other == nil {
		return false
	}

	if len(*in) != len(*other) {
		return false
	} else {
		for i, inElement := range *in {
			if inElement != (*other)[i] {
				return false
			}
		}
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *ClientIPConfig) DeepEqual(other *ClientIPConfig) bool {
	if other == nil {
		return false
	}

	if (in.TimeoutSeconds == nil) != (other.TimeoutSeconds == nil) {
		return false
	} else if in.TimeoutSeconds != nil {
		if *in.TimeoutSeconds != *other.TimeoutSeconds {
			return false
		}
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *Container) DeepEqual(other *Container) bool {
	if other == nil {
		return false
	}

	if in.Name != other.Name {
		return false
	}
	if in.Image != other.Image {
		return false
	}
	if ((in.Ports != nil) && (other.Ports != nil)) || ((in.Ports == nil) != (other.Ports == nil)) {
		in, other := &in.Ports, &other.Ports
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for i, inElement := range *in {
				if !inElement.DeepEqual(&(*other)[i]) {
					return false
				}
			}
		}
	}

	if ((in.VolumeMounts != nil) && (other.VolumeMounts != nil)) || ((in.VolumeMounts == nil) != (other.VolumeMounts == nil)) {
		in, other := &in.VolumeMounts, &other.VolumeMounts
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for i, inElement := range *in {
				if !inElement.DeepEqual(&(*other)[i]) {
					return false
				}
			}
		}
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *ContainerPort) DeepEqual(other *ContainerPort) bool {
	if other == nil {
		return false
	}

	if in.Name != other.Name {
		return false
	}
	if in.HostPort != other.HostPort {
		return false
	}
	if in.ContainerPort != other.ContainerPort {
		return false
	}
	if in.Protocol != other.Protocol {
		return false
	}
	if in.HostIP != other.HostIP {
		return false
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *ContainerState) DeepEqual(other *ContainerState) bool {
	if other == nil {
		return false
	}

	if (in.Waiting == nil) != (other.Waiting == nil) {
		return false
	} else if in.Waiting != nil {
		if !in.Waiting.DeepEqual(other.Waiting) {
			return false
		}
	}

	if (in.Running == nil) != (other.Running == nil) {
		return false
	} else if in.Running != nil {
		if !in.Running.DeepEqual(other.Running) {
			return false
		}
	}

	if (in.Terminated == nil) != (other.Terminated == nil) {
		return false
	} else if in.Terminated != nil {
		if !in.Terminated.DeepEqual(other.Terminated) {
			return false
		}
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *ContainerStateRunning) DeepEqual(other *ContainerStateRunning) bool {
	if other == nil {
		return false
	}

	if !in.StartedAt.DeepEqual(&other.StartedAt) {
		return false
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *ContainerStateTerminated) DeepEqual(other *ContainerStateTerminated) bool {
	if other == nil {
		return false
	}

	if in.ExitCode != other.ExitCode {
		return false
	}
	if in.Signal != other.Signal {
		return false
	}
	if in.Reason != other.Reason {
		return false
	}
	if in.Message != other.Message {
		return false
	}
	if !in.StartedAt.DeepEqual(&other.StartedAt) {
		return false
	}

	if !in.FinishedAt.DeepEqual(&other.FinishedAt) {
		return false
	}

	if in.ContainerID != other.ContainerID {
		return false
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *ContainerStateWaiting) DeepEqual(other *ContainerStateWaiting) bool {
	if other == nil {
		return false
	}

	if in.Reason != other.Reason {
		return false
	}
	if in.Message != other.Message {
		return false
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *ContainerStatus) DeepEqual(other *ContainerStatus) bool {
	if other == nil {
		return false
	}

	if !in.State.DeepEqual(&other.State) {
		return false
	}

	if in.ContainerID != other.ContainerID {
		return false
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *EndpointAddress) DeepEqual(other *EndpointAddress) bool {
	if other == nil {
		return false
	}

	if in.IP != other.IP {
		return false
	}
	if (in.NodeName == nil) != (other.NodeName == nil) {
		return false
	} else if in.NodeName != nil {
		if *in.NodeName != *other.NodeName {
			return false
		}
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *EndpointPort) DeepEqual(other *EndpointPort) bool {
	if other == nil {
		return false
	}

	if in.Name != other.Name {
		return false
	}
	if in.Port != other.Port {
		return false
	}
	if in.Protocol != other.Protocol {
		return false
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *EndpointSubset) DeepEqual(other *EndpointSubset) bool {
	if other == nil {
		return false
	}

	if ((in.Addresses != nil) && (other.Addresses != nil)) || ((in.Addresses == nil) != (other.Addresses == nil)) {
		in, other := &in.Addresses, &other.Addresses
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for i, inElement := range *in {
				if !inElement.DeepEqual(&(*other)[i]) {
					return false
				}
			}
		}
	}

	if ((in.Ports != nil) && (other.Ports != nil)) || ((in.Ports == nil) != (other.Ports == nil)) {
		in, other := &in.Ports, &other.Ports
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for i, inElement := range *in {
				if !inElement.DeepEqual(&(*other)[i]) {
					return false
				}
			}
		}
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *Endpoints) DeepEqual(other *Endpoints) bool {
	if other == nil {
		return false
	}

	if in.TypeMeta != other.TypeMeta {
		return false
	}

	if !in.ObjectMeta.DeepEqual(&other.ObjectMeta) {
		return false
	}

	if ((in.Subsets != nil) && (other.Subsets != nil)) || ((in.Subsets == nil) != (other.Subsets == nil)) {
		in, other := &in.Subsets, &other.Subsets
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for i, inElement := range *in {
				if !inElement.DeepEqual(&(*other)[i]) {
					return false
				}
			}
		}
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *EndpointsList) DeepEqual(other *EndpointsList) bool {
	if other == nil {
		return false
	}

	if in.TypeMeta != other.TypeMeta {
		return false
	}

	if !in.ListMeta.DeepEqual(&other.ListMeta) {
		return false
	}

	if ((in.Items != nil) && (other.Items != nil)) || ((in.Items == nil) != (other.Items == nil)) {
		in, other := &in.Items, &other.Items
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for i, inElement := range *in {
				if !inElement.DeepEqual(&(*other)[i]) {
					return false
				}
			}
		}
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *LoadBalancerIngress) DeepEqual(other *LoadBalancerIngress) bool {
	if other == nil {
		return false
	}

	if in.IP != other.IP {
		return false
	}
	if in.Hostname != other.Hostname {
		return false
	}
	if ((in.Ports != nil) && (other.Ports != nil)) || ((in.Ports == nil) != (other.Ports == nil)) {
		in, other := &in.Ports, &other.Ports
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for i, inElement := range *in {
				if !inElement.DeepEqual(&(*other)[i]) {
					return false
				}
			}
		}
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *LoadBalancerStatus) DeepEqual(other *LoadBalancerStatus) bool {
	if other == nil {
		return false
	}

	if ((in.Ingress != nil) && (other.Ingress != nil)) || ((in.Ingress == nil) != (other.Ingress == nil)) {
		in, other := &in.Ingress, &other.Ingress
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for i, inElement := range *in {
				if !inElement.DeepEqual(&(*other)[i]) {
					return false
				}
			}
		}
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *Namespace) DeepEqual(other *Namespace) bool {
	if other == nil {
		return false
	}

	if in.TypeMeta != other.TypeMeta {
		return false
	}

	if !in.ObjectMeta.DeepEqual(&other.ObjectMeta) {
		return false
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *NamespaceList) DeepEqual(other *NamespaceList) bool {
	if other == nil {
		return false
	}

	if in.TypeMeta != other.TypeMeta {
		return false
	}

	if !in.ListMeta.DeepEqual(&other.ListMeta) {
		return false
	}

	if ((in.Items != nil) && (other.Items != nil)) || ((in.Items == nil) != (other.Items == nil)) {
		in, other := &in.Items, &other.Items
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for i, inElement := range *in {
				if !inElement.DeepEqual(&(*other)[i]) {
					return false
				}
			}
		}
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *Node) DeepEqual(other *Node) bool {
	if other == nil {
		return false
	}

	if in.TypeMeta != other.TypeMeta {
		return false
	}

	if !in.ObjectMeta.DeepEqual(&other.ObjectMeta) {
		return false
	}

	if !in.Spec.DeepEqual(&other.Spec) {
		return false
	}

	if !in.Status.DeepEqual(&other.Status) {
		return false
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *NodeAddress) DeepEqual(other *NodeAddress) bool {
	if other == nil {
		return false
	}

	if in.Type != other.Type {
		return false
	}
	if in.Address != other.Address {
		return false
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *NodeCondition) DeepEqual(other *NodeCondition) bool {
	if other == nil {
		return false
	}

	if in.Type != other.Type {
		return false
	}
	if in.Status != other.Status {
		return false
	}
	if in.Reason != other.Reason {
		return false
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *NodeList) DeepEqual(other *NodeList) bool {
	if other == nil {
		return false
	}

	if in.TypeMeta != other.TypeMeta {
		return false
	}

	if !in.ListMeta.DeepEqual(&other.ListMeta) {
		return false
	}

	if ((in.Items != nil) && (other.Items != nil)) || ((in.Items == nil) != (other.Items == nil)) {
		in, other := &in.Items, &other.Items
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for i, inElement := range *in {
				if !inElement.DeepEqual(&(*other)[i]) {
					return false
				}
			}
		}
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *NodeSpec) DeepEqual(other *NodeSpec) bool {
	if other == nil {
		return false
	}

	if in.PodCIDR != other.PodCIDR {
		return false
	}
	if ((in.PodCIDRs != nil) && (other.PodCIDRs != nil)) || ((in.PodCIDRs == nil) != (other.PodCIDRs == nil)) {
		in, other := &in.PodCIDRs, &other.PodCIDRs
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for i, inElement := range *in {
				if inElement != (*other)[i] {
					return false
				}
			}
		}
	}

	if ((in.Taints != nil) && (other.Taints != nil)) || ((in.Taints == nil) != (other.Taints == nil)) {
		in, other := &in.Taints, &other.Taints
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for i, inElement := range *in {
				if !inElement.DeepEqual(&(*other)[i]) {
					return false
				}
			}
		}
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *NodeStatus) DeepEqual(other *NodeStatus) bool {
	if other == nil {
		return false
	}

	if ((in.Conditions != nil) && (other.Conditions != nil)) || ((in.Conditions == nil) != (other.Conditions == nil)) {
		in, other := &in.Conditions, &other.Conditions
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for i, inElement := range *in {
				if !inElement.DeepEqual(&(*other)[i]) {
					return false
				}
			}
		}
	}

	if ((in.Addresses != nil) && (other.Addresses != nil)) || ((in.Addresses == nil) != (other.Addresses == nil)) {
		in, other := &in.Addresses, &other.Addresses
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for i, inElement := range *in {
				if !inElement.DeepEqual(&(*other)[i]) {
					return false
				}
			}
		}
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *Pod) DeepEqual(other *Pod) bool {
	if other == nil {
		return false
	}

	if in.TypeMeta != other.TypeMeta {
		return false
	}

	if !in.ObjectMeta.DeepEqual(&other.ObjectMeta) {
		return false
	}

	if !in.Spec.DeepEqual(&other.Spec) {
		return false
	}

	if !in.Status.DeepEqual(&other.Status) {
		return false
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *PodCondition) DeepEqual(other *PodCondition) bool {
	if other == nil {
		return false
	}

	if in.Type != other.Type {
		return false
	}
	if in.Status != other.Status {
		return false
	}
	if !in.LastProbeTime.DeepEqual(&other.LastProbeTime) {
		return false
	}

	if !in.LastTransitionTime.DeepEqual(&other.LastTransitionTime) {
		return false
	}

	if in.Reason != other.Reason {
		return false
	}
	if in.Message != other.Message {
		return false
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *PodIP) DeepEqual(other *PodIP) bool {
	if other == nil {
		return false
	}

	if in.IP != other.IP {
		return false
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *PodList) DeepEqual(other *PodList) bool {
	if other == nil {
		return false
	}

	if in.TypeMeta != other.TypeMeta {
		return false
	}

	if !in.ListMeta.DeepEqual(&other.ListMeta) {
		return false
	}

	if ((in.Items != nil) && (other.Items != nil)) || ((in.Items == nil) != (other.Items == nil)) {
		in, other := &in.Items, &other.Items
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for i, inElement := range *in {
				if !inElement.DeepEqual(&(*other)[i]) {
					return false
				}
			}
		}
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *PodSpec) DeepEqual(other *PodSpec) bool {
	if other == nil {
		return false
	}

	if ((in.InitContainers != nil) && (other.InitContainers != nil)) || ((in.InitContainers == nil) != (other.InitContainers == nil)) {
		in, other := &in.InitContainers, &other.InitContainers
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for i, inElement := range *in {
				if !inElement.DeepEqual(&(*other)[i]) {
					return false
				}
			}
		}
	}

	if ((in.Containers != nil) && (other.Containers != nil)) || ((in.Containers == nil) != (other.Containers == nil)) {
		in, other := &in.Containers, &other.Containers
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for i, inElement := range *in {
				if !inElement.DeepEqual(&(*other)[i]) {
					return false
				}
			}
		}
	}

	if in.ServiceAccountName != other.ServiceAccountName {
		return false
	}
	if in.NodeName != other.NodeName {
		return false
	}
	if in.HostNetwork != other.HostNetwork {
		return false
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *PodStatus) DeepEqual(other *PodStatus) bool {
	if other == nil {
		return false
	}

	if in.Phase != other.Phase {
		return false
	}
	if ((in.Conditions != nil) && (other.Conditions != nil)) || ((in.Conditions == nil) != (other.Conditions == nil)) {
		in, other := &in.Conditions, &other.Conditions
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for i, inElement := range *in {
				if !inElement.DeepEqual(&(*other)[i]) {
					return false
				}
			}
		}
	}

	if in.HostIP != other.HostIP {
		return false
	}
	if in.PodIP != other.PodIP {
		return false
	}
	if ((in.PodIPs != nil) && (other.PodIPs != nil)) || ((in.PodIPs == nil) != (other.PodIPs == nil)) {
		in, other := &in.PodIPs, &other.PodIPs
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for i, inElement := range *in {
				if !inElement.DeepEqual(&(*other)[i]) {
					return false
				}
			}
		}
	}

	if (in.StartTime == nil) != (other.StartTime == nil) {
		return false
	} else if in.StartTime != nil {
		if !in.StartTime.DeepEqual(other.StartTime) {
			return false
		}
	}

	if ((in.ContainerStatuses != nil) && (other.ContainerStatuses != nil)) || ((in.ContainerStatuses == nil) != (other.ContainerStatuses == nil)) {
		in, other := &in.ContainerStatuses, &other.ContainerStatuses
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for i, inElement := range *in {
				if !inElement.DeepEqual(&(*other)[i]) {
					return false
				}
			}
		}
	}

	if in.QOSClass != other.QOSClass {
		return false
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *PortStatus) DeepEqual(other *PortStatus) bool {
	if other == nil {
		return false
	}

	if in.Port != other.Port {
		return false
	}
	if in.Protocol != other.Protocol {
		return false
	}
	if (in.Error == nil) != (other.Error == nil) {
		return false
	} else if in.Error != nil {
		if *in.Error != *other.Error {
			return false
		}
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *Secret) DeepEqual(other *Secret) bool {
	if other == nil {
		return false
	}

	if in.TypeMeta != other.TypeMeta {
		return false
	}

	if !in.ObjectMeta.DeepEqual(&other.ObjectMeta) {
		return false
	}

	if (in.Immutable == nil) != (other.Immutable == nil) {
		return false
	} else if in.Immutable != nil {
		if *in.Immutable != *other.Immutable {
			return false
		}
	}

	if ((in.Data != nil) && (other.Data != nil)) || ((in.Data == nil) != (other.Data == nil)) {
		in, other := &in.Data, &other.Data
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for key, inValue := range *in {
				if otherValue, present := (*other)[key]; !present {
					return false
				} else {
					if !inValue.DeepEqual(&otherValue) {
						return false
					}
				}
			}
		}
	}

	if ((in.StringData != nil) && (other.StringData != nil)) || ((in.StringData == nil) != (other.StringData == nil)) {
		in, other := &in.StringData, &other.StringData
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for key, inValue := range *in {
				if otherValue, present := (*other)[key]; !present {
					return false
				} else {
					if inValue != otherValue {
						return false
					}
				}
			}
		}
	}

	if in.Type != other.Type {
		return false
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *SecretList) DeepEqual(other *SecretList) bool {
	if other == nil {
		return false
	}

	if in.TypeMeta != other.TypeMeta {
		return false
	}

	if !in.ListMeta.DeepEqual(&other.ListMeta) {
		return false
	}

	if ((in.Items != nil) && (other.Items != nil)) || ((in.Items == nil) != (other.Items == nil)) {
		in, other := &in.Items, &other.Items
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for i, inElement := range *in {
				if !inElement.DeepEqual(&(*other)[i]) {
					return false
				}
			}
		}
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *Service) DeepEqual(other *Service) bool {
	if other == nil {
		return false
	}

	if in.TypeMeta != other.TypeMeta {
		return false
	}

	if !in.ObjectMeta.DeepEqual(&other.ObjectMeta) {
		return false
	}

	if !in.Spec.DeepEqual(&other.Spec) {
		return false
	}

	if !in.Status.DeepEqual(&other.Status) {
		return false
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *ServiceList) DeepEqual(other *ServiceList) bool {
	if other == nil {
		return false
	}

	if in.TypeMeta != other.TypeMeta {
		return false
	}

	if !in.ListMeta.DeepEqual(&other.ListMeta) {
		return false
	}

	if ((in.Items != nil) && (other.Items != nil)) || ((in.Items == nil) != (other.Items == nil)) {
		in, other := &in.Items, &other.Items
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for i, inElement := range *in {
				if !inElement.DeepEqual(&(*other)[i]) {
					return false
				}
			}
		}
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *ServicePort) DeepEqual(other *ServicePort) bool {
	if other == nil {
		return false
	}

	if in.Name != other.Name {
		return false
	}
	if in.Protocol != other.Protocol {
		return false
	}
	if in.Port != other.Port {
		return false
	}
	if in.NodePort != other.NodePort {
		return false
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *ServiceSpec) DeepEqual(other *ServiceSpec) bool {
	if other == nil {
		return false
	}

	if ((in.Ports != nil) && (other.Ports != nil)) || ((in.Ports == nil) != (other.Ports == nil)) {
		in, other := &in.Ports, &other.Ports
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for i, inElement := range *in {
				if !inElement.DeepEqual(&(*other)[i]) {
					return false
				}
			}
		}
	}

	if ((in.Selector != nil) && (other.Selector != nil)) || ((in.Selector == nil) != (other.Selector == nil)) {
		in, other := &in.Selector, &other.Selector
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for key, inValue := range *in {
				if otherValue, present := (*other)[key]; !present {
					return false
				} else {
					if inValue != otherValue {
						return false
					}
				}
			}
		}
	}

	if in.ClusterIP != other.ClusterIP {
		return false
	}
	if ((in.ClusterIPs != nil) && (other.ClusterIPs != nil)) || ((in.ClusterIPs == nil) != (other.ClusterIPs == nil)) {
		in, other := &in.ClusterIPs, &other.ClusterIPs
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for i, inElement := range *in {
				if inElement != (*other)[i] {
					return false
				}
			}
		}
	}

	if in.Type != other.Type {
		return false
	}
	if ((in.ExternalIPs != nil) && (other.ExternalIPs != nil)) || ((in.ExternalIPs == nil) != (other.ExternalIPs == nil)) {
		in, other := &in.ExternalIPs, &other.ExternalIPs
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for i, inElement := range *in {
				if inElement != (*other)[i] {
					return false
				}
			}
		}
	}

	if in.SessionAffinity != other.SessionAffinity {
		return false
	}
	if in.LoadBalancerIP != other.LoadBalancerIP {
		return false
	}
	if ((in.LoadBalancerSourceRanges != nil) && (other.LoadBalancerSourceRanges != nil)) || ((in.LoadBalancerSourceRanges == nil) != (other.LoadBalancerSourceRanges == nil)) {
		in, other := &in.LoadBalancerSourceRanges, &other.LoadBalancerSourceRanges
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for i, inElement := range *in {
				if inElement != (*other)[i] {
					return false
				}
			}
		}
	}

	if in.ExternalTrafficPolicy != other.ExternalTrafficPolicy {
		return false
	}
	if in.HealthCheckNodePort != other.HealthCheckNodePort {
		return false
	}
	if (in.SessionAffinityConfig == nil) != (other.SessionAffinityConfig == nil) {
		return false
	} else if in.SessionAffinityConfig != nil {
		if !in.SessionAffinityConfig.DeepEqual(other.SessionAffinityConfig) {
			return false
		}
	}

	if ((in.IPFamilies != nil) && (other.IPFamilies != nil)) || ((in.IPFamilies == nil) != (other.IPFamilies == nil)) {
		in, other := &in.IPFamilies, &other.IPFamilies
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for i, inElement := range *in {
				if inElement != (*other)[i] {
					return false
				}
			}
		}
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *ServiceStatus) DeepEqual(other *ServiceStatus) bool {
	if other == nil {
		return false
	}

	if !in.LoadBalancer.DeepEqual(&other.LoadBalancer) {
		return false
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *SessionAffinityConfig) DeepEqual(other *SessionAffinityConfig) bool {
	if other == nil {
		return false
	}

	if (in.ClientIP == nil) != (other.ClientIP == nil) {
		return false
	} else if in.ClientIP != nil {
		if !in.ClientIP.DeepEqual(other.ClientIP) {
			return false
		}
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *Taint) DeepEqual(other *Taint) bool {
	if other == nil {
		return false
	}

	if in.Key != other.Key {
		return false
	}
	if in.Value != other.Value {
		return false
	}
	if in.Effect != other.Effect {
		return false
	}
	if (in.TimeAdded == nil) != (other.TimeAdded == nil) {
		return false
	} else if in.TimeAdded != nil {
		if !in.TimeAdded.DeepEqual(other.TimeAdded) {
			return false
		}
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *TypedLocalObjectReference) DeepEqual(other *TypedLocalObjectReference) bool {
	if other == nil {
		return false
	}

	if (in.APIGroup == nil) != (other.APIGroup == nil) {
		return false
	} else if in.APIGroup != nil {
		if *in.APIGroup != *other.APIGroup {
			return false
		}
	}

	if in.Kind != other.Kind {
		return false
	}
	if in.Name != other.Name {
		return false
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *VolumeMount) DeepEqual(other *VolumeMount) bool {
	if other == nil {
		return false
	}

	if in.MountPath != other.MountPath {
		return false
	}

	return true
}
