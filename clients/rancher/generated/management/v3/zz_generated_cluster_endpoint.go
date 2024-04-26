package client

const (
	ClusterEndpointType                     = "clusterEndpoint"
	ClusterEndpointFieldDomain              = "domain"
	ClusterEndpointFieldEnable              = "enable"
	ClusterEndpointFieldExtensiveParameters = "extensiveParameters"
	ClusterEndpointFieldSecurityGroup       = "securityGroup"
	ClusterEndpointFieldSubnetID            = "subnetId"
)

type ClusterEndpoint struct {
	Domain              string `json:"domain,omitempty" yaml:"domain,omitempty"`
	Enable              bool   `json:"enable,omitempty" yaml:"enable,omitempty"`
	ExtensiveParameters string `json:"extensiveParameters,omitempty" yaml:"extensiveParameters,omitempty"`
	SecurityGroup       string `json:"securityGroup,omitempty" yaml:"securityGroup,omitempty"`
	SubnetID            string `json:"subnetId,omitempty" yaml:"subnetId,omitempty"`
}
