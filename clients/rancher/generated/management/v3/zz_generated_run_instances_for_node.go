package client

const (
	RunInstancesForNodeType                         = "runInstancesForNode"
	RunInstancesForNodeFieldImageID                 = "imageId"
	RunInstancesForNodeFieldInstanceChargeType      = "instanceChargeType"
	RunInstancesForNodeFieldInstanceCount           = "instanceCount"
	RunInstancesForNodeFieldInstanceName            = "instanceName"
	RunInstancesForNodeFieldInstanceType            = "instanceType"
	RunInstancesForNodeFieldInternetChargeType      = "internetChargeType"
	RunInstancesForNodeFieldInternetMaxBandwidthOut = "internetMaxBandwidthOut"
	RunInstancesForNodeFieldKeyIDs                  = "keyIds"
	RunInstancesForNodeFieldMonitorService          = "monitorService"
	RunInstancesForNodeFieldNodeRole                = "nodeRole"
	RunInstancesForNodeFieldProjectID               = "projectId"
	RunInstancesForNodeFieldPublicIpAssigned        = "publicIpAssigned"
	RunInstancesForNodeFieldSecurityService         = "securityService"
	RunInstancesForNodeFieldSubnetID                = "subnetId"
	RunInstancesForNodeFieldSystemDisk              = "systemDisk"
	RunInstancesForNodeFieldUserData                = "userData"
	RunInstancesForNodeFieldVpcID                   = "vpcId"
	RunInstancesForNodeFieldZone                    = "zone"
)

type RunInstancesForNode struct {
	ImageID                 string    `json:"imageId,omitempty" yaml:"imageId,omitempty"`
	InstanceChargeType      string    `json:"instanceChargeType,omitempty" yaml:"instanceChargeType,omitempty"`
	InstanceCount           int64     `json:"instanceCount,omitempty" yaml:"instanceCount,omitempty"`
	InstanceName            string    `json:"instanceName,omitempty" yaml:"instanceName,omitempty"`
	InstanceType            string    `json:"instanceType,omitempty" yaml:"instanceType,omitempty"`
	InternetChargeType      string    `json:"internetChargeType,omitempty" yaml:"internetChargeType,omitempty"`
	InternetMaxBandwidthOut int64     `json:"internetMaxBandwidthOut,omitempty" yaml:"internetMaxBandwidthOut,omitempty"`
	KeyIDs                  []string  `json:"keyIds,omitempty" yaml:"keyIds,omitempty"`
	MonitorService          bool      `json:"monitorService,omitempty" yaml:"monitorService,omitempty"`
	NodeRole                string    `json:"nodeRole,omitempty" yaml:"nodeRole,omitempty"`
	ProjectID               int64     `json:"projectId,omitempty" yaml:"projectId,omitempty"`
	PublicIpAssigned        bool      `json:"publicIpAssigned,omitempty" yaml:"publicIpAssigned,omitempty"`
	SecurityService         bool      `json:"securityService,omitempty" yaml:"securityService,omitempty"`
	SubnetID                string    `json:"subnetId,omitempty" yaml:"subnetId,omitempty"`
	SystemDisk              *DataDisk `json:"systemDisk,omitempty" yaml:"systemDisk,omitempty"`
	UserData                string    `json:"userData,omitempty" yaml:"userData,omitempty"`
	VpcID                   string    `json:"vpcId,omitempty" yaml:"vpcId,omitempty"`
	Zone                    string    `json:"zone,omitempty" yaml:"zone,omitempty"`
}
