package client

const (
	ACKStatusType                       = "ackStatus"
	ACKStatusFieldPrivateRequiresTunnel = "privateRequiresTunnel"
	ACKStatusFieldUpstreamSpec          = "upstreamSpec"
)

type ACKStatus struct {
	PrivateRequiresTunnel *bool                 `json:"privateRequiresTunnel,omitempty" yaml:"privateRequiresTunnel,omitempty"`
	UpstreamSpec          *ACKClusterConfigSpec `json:"upstreamSpec,omitempty" yaml:"upstreamSpec,omitempty"`
}
