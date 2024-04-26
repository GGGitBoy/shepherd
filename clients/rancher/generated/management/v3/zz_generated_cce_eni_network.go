package client

const (
	CCEEniNetworkType         = "cceEniNetwork"
	CCEEniNetworkFieldSubnets = "subnets"
)

type CCEEniNetwork struct {
	Subnets []string `json:"subnets,omitempty" yaml:"subnets,omitempty"`
}
