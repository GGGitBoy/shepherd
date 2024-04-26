package ack

import (
	management "github.com/rancher/shepherd/clients/rancher/generated/management/v3"
	"github.com/rancher/shepherd/pkg/config"
)

const (
	// The json/yaml config key for the ACK hosted cluster config
	ACKClusterConfigConfigurationFileKey = "ackClusterConfig"
)

// ClusterConfig is the configuration needed to create an ACK host cluster
type ClusterConfig struct {
	AliyunCredentialSecret   string         `json:"aliyun_credential_secret,omitempty" yaml:"aliyun_credential_secret,omitempty"`
	CloudMonitorFlags        bool           `json:"cloudMonitorFlags,omitempty" yaml:"cloudMonitorFlags,omitempty"`
	ClusterID                string         `json:"cluster_id,omitempty" yaml:"cluster_id,omitempty"`
	ClusterIsUpgrading       bool           `json:"clusterIsUpgrading,omitempty" yaml:"clusterIsUpgrading,omitempty"`
	ClusterType              string         `json:"clusterType,omitempty" yaml:"clusterType,omitempty"`
	ContainerCidr            string         `json:"containerCidr,omitempty" yaml:"containerCidr,omitempty"`
	DisableRollback          bool           `json:"disableRollback,omitempty" yaml:"disableRollback,omitempty"`
	EndpointPublicAccess     bool           `json:"endpointPublicAccess,omitempty" yaml:"endpointPublicAccess,omitempty"`
	Imported                 bool           `json:"imported,omitempty" yaml:"imported,omitempty"`
	KeyPair                  string         `json:"keyPair,omitempty" yaml:"keyPair,omitempty"`
	KubernetesVersion        string         `json:"kubernetesVersion,omitempty" yaml:"kubernetesVersion,omitempty"`
	LoginPassword            string         `json:"loginPassword,omitempty" yaml:"loginPassword,omitempty"`
	MasterAutoRenew          bool           `json:"masterAutoRenew,omitempty" yaml:"masterAutoRenew,omitempty"`
	MasterAutoRenewPeriod    int64          `json:"masterAutoRenewPeriod,omitempty" yaml:"masterAutoRenewPeriod,omitempty"`
	MasterCount              int64          `json:"masterCount,omitempty" yaml:"masterCount,omitempty"`
	MasterInstanceChargeType string         `json:"masterInstanceChargeType,omitempty" yaml:"masterInstanceChargeType,omitempty"`
	MasterInstanceTypes      []string       `json:"masterInstanceTypes,omitempty" yaml:"masterInstanceTypes,omitempty"`
	MasterPeriod             int64          `json:"masterPeriod,omitempty" yaml:"masterPeriod,omitempty"`
	MasterPeriodUnit         string         `json:"masterPeriodUnit,omitempty" yaml:"masterPeriodUnit,omitempty"`
	MasterSystemDiskCategory string         `json:"masterSystemDiskCategory,omitempty" yaml:"masterSystemDiskCategory,omitempty"`
	MasterSystemDiskSize     int64          `json:"masterSystemDiskSize,omitempty" yaml:"masterSystemDiskSize,omitempty"`
	MasterVswitchIds         []string       `json:"masterVswitchIds,omitempty" yaml:"masterVswitchIds,omitempty"`
	Name                     string         `json:"name,omitempty" yaml:"name,omitempty"`
	NodeCidrMask             int64          `json:"nodeCidrMask,omitempty" yaml:"nodeCidrMask,omitempty"`
	NodePoolList             []NodePoolInfo `json:"node_pool_list,omitempty" yaml:"node_pool_list,omitempty"`
	OsType                   string         `json:"osType,omitempty" yaml:"osType,omitempty"`
	PauseClusterUpgrade      bool           `json:"pauseClusterUpgrade,omitempty" yaml:"pauseClusterUpgrade,omitempty"`
	Platform                 string         `json:"platform,omitempty" yaml:"platform,omitempty"`
	ProxyMode                string         `json:"proxyMode,omitempty" yaml:"proxyMode,omitempty"`
	RegionID                 string         `json:"regionId,omitempty" yaml:"regionId,omitempty"`
	ResourceGroupID          string         `json:"resourceGroupId,omitempty" yaml:"resourceGroupId,omitempty"`
	SSHFlags                 bool           `json:"sshFlags,omitempty" yaml:"sshFlags,omitempty"`
	SecurityGroupID          string         `json:"securityGroupId,omitempty" yaml:"securityGroupId,omitempty"`
	ServiceCidr              string         `json:"serviceCidr,omitempty" yaml:"serviceCidr,omitempty"`
	SnatEntry                bool           `json:"snatEntry,omitempty" yaml:"snatEntry,omitempty"`
	TimeoutMins              int64          `json:"timeoutMins,omitempty" yaml:"timeoutMins,omitempty"`
	VpcID                    string         `json:"vpcId,omitempty" yaml:"vpcId,omitempty"`
	VswitchIds               []string       `json:"vswitchIds,omitempty" yaml:"vswitchIds,omitempty"`
}

