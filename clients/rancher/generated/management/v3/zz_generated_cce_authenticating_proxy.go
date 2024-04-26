package client

const (
	CCEAuthenticatingProxyType            = "cceAuthenticatingProxy"
	CCEAuthenticatingProxyFieldCa         = "ca"
	CCEAuthenticatingProxyFieldCert       = "cert"
	CCEAuthenticatingProxyFieldPrivateKey = "privateKey"
)

type CCEAuthenticatingProxy struct {
	Ca         string `json:"ca,omitempty" yaml:"ca,omitempty"`
	Cert       string `json:"cert,omitempty" yaml:"cert,omitempty"`
	PrivateKey string `json:"privateKey,omitempty" yaml:"privateKey,omitempty"`
}
