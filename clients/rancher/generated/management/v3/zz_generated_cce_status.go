package client

const (
	CCEStatusType                       = "cceStatus"
	CCEStatusFieldClusterExternalIP     = "clusterExternalIP"
	CCEStatusFieldEndpoints             = "endpoints"
	CCEStatusFieldPrivateRequiresTunnel = "privateRequiresTunnel"
	CCEStatusFieldUpstreamSpec          = "upstreamSpec"
)

type CCEStatus struct {
	ClusterExternalIP     string                `json:"clusterExternalIP,omitempty" yaml:"clusterExternalIP,omitempty"`
	Endpoints             []CCEClusterEndpoints `json:"endpoints,omitempty" yaml:"endpoints,omitempty"`
	PrivateRequiresTunnel *bool                 `json:"privateRequiresTunnel,omitempty" yaml:"privateRequiresTunnel,omitempty"`
	UpstreamSpec          *CCEClusterConfigSpec `json:"upstreamSpec,omitempty" yaml:"upstreamSpec,omitempty"`
}
