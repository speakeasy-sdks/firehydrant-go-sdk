// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

// PostV1SignalsEmailTargetsType - The type of target that the inbound email will notify when matched.
type PostV1SignalsEmailTargetsType string

const (
	PostV1SignalsEmailTargetsTypeTeam             PostV1SignalsEmailTargetsType = "Team"
	PostV1SignalsEmailTargetsTypeEntireTeam       PostV1SignalsEmailTargetsType = "EntireTeam"
	PostV1SignalsEmailTargetsTypeEscalationPolicy PostV1SignalsEmailTargetsType = "EscalationPolicy"
	PostV1SignalsEmailTargetsTypeOnCallSchedule   PostV1SignalsEmailTargetsType = "OnCallSchedule"
	PostV1SignalsEmailTargetsTypeUser             PostV1SignalsEmailTargetsType = "User"
	PostV1SignalsEmailTargetsTypeSlackChannel     PostV1SignalsEmailTargetsType = "SlackChannel"
	PostV1SignalsEmailTargetsTypeWebhook          PostV1SignalsEmailTargetsType = "Webhook"
)

func (e PostV1SignalsEmailTargetsType) ToPointer() *PostV1SignalsEmailTargetsType {
	return &e
}
func (e *PostV1SignalsEmailTargetsType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Team":
		fallthrough
	case "EntireTeam":
		fallthrough
	case "EscalationPolicy":
		fallthrough
	case "OnCallSchedule":
		fallthrough
	case "User":
		fallthrough
	case "SlackChannel":
		fallthrough
	case "Webhook":
		*e = PostV1SignalsEmailTargetsType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PostV1SignalsEmailTargetsType: %v", v)
	}
}

// Target - The target that the email target will notify. This object must contain a `type`
// field that specifies the type of target and an `id` field that specifies the ID of
// the target. The `type` field must be one of "escalation_policy", "on_call_schedule",
// "team", "user", or "slack_channel".
type Target struct {
	// The type of target that the inbound email will notify when matched.
	Type PostV1SignalsEmailTargetsType `json:"type"`
	// The ID of the target that the inbound email will notify when matched.
	ID string `json:"id"`
}

func (o *Target) GetType() PostV1SignalsEmailTargetsType {
	if o == nil {
		return PostV1SignalsEmailTargetsType("")
	}
	return o.Type
}

func (o *Target) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

// RuleMatchingStrategy - Whether or not all rules must match, or if only one rule must match.
type RuleMatchingStrategy string

const (
	RuleMatchingStrategyAll RuleMatchingStrategy = "all"
	RuleMatchingStrategyAny RuleMatchingStrategy = "any"
)

func (e RuleMatchingStrategy) ToPointer() *RuleMatchingStrategy {
	return &e
}
func (e *RuleMatchingStrategy) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "all":
		fallthrough
	case "any":
		*e = RuleMatchingStrategy(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RuleMatchingStrategy: %v", v)
	}
}

// PostV1SignalsEmailTargets - Create a Signals email target for a team.
type PostV1SignalsEmailTargets struct {
	// The email target's name.
	Name string `json:"name"`
	// The email address that will be listening to events.
	Slug *string `json:"slug,omitempty"`
	// A detailed description of the email target.
	Description *string `json:"description,omitempty"`
	// The target that the email target will notify. This object must contain a `type`
	// field that specifies the type of target and an `id` field that specifies the ID of
	// the target. The `type` field must be one of "escalation_policy", "on_call_schedule",
	// "team", "user", or "slack_channel".
	//
	Target *Target `json:"target,omitempty"`
	// A list of email addresses that are allowed to send events to the target. Must be exact match.
	AllowedSenders []string `json:"allowed_senders,omitempty"`
	// A list of CEL expressions that should be evaluated and matched to determine if the target should be notified.
	Rules []string `json:"rules,omitempty"`
	// Whether or not all rules must match, or if only one rule must match.
	RuleMatchingStrategy *RuleMatchingStrategy `json:"rule_matching_strategy,omitempty"`
	// The CEL expression that defines the status of an incoming email that is sent to the target.
	StatusCel *string `json:"status_cel,omitempty"`
	// The CEL expression that defines the level of an incoming email that is sent to the target.
	LevelCel *string `json:"level_cel,omitempty"`
}

func (o *PostV1SignalsEmailTargets) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *PostV1SignalsEmailTargets) GetSlug() *string {
	if o == nil {
		return nil
	}
	return o.Slug
}

func (o *PostV1SignalsEmailTargets) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *PostV1SignalsEmailTargets) GetTarget() *Target {
	if o == nil {
		return nil
	}
	return o.Target
}

func (o *PostV1SignalsEmailTargets) GetAllowedSenders() []string {
	if o == nil {
		return nil
	}
	return o.AllowedSenders
}

func (o *PostV1SignalsEmailTargets) GetRules() []string {
	if o == nil {
		return nil
	}
	return o.Rules
}

func (o *PostV1SignalsEmailTargets) GetRuleMatchingStrategy() *RuleMatchingStrategy {
	if o == nil {
		return nil
	}
	return o.RuleMatchingStrategy
}

func (o *PostV1SignalsEmailTargets) GetStatusCel() *string {
	if o == nil {
		return nil
	}
	return o.StatusCel
}

func (o *PostV1SignalsEmailTargets) GetLevelCel() *string {
	if o == nil {
		return nil
	}
	return o.LevelCel
}
