// *** WARNING: this file was generated by pulumi-gen-eks. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package eks

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Defines how Kubernetes pods are executed in Fargate. See aws.eks.FargateProfileArgs for reference.
type FargateProfile struct {
	// Specify a custom role to use for executing pods in Fargate. Defaults to creating a new role with the `arn:aws:iam::aws:policy/AmazonEKSFargatePodExecutionRolePolicy` policy attached.
	PodExecutionRoleArn *string `pulumi:"podExecutionRoleArn"`
	// Specify the subnets in which to execute Fargate tasks for pods. Defaults to the private subnets associated with the cluster.
	SubnetIds []string `pulumi:"subnetIds"`
}

// FargateProfileInput is an input type that accepts FargateProfileArgs and FargateProfileOutput values.
// You can construct a concrete instance of `FargateProfileInput` via:
//
//          FargateProfileArgs{...}
type FargateProfileInput interface {
	pulumi.Input

	ToFargateProfileOutput() FargateProfileOutput
	ToFargateProfileOutputWithContext(context.Context) FargateProfileOutput
}

// Defines how Kubernetes pods are executed in Fargate. See aws.eks.FargateProfileArgs for reference.
type FargateProfileArgs struct {
	// Specify a custom role to use for executing pods in Fargate. Defaults to creating a new role with the `arn:aws:iam::aws:policy/AmazonEKSFargatePodExecutionRolePolicy` policy attached.
	PodExecutionRoleArn pulumi.StringPtrInput `pulumi:"podExecutionRoleArn"`
	// Specify the subnets in which to execute Fargate tasks for pods. Defaults to the private subnets associated with the cluster.
	SubnetIds pulumi.StringArrayInput `pulumi:"subnetIds"`
}

func (FargateProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FargateProfile)(nil)).Elem()
}

func (i FargateProfileArgs) ToFargateProfileOutput() FargateProfileOutput {
	return i.ToFargateProfileOutputWithContext(context.Background())
}

func (i FargateProfileArgs) ToFargateProfileOutputWithContext(ctx context.Context) FargateProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FargateProfileOutput)
}

// Defines how Kubernetes pods are executed in Fargate. See aws.eks.FargateProfileArgs for reference.
type FargateProfileOutput struct{ *pulumi.OutputState }

func (FargateProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FargateProfile)(nil)).Elem()
}

func (o FargateProfileOutput) ToFargateProfileOutput() FargateProfileOutput {
	return o
}

func (o FargateProfileOutput) ToFargateProfileOutputWithContext(ctx context.Context) FargateProfileOutput {
	return o
}

// Specify a custom role to use for executing pods in Fargate. Defaults to creating a new role with the `arn:aws:iam::aws:policy/AmazonEKSFargatePodExecutionRolePolicy` policy attached.
func (o FargateProfileOutput) PodExecutionRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FargateProfile) *string { return v.PodExecutionRoleArn }).(pulumi.StringPtrOutput)
}

// Specify the subnets in which to execute Fargate tasks for pods. Defaults to the private subnets associated with the cluster.
func (o FargateProfileOutput) SubnetIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v FargateProfile) []string { return v.SubnetIds }).(pulumi.StringArrayOutput)
}

// Represents the AWS credentials to scope a given kubeconfig when using a non-default credential chain.
//
// The options can be used independently, or additively.
//
// A scoped kubeconfig is necessary for certain auth scenarios. For example:
//   1. Assume a role on the default account caller,
//   2. Use an AWS creds profile instead of the default account caller,
//   3. Use an AWS creds creds profile instead of the default account caller,
//      and then assume a given role on the profile. This scenario is also
//      possible by only using a profile, iff the profile includes a role to
//      assume in its settings.
//
// See for more details:
// - https://docs.aws.amazon.com/eks/latest/userguide/create-kubeconfig.html
// - https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-role.html
// - https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-profiles.html
type KubeconfigOptions struct {
	// AWS credential profile name to always use instead of the default AWS credential provider chain.
	//
	// The profile is passed to kubeconfig as an authentication environment setting.
	ProfileName *string `pulumi:"profileName"`
	// Role ARN to assume instead of the default AWS credential provider chain.
	//
	// The role is passed to kubeconfig as an authentication exec argument.
	RoleArn *string `pulumi:"roleArn"`
}

// KubeconfigOptionsInput is an input type that accepts KubeconfigOptionsArgs and KubeconfigOptionsOutput values.
// You can construct a concrete instance of `KubeconfigOptionsInput` via:
//
//          KubeconfigOptionsArgs{...}
type KubeconfigOptionsInput interface {
	pulumi.Input

	ToKubeconfigOptionsOutput() KubeconfigOptionsOutput
	ToKubeconfigOptionsOutputWithContext(context.Context) KubeconfigOptionsOutput
}

