//go:build !ignore_autogenerated

/*
Copyright 2023 zncdatadev.

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
	commonsv1alpha1 "github.com/zncdatadev/operator-go/pkg/apis/commons/v1alpha1"
	s3v1alpha1 "github.com/zncdatadev/operator-go/pkg/apis/s3/v1alpha1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthenticationSpec) DeepCopyInto(out *AuthenticationSpec) {
	*out = *in
	if in.Oidc != nil {
		in, out := &in.Oidc, &out.Oidc
		*out = new(OidcSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthenticationSpec.
func (in *AuthenticationSpec) DeepCopy() *AuthenticationSpec {
	if in == nil {
		return nil
	}
	out := new(AuthenticationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BaseRoleSpec) DeepCopyInto(out *BaseRoleSpec) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	if in.OverridesSpec != nil {
		in, out := &in.OverridesSpec, &out.OverridesSpec
		*out = new(commonsv1alpha1.OverridesSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.RoleConfig != nil {
		in, out := &in.RoleConfig, &out.RoleConfig
		*out = new(commonsv1alpha1.RoleConfigSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = new(ConfigSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BaseRoleSpec.
func (in *BaseRoleSpec) DeepCopy() *BaseRoleSpec {
	if in == nil {
		return nil
	}
	out := new(BaseRoleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CatalogLabelSelectorSpec) DeepCopyInto(out *CatalogLabelSelectorSpec) {
	*out = *in
	if in.MatchLabels != nil {
		in, out := &in.MatchLabels, &out.MatchLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.MatchExpressions != nil {
		in, out := &in.MatchExpressions, &out.MatchExpressions
		*out = make([]v1.LabelSelectorRequirement, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CatalogLabelSelectorSpec.
func (in *CatalogLabelSelectorSpec) DeepCopy() *CatalogLabelSelectorSpec {
	if in == nil {
		return nil
	}
	out := new(CatalogLabelSelectorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterConfigSpec) DeepCopyInto(out *ClusterConfigSpec) {
	*out = *in
	if in.Authentication != nil {
		in, out := &in.Authentication, &out.Authentication
		*out = make([]AuthenticationSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.CatalogLabelSelector != nil {
		in, out := &in.CatalogLabelSelector, &out.CatalogLabelSelector
		*out = new(CatalogLabelSelectorSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Tls != nil {
		in, out := &in.Tls, &out.Tls
		*out = new(TlsSpec)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterConfigSpec.
func (in *ClusterConfigSpec) DeepCopy() *ClusterConfigSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigOverridesSpec) DeepCopyInto(out *ConfigOverridesSpec) {
	*out = *in
	if in.Node != nil {
		in, out := &in.Node, &out.Node
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Log != nil {
		in, out := &in.Log, &out.Log
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ExchangeManager != nil {
		in, out := &in.ExchangeManager, &out.ExchangeManager
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigOverridesSpec.
func (in *ConfigOverridesSpec) DeepCopy() *ConfigOverridesSpec {
	if in == nil {
		return nil
	}
	out := new(ConfigOverridesSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigSpec) DeepCopyInto(out *ConfigSpec) {
	*out = *in
	if in.RoleGroupConfigSpec != nil {
		in, out := &in.RoleGroupConfigSpec, &out.RoleGroupConfigSpec
		*out = new(commonsv1alpha1.RoleGroupConfigSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigSpec.
func (in *ConfigSpec) DeepCopy() *ConfigSpec {
	if in == nil {
		return nil
	}
	out := new(ConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConnectorSpec) DeepCopyInto(out *ConnectorSpec) {
	*out = *in
	if in.Generic != nil {
		in, out := &in.Generic, &out.Generic
		*out = new(GenericConnectorSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Hive != nil {
		in, out := &in.Hive, &out.Hive
		*out = new(HiveConnectorSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.IceBerg != nil {
		in, out := &in.IceBerg, &out.IceBerg
		*out = new(IcebergConnectorSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Tpcds != nil {
		in, out := &in.Tpcds, &out.Tpcds
		*out = new(TpcdsConnectorSpec)
		**out = **in
	}
	if in.Tpch != nil {
		in, out := &in.Tpch, &out.Tpch
		*out = new(TpchConnectorSpec)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConnectorSpec.
func (in *ConnectorSpec) DeepCopy() *ConnectorSpec {
	if in == nil {
		return nil
	}
	out := new(ConnectorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CoordinatorsSpec) DeepCopyInto(out *CoordinatorsSpec) {
	*out = *in
	if in.RoleGroups != nil {
		in, out := &in.RoleGroups, &out.RoleGroups
		*out = make(map[string]*RoleGroupSpec, len(*in))
		for key, val := range *in {
			var outVal *RoleGroupSpec
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(RoleGroupSpec)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
	in.BaseRoleSpec.DeepCopyInto(&out.BaseRoleSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CoordinatorsSpec.
func (in *CoordinatorsSpec) DeepCopy() *CoordinatorsSpec {
	if in == nil {
		return nil
	}
	out := new(CoordinatorsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GenericConnectorSpec) DeepCopyInto(out *GenericConnectorSpec) {
	*out = *in
	if in.Properties != nil {
		in, out := &in.Properties, &out.Properties
		*out = new(PropertiesSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GenericConnectorSpec.
func (in *GenericConnectorSpec) DeepCopy() *GenericConnectorSpec {
	if in == nil {
		return nil
	}
	out := new(GenericConnectorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HdfsConnectionSpec) DeepCopyInto(out *HdfsConnectionSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HdfsConnectionSpec.
func (in *HdfsConnectionSpec) DeepCopy() *HdfsConnectionSpec {
	if in == nil {
		return nil
	}
	out := new(HdfsConnectionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HiveConnectorSpec) DeepCopyInto(out *HiveConnectorSpec) {
	*out = *in
	if in.Metastore != nil {
		in, out := &in.Metastore, &out.Metastore
		*out = new(MetastoreConnectionSpec)
		**out = **in
	}
	if in.S3 != nil {
		in, out := &in.S3, &out.S3
		*out = new(s3v1alpha1.S3BucketSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Hdfs != nil {
		in, out := &in.Hdfs, &out.Hdfs
		*out = new(HdfsConnectionSpec)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HiveConnectorSpec.
func (in *HiveConnectorSpec) DeepCopy() *HiveConnectorSpec {
	if in == nil {
		return nil
	}
	out := new(HiveConnectorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IcebergConnectorSpec) DeepCopyInto(out *IcebergConnectorSpec) {
	*out = *in
	if in.Metastore != nil {
		in, out := &in.Metastore, &out.Metastore
		*out = new(MetastoreConnectionSpec)
		**out = **in
	}
	if in.S3 != nil {
		in, out := &in.S3, &out.S3
		*out = new(s3v1alpha1.S3BucketSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Hdfs != nil {
		in, out := &in.Hdfs, &out.Hdfs
		*out = new(HdfsConnectionSpec)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IcebergConnectorSpec.
func (in *IcebergConnectorSpec) DeepCopy() *IcebergConnectorSpec {
	if in == nil {
		return nil
	}
	out := new(IcebergConnectorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageSpec) DeepCopyInto(out *ImageSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageSpec.
func (in *ImageSpec) DeepCopy() *ImageSpec {
	if in == nil {
		return nil
	}
	out := new(ImageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JvmPropertiesRoleConfigSpec) DeepCopyInto(out *JvmPropertiesRoleConfigSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JvmPropertiesRoleConfigSpec.
func (in *JvmPropertiesRoleConfigSpec) DeepCopy() *JvmPropertiesRoleConfigSpec {
	if in == nil {
		return nil
	}
	out := new(JvmPropertiesRoleConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoggingSpec) DeepCopyInto(out *LoggingSpec) {
	*out = *in
	if in.Containers != nil {
		in, out := &in.Containers, &out.Containers
		*out = make(map[string]commonsv1alpha1.LoggingConfigSpec, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoggingSpec.
func (in *LoggingSpec) DeepCopy() *LoggingSpec {
	if in == nil {
		return nil
	}
	out := new(LoggingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetastoreConnectionSpec) DeepCopyInto(out *MetastoreConnectionSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetastoreConnectionSpec.
func (in *MetastoreConnectionSpec) DeepCopy() *MetastoreConnectionSpec {
	if in == nil {
		return nil
	}
	out := new(MetastoreConnectionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OidcSpec) DeepCopyInto(out *OidcSpec) {
	*out = *in
	if in.ExtraScopes != nil {
		in, out := &in.ExtraScopes, &out.ExtraScopes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OidcSpec.
func (in *OidcSpec) DeepCopy() *OidcSpec {
	if in == nil {
		return nil
	}
	out := new(OidcSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PropertiesSpec) DeepCopyInto(out *PropertiesSpec) {
	*out = *in
	if in.ValueFromConfiguration != nil {
		in, out := &in.ValueFromConfiguration, &out.ValueFromConfiguration
		*out = new(ValueFromConfigurationSpec)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PropertiesSpec.
func (in *PropertiesSpec) DeepCopy() *PropertiesSpec {
	if in == nil {
		return nil
	}
	out := new(PropertiesSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoleGroupSpec) DeepCopyInto(out *RoleGroupSpec) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = new(ConfigSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.OverridesSpec != nil {
		in, out := &in.OverridesSpec, &out.OverridesSpec
		*out = new(commonsv1alpha1.OverridesSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoleGroupSpec.
func (in *RoleGroupSpec) DeepCopy() *RoleGroupSpec {
	if in == nil {
		return nil
	}
	out := new(RoleGroupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TlsSpec) DeepCopyInto(out *TlsSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TlsSpec.
func (in *TlsSpec) DeepCopy() *TlsSpec {
	if in == nil {
		return nil
	}
	out := new(TlsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TpcdsConnectorSpec) DeepCopyInto(out *TpcdsConnectorSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TpcdsConnectorSpec.
func (in *TpcdsConnectorSpec) DeepCopy() *TpcdsConnectorSpec {
	if in == nil {
		return nil
	}
	out := new(TpcdsConnectorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TpchConnectorSpec) DeepCopyInto(out *TpchConnectorSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TpchConnectorSpec.
func (in *TpchConnectorSpec) DeepCopy() *TpchConnectorSpec {
	if in == nil {
		return nil
	}
	out := new(TpchConnectorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TrinoCatalog) DeepCopyInto(out *TrinoCatalog) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrinoCatalog.
func (in *TrinoCatalog) DeepCopy() *TrinoCatalog {
	if in == nil {
		return nil
	}
	out := new(TrinoCatalog)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TrinoCatalog) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TrinoCatalogList) DeepCopyInto(out *TrinoCatalogList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TrinoCatalog, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrinoCatalogList.
func (in *TrinoCatalogList) DeepCopy() *TrinoCatalogList {
	if in == nil {
		return nil
	}
	out := new(TrinoCatalogList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TrinoCatalogList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TrinoCatalogSpec) DeepCopyInto(out *TrinoCatalogSpec) {
	*out = *in
	if in.Connectors != nil {
		in, out := &in.Connectors, &out.Connectors
		*out = make([]ConnectorSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ConfigOverrides != nil {
		in, out := &in.ConfigOverrides, &out.ConfigOverrides
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrinoCatalogSpec.
func (in *TrinoCatalogSpec) DeepCopy() *TrinoCatalogSpec {
	if in == nil {
		return nil
	}
	out := new(TrinoCatalogSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TrinoCatalogStatus) DeepCopyInto(out *TrinoCatalogStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrinoCatalogStatus.
func (in *TrinoCatalogStatus) DeepCopy() *TrinoCatalogStatus {
	if in == nil {
		return nil
	}
	out := new(TrinoCatalogStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TrinoCluster) DeepCopyInto(out *TrinoCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrinoCluster.
func (in *TrinoCluster) DeepCopy() *TrinoCluster {
	if in == nil {
		return nil
	}
	out := new(TrinoCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TrinoCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TrinoClusterList) DeepCopyInto(out *TrinoClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TrinoCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrinoClusterList.
func (in *TrinoClusterList) DeepCopy() *TrinoClusterList {
	if in == nil {
		return nil
	}
	out := new(TrinoClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TrinoClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TrinoClusterSpec) DeepCopyInto(out *TrinoClusterSpec) {
	*out = *in
	if in.ClusterConfig != nil {
		in, out := &in.ClusterConfig, &out.ClusterConfig
		*out = new(ClusterConfigSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.ClusterOperation != nil {
		in, out := &in.ClusterOperation, &out.ClusterOperation
		*out = new(commonsv1alpha1.ClusterOperationSpec)
		**out = **in
	}
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		*out = new(ImageSpec)
		**out = **in
	}
	if in.Coordinators != nil {
		in, out := &in.Coordinators, &out.Coordinators
		*out = new(CoordinatorsSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Workers != nil {
		in, out := &in.Workers, &out.Workers
		*out = new(WorkersSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrinoClusterSpec.
func (in *TrinoClusterSpec) DeepCopy() *TrinoClusterSpec {
	if in == nil {
		return nil
	}
	out := new(TrinoClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ValueFromConfigurationSpec) DeepCopyInto(out *ValueFromConfigurationSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ValueFromConfigurationSpec.
func (in *ValueFromConfigurationSpec) DeepCopy() *ValueFromConfigurationSpec {
	if in == nil {
		return nil
	}
	out := new(ValueFromConfigurationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkersSpec) DeepCopyInto(out *WorkersSpec) {
	*out = *in
	if in.RoleGroups != nil {
		in, out := &in.RoleGroups, &out.RoleGroups
		*out = make(map[string]*RoleGroupSpec, len(*in))
		for key, val := range *in {
			var outVal *RoleGroupSpec
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(RoleGroupSpec)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
	in.BaseRoleSpec.DeepCopyInto(&out.BaseRoleSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkersSpec.
func (in *WorkersSpec) DeepCopy() *WorkersSpec {
	if in == nil {
		return nil
	}
	out := new(WorkersSpec)
	in.DeepCopyInto(out)
	return out
}
