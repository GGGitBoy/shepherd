package ack

import (
	"github.com/rancher/shepherd/clients/rancher"
	management "github.com/rancher/shepherd/clients/rancher/generated/management/v3"
)

// CreateACKHostedCluster is a helper function that creates an ACK hosted cluster.
func CreateACKHostedCluster(client *rancher.Client, displayName, cloudCredentialID string, enableClusterAlerting, enableClusterMonitoring, enableNetworkPolicy, windowsPreferedCluster bool, labels map[string]string) (*management.Cluster, error) {
	ackHostCluster := HostClusterConfig(displayName, cloudCredentialID)
	cluster := &management.Cluster{
		ACKConfig:               ackHostCluster,
		DockerRootDir:           "/var/lib/docker",
		EnableClusterAlerting:   enableClusterAlerting,
		EnableClusterMonitoring: enableClusterMonitoring,
		EnableNetworkPolicy:     &enableNetworkPolicy,
		Labels:                  labels,
		Name:                    displayName,
		WindowsPreferedCluster:  windowsPreferedCluster,
	}

	clusterResp, err := client.Management.Cluster.Create(cluster)
	if err != nil {
		return nil, err
	}

	return clusterResp, err
}
