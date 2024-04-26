package client

const (
	DiskInfoType                      = "diskInfo"
	DiskInfoFieldAutoSnapshotPolicyID = "auto_snapshot_policy_id"
	DiskInfoFieldCategory             = "category"
	DiskInfoFieldEncrypted            = "encrypted"
	DiskInfoFieldSize                 = "size"
)

type DiskInfo struct {
	AutoSnapshotPolicyID string `json:"auto_snapshot_policy_id,omitempty" yaml:"auto_snapshot_policy_id,omitempty"`
	Category             string `json:"category,omitempty" yaml:"category,omitempty"`
	Encrypted            string `json:"encrypted,omitempty" yaml:"encrypted,omitempty"`
	Size                 int64  `json:"size,omitempty" yaml:"size,omitempty"`
}
