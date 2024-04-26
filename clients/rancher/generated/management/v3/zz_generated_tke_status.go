package client

const (
	TKEStatusType                       = "tkeStatus"
	TKEStatusFieldPrivateRequiresTunnel = "privateRequiresTunnel"
	TKEStatusFieldUpstreamSpec          = "upstreamSpec"
)

type TKEStatus struct {
	PrivateRequiresTunnel *bool                 `json:"privateRequiresTunnel,omitempty" yaml:"privateRequiresTunnel,omitempty"`
	UpstreamSpec          *TKEClusterConfigSpec `json:"upstreamSpec,omitempty" yaml:"upstreamSpec,omitempty"`
}
