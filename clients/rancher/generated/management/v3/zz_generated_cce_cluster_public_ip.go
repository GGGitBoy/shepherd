package client

const (
	CCEClusterPublicIPType           = "cceClusterPublicIP"
	CCEClusterPublicIPFieldCreateEIP = "createEIP"
	CCEClusterPublicIPFieldEip       = "eip"
)

type CCEClusterPublicIP struct {
	CreateEIP bool    `json:"createEIP,omitempty" yaml:"createEIP,omitempty"`
	Eip       *CCEEip `json:"eip,omitempty" yaml:"eip,omitempty"`
}
