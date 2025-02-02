/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type AksObservation struct {

	// Maximum number of pods that can be run on a node, which affects how many IP addresses you will need for each node. Defaults to 30
	MaxPodsPerNode *float64 `json:"maxPodsPerNode,omitempty" tf:"max_pods_per_node,omitempty"`
}

type AksParameters struct {

	// Maximum number of pods that can be run on a node, which affects how many IP addresses you will need for each node. Defaults to 30
	// +kubebuilder:validation:Optional
	MaxPodsPerNode *float64 `json:"maxPodsPerNode,omitempty" tf:"max_pods_per_node,omitempty"`
}

type EksObservation struct {

	// IP address to use for DNS queries within the cluster
	DNSClusterIP *string `json:"dnsClusterIp,omitempty" tf:"dns_cluster_ip,omitempty"`

	// Allow configure the IMDSv2 hop limit, the default is 2
	ImdsHopLimit *float64 `json:"imdsHopLimit,omitempty" tf:"imds_hop_limit,omitempty"`

	// When the value is true both IMDSv1 and IMDSv2 are enabled. Setting the value to false disables permanently IMDSv1 and might affect legacy workloads running on the node created with this configuration. The default is true if the flag isn't provided
	ImdsV1 *bool `json:"imdsV1,omitempty" tf:"imds_v1,omitempty"`

	// Cluster's instance profile ARN used for CAST provisioned nodes
	InstanceProfileArn *string `json:"instanceProfileArn,omitempty" tf:"instance_profile_arn,omitempty"`

	// AWS key pair ID to be used for CAST provisioned nodes. Has priority over ssh_public_key
	KeyPairID *string `json:"keyPairId,omitempty" tf:"key_pair_id,omitempty"`

	// Cluster's security groups configuration for CAST provisioned nodes
	SecurityGroups []*string `json:"securityGroups,omitempty" tf:"security_groups,omitempty"`

	// AWS EBS volume IOPS to be used for CAST provisioned nodes
	VolumeIops *float64 `json:"volumeIops,omitempty" tf:"volume_iops,omitempty"`

	// AWS KMS key ARN for encrypting EBS volume attached to the node
	VolumeKMSKeyArn *string `json:"volumeKmsKeyArn,omitempty" tf:"volume_kms_key_arn,omitempty"`

	// AWS EBS volume throughput in MiB/s to be used for CAST provisioned nodes
	VolumeThroughput *float64 `json:"volumeThroughput,omitempty" tf:"volume_throughput,omitempty"`

	// AWS EBS volume type to be used for CAST provisioned nodes. One of: gp3, io1, io2
	VolumeType *string `json:"volumeType,omitempty" tf:"volume_type,omitempty"`
}

type EksParameters struct {

	// IP address to use for DNS queries within the cluster
	// +kubebuilder:validation:Optional
	DNSClusterIP *string `json:"dnsClusterIp,omitempty" tf:"dns_cluster_ip,omitempty"`

	// Allow configure the IMDSv2 hop limit, the default is 2
	// +kubebuilder:validation:Optional
	ImdsHopLimit *float64 `json:"imdsHopLimit,omitempty" tf:"imds_hop_limit,omitempty"`

	// When the value is true both IMDSv1 and IMDSv2 are enabled. Setting the value to false disables permanently IMDSv1 and might affect legacy workloads running on the node created with this configuration. The default is true if the flag isn't provided
	// +kubebuilder:validation:Optional
	ImdsV1 *bool `json:"imdsV1,omitempty" tf:"imds_v1,omitempty"`

	// Cluster's instance profile ARN used for CAST provisioned nodes
	// +kubebuilder:validation:Required
	InstanceProfileArn *string `json:"instanceProfileArn" tf:"instance_profile_arn,omitempty"`

	// AWS key pair ID to be used for CAST provisioned nodes. Has priority over ssh_public_key
	// +kubebuilder:validation:Optional
	KeyPairID *string `json:"keyPairId,omitempty" tf:"key_pair_id,omitempty"`

	// Cluster's security groups configuration for CAST provisioned nodes
	// +kubebuilder:validation:Required
	SecurityGroups []*string `json:"securityGroups" tf:"security_groups,omitempty"`

	// AWS EBS volume IOPS to be used for CAST provisioned nodes
	// +kubebuilder:validation:Optional
	VolumeIops *float64 `json:"volumeIops,omitempty" tf:"volume_iops,omitempty"`

	// AWS KMS key ARN for encrypting EBS volume attached to the node
	// +kubebuilder:validation:Optional
	VolumeKMSKeyArn *string `json:"volumeKmsKeyArn,omitempty" tf:"volume_kms_key_arn,omitempty"`

	// AWS EBS volume throughput in MiB/s to be used for CAST provisioned nodes
	// +kubebuilder:validation:Optional
	VolumeThroughput *float64 `json:"volumeThroughput,omitempty" tf:"volume_throughput,omitempty"`

	// AWS EBS volume type to be used for CAST provisioned nodes. One of: gp3, io1, io2
	// +kubebuilder:validation:Optional
	VolumeType *string `json:"volumeType,omitempty" tf:"volume_type,omitempty"`
}

