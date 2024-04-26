package client

const (
	CCEAuthenticationType                     = "cceAuthentication"
	CCEAuthenticationFieldAuthenticatingProxy = "authenticatingProxy"
	CCEAuthenticationFieldMode                = "mode"
)

type CCEAuthentication struct {
	AuthenticatingProxy *CCEAuthenticatingProxy `json:"authenticatingProxy,omitempty" yaml:"authenticatingProxy,omitempty"`
	Mode                string                  `json:"mode,omitempty" yaml:"mode,omitempty"`
}
