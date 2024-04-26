package client

const (
	CCEHostNetworkType               = "cceHostNetwork"
	CCEHostNetworkFieldSecurityGroup = "securityGroup"
	CCEHostNetworkFieldSubnetID      = "subnetID"
	CCEHostNetworkFieldVpcID         = "vpcID"
)

type CCEHostNetwork struct {
	SecurityGroup string `json:"securityGroup,omitempty" yaml:"securityGroup,omitempty"`
	SubnetID      string `json:"subnetID,omitempty" yaml:"subnetID,omitempty"`
	VpcID         string `json:"vpcID,omitempty" yaml:"vpcID,omitempty"`
}
