package client

const (
	CCENatGatewayType               = "cceNatGateway"
	CCENatGatewayFieldEnabled       = "enabled"
	CCENatGatewayFieldExistingEIPID = "existingEIPID"
	CCENatGatewayFieldSNatRuleEIP   = "snatRuleEIP"
)

type CCENatGateway struct {
	Enabled       bool    `json:"enabled,omitempty" yaml:"enabled,omitempty"`
	ExistingEIPID string  `json:"existingEIPID,omitempty" yaml:"existingEIPID,omitempty"`
	SNatRuleEIP   *CCEEip `json:"snatRuleEIP,omitempty" yaml:"snatRuleEIP,omitempty"`
}