// Represents the AWS credentials to scope a given kubeconfig when using a non-default credential chain.
//
// The options can be used independently, or additively.
//
// A scoped kubeconfig is necessary for certain auth scenarios. For example:
//   1. Assume a role on the default account caller,
//   2. Use an AWS creds profile instead of the default account caller,
//   3. Use an AWS creds creds profile instead of the default account caller,
//      and then assume a given role on the profile. This scenario is also
//      possible by only using a profile, iff the profile includes a role to
//      assume in its settings.
//
// See for more details:
// - https://docs.aws.amazon.com/eks/latest/userguide/create-kubeconfig.html
// - https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-role.html
// - https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-profiles.html
type KubeconfigOptionsArgs struct {
	// AWS credential profile name to always use instead of the default AWS credential provider chain.
	//
	// The profile is passed to kubeconfig as an authentication environment setting.
	ProfileName pulumi.StringPtrInput `pulumi:"profileName"`
	// Role ARN to assume instead of the default AWS credential provider chain.
	//
	// The role is passed to kubeconfig as an authentication exec argument.
	RoleArn pulumi.StringPtrInput `pulumi:"roleArn"`
}

func (KubeconfigOptionsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KubeconfigOptions)(nil)).Elem()
}

func (i KubeconfigOptionsArgs) ToKubeconfigOptionsOutput() KubeconfigOptionsOutput {
	return i.ToKubeconfigOptionsOutputWithContext(context.Background())
}

func (i KubeconfigOptionsArgs) ToKubeconfigOptionsOutputWithContext(ctx context.Context) KubeconfigOptionsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubeconfigOptionsOutput)
}

func (i KubeconfigOptionsArgs) ToKubeconfigOptionsPtrOutput() KubeconfigOptionsPtrOutput {
	return i.ToKubeconfigOptionsPtrOutputWithContext(context.Background())
}

func (i KubeconfigOptionsArgs) ToKubeconfigOptionsPtrOutputWithContext(ctx context.Context) KubeconfigOptionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubeconfigOptionsOutput).ToKubeconfigOptionsPtrOutputWithContext(ctx)
}

// KubeconfigOptionsPtrInput is an input type that accepts KubeconfigOptionsArgs, KubeconfigOptionsPtr and KubeconfigOptionsPtrOutput values.
// You can construct a concrete instance of `KubeconfigOptionsPtrInput` via:
//
//          KubeconfigOptionsArgs{...}
//
//  or:
//
//          nil
type KubeconfigOptionsPtrInput interface {
	pulumi.Input

	ToKubeconfigOptionsPtrOutput() KubeconfigOptionsPtrOutput
	ToKubeconfigOptionsPtrOutputWithContext(context.Context) KubeconfigOptionsPtrOutput
}

type kubeconfigOptionsPtrType KubeconfigOptionsArgs

func KubeconfigOptionsPtr(v *KubeconfigOptionsArgs) KubeconfigOptionsPtrInput {
	return (*kubeconfigOptionsPtrType)(v)
}

func (*kubeconfigOptionsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**KubeconfigOptions)(nil)).Elem()
}

func (i *kubeconfigOptionsPtrType) ToKubeconfigOptionsPtrOutput() KubeconfigOptionsPtrOutput {
	return i.ToKubeconfigOptionsPtrOutputWithContext(context.Background())
}

func (i *kubeconfigOptionsPtrType) ToKubeconfigOptionsPtrOutputWithContext(ctx context.Context) KubeconfigOptionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubeconfigOptionsPtrOutput)
}

// Represents the AWS credentials to scope a given kubeconfig when using a non-default credential chain.
//
// The options can be used independently, or additively.
//
// A scoped kubeconfig is necessary for certain auth scenarios. For example:
//   1. Assume a role on the default account caller,
//   2. Use an AWS creds profile instead of the default account caller,
//   3. Use an AWS creds creds profile instead of the default account caller,
//      and then assume a given role on the profile. This scenario is also
//      possible by only using a profile, iff the profile includes a role to
//      assume in its settings.
//
// See for more details:
// - https://docs.aws.amazon.com/eks/latest/userguide/create-kubeconfig.html
// - https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-role.html
// - https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-profiles.html
type KubeconfigOptionsOutput struct{ *pulumi.OutputState }

func (KubeconfigOptionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KubeconfigOptions)(nil)).Elem()
}

func (o KubeconfigOptionsOutput) ToKubeconfigOptionsOutput() KubeconfigOptionsOutput {
	return o
}

func (o KubeconfigOptionsOutput) ToKubeconfigOptionsOutputWithContext(ctx context.Context) KubeconfigOptionsOutput {
	return o
}

func (o KubeconfigOptionsOutput) ToKubeconfigOptionsPtrOutput() KubeconfigOptionsPtrOutput {
	return o.ToKubeconfigOptionsPtrOutputWithContext(context.Background())
}

func (o KubeconfigOptionsOutput) ToKubeconfigOptionsPtrOutputWithContext(ctx context.Context) KubeconfigOptionsPtrOutput {
	return o.ApplyT(func(v KubeconfigOptions) *KubeconfigOptions {
		return &v
	}).(KubeconfigOptionsPtrOutput)
}

