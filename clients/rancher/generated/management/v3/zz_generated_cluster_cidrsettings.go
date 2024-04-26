package client

const (
	ClusterCIDRSettingsType                           = "clusterCIDRSettings"
	ClusterCIDRSettingsFieldClaimExpiredSeconds       = "claimExpiredSeconds"
	ClusterCIDRSettingsFieldClusterCIDR               = "clusterCIDR"
	ClusterCIDRSettingsFieldEniSubnetIDs              = "eniSubnetIds"
	ClusterCIDRSettingsFieldIgnoreClusterCIDRConflict = "ignoreClusterCIDRConflict"
	ClusterCIDRSettingsFieldIgnoreServiceCIDRConflict = "ignoreServiceCIDRConflict"
	ClusterCIDRSettingsFieldMaxClusterServiceNum      = "maxClusterServiceNum"
	ClusterCIDRSettingsFieldMaxNodePodNum             = "maxNodePodNum"
	ClusterCIDRSettingsFieldOsCustomizeType           = "osCustomizeType"
	ClusterCIDRSettingsFieldServiceCIDR               = "serviceCIDR"
	ClusterCIDRSettingsFieldSubnetID                  = "subnetId"
)

type ClusterCIDRSettings struct {
	ClaimExpiredSeconds       int64    `json:"claimExpiredSeconds,omitempty" yaml:"claimExpiredSeconds,omitempty"`
	ClusterCIDR               string   `json:"clusterCIDR,omitempty" yaml:"clusterCIDR,omitempty"`
	EniSubnetIDs              []string `json:"eniSubnetIds,omitempty" yaml:"eniSubnetIds,omitempty"`
	IgnoreClusterCIDRConflict bool     `json:"ignoreClusterCIDRConflict,omitempty" yaml:"ignoreClusterCIDRConflict,omitempty"`
	IgnoreServiceCIDRConflict bool     `json:"ignoreServiceCIDRConflict,omitempty" yaml:"ignoreServiceCIDRConflict,omitempty"`
	MaxClusterServiceNum      int64    `json:"maxClusterServiceNum,omitempty" yaml:"maxClusterServiceNum,omitempty"`
	MaxNodePodNum             int64    `json:"maxNodePodNum,omitempty" yaml:"maxNodePodNum,omitempty"`
	OsCustomizeType           string   `json:"osCustomizeType,omitempty" yaml:"osCustomizeType,omitempty"`
	ServiceCIDR               string   `json:"serviceCIDR,omitempty" yaml:"serviceCIDR,omitempty"`
	SubnetID                  string   `json:"subnetId,omitempty" yaml:"subnetId,omitempty"`
}
