package client

const (
	CCENodeVolumeType      = "cceNodeVolume"
	CCENodeVolumeFieldSize = "size"
	CCENodeVolumeFieldType = "type"
)

type CCENodeVolume struct {
	Size int64  `json:"size,omitempty" yaml:"size,omitempty"`
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
