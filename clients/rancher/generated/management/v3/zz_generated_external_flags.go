package client

const (
	ExternalFlagsType                  = "externalFlags"
	ExternalFlagsFieldClusterFilePath  = "clusterFilePath"
	ExternalFlagsFieldDisablePortCheck = "disablePortCheck"
	ExternalFlagsFieldLocal            = "local"
	ExternalFlagsFieldUpdateOnly       = "updateOnly"
	ExternalFlagsFieldUseLocalState    = "useLocalState"
)

type ExternalFlags struct {
	ClusterFilePath  string `json:"clusterFilePath,omitempty" yaml:"clusterFilePath,omitempty"`
	DisablePortCheck bool   `json:"disablePortCheck,omitempty" yaml:"disablePortCheck,omitempty"`
	Local            bool   `json:"local,omitempty" yaml:"local,omitempty"`
	UpdateOnly       bool   `json:"updateOnly,omitempty" yaml:"updateOnly,omitempty"`
	UseLocalState    bool   `json:"useLocalState,omitempty" yaml:"useLocalState,omitempty"`
}
