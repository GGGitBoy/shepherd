package client

const (
	CCEContainerNetworkType      = "cceContainerNetwork"
	CCEContainerNetworkFieldCIDR = "cidr"
	CCEContainerNetworkFieldMode = "mode"
)

type CCEContainerNetwork struct {
	CIDR string `json:"cidr,omitempty" yaml:"cidr,omitempty"`
	Mode string `json:"mode,omitempty" yaml:"mode,omitempty"`
}
