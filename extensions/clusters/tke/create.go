package tke

import (
	"github.com/rancher/shepherd/clients/rancher"
	management "github.com/rancher/shepherd/clients/rancher/generated/management/v3"
)

// CreateTKEHostedCluster is a helper function that creates an TKE hosted cluster
func CreateTKEHostedCluster(client *rancher.Client, displayName, cloudCredentialID string, enableClusterAlerting, enableClusterMonitoring, enableNetworkPolicy, windowsPreferedCluster bool, labels map[string]string) (*management.Cluster, error) {
	tkeHostCluster := HostClusterConfig(displayName, cloudCredentialID)
	cluster := &management.Cluster{
		DockerRootDir:           "/var/lib/docker",
		TKEConfig:               tkeHostCluster,
		Name:                    displayName,
		EnableClusterAlerting:   enableClusterAlerting,
		EnableClusterMonitoring: enableClusterMonitoring,
		EnableNetworkPolicy:     &enableNetworkPolicy,
		Labels:                  labels,
		WindowsPreferedCluster:  windowsPreferedCluster,
	}

	clusterResp, err := client.Management.Cluster.Create(cluster)
	if err != nil {
		return nil, err
	}
	return clusterResp, err
}
