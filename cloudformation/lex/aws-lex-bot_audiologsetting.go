// Code generated by "go generate". Please don't change this file directly.

package lex

import (
	"github.com/awslabs/goformation/v6/cloudformation/policies"
)

// Bot_AudioLogSetting AWS CloudFormation Resource (AWS::Lex::Bot.AudioLogSetting)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-audiologsetting.html
type Bot_AudioLogSetting struct {

	// Destination AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-audiologsetting.html#cfn-lex-bot-audiologsetting-destination
	Destination *Bot_AudioLogDestination `json:"Destination"`

	// Enabled AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-audiologsetting.html#cfn-lex-bot-audiologsetting-enabled
	Enabled bool `json:"Enabled"`

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
func (r *Bot_AudioLogSetting) AWSCloudFormationType() string {
	return "AWS::Lex::Bot.AudioLogSetting"
}
