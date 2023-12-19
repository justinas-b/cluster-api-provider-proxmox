package v1alpha1

import (
	apiconversion "k8s.io/apimachinery/pkg/conversion"
	utilconversion "sigs.k8s.io/cluster-api/util/conversion"
	"sigs.k8s.io/controller-runtime/pkg/conversion"

	"github.com/ionos-cloud/cluster-api-provider-proxmox/api/v1alpha2"
)

// ConvertTo converts this ProxmoxCluster to the Hub version (v1alpha2).
func (src *ProxmoxCluster) ConvertTo(dstRaw conversion.Hub) error { // nolint
	dst := dstRaw.(*v1alpha2.ProxmoxCluster)
	if err := Convert_v1alpha1_ProxmoxCluster_To_v1alpha2_ProxmoxCluster(src, dst, nil); err != nil {
		return err
	}

	// Manually restore data.
	restored := &v1alpha2.ProxmoxCluster{}
	if ok, err := utilconversion.UnmarshalData(src, restored); err != nil || !ok {
		return err
	}

	return nil
}

// ConvertFrom converts from the Hub version (v1alpha1) to this version.
func (dst *ProxmoxCluster) ConvertFrom(srcRaw conversion.Hub) error { // nolint
	src := srcRaw.(*v1alpha2.ProxmoxCluster)
	if err := Convert_v1alpha2_ProxmoxCluster_To_v1alpha1_ProxmoxCluster(src, dst, nil); err != nil {
		return err
	}

	// Preserve Hub data on down-conversion.
	if err := utilconversion.MarshalData(src, dst); err != nil {
		return err
	}
	return nil
}

// ConvertTo converts this ProxmoxClusterList to the Hub version (v1alpha2).
func (src *ProxmoxClusterList) ConvertTo(dstRaw conversion.Hub) error { // nolint
	dst := dstRaw.(*v1alpha2.ProxmoxClusterList)
	return Convert_v1alpha1_ProxmoxClusterList_To_v1alpha2_ProxmoxClusterList(src, dst, nil)
}

// ConvertFrom converts from the Hub version (v1alpha1) to this version.
func (dst *ProxmoxClusterList) ConvertFrom(srcRaw conversion.Hub) error { // nolint
	src := srcRaw.(*v1alpha2.ProxmoxClusterList)
	return Convert_v1alpha2_ProxmoxClusterList_To_v1alpha1_ProxmoxClusterList(src, dst, nil)
}

// Convert_v1alpha2_NetworkConfig_To_v1alpha1_NetworkConfig is an autogenerated conversion function.
func Convert_v1alpha2_NetworkConfig_To_v1alpha1_NetworkConfig(in *v1alpha2.NetworkConfig, out *NetworkConfig, s apiconversion.Scope) error {
	return autoConvert_v1alpha2_NetworkConfig_To_v1alpha1_NetworkConfig(in, out, s)
}

// Convert_v1alpha1_NetworkConfig_To_v1alpha2_NetworkConfig is an autogenerated conversion function.
func Convert_v1alpha1_NetworkConfig_To_v1alpha2_NetworkConfig(in *NetworkConfig, out *v1alpha2.NetworkConfig, s apiconversion.Scope) error {
	return autoConvert_v1alpha1_NetworkConfig_To_v1alpha2_NetworkConfig(in, out, s)
}

// Convert_v1alpha2_IPConfig_To_v1alpha1_IPConfig is an autogenerated conversion function.
func Convert_v1alpha2_IPConfig_To_v1alpha1_IPConfig(in *v1alpha2.IPConfig, out *IPConfig, s apiconversion.Scope) error {
	return autoConvert_v1alpha2_IPConfig_To_v1alpha1_IPConfig(in, out, s)
}
