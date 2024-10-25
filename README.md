<div align="center">
    <picture>
        <img src="https://github.com/user-attachments/assets/708ff384-c842-4e04-a928-6a5fdce7e756" width="400">
    </picture>
   <p>Manage incidents from ring to retro</p>
	<p>Developer-friendly & type-safe Go SDK specifically catered to leverage <strong>FireHydrant</strong> API.</p>
   <a href=""><img src="https://img.shields.io/static/v1?label=Docs&message=API Ref&color=000000&style=for-the-badge" /></a>
  <a href="https://opensource.org/licenses/MIT"><img src="https://img.shields.io/badge/License-MIT-blue.svg?style=for-the-badge" /></a>
</div>


<br /><br />
> [!IMPORTANT]
> This SDK is not yet ready for production use. To complete setup please follow the steps outlined in your [workspace](https://app.speakeasy.com/org/speakeasy-self/speakeasy-self). Delete this section before > publishing to a package manager.

<!-- Start Summary [summary] -->
## Summary


<!-- End Summary [summary] -->

<!-- Start Table of Contents [toc] -->
## Table of Contents

* [SDK Installation](#sdk-installation)
* [SDK Example Usage](#sdk-example-usage)
* [Available Resources and Operations](#available-resources-and-operations)
* [Retries](#retries)
* [Error Handling](#error-handling)
* [Server Selection](#server-selection)
* [Custom HTTP Client](#custom-http-client)
* [Authentication](#authentication)
* [Special Types](#special-types)
<!-- End Table of Contents [toc] -->

<!-- Start SDK Installation [installation] -->
## SDK Installation

To add the SDK as a dependency to your project:
```bash
go get firehydrant
```
<!-- End SDK Installation [installation] -->

<!-- Start SDK Example Usage [usage] -->
## SDK Example Usage

### Example

```go
package main

import (
	"context"
	"firehydrant"
	"firehydrant/models/components"
	"log"
)

func main() {
	s := firehydrant.New(
		firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.Incidents.Create(ctx, components.CreateIncident{
		Name: "<value>",
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.IncidentEntity != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

<details open>
<summary>Available methods</summary>

### [Ai](docs/sdks/ai/README.md)

* [GetSummarizeIncidentVote](docs/sdks/ai/README.md#getsummarizeincidentvote)
* [VoteSummary](docs/sdks/ai/README.md#votesummary)
* [UpdatePreferences](docs/sdks/ai/README.md#updatepreferences) - Update preferences
* [GetPreferences](docs/sdks/ai/README.md#getpreferences) - Preferences

### [Alerts](docs/sdks/alerts/README.md)

* [List](docs/sdks/alerts/README.md#list) - Retrieve all alerts
* [Get](docs/sdks/alerts/README.md#get) - Retrieve a single alert

### [Bootstrap](docs/sdks/bootstrap/README.md)

* [Get](docs/sdks/bootstrap/README.md#get)

### [Catalogs](docs/sdks/catalogs/README.md)

* [Refresh](docs/sdks/catalogs/README.md#refresh) - Refresh a catalog.
* [Ingest](docs/sdks/catalogs/README.md#ingest) - Accept catalog data in the configured format.

### [Changes](docs/sdks/changes/README.md)

* [Create](docs/sdks/changes/README.md#create) - Create a new change entry
* [List](docs/sdks/changes/README.md#list) - Lists all changes
* [Delete](docs/sdks/changes/README.md#delete) - Archive a change entry
* [Update](docs/sdks/changes/README.md#update) - Update a change entry
* [CreateIdentity](docs/sdks/changes/README.md#createidentity) - Create an identity for this change
* [ListIdentities](docs/sdks/changes/README.md#listidentities) - Retrieve all identities for the change
* [DeleteIdentity](docs/sdks/changes/README.md#deleteidentity) - Delete an identity
* [UpdateIdentity](docs/sdks/changes/README.md#updateidentity) - Update an identity
* [CreateEvent](docs/sdks/changes/README.md#createevent) - Create a change event
* [ListEvents](docs/sdks/changes/README.md#listevents) - List change events
* [DeleteEvent](docs/sdks/changes/README.md#deleteevent) - Delete a change event
* [UpdateEvent](docs/sdks/changes/README.md#updateevent) - Update a change event
* [GetEvent](docs/sdks/changes/README.md#getevent) - Retrieve a change event

### [ChangeTypes](docs/sdks/changetypes/README.md)

* [List](docs/sdks/changetypes/README.md#list) - Lists all change types

### [ChecklistTemplates](docs/sdks/checklisttemplates/README.md)

* [Create](docs/sdks/checklisttemplates/README.md#create) - Create a checklist template
* [Delete](docs/sdks/checklisttemplates/README.md#delete) - Archive a checklist template
* [List](docs/sdks/checklisttemplates/README.md#list) - List all checklist templates
* [Update](docs/sdks/checklisttemplates/README.md#update) - Update a checklist template
* [Get](docs/sdks/checklisttemplates/README.md#get) - Retrieve a single checklist template

### [Conversations](docs/sdks/conversations/README.md)

* [DeleteReaction](docs/sdks/conversations/README.md#deletereaction) - Archive a reaction
* [CreateReaction](docs/sdks/conversations/README.md#createreaction) - Create a reaction
* [ListReactions](docs/sdks/conversations/README.md#listreactions) - List all reactions for a comment
* [DeleteComment](docs/sdks/conversations/README.md#deletecomment) - Archive a comment
* [UpdateComment](docs/sdks/conversations/README.md#updatecomment) - Update a comment
* [GetComment](docs/sdks/conversations/README.md#getcomment) - Retrieve a single comment
* [CreateComment](docs/sdks/conversations/README.md#createcomment) - Create a comment
* [ListComments](docs/sdks/conversations/README.md#listcomments) - List all comments

### [CurrentUser](docs/sdks/currentuser/README.md)

* [Get](docs/sdks/currentuser/README.md#get) - Retrieve the current user

### [CustomFields](docs/sdks/customfields/README.md)

* [DeleteDefinition](docs/sdks/customfields/README.md#deletedefinition) - Delete a custom field definition
* [UpdateDefinition](docs/sdks/customfields/README.md#updatedefinition) - Update custom field definition
* [CreateDefinition](docs/sdks/customfields/README.md#createdefinition) - Create custom field definition
* [ListDefinitions](docs/sdks/customfields/README.md#listdefinitions) - List custom field definitions
* [GetSelectOptions](docs/sdks/customfields/README.md#getselectoptions) - Get custom field permissible values

### [Entitlements](docs/sdks/entitlements/README.md)

* [List](docs/sdks/entitlements/README.md#list) - Retrieve all entitlements

### [Environments](docs/sdks/environments/README.md)

* [Create](docs/sdks/environments/README.md#create) - Create an environment
* [List](docs/sdks/environments/README.md#list) - List all environments
* [Archive](docs/sdks/environments/README.md#archive) - Archive an environment
* [Update](docs/sdks/environments/README.md#update) - Update an environment
* [Get](docs/sdks/environments/README.md#get) - Retrieve a single environment


### [FormConfigurations](docs/sdks/formconfigurations/README.md)

* [Get](docs/sdks/formconfigurations/README.md#get)

### [Functionalities](docs/sdks/functionalities/README.md)

* [Create](docs/sdks/functionalities/README.md#create) - Create a functionality
* [List](docs/sdks/functionalities/README.md#list) - List all functionalities
* [Archive](docs/sdks/functionalities/README.md#archive) - Archive a functionality
* [Update](docs/sdks/functionalities/README.md#update) - Update a functionality
* [Get](docs/sdks/functionalities/README.md#get) - Retrieve a single functionality
* [ListServices](docs/sdks/functionalities/README.md#listservices)

### [IncidentRoles](docs/sdks/incidentroles/README.md)

* [List](docs/sdks/incidentroles/README.md#list) - List all incident roles
* [Create](docs/sdks/incidentroles/README.md#create) - Create an incident role
* [Archive](docs/sdks/incidentroles/README.md#archive) - Archive an incident role
* [Update](docs/sdks/incidentroles/README.md#update) - Update an incident role
* [Get](docs/sdks/incidentroles/README.md#get) - Retrieve an incident role

### [Incidents](docs/sdks/incidents/README.md)

* [Create](docs/sdks/incidents/README.md#create) - Create an incident
* [List](docs/sdks/incidents/README.md#list) - List all incidents
* [GetChannel](docs/sdks/incidents/README.md#getchannel) - Retrieve chat channel information for an incident
* [Close](docs/sdks/incidents/README.md#close) - Close an incident
* [Resolve](docs/sdks/incidents/README.md#resolve) - Resolve an active incident
* [Archive](docs/sdks/incidents/README.md#archive) - Archive an incident
* [Update](docs/sdks/incidents/README.md#update) - Update an incident
* [Get](docs/sdks/incidents/README.md#get) - Retrieve an incident
* [Unarchive](docs/sdks/incidents/README.md#unarchive) - Unarchives an incident
* [AttachAlert](docs/sdks/incidents/README.md#attachalert) - Attach an alert to an incident
* [ListAlerts](docs/sdks/incidents/README.md#listalerts) - List alerts on an incident
* [SetPrimaryAlert](docs/sdks/incidents/README.md#setprimaryalert) - Set an alert as primary
* [RemoveAlert](docs/sdks/incidents/README.md#removealert) - Remove an alert
* [BulkUpdateMilestones](docs/sdks/incidents/README.md#bulkupdatemilestones) - Update milestone times
* [ListMilestones](docs/sdks/incidents/README.md#listmilestones) - List milestones on an incident
* [AddRelatedChange](docs/sdks/incidents/README.md#addrelatedchange) - Add a related change to an incident
* [GetRelatedChanges](docs/sdks/incidents/README.md#getrelatedchanges) - List related changes on an incident
* [UpdateRelatedChangeEvent](docs/sdks/incidents/README.md#updaterelatedchangeevent) - Update a change attached to an incident
* [AddStatusPage](docs/sdks/incidents/README.md#addstatuspage) - Add a status page to an incident
* [ListStatusPages](docs/sdks/incidents/README.md#liststatuspages) - List status pages on an incident
* [DeleteStatusPage](docs/sdks/incidents/README.md#deletestatuspage) - Remove a status page incident attached to an incident
* [AddTasksFromList](docs/sdks/incidents/README.md#addtasksfromlist) - Adds tasks to incident from task list
* [CreateTask](docs/sdks/incidents/README.md#createtask) - Creates a new task
* [GetTasks](docs/sdks/incidents/README.md#gettasks) - Retrieve tasks
* [DeleteTask](docs/sdks/incidents/README.md#deletetask) - Deletes a task
* [UpdateTask](docs/sdks/incidents/README.md#updatetask) - Update a task
* [GetTask](docs/sdks/incidents/README.md#gettask) - Retrieve a single task for an incident
* [ConvertTask](docs/sdks/incidents/README.md#converttask) - Creates a follow-up from a task, removing the task in the process
* [AddLink](docs/sdks/incidents/README.md#addlink) - Add a link to the incident
* [DeleteLink](docs/sdks/incidents/README.md#deletelink) - Deletes the external incident link
* [UpdateLink](docs/sdks/incidents/README.md#updatelink) - Updates the external incident link
* [GetTranscript](docs/sdks/incidents/README.md#gettranscript) - Lists all of the messages in the incident's transcript
* [DeleteTranscript](docs/sdks/incidents/README.md#deletetranscript)
* [GetSimilarIncidents](docs/sdks/incidents/README.md#getsimilarincidents)
* [GetAttachments](docs/sdks/incidents/README.md#getattachments)
* [AddAttachment](docs/sdks/incidents/README.md#addattachment) - Add an attachment to the incident timeline
* [ListEvents](docs/sdks/incidents/README.md#listevents) - List events for an incident
* [UpdateEvent](docs/sdks/incidents/README.md#updateevent) - Update a single event for an incident
* [GetEvent](docs/sdks/incidents/README.md#getevent) - Retrieve a single event for an incident
* [UpdateVotes](docs/sdks/incidents/README.md#updatevotes) - Update the votes on an object
* [UpdateImpact](docs/sdks/incidents/README.md#updateimpact) - Create a status update for an incident
* [ReplaceImpact](docs/sdks/incidents/README.md#replaceimpact) - Create a status update for an incident
* [AddImpact](docs/sdks/incidents/README.md#addimpact) - Add impacted infrastructure
* [ListImpact](docs/sdks/incidents/README.md#listimpact) - List impacted infrastructure
* [RemoveImpact](docs/sdks/incidents/README.md#removeimpact) - Remove impacted infrastructure on an incident
* [AddNote](docs/sdks/incidents/README.md#addnote) - Add a note to an incident
* [UpdateNote](docs/sdks/incidents/README.md#updatenote) - Update a note
* [AddGenericChatMessage](docs/sdks/incidents/README.md#addgenericchatmessage) - Add a generic chat message to the incident timeline.
* [DeleteChatMessage](docs/sdks/incidents/README.md#deletechatmessage) - Delete an existing generic chat message on an incident.
* [UpdateChatMessage](docs/sdks/incidents/README.md#updatechatmessage) - Update an existing generic chat message on an incident.
* [ListRoleAssignments](docs/sdks/incidents/README.md#listroleassignments) - List all assignees
* [UnassignRole](docs/sdks/incidents/README.md#unassignrole) - Unassign a role
* [AssignTeam](docs/sdks/incidents/README.md#assignteam) - Assign a team
* [DeleteTeamAssignment](docs/sdks/incidents/README.md#deleteteamassignment) - Unassign a team
* [GetUserRole](docs/sdks/incidents/README.md#getuserrole) - Get current role for user
* [GetRelationships](docs/sdks/incidents/README.md#getrelationships) - List any parent/child relationships for an incident

#### [Incidents.Events](docs/sdks/events/README.md)

* [Delete](docs/sdks/events/README.md#delete) - Delete a single event for an incident

#### [Incidents.Events.Votes](docs/sdks/votes/README.md)

* [GetStatus](docs/sdks/votes/README.md#getstatus) - Returns the current vote counts for an object

#### [Incidents.Links](docs/sdks/links/README.md)

* [List](docs/sdks/links/README.md#list) - List all the editable links on an incident

#### [Incidents.RoleAssignments](docs/sdks/roleassignments/README.md)

* [Assign](docs/sdks/roleassignments/README.md#assign) - Assign a role

### [IncidentTags](docs/sdks/incidenttags/README.md)

* [Validate](docs/sdks/incidenttags/README.md#validate) - Validate a list of tags
* [List](docs/sdks/incidenttags/README.md#list) - List all incident tags

### [IncidentTypes](docs/sdks/incidenttypes/README.md)

* [Create](docs/sdks/incidenttypes/README.md#create) - Create an incident type
* [List](docs/sdks/incidenttypes/README.md#list) - List all incident types
* [Archive](docs/sdks/incidenttypes/README.md#archive) - Archive an incident type
* [Update](docs/sdks/incidenttypes/README.md#update) - Update an incident type
* [Get](docs/sdks/incidenttypes/README.md#get) - Retrieve an incident type

### [Infrastructures](docs/sdks/infrastructures/README.md)

* [List](docs/sdks/infrastructures/README.md#list) - Lists functionality, service and environment objects

### [Integrations](docs/sdks/integrations/README.md)

* [List](docs/sdks/integrations/README.md#list)
* [Get](docs/sdks/integrations/README.md#get) - Retrieve a single integration
* [UpdateFieldMap](docs/sdks/integrations/README.md#updatefieldmap) - Update field map
* [GetAvailableFields](docs/sdks/integrations/README.md#getavailablefields) - Get mappable fields
* [GetConnections](docs/sdks/integrations/README.md#getconnections)
* [CreateConnection](docs/sdks/integrations/README.md#createconnection)
* [RefreshConnection](docs/sdks/integrations/README.md#refreshconnection)
* [UpdateConnection](docs/sdks/integrations/README.md#updateconnection)
* [GetStatus](docs/sdks/integrations/README.md#getstatus)
* [UpdateAwsConnection](docs/sdks/integrations/README.md#updateawsconnection) - Update an AWS connection
* [GetAwsConnection](docs/sdks/integrations/README.md#getawsconnection) - Retrieve an AWS connection
* [ListCloudtrailBatches](docs/sdks/integrations/README.md#listcloudtrailbatches) - List CloudTrail batches
* [UpdateCloudtrailBatch](docs/sdks/integrations/README.md#updatecloudtrailbatch) - Update a CloudTrail batch
* [GetCloudtrailBatch](docs/sdks/integrations/README.md#getcloudtrailbatch) - Retrieve a CloudTrail batch
* [GetAwsCloudtrailEvents](docs/sdks/integrations/README.md#getawscloudtrailevents)
* [GetConfluenceCloudSpaceKeys](docs/sdks/integrations/README.md#getconfluencecloudspacekeys) - Lists available space keys
* [GetSlackConnectionWorkspaces](docs/sdks/integrations/README.md#getslackconnectionworkspaces)
* [GetSlackUsergroups](docs/sdks/integrations/README.md#getslackusergroups)
* [CreateEmojiAction](docs/sdks/integrations/README.md#createemojiaction) - Creates a new slack emoji action
* [ListSlackEmojiActions](docs/sdks/integrations/README.md#listslackemojiactions) - Lists all slack emoji actions
* [DeleteSlackEmojiAction](docs/sdks/integrations/README.md#deleteslackemojiaction) - Deletes a slack emoji action
* [UpdateEmojiAction](docs/sdks/integrations/README.md#updateemojiaction) - Updates a slack emoji action
* [GetSlackEmojiAction](docs/sdks/integrations/README.md#getslackemojiaction) - Retrieves a slack emoji action
* [ListStatuspageConnections](docs/sdks/integrations/README.md#liststatuspageconnections) - List Statuspage connections
* [DeleteStatuspageConnection](docs/sdks/integrations/README.md#deletestatuspageconnection) - Delete a Statuspage connection
* [UpdateStatuspageConnection](docs/sdks/integrations/README.md#updatestatuspageconnection) - Update a Statuspage connection
* [GetStatuspageConnection](docs/sdks/integrations/README.md#getstatuspageconnection) - Retrieve a Statuspage connection
* [GetStatuspagePages](docs/sdks/integrations/README.md#getstatuspagepages)

#### [Integrations.Aws](docs/sdks/aws/README.md)


#### [Integrations.Aws.Connections](docs/sdks/firehydrantconnections/README.md)

* [List](docs/sdks/firehydrantconnections/README.md#list) - List AWS connections

#### [Integrations.Zendesk](docs/sdks/zendesk/README.md)

* [Search](docs/sdks/zendesk/README.md#search)

### [Lifecycles](docs/sdks/lifecycles/README.md)

* [CreateMeasurementDefinition](docs/sdks/lifecycles/README.md#createmeasurementdefinition) - Create a measurement definition
* [ListMeasurementDefinitions](docs/sdks/lifecycles/README.md#listmeasurementdefinitions) - List all measurement definitions
* [ArchiveMeasurementDefinition](docs/sdks/lifecycles/README.md#archivemeasurementdefinition) - Archive a measurement definition
* [UpdateMeasurementDefinition](docs/sdks/lifecycles/README.md#updatemeasurementdefinition) - Update a measurement definition
* [GetMeasurementDefinition](docs/sdks/lifecycles/README.md#getmeasurementdefinition) - Retrieve a measurement definition
* [ListPhases](docs/sdks/lifecycles/README.md#listphases) - List all phases and milestones
* [CreateMilestone](docs/sdks/lifecycles/README.md#createmilestone) - Create a milestone
* [DeleteMilestone](docs/sdks/lifecycles/README.md#deletemilestone) - Delete a milestone
* [UpdateMilestone](docs/sdks/lifecycles/README.md#updatemilestone) - Update a milestone

### [Metrics](docs/sdks/metrics/README.md)

* [GetTicketFunnel](docs/sdks/metrics/README.md#getticketfunnel) - List ticket task and follow up creation and completion metrics
* [ListRetrospectives](docs/sdks/metrics/README.md#listretrospectives) - List retrospective metrics
* [ListMilestoneFunnel](docs/sdks/metrics/README.md#listmilestonefunnel) - List funnel metrics
* [GetUserInvolvements](docs/sdks/metrics/README.md#getuserinvolvements) - List user metrics
* [GetIncidentMetrics](docs/sdks/metrics/README.md#getincidentmetrics) - List incident metrics
* [GetMttx](docs/sdks/metrics/README.md#getmttx) - Fetch infrastructure metrics based on custom query
* [List](docs/sdks/metrics/README.md#list) - List metrics for a component type
* [Get](docs/sdks/metrics/README.md#get) - Show metrics for a component

### [Noauth](docs/sdks/noauth/README.md)

* [Ping](docs/sdks/noauth/README.md#ping) - Ping

### [Nunc](docs/sdks/nunc/README.md)

* [DeleteSubscription](docs/sdks/nunc/README.md#deletesubscription) - Unsubscribe from status page updates
* [Subscribe](docs/sdks/nunc/README.md#subscribe) - Subscribe to status page updates

#### [Nunc.Connections](docs/sdks/connections/README.md)

* [Create](docs/sdks/connections/README.md#create) - Create a FireHydrant hosted status page

### [NuncConnections](docs/sdks/nuncconnections/README.md)

* [List](docs/sdks/nuncconnections/README.md#list) - Lists the FireHydrant hosted status pages
* [Subscribe](docs/sdks/nuncconnections/README.md#subscribe) - Subscribes a comma-separated string of emails to status page updates
* [ListSubscribers](docs/sdks/nuncconnections/README.md#listsubscribers) - Retrieves the list of subscribers for a status page
* [Delete](docs/sdks/nuncconnections/README.md#delete) - Delete a FireHydrant hosted status page
* [Get](docs/sdks/nuncconnections/README.md#get) - Retrieve information about a FireHydrant hosted status page
* [DeleteComponentGroup](docs/sdks/nuncconnections/README.md#deletecomponentgroup) - Delete a component group displayed on a FireHydrant status page
* [UpdateComponentGroup](docs/sdks/nuncconnections/README.md#updatecomponentgroup) - Update a component group to be displayed on a FireHydrant status page
* [AddComponentGroup](docs/sdks/nuncconnections/README.md#addcomponentgroup) - Add a component group to be displayed on a FireHydrant status page
* [UpdateLink](docs/sdks/nuncconnections/README.md#updatelink) - Update a link to be displayed on a FireHydrant status page
* [AddLink](docs/sdks/nuncconnections/README.md#addlink) - Update a link to be displayed on a FireHydrant status page
* [UpdateImage](docs/sdks/nuncconnections/README.md#updateimage) - Add or replace an image attached to a FireHydrant status page
* [UnsubscribeSubscribers](docs/sdks/nuncconnections/README.md#unsubscribesubscribers) - Unsubscribes one or more status page subscribers.
* [Update](docs/sdks/nuncconnections/README.md#update) - Update a FireHydrant hosted status page
* [DeleteLink](docs/sdks/nuncconnections/README.md#deletelink) - Delete a link displayed on a FireHydrant status page
* [DeleteImage](docs/sdks/nuncconnections/README.md#deleteimage) - Delete an image attached to a FireHydrant status page

### [Ping](docs/sdks/ping/README.md)

* [Get](docs/sdks/ping/README.md#get) - Ping

### [PostMortems](docs/sdks/postmortems/README.md)

* [CreateReport](docs/sdks/postmortems/README.md#createreport) - Create a report
* [GetReport](docs/sdks/postmortems/README.md#getreport) - Get a report
* [CreateReason](docs/sdks/postmortems/README.md#createreason) - Create contributing factor
* [DeleteReason](docs/sdks/postmortems/README.md#deletereason) - Delete a contributing factor
* [ReorderContributingFactor](docs/sdks/postmortems/README.md#reordercontributingfactor) - Reorder a contributing factor
* [UpdateField](docs/sdks/postmortems/README.md#updatefield) - Update a field
* [ListReports](docs/sdks/postmortems/README.md#listreports) - List all reports
* [UpdateReport](docs/sdks/postmortems/README.md#updatereport) - Update a report
* [ListContributingFactors](docs/sdks/postmortems/README.md#listcontributingfactors) - List contributing factors
* [UpdateReason](docs/sdks/postmortems/README.md#updatereason) - Update a contributing factor
* [PublishReport](docs/sdks/postmortems/README.md#publishreport) - Publish a retrospective report
* [UpdateQuestions](docs/sdks/postmortems/README.md#updatequestions) - Update incident retrospective questions
* [ListQuestions](docs/sdks/postmortems/README.md#listquestions) - List incident retrospective questions
* [GetQuestion](docs/sdks/postmortems/README.md#getquestion) - Get an incident retrospective question configured to be provided and filled out on each retrospective report.

### [Priorities](docs/sdks/priorities/README.md)

* [Create](docs/sdks/priorities/README.md#create) - Create priority
* [List](docs/sdks/priorities/README.md#list) - Lists priorities
* [Delete](docs/sdks/priorities/README.md#delete) - Delete a specific priority
* [Update](docs/sdks/priorities/README.md#update) - Update a specific priority
* [Get](docs/sdks/priorities/README.md#get) - Retrieve a specific priority

### [ProcessingLogEntries](docs/sdks/processinglogentries/README.md)

* [Get](docs/sdks/processinglogentries/README.md#get) - Processing Log Entries for a specific alert

### [Reports](docs/sdks/reports/README.md)

* [GetMeanTime](docs/sdks/reports/README.md#getmeantime) - Get mean time report

### [RunbookAudits](docs/sdks/runbookaudits/README.md)

* [List](docs/sdks/runbookaudits/README.md#list) - List runbook audits

### [Runbooks](docs/sdks/runbooks/README.md)

* [ListActions](docs/sdks/runbooks/README.md#listactions) - List all Runbook actions
* [Execute](docs/sdks/runbooks/README.md#execute) - Create a runbook execution
* [ListExecutions](docs/sdks/runbooks/README.md#listexecutions) - List all executions of Runbooks
* [DeleteExecution](docs/sdks/runbooks/README.md#deleteexecution) - Terminate a runbook execution
* [GetExecution](docs/sdks/runbooks/README.md#getexecution) - Retrieve a runbook execution
* [GetVotingStatus](docs/sdks/runbooks/README.md#getvotingstatus) - Returns the current vote counts for an object
* [UpdateVotes](docs/sdks/runbooks/README.md#updatevotes) - Update the votes on an object
* [UpdateStepExecution](docs/sdks/runbooks/README.md#updatestepexecution) - Updates a runbook step execution
* [GetVoteStatus](docs/sdks/runbooks/README.md#getvotestatus) - Returns the current vote counts for an object
* [GetScript](docs/sdks/runbooks/README.md#getscript) - Retrieves the bash script from a "script" step.
* [UpdateScriptState](docs/sdks/runbooks/README.md#updatescriptstate) - Updates the execution's step.
* [GetSelectOptions](docs/sdks/runbooks/README.md#getselectoptions)
* [Create](docs/sdks/runbooks/README.md#create) - Create a runbook
* [List](docs/sdks/runbooks/README.md#list) - List runbooks
* [Delete](docs/sdks/runbooks/README.md#delete) - Delete a runbook
* [Update](docs/sdks/runbooks/README.md#update) - Update a runbook
* [Get](docs/sdks/runbooks/README.md#get) - Retrieve a runbook

#### [Runbooks.Steps](docs/sdks/steps/README.md)

* [UpdateVotes](docs/sdks/steps/README.md#updatevotes) - Update the votes on an object

### [SavedSearches](docs/sdks/savedsearches/README.md)

* [Delete](docs/sdks/savedsearches/README.md#delete) - Delete a specific saved search
* [Update](docs/sdks/savedsearches/README.md#update) - Update a specific saved search
* [List](docs/sdks/savedsearches/README.md#list) - Lists save searches
* [Get](docs/sdks/savedsearches/README.md#get) - Retrieve a specific save search
* [Create](docs/sdks/savedsearches/README.md#create) - Create saved search

### [ScheduledMaintenances](docs/sdks/scheduledmaintenances/README.md)

* [Create](docs/sdks/scheduledmaintenances/README.md#create) - Create a scheduled maintenance event
* [Update](docs/sdks/scheduledmaintenances/README.md#update) - Update a scheduled maintenance event
* [Get](docs/sdks/scheduledmaintenances/README.md#get) - Retrieve a scheduled maintenance event
* [List](docs/sdks/scheduledmaintenances/README.md#list) - List scheduled maintenance events
* [Delete](docs/sdks/scheduledmaintenances/README.md#delete) - Delete a scheduled maintenance event

### [Schedules](docs/sdks/schedules/README.md)

* [List](docs/sdks/schedules/README.md#list) - List all schedules

### [Scim](docs/sdks/scim/README.md)

* [UpdateGroup](docs/sdks/scim/README.md#updategroup) - Updates a Team (via the Group protocol) and assigns members to that team.
* [DeleteGroup](docs/sdks/scim/README.md#deletegroup) - Deletes a Team (via the Group protocol)
* [GetGroup](docs/sdks/scim/README.md#getgroup) - Lists a Team (via the Group protocol)
* [CreateGroup](docs/sdks/scim/README.md#creategroup) - Creates a new Team (via the Group protocol) and assigns members to that team.
* [ListGroups](docs/sdks/scim/README.md#listgroups) - List all Teams (via the Group protocol)
* [DeleteUser](docs/sdks/scim/README.md#deleteuser) - Deletes a User using SCIM protocol
* [UpdateUser](docs/sdks/scim/README.md#updateuser) - Updates a User using SCIM protocol
* [GetUser](docs/sdks/scim/README.md#getuser) - Lists a User (via the User protocol)
* [CreateUser](docs/sdks/scim/README.md#createuser) - Creates a new User using SCIM protocol
* [ListUsers](docs/sdks/scim/README.md#listusers) - Gets a list of Users using SCIM protocol

### [ServiceDependencies](docs/sdks/servicedependencies/README.md)

* [Create](docs/sdks/servicedependencies/README.md#create) - Create a service dependency
* [Get](docs/sdks/servicedependencies/README.md#get) - Retrieve a single service dependency
* [Delete](docs/sdks/servicedependencies/README.md#delete) - Delete a service dependency
* [Update](docs/sdks/servicedependencies/README.md#update) - Update a service dependency

### [Services](docs/sdks/services/README.md)

* [Create](docs/sdks/services/README.md#create) - Create a service
* [List](docs/sdks/services/README.md#list) - List all services
* [BulkCreateServiceLinks](docs/sdks/services/README.md#bulkcreateservicelinks) - Bulk create services and service links
* [Delete](docs/sdks/services/README.md#delete) - Delete a service
* [Update](docs/sdks/services/README.md#update) - Update a service
* [Get](docs/sdks/services/README.md#get) - Retrieve a single service
* [ListDependencies](docs/sdks/services/README.md#listdependencies) - Retrieve a service's dependencies
* [ListAvailableUpstreamDependencies](docs/sdks/services/README.md#listavailableupstreamdependencies) - Retrieve all available upstream dependencies
* [ListAvailableDownstreamDependencies](docs/sdks/services/README.md#listavailabledownstreamdependencies) - Retrieve all available downstream dependencies
* [DeleteServiceLink](docs/sdks/services/README.md#deleteservicelink) - Delete a service link
* [CreateChecklistResponse](docs/sdks/services/README.md#createchecklistresponse) - Record a response for a checklist item

### [Severities](docs/sdks/severities/README.md)

* [Create](docs/sdks/severities/README.md#create) - Create severity
* [List](docs/sdks/severities/README.md#list) - Lists severities
* [Delete](docs/sdks/severities/README.md#delete) - Delete a specific severity
* [Update](docs/sdks/severities/README.md#update) - Update a specific severity
* [Get](docs/sdks/severities/README.md#get) - Retrieve a specific severity

### [SeverityMatrix](docs/sdks/severitymatrix/README.md)

* [Update](docs/sdks/severitymatrix/README.md#update) - Update your severity matrix
* [DeleteCondition](docs/sdks/severitymatrix/README.md#deletecondition) - Delete a specific condition
* [GetCondition](docs/sdks/severitymatrix/README.md#getcondition)
* [Get](docs/sdks/severitymatrix/README.md#get) - Retrieve your severity matrix
* [ListConditions](docs/sdks/severitymatrix/README.md#listconditions) - Lists conditions
* [CreateCondition](docs/sdks/severitymatrix/README.md#createcondition) - Create condition
* [UpdateCondition](docs/sdks/severitymatrix/README.md#updatecondition) - Update a specific condition
* [ListImpacts](docs/sdks/severitymatrix/README.md#listimpacts) - Lists impacts
* [CreateImpact](docs/sdks/severitymatrix/README.md#createimpact) - Create impact
* [DeleteImpact](docs/sdks/severitymatrix/README.md#deleteimpact) - Delete a specific impact
* [UpdateImpact](docs/sdks/severitymatrix/README.md#updateimpact) - Update a specific impact

### [Signals](docs/sdks/signals/README.md)

* [GetAnalyticsTimeseries](docs/sdks/signals/README.md#getanalyticstimeseries) - Generate timeseries alert metrics
* [ReportGroupedMetrics](docs/sdks/signals/README.md#reportgroupedmetrics) - Generate grouped alert metrics
* [GetMttxAnalytics](docs/sdks/signals/README.md#getmttxanalytics) - Get MTTX metrics for Signals Alerts
* [ListEventSources](docs/sdks/signals/README.md#listeventsources)
* [CreateEmailTarget](docs/sdks/signals/README.md#createemailtarget) - Create an email target
* [ListEmailTargets](docs/sdks/signals/README.md#listemailtargets) - List email targets
* [DeleteEmailTarget](docs/sdks/signals/README.md#deleteemailtarget) - Delete an email target
* [UpdateEmailTarget](docs/sdks/signals/README.md#updateemailtarget) - Update an email target
* [CreateWebhookTarget](docs/sdks/signals/README.md#createwebhooktarget) - Create an webhook target
* [GetWebhookTargets](docs/sdks/signals/README.md#getwebhooktargets) - List webhook targets
* [DeleteWebhookTarget](docs/sdks/signals/README.md#deletewebhooktarget) - Delete an webhook target
* [UpdateWebhookTarget](docs/sdks/signals/README.md#updatewebhooktarget) - Update an webhook target
* [GetWebhookTarget](docs/sdks/signals/README.md#getwebhooktarget) - Get an webhook target
* [GetTransposers](docs/sdks/signals/README.md#gettransposers)
* [GetIngestURL](docs/sdks/signals/README.md#getingesturl) - Retrieve the url for ingesting signals
* [PostDebugger](docs/sdks/signals/README.md#postdebugger)

#### [Signals.EmailTargets](docs/sdks/emailtargets/README.md)

* [Get](docs/sdks/emailtargets/README.md#get) - Get an email target

### [SignalsOnCall](docs/sdks/signalsoncall/README.md)

* [List](docs/sdks/signalsoncall/README.md#list) - List all on-call schedules

### [StatusUpdateTemplates](docs/sdks/statusupdatetemplates/README.md)

* [Create](docs/sdks/statusupdatetemplates/README.md#create) - Create a status update template
* [Delete](docs/sdks/statusupdatetemplates/README.md#delete) - Delete a status update template
* [List](docs/sdks/statusupdatetemplates/README.md#list) - List all status update templates
* [Update](docs/sdks/statusupdatetemplates/README.md#update) - Update a status update template
* [Get](docs/sdks/statusupdatetemplates/README.md#get) - Get a status update template

### [TaskLists](docs/sdks/tasklists/README.md)

* [Create](docs/sdks/tasklists/README.md#create) - Create a task list
* [List](docs/sdks/tasklists/README.md#list) - List all task lists
* [Update](docs/sdks/tasklists/README.md#update) - Update a task list
* [Get](docs/sdks/tasklists/README.md#get) - Retrieve a task list
* [Delete](docs/sdks/tasklists/README.md#delete) - Delete a task list

### [Teams](docs/sdks/teams/README.md)

* [Create](docs/sdks/teams/README.md#create) - Create a team
* [List](docs/sdks/teams/README.md#list) - List all teams
* [Archive](docs/sdks/teams/README.md#archive) - Archive a team
* [Update](docs/sdks/teams/README.md#update) - Update a team
* [Get](docs/sdks/teams/README.md#get) - Retrieve a team
* [CreateEscalationPolicy](docs/sdks/teams/README.md#createescalationpolicy) - Create an escalation policy
* [ListEscalationPolicies](docs/sdks/teams/README.md#listescalationpolicies) - List escalation policies
* [DeleteEscalationPolicy](docs/sdks/teams/README.md#deleteescalationpolicy) - Delete an escalation policy
* [UpdateEscalationPolicy](docs/sdks/teams/README.md#updateescalationpolicy) - Update an escalation policy
* [GetEscalationPolicy](docs/sdks/teams/README.md#getescalationpolicy) - Get an escalation policy
* [CreateOnCallSchedule](docs/sdks/teams/README.md#createoncallschedule) - Create an on-call schedule
* [ListOnCallSchedules](docs/sdks/teams/README.md#listoncallschedules) - List on-call schedules
* [DeleteOnCallSchedule](docs/sdks/teams/README.md#deleteoncallschedule) - Delete an on-call schedule
* [UpdateOnCallSchedule](docs/sdks/teams/README.md#updateoncallschedule) - Update an on-call schedule
* [GetOnCallSchedule](docs/sdks/teams/README.md#getoncallschedule) - Get an on-call schedule
* [CreateShift](docs/sdks/teams/README.md#createshift) - Create an on-call shift
* [DeleteOnCallShift](docs/sdks/teams/README.md#deleteoncallshift) - Delete an on-call shift
* [UpdateOnCallShift](docs/sdks/teams/README.md#updateoncallshift) - Update an on-call shift
* [GetShift](docs/sdks/teams/README.md#getshift) - Get an on-call shift
* [CreateSignalRule](docs/sdks/teams/README.md#createsignalrule) - Create a Signals rule
* [GetSignalRules](docs/sdks/teams/README.md#getsignalrules) - List Signals rules
* [DeleteSignalRule](docs/sdks/teams/README.md#deletesignalrule) - Delete a Signals rule
* [UpdateSignalRule](docs/sdks/teams/README.md#updatesignalrule) - Update a Signals rule
* [GetSignalRule](docs/sdks/teams/README.md#getsignalrule) - Get a Signals rule

### [Ticketing](docs/sdks/ticketing/README.md)

* [CreateTicket](docs/sdks/ticketing/README.md#createticket) - Create a ticket
* [ListTickets](docs/sdks/ticketing/README.md#listtickets) - List all tickets
* [DeleteTicket](docs/sdks/ticketing/README.md#deleteticket)
* [UpdateTicket](docs/sdks/ticketing/README.md#updateticket) - Update a ticket
* [GetTicket](docs/sdks/ticketing/README.md#getticket) - Retrieve a single ticket
* [ListProjects](docs/sdks/ticketing/README.md#listprojects) - List all ticketing projects
* [GetProject](docs/sdks/ticketing/README.md#getproject) - Retrieve a ticketing projects
* [GetConfigurationOptions](docs/sdks/ticketing/README.md#getconfigurationoptions)
* [GetConfigurationOptionsForField](docs/sdks/ticketing/README.md#getconfigurationoptionsforfield)
* [GetAvailableFields](docs/sdks/ticketing/README.md#getavailablefields) - Get the fields that can be mapped for this project.
* [CreateFieldMap](docs/sdks/ticketing/README.md#createfieldmap) - Create field map for a ticketing project
* [DeleteProjectFieldMap](docs/sdks/ticketing/README.md#deleteprojectfieldmap) - Archive field map for a ticketing project
* [GetFieldMap](docs/sdks/ticketing/README.md#getfieldmap) - Retrieve field map for a ticketing project
* [UpdateFieldMap](docs/sdks/ticketing/README.md#updatefieldmap) - Update field map for a ticketing project
* [CreateProviderProjectConfiguration](docs/sdks/ticketing/README.md#createproviderprojectconfiguration) - Create configuration for a ticketing project
* [DeleteConfiguration](docs/sdks/ticketing/README.md#deleteconfiguration) - Archive configuration for a ticketing project
* [UpdateProviderProjectConfiguration](docs/sdks/ticketing/README.md#updateproviderprojectconfiguration) - Update configuration for a ticketing project
* [GetConfiguration](docs/sdks/ticketing/README.md#getconfiguration) - Retrieve configuration for a ticketing project
* [CreatePriority](docs/sdks/ticketing/README.md#createpriority) - Create a ticketing priority
* [ListPriorities](docs/sdks/ticketing/README.md#listpriorities) - List all ticketing priorities
* [DeletePriority](docs/sdks/ticketing/README.md#deletepriority) - Delete a ticketing priority
* [UpdatePriority](docs/sdks/ticketing/README.md#updatepriority) - Update a ticketing priority
* [GetPriority](docs/sdks/ticketing/README.md#getpriority) - Retrieve a ticketing priority
* [GetTags](docs/sdks/ticketing/README.md#gettags) - List all ticket tags

### [Users](docs/sdks/users/README.md)

* [List](docs/sdks/users/README.md#list) - List users
* [GetServices](docs/sdks/users/README.md#getservices) - Retrieves a list of services owned by the teams a user is on

### [Webhooks](docs/sdks/webhooks/README.md)

* [Create](docs/sdks/webhooks/README.md#create) - Create webhook
* [List](docs/sdks/webhooks/README.md#list) - Lists webhooks
* [GetDeliveries](docs/sdks/webhooks/README.md#getdeliveries)
* [Delete](docs/sdks/webhooks/README.md#delete) - Delete a specific webhook
* [Update](docs/sdks/webhooks/README.md#update) - Update a specific webhook
* [Get](docs/sdks/webhooks/README.md#get) - Retrieve a specific webhook

</details>
<!-- End Available Resources and Operations [operations] -->

<!-- Start Retries [retries] -->
## Retries

Some of the endpoints in this SDK support retries. If you use the SDK without any configuration, it will fall back to the default retry strategy provided by the API. However, the default retry strategy can be overridden on a per-operation basis, or across the entire SDK.

To change the default retry strategy for a single API call, simply provide a `retry.Config` object to the call by using the `WithRetries` option:
```go
package main

import (
	"context"
	"firehydrant"
	"firehydrant/retry"
	"log"
	"models/operations"
)

func main() {
	s := firehydrant.New(
		firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.Ping.Get(ctx, operations.WithRetries(
		retry.Config{
			Strategy: "backoff",
			Backoff: &retry.BackoffStrategy{
				InitialInterval: 1,
				MaxInterval:     50,
				Exponent:        1.1,
				MaxElapsedTime:  100,
			},
			RetryConnectionErrors: false,
		}))
	if err != nil {
		log.Fatal(err)
	}
	if res.PongEntity != nil {
		// handle response
	}
}

```

If you'd like to override the default retry strategy for all operations that support retries, you can use the `WithRetryConfig` option at SDK initialization:
```go
package main

import (
	"context"
	"firehydrant"
	"firehydrant/retry"
	"log"
)

func main() {
	s := firehydrant.New(
		firehydrant.WithRetryConfig(
			retry.Config{
				Strategy: "backoff",
				Backoff: &retry.BackoffStrategy{
					InitialInterval: 1,
					MaxInterval:     50,
					Exponent:        1.1,
					MaxElapsedTime:  100,
				},
				RetryConnectionErrors: false,
			}),
		firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.Ping.Get(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if res.PongEntity != nil {
		// handle response
	}
}

```
<!-- End Retries [retries] -->

<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations. All operations return a response object or an error, they will never return both.

By Default, an API error will return `sdkerrors.SDKError`. When custom error responses are specified for an operation, the SDK may also return their associated error. You can refer to respective *Errors* tables in SDK docs for more details on possible error types for each operation.

For example, the `Create` function may return the following errors:

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| sdkerrors.ErrorEntity | 400                   | application/json      |
| sdkerrors.SDKError    | 4XX, 5XX              | \*/\*                 |

### Example

```go
package main

import (
	"context"
	"errors"
	"firehydrant"
	"firehydrant/models/components"
	"firehydrant/models/sdkerrors"
	"log"
)

func main() {
	s := firehydrant.New(
		firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.Services.Create(ctx, components.PostV1Services{
		Name: "<value>",
	})
	if err != nil {

		var e *sdkerrors.ErrorEntity
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.SDKError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}
	}
}

```
<!-- End Error Handling [errors] -->

<!-- Start Server Selection [server] -->
## Server Selection

### Select Server by Index

You can override the default server globally using the `WithServerIndex` option when initializing the SDK client instance. The selected server will then be used as the default on the operations that use it. This table lists the indexes associated with the available servers:

| # | Server | Variables |
| - | ------ | --------- |
| 0 | `https://api.firehydrant.io` | None |

#### Example

```go
package main

import (
	"context"
	"firehydrant"
	"log"
)

func main() {
	s := firehydrant.New(
		firehydrant.WithServerIndex(0),
		firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.Ping.Get(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if res.PongEntity != nil {
		// handle response
	}
}

```


### Override Server URL Per-Client

The default server can also be overridden globally using the `WithServerURL` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	"firehydrant"
	"log"
)

func main() {
	s := firehydrant.New(
		firehydrant.WithServerURL("https://api.firehydrant.io"),
		firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.Ping.Get(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if res.PongEntity != nil {
		// handle response
	}
}

```
<!-- End Server Selection [server] -->

<!-- Start Custom HTTP Client [http-client] -->
## Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"
	"github.com/myorg/your-go-sdk"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = sdk.New(sdk.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client [http-client] -->

<!-- Start Authentication [security] -->
## Authentication

### Per-Client Security Schemes

This SDK supports the following security scheme globally:

| Name     | Type     | Scheme   |
| -------- | -------- | -------- |
| `APIKey` | apiKey   | API key  |

You can configure it using the `WithSecurity` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	"firehydrant"
	"log"
)

func main() {
	s := firehydrant.New(
		firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.Ping.Get(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if res.PongEntity != nil {
		// handle response
	}
}

```
<!-- End Authentication [security] -->

<!-- Start Special Types [types] -->
## Special Types

This SDK defines the following custom types to assist with marshalling and unmarshalling data.

### Date

`types.Date` is a wrapper around time.Time that allows for JSON marshaling a date string formatted as "2006-01-02".

#### Usage

```go
d1 := types.NewDate(time.Now()) // returns *types.Date

d2 := types.DateFromTime(time.Now()) // returns types.Date

d3, err := types.NewDateFromString("2019-01-01") // returns *types.Date, error

d4, err := types.DateFromString("2019-01-01") // returns types.Date, error

d5 := types.MustNewDateFromString("2019-01-01") // returns *types.Date and panics on error

d6 := types.MustDateFromString("2019-01-01") // returns types.Date and panics on error
```
<!-- End Special Types [types] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->

# Development

## Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

## Contributions

While we value open-source contributions to this SDK, this library is generated programmatically. Any manual changes added to internal files will be overwritten on the next generation. 
We look forward to hearing your feedback. Feel free to open a PR or an issue with a proof of concept and we'll do our best to include it in a future release. 

### SDK Created by [Speakeasy](https://www.speakeasy.com/?utm_source=openapi&utm_campaign=go)
