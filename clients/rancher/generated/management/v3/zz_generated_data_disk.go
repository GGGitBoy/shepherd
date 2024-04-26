package client

const (
	DataDiskType          = "dataDisk"
	DataDiskFieldDiskSize = "diskSize"
	DataDiskFieldDiskType = "diskType"
)

type DataDisk struct {
	DiskSize int64  `json:"diskSize,omitempty" yaml:"diskSize,omitempty"`
	DiskType string `json:"diskType,omitempty" yaml:"diskType,omitempty"`
}
