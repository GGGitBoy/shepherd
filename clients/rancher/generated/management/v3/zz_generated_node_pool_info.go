package client

const (
	NodePoolInfoType                       = "nodePoolInfo"
	NodePoolInfoFieldAutoRenew             = "auto_renew"
	NodePoolInfoFieldAutoRenewPeriod       = "auto_renew_period"
	NodePoolInfoFieldDataDisk              = "data_disk"
	NodePoolInfoFieldEipBandwidth          = "eip_bandwidth"
	NodePoolInfoFieldEipInternetChargeType = "eip_internet_charge_type"
	NodePoolInfoFieldInstanceChargeType    = "instance_charge_type"
	NodePoolInfoFieldInstanceTypes         = "instance_types"
	NodePoolInfoFieldInstancesNum          = "instances_num"
	NodePoolInfoFieldIsBondEip             = "is_bond_eip"
	NodePoolInfoFieldKeyPair               = "key_pair"
	NodePoolInfoFieldLoginPassword         = "login_password"
	NodePoolInfoFieldName                  = "name"
	NodePoolInfoFieldNodepoolId            = "nodepool_id"
	NodePoolInfoFieldPeriod                = "period"
	NodePoolInfoFieldPeriodUnit            = "period_unit"
	NodePoolInfoFieldPlatform              = "platform"
	NodePoolInfoFieldScalingType           = "scaling_type"
	NodePoolInfoFieldSystemDiskCategory    = "system_disk_category"
	NodePoolInfoFieldSystemDiskSize        = "system_disk_size"
	NodePoolInfoFieldVSwitchIds            = "v_switch_ids"
)

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
	NodepoolId            string     `json:"nodepool_id,omitempty" yaml:"nodepool_id,omitempty"`
	Period                int64      `json:"period,omitempty" yaml:"period,omitempty"`
	PeriodUnit            string     `json:"period_unit,omitempty" yaml:"period_unit,omitempty"`
	Platform              string     `json:"platform,omitempty" yaml:"platform,omitempty"`
	ScalingType           string     `json:"scaling_type,omitempty" yaml:"scaling_type,omitempty"`
	SystemDiskCategory    string     `json:"system_disk_category,omitempty" yaml:"system_disk_category,omitempty"`
	SystemDiskSize        int64      `json:"system_disk_size,omitempty" yaml:"system_disk_size,omitempty"`
	VSwitchIds            []string   `json:"v_switch_ids,omitempty" yaml:"v_switch_ids,omitempty"`
}