type GkeObservation struct {

	// Type of boot disk attached to the node. (See [disk types](https://cloud.google.com/compute/docs/disks#pdspecs)). One of: pd-standard, pd-balanced, pd-ssd, pd-extreme
	DiskType *string `json:"diskType,omitempty" tf:"disk_type,omitempty"`

	// Maximum number of pods that can be run on a node, which affects how many IP addresses you will need for each node. Defaults to 110
	MaxPodsPerNode *float64 `json:"maxPodsPerNode,omitempty" tf:"max_pods_per_node,omitempty"`

	// Network tags to be added on a VM. (See [network tags](https://cloud.google.com/vpc/docs/add-remove-network-tags))
	NetworkTags []*string `json:"networkTags,omitempty" tf:"network_tags,omitempty"`
}

type GkeParameters struct {

	// Type of boot disk attached to the node. (See [disk types](https://cloud.google.com/compute/docs/disks#pdspecs)). One of: pd-standard, pd-balanced, pd-ssd, pd-extreme
	// +kubebuilder:validation:Optional
	DiskType *string `json:"diskType,omitempty" tf:"disk_type,omitempty"`

	// Maximum number of pods that can be run on a node, which affects how many IP addresses you will need for each node. Defaults to 110
	// +kubebuilder:validation:Optional
	MaxPodsPerNode *float64 `json:"maxPodsPerNode,omitempty" tf:"max_pods_per_node,omitempty"`

	// Network tags to be added on a VM. (See [network tags](https://cloud.google.com/vpc/docs/add-remove-network-tags))
	// +kubebuilder:validation:Optional
	NetworkTags []*string `json:"networkTags,omitempty" tf:"network_tags,omitempty"`
}

type KopsObservation struct {

	// AWS key pair ID to be used for provisioned nodes. Has priority over sshPublicKey
	KeyPairID *string `json:"keyPairId,omitempty" tf:"key_pair_id,omitempty"`
}

type KopsParameters struct {

	// AWS key pair ID to be used for provisioned nodes. Has priority over sshPublicKey
	// +kubebuilder:validation:Optional
	KeyPairID *string `json:"keyPairId,omitempty" tf:"key_pair_id,omitempty"`
}

