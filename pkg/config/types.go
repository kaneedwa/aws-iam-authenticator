/*
Copyright 2017 by the contributors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package config

type IdentityMapping struct {
	IdentityARN string

	// Username is the username pattern that this instances assuming this
	// role will have in Kubernetes.
	Username string

	// Groups is a list of Kubernetes groups this role will authenticate
	// as (e.g., `system:masters`). Each group name can include placeholders.
	Groups []string
}

// RoleMapping is a mapping of an AWS Role ARN to a Kubernetes username and a
// list of Kubernetes groups. The username and groups are specified as templates
// that may optionally contain two template parameters:
//
//  1) "{{AccountID}}" is the 12 digit AWS ID.
//  2) "{{SessionName}}" is the role session name.
//
// The meaning of SessionName depends on the type of entity assuming the role.
// In the case of an EC2 instance role this will be the EC2 instance ID. In the
// case of a federated role it will be the federated identity (controlled by the
// federated identity provider). In the case of a role assumed directly with
// sts:AssumeRole it will be user controlled.
//
// You can use plain values without parameters to have a more static mapping.
type RoleMapping struct {
	// RoleARN is the AWS Resource Name of the role. (e.g., "arn:aws:iam::000000000000:role/Foo").
	RoleARN string

	// Username is the username pattern that this instances assuming this
	// role will have in Kubernetes.
	Username string

	// Groups is a list of Kubernetes groups this role will authenticate
	// as (e.g., `system:masters`). Each group name can include placeholders.
	Groups []string
}

// UserMapping is a static mapping of a single AWS User ARN to a
// Kubernetes username and a list of Kubernetes groups
type UserMapping struct {
	// UserARN is the AWS Resource Name of the user. (e.g., "arn:aws:iam::000000000000:user/Test").
	UserARN string

	// Username is the Kubernetes username this role will authenticate as (e.g., `mycorp:foo`)
	Username string

	// Groups is a list of Kubernetes groups this role will authenticate as (e.g., `system:masters`)
	Groups []string
}

// Config specifies the configuration for a aws-iam-authenticator server
type Config struct {
	// PartitionID is the AWS partition tokens are valid in. See
	// github.com/aws/aws-sdk-go/aws/endpoints
	// endpoints.DefaultPartitions()
	PartitionID string

	// ClusterID is a unique-per-cluster identifier for your
	// aws-iam-authenticator installation.
	ClusterID string

	// KubeconfigPregenerated is set to `true` when a webhook kubeconfig is
	// pre-generated by running the `init` command, and therefore the
	// `server` shouldn't unnecessarily re-generate a new one.
	KubeconfigPregenerated bool

	// HostPort is the TCP Port on which to listen for authentication checks.
	HostPort int

	// Hostname is the hostname that the server bind to.
	Hostname string

	// GenerateKubeconfigPath is the output path where a generated webhook
	// kubeconfig (for `--authentication-token-webhook-config-file`) will be
	// stored.
	GenerateKubeconfigPath string

	// StateDir is the directory where generated certificates and private keys
	// will be stored. You want these persisted between runs so that your API
	// server webhook configuration doesn't change on restart.
	StateDir string

	// RoleMappings is a list of mappings from AWS IAM Role to
	// Kubernetes username + groups.
	RoleMappings []RoleMapping

	// UserMappings is a list of mappings from AWS IAM User to
	// Kubernetes username + groups.
	UserMappings []UserMapping

	// AutoMappedAWSAccounts is a list of AWS accounts that are allowed without an explicit user/role mapping.
	// IAM ARN from these accounts automatically maps to the Kubernetes username.
	AutoMappedAWSAccounts []string

	// ScrubbedAWSAccounts is a list of AWS accounts that the role ARNs and uids
	// are scrubbed from server log statements
	ScrubbedAWSAccounts []string

	// ServerEC2DescribeInstancesRoleARN is an optional AWS Resource Name for an IAM Role to be assumed
	// before calling ec2:DescribeInstances to determine the private DNS of the calling kubelet (EC2 Instance).
	// If nil, defaults to using the IAM Role attached to the instance where aws-iam-authenticator is
	// running.
	ServerEC2DescribeInstancesRoleARN string

	// Address defines the hostname or IP Address to bind the HTTPS server to listen to. This is useful when creating
	// a local server to handle the authentication request for development.
	Address string

	// Master is an optional param which configures api servers endpoint for listening for new CRDs
	// +optional
	Master string

	// Kubeconfig is an optional param which configures the kubeconfig path for connecting to a specific
	// API server this is useful for local development, allowing you to connect to a remote server.
	// +optional
	Kubeconfig string

	// BackendMode is an ordered list of backends to get mappings from. Comma-delimited list of: MountedFile,EKSConfigMap,CRD
	BackendMode []string

	// Ec2 DescribeInstances rate limiting variables initially set to defaults until we completely
	// understand we don't need to change
	EC2DescribeInstancesQps   int
	EC2DescribeInstancesBurst int
}
