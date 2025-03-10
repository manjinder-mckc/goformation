// Code generated by "go generate". Please don't change this file directly.

package lex

import (
	"github.com/awslabs/goformation/v6/cloudformation/policies"
)

// Bot_GrammarSlotTypeSetting AWS CloudFormation Resource (AWS::Lex::Bot.GrammarSlotTypeSetting)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-grammarslottypesetting.html
type Bot_GrammarSlotTypeSetting struct {

	// Source AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-grammarslottypesetting.html#cfn-lex-bot-grammarslottypesetting-source
	Source *Bot_GrammarSlotTypeSource `json:"Source,omitempty"`

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
func (r *Bot_GrammarSlotTypeSetting) AWSCloudFormationType() string {
	return "AWS::Lex::Bot.GrammarSlotTypeSetting"
}
