# ScheduledPlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of this scheduled plan | [optional] 
**UserId** | **int64** | User Id which owns this scheduled plan | [optional] 
**RunAsRecipient** | **bool** | Whether schedule is run as recipient (only applicable for email recipients) | [optional] 
**Enabled** | **bool** | Whether the ScheduledPlan is enabled | [optional] 
**LookId** | **int64** | Id of a look | [optional] 
**DashboardId** | **int64** | Id of a dashboard | [optional] 
**LookmlDashboardId** | **string** | Id of a LookML dashboard | [optional] 
**FiltersString** | **string** | Query string to run look or dashboard with | [optional] 
**DashboardFilters** | **string** | (DEPRECATED) Alias for filters_string field | [optional] 
**RequireResults** | **bool** | Delivery should occur if running the dashboard or look returns results | [optional] 
**RequireNoResults** | **bool** | Delivery should occur if the dashboard look does not return results | [optional] 
**RequireChange** | **bool** | Delivery should occur if data have changed since the last run | [optional] 
**SendAllResults** | **bool** | Will run an unlimited query and send all results. | [optional] 
**Crontab** | **string** | Vixie-Style crontab specification when to run | [optional] 
**Datagroup** | **string** | Name of a datagroup; if specified will run when datagroup triggered (can&#39;t be used with cron string) | [optional] 
**Timezone** | **string** | Timezone for interpreting the specified crontab (default is Looker instance timezone) | [optional] 
**QueryId** | **string** | Query id | [optional] 
**ScheduledPlanDestination** | [**[]ScheduledPlanDestination**](ScheduledPlanDestination.md) | Scheduled plan destinations | [optional] 
**RunOnce** | **bool** | Whether the plan in question should only be run once (usually for testing) | [optional] 
**IncludeLinks** | **bool** | Whether links back to Looker should be included in this ScheduledPlan | [optional] 
**PdfPaperSize** | **string** | The size of paper a PDF should be rendered for | [optional] 
**PdfLandscape** | **bool** | Whether the paper should be landscape | [optional] 
**Embed** | **bool** | Whether this schedule is in an embed context or not | [optional] 
**ColorTheme** | **string** | Color scheme of the dashboard if applicable | [optional] 
**LongTables** | **bool** | Whether or not to expand table vis to full length | [optional] 
**Id** | **int64** | Unique Id | [optional] [readonly] 
**CreatedAt** | [**time.Time**](time.Time.md) | Date and time when ScheduledPlan was created | [optional] [readonly] 
**UpdatedAt** | [**time.Time**](time.Time.md) | Date and time when ScheduledPlan was last updated | [optional] [readonly] 
**Title** | **string** | Title | [optional] [readonly] 
**User** | [**UserPublic**](UserPublic.md) |  | [optional] 
**NextRunAt** | [**time.Time**](time.Time.md) | When the ScheduledPlan will next run (null if running once) | [optional] [readonly] 
**LastRunAt** | [**time.Time**](time.Time.md) | When the ScheduledPlan was last run | [optional] [readonly] 
**Can** | **map[string]bool** | Operations the current user is able to perform on this object | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


