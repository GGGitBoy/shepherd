package client

const (
	CCENodeExtendParamType             = "cceNodeExtendParam"
	CCENodeExtendParamFieldIsAutoRenew = "isAutoRenew"
	CCENodeExtendParamFieldPeriodNum   = "periodNum"
	CCENodeExtendParamFieldPeriodType  = "periodType"
)

type CCENodeExtendParam struct {
	IsAutoRenew string `json:"isAutoRenew,omitempty" yaml:"isAutoRenew,omitempty"`
	PeriodNum   int64  `json:"periodNum,omitempty" yaml:"periodNum,omitempty"`
	PeriodType  string `json:"periodType,omitempty" yaml:"periodType,omitempty"`
}