type NodeConfigurationObservation struct {
	Aks []AksObservation `json:"aks,omitempty" tf:"aks,omitempty"`

	// CAST AI cluster id
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// Optional container runtime to be used by kubelet. Applicable for EKS only.  Supported values include: `dockerd`, `containerd`
	ContainerRuntime *string `json:"containerRuntime,omitempty" tf:"container_runtime,omitempty"`

	// Disk to CPU ratio. Sets the number of GiBs to be added for every CPU on the node. Defaults to 0
	DiskCPURatio *float64 `json:"diskCpuRatio,omitempty" tf:"disk_cpu_ratio,omitempty"`

	// Optional docker daemon configuration properties in JSON format. Provide only properties that you want to override. Applicable for EKS only. [Available values](https://docs.docker.com/engine/reference/commandline/dockerd/#daemon-configuration-file)
	DockerConfig *string `json:"dockerConfig,omitempty" tf:"docker_config,omitempty"`

	Eks []EksObservation `json:"eks,omitempty" tf:"eks,omitempty"`

	Gke []GkeObservation `json:"gke,omitempty" tf:"gke,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Image to be used while provisioning the node. If nothing is provided will be resolved to latest available image based on Kubernetes version if possible
	Image *string `json:"image,omitempty" tf:"image,omitempty"`

	// Init script to be run on your instance at launch. Should not contain any sensitive data. Value should be base64 encoded
	InitScript *string `json:"initScript,omitempty" tf:"init_script,omitempty"`

	Kops []KopsObservation `json:"kops,omitempty" tf:"kops,omitempty"`

	// Optional kubelet configuration properties in JSON format. Provide only properties that you want to override. Applicable for EKS only. [Available values](https://kubernetes.io/docs/reference/config-api/kubelet-config.v1beta1/)
	KubeletConfig *string `json:"kubeletConfig,omitempty" tf:"kubelet_config,omitempty"`

	// Minimal disk size in GiB. Defaults to 100, min 30, max 1000
	MinDiskSize *float64 `json:"minDiskSize,omitempty" tf:"min_disk_size,omitempty"`

	// Name of the node configuration
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// SSH public key to be used for provisioned nodes
	SSHPublicKey *string `json:"sshPublicKey,omitempty" tf:"ssh_public_key,omitempty"`

	// Subnet ids to be used for provisioned nodes
	Subnets []*string `json:"subnets,omitempty" tf:"subnets,omitempty"`

	// Tags to be added on cloud instances for provisioned nodes
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type NodeConfigurationParameters struct {

	// +kubebuilder:validation:Optional
	Aks []AksParameters `json:"aks,omitempty" tf:"aks,omitempty"`

	// CAST AI cluster id
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/crossplane-provider-castai/apis/castai/v1alpha1.EksClusterId
	// +kubebuilder:validation:Optional
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// Reference to a EksClusterId in castai to populate clusterId.
	// +kubebuilder:validation:Optional
	ClusterIDRef *v1.Reference `json:"clusterIdRef,omitempty" tf:"-"`

	// Selector for a EksClusterId in castai to populate clusterId.
	// +kubebuilder:validation:Optional
	ClusterIDSelector *v1.Selector `json:"clusterIdSelector,omitempty" tf:"-"`

	// Optional container runtime to be used by kubelet. Applicable for EKS only.  Supported values include: `dockerd`, `containerd`
	// +kubebuilder:validation:Optional
	ContainerRuntime *string `json:"containerRuntime,omitempty" tf:"container_runtime,omitempty"`

	// Disk to CPU ratio. Sets the number of GiBs to be added for every CPU on the node. Defaults to 0
	// +kubebuilder:validation:Optional
	DiskCPURatio *float64 `json:"diskCpuRatio,omitempty" tf:"disk_cpu_ratio,omitempty"`

	// Optional docker daemon configuration properties in JSON format. Provide only properties that you want to override. Applicable for EKS only. [Available values](https://docs.docker.com/engine/reference/commandline/dockerd/#daemon-configuration-file)
	// +kubebuilder:validation:Optional
	DockerConfig *string `json:"dockerConfig,omitempty" tf:"docker_config,omitempty"`

	// +kubebuilder:validation:Optional
	Eks []EksParameters `json:"eks,omitempty" tf:"eks,omitempty"`

	// +kubebuilder:validation:Optional
	Gke []GkeParameters `json:"gke,omitempty" tf:"gke,omitempty"`

	// Image to be used while provisioning the node. If nothing is provided will be resolved to latest available image based on Kubernetes version if possible
	// +kubebuilder:validation:Optional
	Image *string `json:"image,omitempty" tf:"image,omitempty"`

	// Init script to be run on your instance at launch. Should not contain any sensitive data. Value should be base64 encoded
	// +kubebuilder:validation:Optional
	InitScript *string `json:"initScript,omitempty" tf:"init_script,omitempty"`

	// +kubebuilder:validation:Optional
	Kops []KopsParameters `json:"kops,omitempty" tf:"kops,omitempty"`

	// Optional kubelet configuration properties in JSON format. Provide only properties that you want to override. Applicable for EKS only. [Available values](https://kubernetes.io/docs/reference/config-api/kubelet-config.v1beta1/)
	// +kubebuilder:validation:Optional
	KubeletConfig *string `json:"kubeletConfig,omitempty" tf:"kubelet_config,omitempty"`

	// Minimal disk size in GiB. Defaults to 100, min 30, max 1000
	// +kubebuilder:validation:Optional
	MinDiskSize *float64 `json:"minDiskSize,omitempty" tf:"min_disk_size,omitempty"`

	// Name of the node configuration
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// SSH public key to be used for provisioned nodes
	// +kubebuilder:validation:Optional
	SSHPublicKey *string `json:"sshPublicKey,omitempty" tf:"ssh_public_key,omitempty"`

	// Subnet ids to be used for provisioned nodes
	// +kubebuilder:validation:Optional
	Subnets []*string `json:"subnets,omitempty" tf:"subnets,omitempty"`

	// Tags to be added on cloud instances for provisioned nodes
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// NodeConfigurationSpec defines the desired state of NodeConfiguration
type NodeConfigurationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NodeConfigurationParameters `json:"forProvider"`
}

// NodeConfigurationStatus defines the observed state of NodeConfiguration.
type NodeConfigurationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NodeConfigurationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// NodeConfiguration is the Schema for the NodeConfigurations API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,castai}
type NodeConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.name)",message="name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.subnets)",message="subnets is a required parameter"
	Spec   NodeConfigurationSpec   `json:"spec"`
	Status NodeConfigurationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NodeConfigurationList contains a list of NodeConfigurations
type NodeConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NodeConfiguration `json:"items"`
}

// Repository type metadata.
var (
	NodeConfiguration_Kind             = "NodeConfiguration"
	NodeConfiguration_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NodeConfiguration_Kind}.String()
	NodeConfiguration_KindAPIVersion   = NodeConfiguration_Kind + "." + CRDGroupVersion.String()
	NodeConfiguration_GroupVersionKind = CRDGroupVersion.WithKind(NodeConfiguration_Kind)
)

func init() {
	SchemeBuilder.Register(&NodeConfiguration{}, &NodeConfigurationList{})
}