// AWS credential profile name to always use instead of the default AWS credential provider chain.
//
// The profile is passed to kubeconfig as an authentication environment setting.
func (o KubeconfigOptionsOutput) ProfileName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KubeconfigOptions) *string { return v.ProfileName }).(pulumi.StringPtrOutput)
}

// Role ARN to assume instead of the default AWS credential provider chain.
//
// The role is passed to kubeconfig as an authentication exec argument.
func (o KubeconfigOptionsOutput) RoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KubeconfigOptions) *string { return v.RoleArn }).(pulumi.StringPtrOutput)
}

type KubeconfigOptionsPtrOutput struct{ *pulumi.OutputState }

func (KubeconfigOptionsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KubeconfigOptions)(nil)).Elem()
}

func (o KubeconfigOptionsPtrOutput) ToKubeconfigOptionsPtrOutput() KubeconfigOptionsPtrOutput {
	return o
}

func (o KubeconfigOptionsPtrOutput) ToKubeconfigOptionsPtrOutputWithContext(ctx context.Context) KubeconfigOptionsPtrOutput {
	return o
}

func (o KubeconfigOptionsPtrOutput) Elem() KubeconfigOptionsOutput {
	return o.ApplyT(func(v *KubeconfigOptions) KubeconfigOptions { return *v }).(KubeconfigOptionsOutput)
}

// AWS credential profile name to always use instead of the default AWS credential provider chain.
//
// The profile is passed to kubeconfig as an authentication environment setting.
func (o KubeconfigOptionsPtrOutput) ProfileName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KubeconfigOptions) *string {
		if v == nil {
			return nil
		}
		return v.ProfileName
	}).(pulumi.StringPtrOutput)
}

// Role ARN to assume instead of the default AWS credential provider chain.
//
// The role is passed to kubeconfig as an authentication exec argument.
func (o KubeconfigOptionsPtrOutput) RoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KubeconfigOptions) *string {
		if v == nil {
			return nil
		}
		return v.RoleArn
	}).(pulumi.StringPtrOutput)
}

// Describes a mapping from an AWS IAM role to a Kubernetes user and groups.
type RoleMapping struct {
	// A list of groups within Kubernetes to which the role is mapped.
	Groups []string `pulumi:"groups"`
	// The ARN of the IAM role to add.
	RoleArn string `pulumi:"roleArn"`
	// The user name within Kubernetes to map to the IAM role. By default, the user name is the ARN of the IAM role.
	Username string `pulumi:"username"`
}

// RoleMappingInput is an input type that accepts RoleMappingArgs and RoleMappingOutput values.
// You can construct a concrete instance of `RoleMappingInput` via:
//
//          RoleMappingArgs{...}
type RoleMappingInput interface {
	pulumi.Input

	ToRoleMappingOutput() RoleMappingOutput
	ToRoleMappingOutputWithContext(context.Context) RoleMappingOutput
}

// Describes a mapping from an AWS IAM role to a Kubernetes user and groups.
type RoleMappingArgs struct {
	// A list of groups within Kubernetes to which the role is mapped.
	Groups pulumi.StringArrayInput `pulumi:"groups"`
	// The ARN of the IAM role to add.
	RoleArn pulumi.StringInput `pulumi:"roleArn"`
	// The user name within Kubernetes to map to the IAM role. By default, the user name is the ARN of the IAM role.
	Username pulumi.StringInput `pulumi:"username"`
}

func (RoleMappingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RoleMapping)(nil)).Elem()
}

func (i RoleMappingArgs) ToRoleMappingOutput() RoleMappingOutput {
	return i.ToRoleMappingOutputWithContext(context.Background())
}

func (i RoleMappingArgs) ToRoleMappingOutputWithContext(ctx context.Context) RoleMappingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoleMappingOutput)
}

// RoleMappingArrayInput is an input type that accepts RoleMappingArray and RoleMappingArrayOutput values.
// You can construct a concrete instance of `RoleMappingArrayInput` via:
//
//          RoleMappingArray{ RoleMappingArgs{...} }
type RoleMappingArrayInput interface {
	pulumi.Input

	ToRoleMappingArrayOutput() RoleMappingArrayOutput
	ToRoleMappingArrayOutputWithContext(context.Context) RoleMappingArrayOutput
}

type RoleMappingArray []RoleMappingInput

func (RoleMappingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RoleMapping)(nil)).Elem()
}

func (i RoleMappingArray) ToRoleMappingArrayOutput() RoleMappingArrayOutput {
	return i.ToRoleMappingArrayOutputWithContext(context.Background())
}

func (i RoleMappingArray) ToRoleMappingArrayOutputWithContext(ctx context.Context) RoleMappingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoleMappingArrayOutput)
}

// Describes a mapping from an AWS IAM role to a Kubernetes user and groups.
type RoleMappingOutput struct{ *pulumi.OutputState }

