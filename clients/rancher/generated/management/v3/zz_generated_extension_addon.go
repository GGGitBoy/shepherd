package client

const (
	ExtensionAddonType            = "extensionAddon"
	ExtensionAddonFieldAddonName  = "addonName"
	ExtensionAddonFieldAddonParam = "addonParam"
)

type ExtensionAddon struct {
	AddonName  string `json:"addonName,omitempty" yaml:"addonName,omitempty"`
	AddonParam string `json:"addonParam,omitempty" yaml:"addonParam,omitempty"`
}
