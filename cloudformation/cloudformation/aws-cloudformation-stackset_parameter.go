// Code generated by "go generate". Please don't change this file directly.

package cloudformation

import (
	"github.com/awslabs/goformation/v6/cloudformation/policies"
)

// StackSet_Parameter AWS CloudFormation Resource (AWS::CloudFormation::StackSet.Parameter)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-stackset-parameter.html
type StackSet_Parameter struct {

	// ParameterKey AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-stackset-parameter.html#cfn-cloudformation-stackset-parameter-parameterkey
	ParameterKey string `json:"ParameterKey"`

	// ParameterValue AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-stackset-parameter.html#cfn-cloudformation-stackset-parameter-parametervalue
	ParameterValue string `json:"ParameterValue"`

	// AWSCloudFormationDeletionPolicy represents a CloudFormation DeletionPolicy
	AWSCloudFormationDeletionPolicy policies.DeletionPolicy `json:"-"`

	// AWSCloudFormationUpdateReplacePolicy represents a CloudFormation UpdateReplacePolicy
	AWSCloudFormationUpdateReplacePolicy policies.UpdateReplacePolicy `json:"-"`

	// AWSCloudFormationDependsOn stores the logical ID of the resources to be created before this resource
	AWSCloudFormationDependsOn []string `json:"-"`

	// AWSCloudFormationMetadata stores structured data associated with this resource
	AWSCloudFormationMetadata map[string]interface{} `json:"-"`

	// AWSCloudFormationCondition stores the logical ID of the condition that must be satisfied for this resource to be created
	AWSCloudFormationCondition string `json:"-"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *StackSet_Parameter) AWSCloudFormationType() string {
	return "AWS::CloudFormation::StackSet.Parameter"
}