func (RoleMappingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RoleMapping)(nil)).Elem()
}

func (o RoleMappingOutput) ToRoleMappingOutput() RoleMappingOutput {
	return o
}

func (o RoleMappingOutput) ToRoleMappingOutputWithContext(ctx context.Context) RoleMappingOutput {
	return o
}

// A list of groups within Kubernetes to which the role is mapped.
func (o RoleMappingOutput) Groups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v RoleMapping) []string { return v.Groups }).(pulumi.StringArrayOutput)
}

// The ARN of the IAM role to add.
func (o RoleMappingOutput) RoleArn() pulumi.StringOutput {
	return o.ApplyT(func(v RoleMapping) string { return v.RoleArn }).(pulumi.StringOutput)
}

// The user name within Kubernetes to map to the IAM role. By default, the user name is the ARN of the IAM role.
func (o RoleMappingOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v RoleMapping) string { return v.Username }).(pulumi.StringOutput)
}

type RoleMappingArrayOutput struct{ *pulumi.OutputState }

func (RoleMappingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RoleMapping)(nil)).Elem()
}

func (o RoleMappingArrayOutput) ToRoleMappingArrayOutput() RoleMappingArrayOutput {
	return o
}

func (o RoleMappingArrayOutput) ToRoleMappingArrayOutputWithContext(ctx context.Context) RoleMappingArrayOutput {
	return o
}

func (o RoleMappingArrayOutput) Index(i pulumi.IntInput) RoleMappingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RoleMapping {
		return vs[0].([]RoleMapping)[vs[1].(int)]
	}).(RoleMappingOutput)
}

// Describes a mapping from an AWS IAM user to a Kubernetes user and groups.
type UserMapping struct {
	// A list of groups within Kubernetes to which the user is mapped to.
	Groups []string `pulumi:"groups"`
	// The ARN of the IAM user to add.
	UserArn string `pulumi:"userArn"`
	// The user name within Kubernetes to map to the IAM user. By default, the user name is the ARN of the IAM user.
	Username string `pulumi:"username"`
}

// UserMappingInput is an input type that accepts UserMappingArgs and UserMappingOutput values.
// You can construct a concrete instance of `UserMappingInput` via:
//
//          UserMappingArgs{...}
type UserMappingInput interface {
	pulumi.Input

	ToUserMappingOutput() UserMappingOutput
	ToUserMappingOutputWithContext(context.Context) UserMappingOutput
}

// Describes a mapping from an AWS IAM user to a Kubernetes user and groups.
type UserMappingArgs struct {
	// A list of groups within Kubernetes to which the user is mapped to.
	Groups pulumi.StringArrayInput `pulumi:"groups"`
	// The ARN of the IAM user to add.
	UserArn pulumi.StringInput `pulumi:"userArn"`
	// The user name within Kubernetes to map to the IAM user. By default, the user name is the ARN of the IAM user.
	Username pulumi.StringInput `pulumi:"username"`
}

func (UserMappingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UserMapping)(nil)).Elem()
}

func (i UserMappingArgs) ToUserMappingOutput() UserMappingOutput {
	return i.ToUserMappingOutputWithContext(context.Background())
}

func (i UserMappingArgs) ToUserMappingOutputWithContext(ctx context.Context) UserMappingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserMappingOutput)
}

// UserMappingArrayInput is an input type that accepts UserMappingArray and UserMappingArrayOutput values.
// You can construct a concrete instance of `UserMappingArrayInput` via:
//
//          UserMappingArray{ UserMappingArgs{...} }
type UserMappingArrayInput interface {
	pulumi.Input

	ToUserMappingArrayOutput() UserMappingArrayOutput
	ToUserMappingArrayOutputWithContext(context.Context) UserMappingArrayOutput
}

type UserMappingArray []UserMappingInput

func (UserMappingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]UserMapping)(nil)).Elem()
}

func (i UserMappingArray) ToUserMappingArrayOutput() UserMappingArrayOutput {
	return i.ToUserMappingArrayOutputWithContext(context.Background())
}

func (i UserMappingArray) ToUserMappingArrayOutputWithContext(ctx context.Context) UserMappingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserMappingArrayOutput)
}

// Describes a mapping from an AWS IAM user to a Kubernetes user and groups.
type UserMappingOutput struct{ *pulumi.OutputState }

func (UserMappingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserMapping)(nil)).Elem()
}

func (o UserMappingOutput) ToUserMappingOutput() UserMappingOutput {
	return o
}

func (o UserMappingOutput) ToUserMappingOutputWithContext(ctx context.Context) UserMappingOutput {
	return o
}

// A list of groups within Kubernetes to which the user is mapped to.
func (o UserMappingOutput) Groups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v UserMapping) []string { return v.Groups }).(pulumi.StringArrayOutput)
}

// The ARN of the IAM user to add.
func (o UserMappingOutput) UserArn() pulumi.StringOutput {
	return o.ApplyT(func(v UserMapping) string { return v.UserArn }).(pulumi.StringOutput)
}

