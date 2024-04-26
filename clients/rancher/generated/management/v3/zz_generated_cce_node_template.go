package client

const (
	CCENodeTemplateType                 = "cceNodeTemplate"
	CCENodeTemplateFieldAvailableZone   = "availableZone"
	CCENodeTemplateFieldBillingMode     = "billingMode"
	CCENodeTemplateFieldDataVolumes     = "dataVolumes"
	CCENodeTemplateFieldExtendParam     = "extendParam"
	CCENodeTemplateFieldFlavor          = "flavor"
	CCENodeTemplateFieldOperatingSystem = "operatingSystem"
	CCENodeTemplateFieldPublicIP        = "publicIP"
	CCENodeTemplateFieldRootVolume      = "rootVolume"
	CCENodeTemplateFieldRuntime         = "runtime"
	CCENodeTemplateFieldSSHKey          = "sshKey"
)

type CCENodeTemplate struct {
	AvailableZone   string              `json:"availableZone,omitempty" yaml:"availableZone,omitempty"`
	BillingMode     int64               `json:"billingMode,omitempty" yaml:"billingMode,omitempty"`
	DataVolumes     []CCENodeVolume     `json:"dataVolumes,omitempty" yaml:"dataVolumes,omitempty"`
	ExtendParam     *CCENodeExtendParam `json:"extendParam,omitempty" yaml:"extendParam,omitempty"`
	Flavor          string              `json:"flavor,omitempty" yaml:"flavor,omitempty"`
	OperatingSystem string              `json:"operatingSystem,omitempty" yaml:"operatingSystem,omitempty"`
	PublicIP        *CCENodePublicIP    `json:"publicIP,omitempty" yaml:"publicIP,omitempty"`
	RootVolume      *CCENodeVolume      `json:"rootVolume,omitempty" yaml:"rootVolume,omitempty"`
	Runtime         string              `json:"runtime,omitempty" yaml:"runtime,omitempty"`
	SSHKey          string              `json:"sshKey,omitempty" yaml:"sshKey,omitempty"`
}
