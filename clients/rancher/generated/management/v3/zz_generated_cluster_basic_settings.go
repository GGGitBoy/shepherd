package client

const (
	ClusterBasicSettingsType                    = "clusterBasicSettings"
	ClusterBasicSettingsFieldClusterDescription = "clusterDescription"
	ClusterBasicSettingsFieldClusterLevel       = "clusterLevel"
	ClusterBasicSettingsFieldClusterName        = "clusterName"
	ClusterBasicSettingsFieldClusterOs          = "clusterOs"
	ClusterBasicSettingsFieldClusterType        = "clusterType"
	ClusterBasicSettingsFieldClusterVersion     = "clusterVersion"
	ClusterBasicSettingsFieldIsAutoUpgrade      = "isAutoUpgrade"
	ClusterBasicSettingsFieldProjectID          = "projectId"
	ClusterBasicSettingsFieldTags               = "tags"
	ClusterBasicSettingsFieldVpcID              = "vpcId"
)

type ClusterBasicSettings struct {
	ClusterDescription string   `json:"clusterDescription,omitempty" yaml:"clusterDescription,omitempty"`
	ClusterLevel       string   `json:"clusterLevel,omitempty" yaml:"clusterLevel,omitempty"`
	ClusterName        string   `json:"clusterName,omitempty" yaml:"clusterName,omitempty"`
	ClusterOs          string   `json:"clusterOs,omitempty" yaml:"clusterOs,omitempty"`
	ClusterType        string   `json:"clusterType,omitempty" yaml:"clusterType,omitempty"`
	ClusterVersion     string   `json:"clusterVersion,omitempty" yaml:"clusterVersion,omitempty"`
	IsAutoUpgrade      bool     `json:"isAutoUpgrade,omitempty" yaml:"isAutoUpgrade,omitempty"`
	ProjectID          int64    `json:"projectId,omitempty" yaml:"projectId,omitempty"`
	Tags               []string `json:"tags,omitempty" yaml:"tags,omitempty"`
	VpcID              string   `json:"vpcId,omitempty" yaml:"vpcId,omitempty"`
}
