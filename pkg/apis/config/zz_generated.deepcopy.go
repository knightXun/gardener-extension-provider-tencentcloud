// +build !ignore_autogenerated

/*
Copyright (c) 2020 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package config

import (
	healthcheckconfig "github.com/gardener/gardener/extensions/pkg/controller/healthcheck/config"
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	componentbaseconfig "k8s.io/component-base/config"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerConfiguration) DeepCopyInto(out *ControllerConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.ClientConnection != nil {
		in, out := &in.ClientConnection, &out.ClientConnection
		*out = new(componentbaseconfig.ClientConnectionConfiguration)
		**out = **in
	}
	if in.MachineImageOwnerSecretRef != nil {
		in, out := &in.MachineImageOwnerSecretRef, &out.MachineImageOwnerSecretRef
		*out = new(v1.SecretReference)
		**out = **in
	}
	if in.WhitelistedImageIDs != nil {
		in, out := &in.WhitelistedImageIDs, &out.WhitelistedImageIDs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.KubeAPIServer != nil {
		in, out := &in.KubeAPIServer, &out.KubeAPIServer
		*out = new(KubeAPIServer)
		**out = **in
	}
	in.ETCD.DeepCopyInto(&out.ETCD)
	if in.HealthCheckConfig != nil {
		in, out := &in.HealthCheckConfig, &out.HealthCheckConfig
		*out = new(healthcheckconfig.HealthCheckConfig)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerConfiguration.
func (in *ControllerConfiguration) DeepCopy() *ControllerConfiguration {
	if in == nil {
		return nil
	}
	out := new(ControllerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ControllerConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ETCD) DeepCopyInto(out *ETCD) {
	*out = *in
	in.Storage.DeepCopyInto(&out.Storage)
	in.Backup.DeepCopyInto(&out.Backup)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ETCD.
func (in *ETCD) DeepCopy() *ETCD {
	if in == nil {
		return nil
	}
	out := new(ETCD)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ETCDBackup) DeepCopyInto(out *ETCDBackup) {
	*out = *in
	if in.Schedule != nil {
		in, out := &in.Schedule, &out.Schedule
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ETCDBackup.
func (in *ETCDBackup) DeepCopy() *ETCDBackup {
	if in == nil {
		return nil
	}
	out := new(ETCDBackup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ETCDStorage) DeepCopyInto(out *ETCDStorage) {
	*out = *in
	if in.ClassName != nil {
		in, out := &in.ClassName, &out.ClassName
		*out = new(string)
		**out = **in
	}
	if in.Capacity != nil {
		in, out := &in.Capacity, &out.Capacity
		x := (*in).DeepCopy()
		*out = &x
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ETCDStorage.
func (in *ETCDStorage) DeepCopy() *ETCDStorage {
	if in == nil {
		return nil
	}
	out := new(ETCDStorage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeAPIServer) DeepCopyInto(out *KubeAPIServer) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeAPIServer.
func (in *KubeAPIServer) DeepCopy() *KubeAPIServer {
	if in == nil {
		return nil
	}
	out := new(KubeAPIServer)
	in.DeepCopyInto(out)
	return out
}