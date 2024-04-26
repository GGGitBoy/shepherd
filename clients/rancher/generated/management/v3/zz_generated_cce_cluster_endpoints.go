package client

const (
	CCEClusterEndpointsType      = "cceClusterEndpoints"
	CCEClusterEndpointsFieldType = "type"
	CCEClusterEndpointsFieldUrl  = "url"
)

type CCEClusterEndpoints struct {
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
	Url  string `json:"url,omitempty" yaml:"url,omitempty"`
}
