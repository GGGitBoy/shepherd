package client

const (
	CCEEipType           = "cceEip"
	CCEEipFieldBandwidth = "bandwidth"
	CCEEipFieldIptype    = "ipType"
)

type CCEEip struct {
	Bandwidth *CCEEipBandwidth `json:"bandwidth,omitempty" yaml:"bandwidth,omitempty"`
	Iptype    string           `json:"ipType,omitempty" yaml:"ipType,omitempty"`
}
