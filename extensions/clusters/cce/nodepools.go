package cce

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

// updateNodePoolQuantity is a helper method that will update the node pool with the desired quantity.
func updateNodePoolQuantity(client *rancher.Client, cluster *management.Cluster, nodePool *NodePool) (*management.Cluster, error) {
	clusterResp, err := client.Management.Cluster.ByID(cluster.ID)
	if err != nil {
		return nil, err
	}

	var cceConfig = clusterResp.CCEConfig
	cceConfig.NodePools[0].InitialNodeCount += nodePool.InitialNodeCount

	cceHostCluster := &management.Cluster{
		DockerRootDir:           "/var/lib/docker",
		CCEConfig:               cceConfig,
		EnableClusterAlerting:   clusterResp.EnableClusterAlerting,
		EnableClusterMonitoring: clusterResp.EnableClusterMonitoring,
		EnableNetworkPolicy:     clusterResp.EnableNetworkPolicy,
		Labels:                  clusterResp.Labels,
		Name:                    clusterResp.Name,
		WindowsPreferedCluster:  clusterResp.WindowsPreferedCluster,
	}

	logrus.Infof("Scaling the node pool to %v total nodes", cceConfig.NodePools[0].InitialNodeCount)
	updatedCluster, err := client.Management.Cluster.Update(clusterResp, cceHostCluster)
	if err != nil {
		return nil, err
	}

	err = kwait.Poll(500*time.Millisecond, 10*time.Minute, func() (done bool, err error) {
		clusterResp, err := client.Management.Cluster.ByID(updatedCluster.ID)
		if err != nil {
			return false, err
		}

		if clusterResp.State == active && clusterResp.NodeCount == cceConfig.NodePools[0].InitialNodeCount {
			return true, nil
		}

		return false, nil
	})
	if err != nil {
		return nil, err
	}

	return updatedCluster, nil
}

// ScalingCCENodePoolsNodes is a helper function that tests scaling of an CCE node pool by adding a new one and then deleting it.
func ScalingCCENodePoolsNodes(client *rancher.Client, cluster *management.Cluster, nodePool *NodePool) (*management.Cluster, error) {
	updatedCluster, err := updateNodePoolQuantity(client, cluster, nodePool)
	if err != nil {
		return nil, err
	}

	logrus.Infof("Node group has been scaled!")

	return updatedCluster, nil
}