// The user name within Kubernetes to map to the IAM user. By default, the user name is the ARN of the IAM user.
func (o UserMappingOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v UserMapping) string { return v.Username }).(pulumi.StringOutput)
}

type UserMappingArrayOutput struct{ *pulumi.OutputState }

func (UserMappingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]UserMapping)(nil)).Elem()
}

func (o UserMappingArrayOutput) ToUserMappingArrayOutput() UserMappingArrayOutput {
	return o
}

func (o UserMappingArrayOutput) ToUserMappingArrayOutputWithContext(ctx context.Context) UserMappingArrayOutput {
	return o
}

func (o UserMappingArrayOutput) Index(i pulumi.IntInput) UserMappingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) UserMapping {
		return vs[0].([]UserMapping)[vs[1].(int)]
	}).(UserMappingOutput)
}

// Describes the configuration options available for the Amazon VPC CNI plugin for Kubernetes.
type VpcCniOptions struct {
	// Specifies that your pods may use subnets and security groups (within the same VPC as your control plane resources) that are independent of your cluster's `resourcesVpcConfig`.
	//
	// Defaults to false.
	CustomNetworkConfig *bool `pulumi:"customNetworkConfig"`
	// Specifies the ENI_CONFIG_LABEL_DEF environment variable value for worker nodes. This is used to tell Kubernetes to automatically apply the ENIConfig for each Availability Zone
	// Ref: https://docs.aws.amazon.com/eks/latest/userguide/cni-custom-network.html (step 5(c))
	//
	// Defaults to the official AWS CNI image in ECR.
	EniConfigLabelDef *string `pulumi:"eniConfigLabelDef"`
	// Used to configure the MTU size for attached ENIs. The valid range is from 576 to 9001.
	//
	// Defaults to 9001.
	EniMtu *int `pulumi:"eniMtu"`
	// Specifies whether an external NAT gateway should be used to provide SNAT of secondary ENI IP addresses. If set to true, the SNAT iptables rule and off-VPC IP rule are not applied, and these rules are removed if they have already been applied.
	//
	// Defaults to false.
	ExternalSnat *bool `pulumi:"externalSnat"`
	// Specifies the container image to use in the AWS CNI cluster DaemonSet.
	//
	// Defaults to the official AWS CNI image in ECR.
	Image *string `pulumi:"image"`
	// Specifies the file path used for logs.
	//
	// Defaults to "stdout" to emit Pod logs for `kubectl logs`.
	LogFile *string `pulumi:"logFile"`
	// Specifies the log level used for logs.
	//
	// Defaults to "DEBUG"
	// See more options: https://git.io/fj92K
	LogLevel *string `pulumi:"logLevel"`
	// Specifies whether NodePort services are enabled on a worker node's primary network interface. This requires additional iptables rules and that the kernel's reverse path filter on the primary interface is set to loose.
	//
	// Defaults to true.
	NodePortSupport *bool `pulumi:"nodePortSupport"`
	// Specifies the veth prefix used to generate the host-side veth device name for the CNI.
	//
	// The prefix can be at most 4 characters long.
	//
	// Defaults to "eni".
	VethPrefix *string `pulumi:"vethPrefix"`
	// Specifies the number of free elastic network interfaces (and all of their available IP addresses) that the ipamD daemon should attempt to keep available for pod assignment on the node.
	//
	// Defaults to 1.
	WarmEniTarget *int `pulumi:"warmEniTarget"`
	// Specifies the number of free IP addresses that the ipamD daemon should attempt to keep available for pod assignment on the node.
	WarmIpTarget *int `pulumi:"warmIpTarget"`
}

// VpcCniOptionsInput is an input type that accepts VpcCniOptionsArgs and VpcCniOptionsOutput values.
// You can construct a concrete instance of `VpcCniOptionsInput` via:
//
//          VpcCniOptionsArgs{...}
type VpcCniOptionsInput interface {
	pulumi.Input

	ToVpcCniOptionsOutput() VpcCniOptionsOutput
	ToVpcCniOptionsOutputWithContext(context.Context) VpcCniOptionsOutput
}

