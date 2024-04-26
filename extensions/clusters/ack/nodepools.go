package ack

import (
	"time"

	"github.com/rancher/shepherd/clients/rancher"
	management "github.com/rancher/shepherd/clients/rancher/generated/management/v3"
	"github.com/sirupsen/logrus"
	kwait "k8s.io/apimachinery/pkg/util/wait"
)

const (
	active = "active"
)

func ScalingACKNodePoolsNodes(client *rancher.Client, cluster *management.Cluster, nodePool *NodePoolInfo) (*management.Cluster, error) {
	updatedCluster, err := updateNodePoolQuantity(client, cluster, nodePool)
	if err != nil {
		return nil, err
	}

	logrus.Infof("Node group has been scaled!")

	return updatedCluster, nil
}

// ScalingACKNodePools is a helper function that tests scaling of an ACK node pool by adding a new one and then deleting it.
func updateNodePoolQuantity(client *rancher.Client, cluster *management.Cluster, nodePool *NodePoolInfo) (*management.Cluster, error) {
	clusterResp, err := client.Management.Cluster.ByID(cluster.ID)
	if err != nil {
		return nil, err
	}

	var ackConfig = clusterResp.ACKConfig
	ackConfig.NodePoolList[0].InstancesNum += nodePool.InstancesNum

	ackHostCluster := &management.Cluster{
		DockerRootDir:           "/var/lib/docker",
		ACKConfig:               ackConfig,
		EnableClusterAlerting:   clusterResp.EnableClusterAlerting,
		EnableClusterMonitoring: clusterResp.EnableClusterMonitoring,
		EnableNetworkPolicy:     clusterResp.EnableNetworkPolicy,
		Labels:                  clusterResp.Labels,
		Name:                    clusterResp.Name,
		WindowsPreferedCluster:  clusterResp.WindowsPreferedCluster,
	}

	logrus.Infof("Scaling the node pool to %v total nodes", ackConfig.NodePoolList[0].InstancesNum)
	updatedCluster, err := client.Management.Cluster.Update(clusterResp, ackHostCluster)
	if err != nil {
		return nil, err
	}

	err = kwait.Poll(500*time.Millisecond, 10*time.Minute, func() (done bool, err error) {
		clusterResp, err := client.Management.Cluster.ByID(updatedCluster.ID)
		if err != nil {
			return false, err
		}

		if clusterResp.State == active && clusterResp.NodeCount == ackConfig.NodePoolList[0].InstancesNum {
			return true, nil
		}

		return false, nil
	})
	if err != nil {
		return nil, err
	}

	return updatedCluster, nil
}
