package client

const (
	CCEClusterExtendParamType                   = "cceClusterExtendParam"
	CCEClusterExtendParamFieldClusterAZ         = "clusterAZ"
	CCEClusterExtendParamFieldClusterExternalIP = "clusterExternalIP"
	CCEClusterExtendParamFieldIsAutoPay         = "isAutoPay"
	CCEClusterExtendParamFieldIsAutoRenew       = "isAutoRenew"
	CCEClusterExtendParamFieldPeriodNum         = "periodNum"
	CCEClusterExtendParamFieldPeriodType        = "periodType"
)

type CCEClusterExtendParam struct {
	ClusterAZ         string `json:"clusterAZ,omitempty" yaml:"clusterAZ,omitempty"`
	ClusterExternalIP string `json:"clusterExternalIP,omitempty" yaml:"clusterExternalIP,omitempty"`
	IsAutoPay         string `json:"isAutoPay,omitempty" yaml:"isAutoPay,omitempty"`
	IsAutoRenew       string `json:"isAutoRenew,omitempty" yaml:"isAutoRenew,omitempty"`
	PeriodNum         int64  `json:"periodNum,omitempty" yaml:"periodNum,omitempty"`
	PeriodType        string `json:"periodType,omitempty" yaml:"periodType,omitempty"`
}