// Describes the configuration options available for the Amazon VPC CNI plugin for Kubernetes.
type VpcCniOptionsArgs struct {
	// Specifies that your pods may use subnets and security groups (within the same VPC as your control plane resources) that are independent of your cluster's `resourcesVpcConfig`.
	//
	// Defaults to false.
	CustomNetworkConfig pulumi.BoolPtrInput `pulumi:"customNetworkConfig"`
	// Specifies the ENI_CONFIG_LABEL_DEF environment variable value for worker nodes. This is used to tell Kubernetes to automatically apply the ENIConfig for each Availability Zone
	// Ref: https://docs.aws.amazon.com/eks/latest/userguide/cni-custom-network.html (step 5(c))
	//
	// Defaults to the official AWS CNI image in ECR.
	EniConfigLabelDef pulumi.StringPtrInput `pulumi:"eniConfigLabelDef"`
	// Used to configure the MTU size for attached ENIs. The valid range is from 576 to 9001.
	//
	// Defaults to 9001.
	EniMtu pulumi.IntPtrInput `pulumi:"eniMtu"`
	// Specifies whether an external NAT gateway should be used to provide SNAT of secondary ENI IP addresses. If set to true, the SNAT iptables rule and off-VPC IP rule are not applied, and these rules are removed if they have already been applied.
	//
	// Defaults to false.
	ExternalSnat pulumi.BoolPtrInput `pulumi:"externalSnat"`
	// Specifies the container image to use in the AWS CNI cluster DaemonSet.
	//
	// Defaults to the official AWS CNI image in ECR.
	Image pulumi.StringPtrInput `pulumi:"image"`
	// Specifies the file path used for logs.
	//
	// Defaults to "stdout" to emit Pod logs for `kubectl logs`.
	LogFile pulumi.StringPtrInput `pulumi:"logFile"`
	// Specifies the log level used for logs.
	//
	// Defaults to "DEBUG"
	// See more options: https://git.io/fj92K
	LogLevel pulumi.StringPtrInput `pulumi:"logLevel"`
	// Specifies whether NodePort services are enabled on a worker node's primary network interface. This requires additional iptables rules and that the kernel's reverse path filter on the primary interface is set to loose.
	//
	// Defaults to true.
	NodePortSupport pulumi.BoolPtrInput `pulumi:"nodePortSupport"`
	// Specifies the veth prefix used to generate the host-side veth device name for the CNI.
	//
	// The prefix can be at most 4 characters long.
	//
	// Defaults to "eni".
	VethPrefix pulumi.StringPtrInput `pulumi:"vethPrefix"`
	// Specifies the number of free elastic network interfaces (and all of their available IP addresses) that the ipamD daemon should attempt to keep available for pod assignment on the node.
	//
	// Defaults to 1.
	WarmEniTarget pulumi.IntPtrInput `pulumi:"warmEniTarget"`
	// Specifies the number of free IP addresses that the ipamD daemon should attempt to keep available for pod assignment on the node.
	WarmIpTarget pulumi.IntPtrInput `pulumi:"warmIpTarget"`
}

func (VpcCniOptionsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VpcCniOptions)(nil)).Elem()
}

func (i VpcCniOptionsArgs) ToVpcCniOptionsOutput() VpcCniOptionsOutput {
	return i.ToVpcCniOptionsOutputWithContext(context.Background())
}

func (i VpcCniOptionsArgs) ToVpcCniOptionsOutputWithContext(ctx context.Context) VpcCniOptionsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcCniOptionsOutput)
}

func (i VpcCniOptionsArgs) ToVpcCniOptionsPtrOutput() VpcCniOptionsPtrOutput {
	return i.ToVpcCniOptionsPtrOutputWithContext(context.Background())
}

func (i VpcCniOptionsArgs) ToVpcCniOptionsPtrOutputWithContext(ctx context.Context) VpcCniOptionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcCniOptionsOutput).ToVpcCniOptionsPtrOutputWithContext(ctx)
}

// VpcCniOptionsPtrInput is an input type that accepts VpcCniOptionsArgs, VpcCniOptionsPtr and VpcCniOptionsPtrOutput values.
// You can construct a concrete instance of `VpcCniOptionsPtrInput` via:
//
//          VpcCniOptionsArgs{...}
//
//  or:
//
//          nil
type VpcCniOptionsPtrInput interface {
	pulumi.Input

	ToVpcCniOptionsPtrOutput() VpcCniOptionsPtrOutput
	ToVpcCniOptionsPtrOutputWithContext(context.Context) VpcCniOptionsPtrOutput
}

type vpcCniOptionsPtrType VpcCniOptionsArgs

func VpcCniOptionsPtr(v *VpcCniOptionsArgs) VpcCniOptionsPtrInput {
	return (*vpcCniOptionsPtrType)(v)
}

func (*vpcCniOptionsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcCniOptions)(nil)).Elem()
}

func (i *vpcCniOptionsPtrType) ToVpcCniOptionsPtrOutput() VpcCniOptionsPtrOutput {
	return i.ToVpcCniOptionsPtrOutputWithContext(context.Background())
}

func (i *vpcCniOptionsPtrType) ToVpcCniOptionsPtrOutputWithContext(ctx context.Context) VpcCniOptionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcCniOptionsPtrOutput)
}

// Describes the configuration options available for the Amazon VPC CNI plugin for Kubernetes.
type VpcCniOptionsOutput struct{ *pulumi.OutputState }

func (VpcCniOptionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpcCniOptions)(nil)).Elem()
}

func (o VpcCniOptionsOutput) ToVpcCniOptionsOutput() VpcCniOptionsOutput {
	return o
}