// NodePool is the configuration needed to an ACK node pool
type NodePoolInfo struct {
	AutoRenew             bool       `json:"auto_renew,omitempty" yaml:"auto_renew,omitempty"`
	AutoRenewPeriod       int64      `json:"auto_renew_period,omitempty" yaml:"auto_renew_period,omitempty"`
	DataDisk              []DiskInfo `json:"data_disk,omitempty" yaml:"data_disk,omitempty"`
	EipBandwidth          int64      `json:"eip_bandwidth,omitempty" yaml:"eip_bandwidth,omitempty"`
	EipInternetChargeType string     `json:"eip_internet_charge_type,omitempty" yaml:"eip_internet_charge_type,omitempty"`
	InstanceChargeType    string     `json:"instance_charge_type,omitempty" yaml:"instance_charge_type,omitempty"`
	InstanceTypes         []string   `json:"instance_types,omitempty" yaml:"instance_types,omitempty"`
	InstancesNum          int64      `json:"instances_num,omitempty" yaml:"instances_num,omitempty"`
	IsBondEip             bool       `json:"is_bond_eip,omitempty" yaml:"is_bond_eip,omitempty"`
	KeyPair               string     `json:"key_pair,omitempty" yaml:"key_pair,omitempty"`
	LoginPassword         string     `json:"login_password,omitempty" yaml:"login_password,omitempty"`
	Name                  string     `json:"name,omitempty" yaml:"name,omitempty"`
	NodepoolID            string     `json:"nodepool_id,omitempty" yaml:"nodepool_id,omitempty"`
	Period                int64      `json:"period,omitempty" yaml:"period,omitempty"`
	PeriodUnit            string     `json:"period_unit,omitempty" yaml:"period_unit,omitempty"`
	Platform              string     `json:"platform,omitempty" yaml:"platform,omitempty"`
	ScalingType           string     `json:"scaling_type,omitempty" yaml:"scaling_type,omitempty"`
	SystemDiskCategory    string     `json:"system_disk_category,omitempty" yaml:"system_disk_category,omitempty"`
	SystemDiskSize        int64      `json:"system_disk_size,omitempty" yaml:"system_disk_size,omitempty"`
	VSwitchIds            []string   `json:"v_switch_ids,omitempty" yaml:"v_switch_ids,omitempty"`
}

type DiskInfo struct {
	AutoSnapshotPolicyID string `json:"auto_snapshot_policy_id,omitempty" yaml:"auto_snapshot_policy_id,omitempty"`
	Category             string `json:"category,omitempty" yaml:"category,omitempty"`
	Encrypted            string `json:"encrypted,omitempty" yaml:"encrypted,omitempty"`
	Size                 int64  `json:"size,omitempty" yaml:"size,omitempty"`
}

func ackDiskInfoConstructor(dataDisks *[]DiskInfo) []management.DiskInfo {
	var ackDiskInfo []management.DiskInfo
	if len(*dataDisks) > 0 {
		for _, dataDisk := range *dataDisks {
			ackDataDisk := management.DiskInfo{
				Category:             dataDisk.Category,
				Size:                 dataDisk.Size,
				Encrypted:            dataDisk.Encrypted,
				AutoSnapshotPolicyID: dataDisk.AutoSnapshotPolicyID,
			}
			ackDiskInfo = append(ackDiskInfo, ackDataDisk)
		}
	}
	return ackDiskInfo
}

func ackNodePoolConstructor(ackNodePoolConfigs *[]NodePoolInfo) []management.NodePoolInfo {
	var ackNodePools []management.NodePoolInfo
	for _, ackNodePoolConfig := range *ackNodePoolConfigs {
		ackNodePool := management.NodePoolInfo{
			Name:               ackNodePoolConfig.Name,
			InstanceTypes:      ackNodePoolConfig.InstanceTypes,
			InstancesNum:       ackNodePoolConfig.InstancesNum,
			KeyPair:            ackNodePoolConfig.KeyPair,
			Platform:           ackNodePoolConfig.Platform,
			SystemDiskCategory: ackNodePoolConfig.SystemDiskCategory,
			SystemDiskSize:     ackNodePoolConfig.SystemDiskSize,
			DataDisk:           ackDiskInfoConstructor(&ackNodePoolConfig.DataDisk),
			VSwitchIds:         ackNodePoolConfig.VSwitchIds,
		}
		ackNodePools = append(ackNodePools, ackNodePool)
	}
	return ackNodePools
}

func HostClusterConfig(displayName, cloudCredentialID string) *management.ACKClusterConfigSpec {
	var ackClusterConfig ClusterConfig
	config.LoadConfig(ACKClusterConfigConfigurationFileKey, &ackClusterConfig)

	return &management.ACKClusterConfigSpec{
		Name:                     displayName,
		AliyunCredentialSecret:   cloudCredentialID,
		ClusterType:              ackClusterConfig.ClusterType,
		RegionID:                 ackClusterConfig.RegionID,
		ContainerCidr:            ackClusterConfig.ContainerCidr,
		ServiceCidr:              ackClusterConfig.ServiceCidr,
		KubernetesVersion:        ackClusterConfig.KubernetesVersion,
		ProxyMode:                ackClusterConfig.ProxyMode,
		NodeCidrMask:             ackClusterConfig.NodeCidrMask,
		MasterInstanceChargeType: ackClusterConfig.MasterInstanceChargeType,
		MasterAutoRenew:          ackClusterConfig.MasterAutoRenew,
		MasterPeriod:             ackClusterConfig.MasterPeriod,
		SnatEntry:                ackClusterConfig.SnatEntry,
		EndpointPublicAccess:     ackClusterConfig.EndpointPublicAccess,
		MasterAutoRenewPeriod:    ackClusterConfig.MasterAutoRenewPeriod,
		MasterSystemDiskSize:     ackClusterConfig.MasterSystemDiskSize,
		MasterSystemDiskCategory: ackClusterConfig.MasterSystemDiskCategory,
		MasterCount:              ackClusterConfig.MasterCount,
		OsType:                   ackClusterConfig.OsType,
		ResourceGroupID:          ackClusterConfig.ResourceGroupID,
		VpcID:                    ackClusterConfig.VpcID,
		MasterVswitchIds:         ackClusterConfig.MasterVswitchIds,
		KeyPair:                  ackClusterConfig.KeyPair,
		NodePoolList:             ackNodePoolConstructor(&ackClusterConfig.NodePoolList),
	}
}
