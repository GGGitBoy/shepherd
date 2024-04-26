package resources

import (
	management "github.com/rancher/shepherd/clients/rancher/generated/management/v3"
)

// CreateNodePool is a helper function that creates a node pool.
func CreateNodePool(ackHostCluster *management.ACKClusterConfigSpec) (*management.NodePoolInfo, error) {
	nodePoolInfo := management.NodePoolInfo{
		Name:               ackHostCluster.NodePoolList[0].Name,
		InstanceTypes:      ackHostCluster.NodePoolList[0].InstanceTypes,
		InstancesNum:       ackHostCluster.NodePoolList[0].InstancesNum,
		KeyPair:            ackHostCluster.NodePoolList[0].KeyPair,
		Platform:           ackHostCluster.NodePoolList[0].Platform,
		SystemDiskCategory: ackHostCluster.NodePoolList[0].SystemDiskCategory,
		SystemDiskSize:     ackHostCluster.NodePoolList[0].SystemDiskSize,
		DataDisk:           ackHostCluster.NodePoolList[0].DataDisk,
		VSwitchIds:         ackHostCluster.NodePoolList[0].VSwitchIds,
	}

	return &nodePoolInfo, nil
}