func (o VpcCniOptionsOutput) ToVpcCniOptionsOutputWithContext(ctx context.Context) VpcCniOptionsOutput {
	return o
}

func (o VpcCniOptionsOutput) ToVpcCniOptionsPtrOutput() VpcCniOptionsPtrOutput {
	return o.ToVpcCniOptionsPtrOutputWithContext(context.Background())
}

func (o VpcCniOptionsOutput) ToVpcCniOptionsPtrOutputWithContext(ctx context.Context) VpcCniOptionsPtrOutput {
	return o.ApplyT(func(v VpcCniOptions) *VpcCniOptions {
		return &v
	}).(VpcCniOptionsPtrOutput)
}

// Specifies that your pods may use subnets and security groups (within the same VPC as your control plane resources) that are independent of your cluster's `resourcesVpcConfig`.
//
// Defaults to false.
func (o VpcCniOptionsOutput) CustomNetworkConfig() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VpcCniOptions) *bool { return v.CustomNetworkConfig }).(pulumi.BoolPtrOutput)
}

// Specifies the ENI_CONFIG_LABEL_DEF environment variable value for worker nodes. This is used to tell Kubernetes to automatically apply the ENIConfig for each Availability Zone
// Ref: https://docs.aws.amazon.com/eks/latest/userguide/cni-custom-network.html (step 5(c))
//
// Defaults to the official AWS CNI image in ECR.
func (o VpcCniOptionsOutput) EniConfigLabelDef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpcCniOptions) *string { return v.EniConfigLabelDef }).(pulumi.StringPtrOutput)
}

// Used to configure the MTU size for attached ENIs. The valid range is from 576 to 9001.
//
// Defaults to 9001.
func (o VpcCniOptionsOutput) EniMtu() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VpcCniOptions) *int { return v.EniMtu }).(pulumi.IntPtrOutput)
}

// Specifies whether an external NAT gateway should be used to provide SNAT of secondary ENI IP addresses. If set to true, the SNAT iptables rule and off-VPC IP rule are not applied, and these rules are removed if they have already been applied.
//
// Defaults to false.
func (o VpcCniOptionsOutput) ExternalSnat() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VpcCniOptions) *bool { return v.ExternalSnat }).(pulumi.BoolPtrOutput)
}

// Specifies the container image to use in the AWS CNI cluster DaemonSet.
//
// Defaults to the official AWS CNI image in ECR.
func (o VpcCniOptionsOutput) Image() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpcCniOptions) *string { return v.Image }).(pulumi.StringPtrOutput)
}

// Specifies the file path used for logs.
//
// Defaults to "stdout" to emit Pod logs for `kubectl logs`.
func (o VpcCniOptionsOutput) LogFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpcCniOptions) *string { return v.LogFile }).(pulumi.StringPtrOutput)
}

// Specifies the log level used for logs.
//
// Defaults to "DEBUG"
// See more options: https://git.io/fj92K
func (o VpcCniOptionsOutput) LogLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpcCniOptions) *string { return v.LogLevel }).(pulumi.StringPtrOutput)
}

// Specifies whether NodePort services are enabled on a worker node's primary network interface. This requires additional iptables rules and that the kernel's reverse path filter on the primary interface is set to loose.
//
// Defaults to true.
func (o VpcCniOptionsOutput) NodePortSupport() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VpcCniOptions) *bool { return v.NodePortSupport }).(pulumi.BoolPtrOutput)
}

// Specifies the veth prefix used to generate the host-side veth device name for the CNI.
//
// The prefix can be at most 4 characters long.
//
// Defaults to "eni".
func (o VpcCniOptionsOutput) VethPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpcCniOptions) *string { return v.VethPrefix }).(pulumi.StringPtrOutput)
}

// Specifies the number of free elastic network interfaces (and all of their available IP addresses) that the ipamD daemon should attempt to keep available for pod assignment on the node.
//
// Defaults to 1.
func (o VpcCniOptionsOutput) WarmEniTarget() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VpcCniOptions) *int { return v.WarmEniTarget }).(pulumi.IntPtrOutput)
}

// Specifies the number of free IP addresses that the ipamD daemon should attempt to keep available for pod assignment on the node.
func (o VpcCniOptionsOutput) WarmIpTarget() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VpcCniOptions) *int { return v.WarmIpTarget }).(pulumi.IntPtrOutput)
}

type VpcCniOptionsPtrOutput struct{ *pulumi.OutputState }

func (VpcCniOptionsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcCniOptions)(nil)).Elem()
}

func (o VpcCniOptionsPtrOutput) ToVpcCniOptionsPtrOutput() VpcCniOptionsPtrOutput {
	return o
}

func (o VpcCniOptionsPtrOutput) ToVpcCniOptionsPtrOutputWithContext(ctx context.Context) VpcCniOptionsPtrOutput {
	return o
}

func (o VpcCniOptionsPtrOutput) Elem() VpcCniOptionsOutput {
	return o.ApplyT(func(v *VpcCniOptions) VpcCniOptions { return *v }).(VpcCniOptionsOutput)
}

