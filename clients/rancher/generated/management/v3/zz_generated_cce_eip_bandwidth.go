package client

const (
	CCEEipBandwidthType            = "cceEipBandwidth"
	CCEEipBandwidthFieldChargeMode = "chargeMode"
	CCEEipBandwidthFieldShareType  = "shareType"
	CCEEipBandwidthFieldSize       = "size"
)

type CCEEipBandwidth struct {
	ChargeMode string `json:"chargeMode,omitempty" yaml:"chargeMode,omitempty"`
	ShareType  string `json:"shareType,omitempty" yaml:"shareType,omitempty"`
	Size       int64  `json:"size,omitempty" yaml:"size,omitempty"`
}
