// Code generated by "go generate". Please don't change this file directly.

package wafv2

import (
	"github.com/awslabs/goformation/v6/cloudformation/policies"
)

// RuleGroup_SqliMatchStatement AWS CloudFormation Resource (AWS::WAFv2::RuleGroup.SqliMatchStatement)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-sqlimatchstatement.html
type RuleGroup_SqliMatchStatement struct {

	// FieldToMatch AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-sqlimatchstatement.html#cfn-wafv2-rulegroup-sqlimatchstatement-fieldtomatch
	FieldToMatch *RuleGroup_FieldToMatch `json:"FieldToMatch"`

	// TextTransformations AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-sqlimatchstatement.html#cfn-wafv2-rulegroup-sqlimatchstatement-texttransformations
	TextTransformations []RuleGroup_TextTransformation `json:"TextTransformations"`

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
func (r *RuleGroup_SqliMatchStatement) AWSCloudFormationType() string {
	return "AWS::WAFv2::RuleGroup.SqliMatchStatement"
}