// Specifies that your pods may use subnets and security groups (within the same VPC as your control plane resources) that are independent of your cluster's `resourcesVpcConfig`.
//
// Defaults to false.
func (o VpcCniOptionsPtrOutput) CustomNetworkConfig() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VpcCniOptions) *bool {
		if v == nil {
			return nil
		}
		return v.CustomNetworkConfig
	}).(pulumi.BoolPtrOutput)
}

// Specifies the ENI_CONFIG_LABEL_DEF environment variable value for worker nodes. This is used to tell Kubernetes to automatically apply the ENIConfig for each Availability Zone
// Ref: https://docs.aws.amazon.com/eks/latest/userguide/cni-custom-network.html (step 5(c))
//
// Defaults to the official AWS CNI image in ECR.
func (o VpcCniOptionsPtrOutput) EniConfigLabelDef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpcCniOptions) *string {
		if v == nil {
			return nil
		}
		return v.EniConfigLabelDef
	}).(pulumi.StringPtrOutput)
}

// Used to configure the MTU size for attached ENIs. The valid range is from 576 to 9001.
//
// Defaults to 9001.
func (o VpcCniOptionsPtrOutput) EniMtu() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VpcCniOptions) *int {
		if v == nil {
			return nil
		}
		return v.EniMtu
	}).(pulumi.IntPtrOutput)
}

// Specifies whether an external NAT gateway should be used to provide SNAT of secondary ENI IP addresses. If set to true, the SNAT iptables rule and off-VPC IP rule are not applied, and these rules are removed if they have already been applied.
//
// Defaults to false.
func (o VpcCniOptionsPtrOutput) ExternalSnat() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VpcCniOptions) *bool {
		if v == nil {
			return nil
		}
		return v.ExternalSnat
	}).(pulumi.BoolPtrOutput)
}

// Specifies the container image to use in the AWS CNI cluster DaemonSet.
//
// Defaults to the official AWS CNI image in ECR.
func (o VpcCniOptionsPtrOutput) Image() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpcCniOptions) *string {
		if v == nil {
			return nil
		}
		return v.Image
	}).(pulumi.StringPtrOutput)
}

// Specifies the file path used for logs.
//
// Defaults to "stdout" to emit Pod logs for `kubectl logs`.
func (o VpcCniOptionsPtrOutput) LogFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpcCniOptions) *string {
		if v == nil {
			return nil
		}
		return v.LogFile
	}).(pulumi.StringPtrOutput)
}

// Specifies the log level used for logs.
//
// Defaults to "DEBUG"
// See more options: https://git.io/fj92K
func (o VpcCniOptionsPtrOutput) LogLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpcCniOptions) *string {
		if v == nil {
			return nil
		}
		return v.LogLevel
	}).(pulumi.StringPtrOutput)
}

// Specifies whether NodePort services are enabled on a worker node's primary network interface. This requires additional iptables rules and that the kernel's reverse path filter on the primary interface is set to loose.
//
// Defaults to true.
func (o VpcCniOptionsPtrOutput) NodePortSupport() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VpcCniOptions) *bool {
		if v == nil {
			return nil
		}
		return v.NodePortSupport
	}).(pulumi.BoolPtrOutput)
}

// Specifies the veth prefix used to generate the host-side veth device name for the CNI.
//
// The prefix can be at most 4 characters long.
//
// Defaults to "eni".
func (o VpcCniOptionsPtrOutput) VethPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpcCniOptions) *string {
		if v == nil {
			return nil
		}
		return v.VethPrefix
	}).(pulumi.StringPtrOutput)
}

// Specifies the number of free elastic network interfaces (and all of their available IP addresses) that the ipamD daemon should attempt to keep available for pod assignment on the node.
//
// Defaults to 1.
func (o VpcCniOptionsPtrOutput) WarmEniTarget() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VpcCniOptions) *int {
		if v == nil {
			return nil
		}
		return v.WarmEniTarget
	}).(pulumi.IntPtrOutput)
}

// Specifies the number of free IP addresses that the ipamD daemon should attempt to keep available for pod assignment on the node.
func (o VpcCniOptionsPtrOutput) WarmIpTarget() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VpcCniOptions) *int {
		if v == nil {
			return nil
		}
		return v.WarmIpTarget
	}).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(FargateProfileOutput{})
	pulumi.RegisterOutputType(KubeconfigOptionsOutput{})
	pulumi.RegisterOutputType(KubeconfigOptionsPtrOutput{})
	pulumi.RegisterOutputType(RoleMappingOutput{})
	pulumi.RegisterOutputType(RoleMappingArrayOutput{})
	pulumi.RegisterOutputType(UserMappingOutput{})
	pulumi.RegisterOutputType(UserMappingArrayOutput{})
	pulumi.RegisterOutputType(VpcCniOptionsOutput{})
	pulumi.RegisterOutputType(VpcCniOptionsPtrOutput{})
}
