# \ScheduledPlanApi

All URIs are relative to *https://mattermost.looker.com:19999/api/3.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AllScheduledPlans**](ScheduledPlanApi.md#AllScheduledPlans) | **Get** /scheduled_plans | Get All Scheduled Plans
[**CreateScheduledPlan**](ScheduledPlanApi.md#CreateScheduledPlan) | **Post** /scheduled_plans | Create Scheduled Plan
[**DeleteScheduledPlan**](ScheduledPlanApi.md#DeleteScheduledPlan) | **Delete** /scheduled_plans/{scheduled_plan_id} | Delete Scheduled Plan
[**ScheduledPlan**](ScheduledPlanApi.md#ScheduledPlan) | **Get** /scheduled_plans/{scheduled_plan_id} | Get Scheduled Plan
[**ScheduledPlanRunOnce**](ScheduledPlanApi.md#ScheduledPlanRunOnce) | **Post** /scheduled_plans/run_once | Run Scheduled Plan Once
[**ScheduledPlanRunOnceById**](ScheduledPlanApi.md#ScheduledPlanRunOnceById) | **Post** /scheduled_plans/{scheduled_plan_id}/run_once | Run Scheduled Plan Once by Id
[**ScheduledPlansForDashboard**](ScheduledPlanApi.md#ScheduledPlansForDashboard) | **Get** /scheduled_plans/dashboard/{dashboard_id} | Scheduled Plans for Dashboard
[**ScheduledPlansForLook**](ScheduledPlanApi.md#ScheduledPlansForLook) | **Get** /scheduled_plans/look/{look_id} | Scheduled Plans for Look
[**ScheduledPlansForLookmlDashboard**](ScheduledPlanApi.md#ScheduledPlansForLookmlDashboard) | **Get** /scheduled_plans/lookml_dashboard/{lookml_dashboard_id} | Scheduled Plans for LookML Dashboard
[**ScheduledPlansForSpace**](ScheduledPlanApi.md#ScheduledPlansForSpace) | **Get** /scheduled_plans/space/{space_id} | Scheduled Plans for Space
[**UpdateScheduledPlan**](ScheduledPlanApi.md#UpdateScheduledPlan) | **Patch** /scheduled_plans/{scheduled_plan_id} | Update Scheduled Plan



## AllScheduledPlans

> []ScheduledPlan AllScheduledPlans(ctx, optional)

Get All Scheduled Plans

### List All Scheduled Plans  Returns all scheduled plans which belong to the caller or given user.  If no user_id is provided, this function returns the scheduled plans owned by the caller.   To list all schedules for all users, pass `all_users=true`.   The caller must have `see_schedules` permission to see other users' scheduled plans.   

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AllScheduledPlansOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AllScheduledPlansOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **optional.Int64**| Return scheduled plans belonging to this user_id. If not provided, returns scheduled plans owned by the caller. | 
 **fields** | **optional.String**| Comma delimited list of field names. If provided, only the fields specified will be included in the response | 
 **allUsers** | **optional.Bool**| Return scheduled plans belonging to all users (caller needs see_schedules permission) | 

### Return type

[**[]ScheduledPlan**](ScheduledPlan.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateScheduledPlan

> ScheduledPlan CreateScheduledPlan(ctx, body)

Create Scheduled Plan

### Create a Scheduled Plan  Create a scheduled plan to render a Look or Dashboard on a recurring schedule.  To create a scheduled plan, you MUST provide values for the following fields: `name` and `look_id`, `dashboard_id`, `lookml_dashboard_id`, or `query_id` and `cron_tab` or `datagroup` and at least one scheduled_plan_destination  A scheduled plan MUST have at least one scheduled_plan_destination defined.  When `look_id` is set, `require_no_results`, `require_results`, and `require_change` are all required.  If `create_scheduled_plan` fails with a 422 error, be sure to look at the error messages in the response which will explain exactly what fields are missing or values that are incompatible.  The queries that provide the data for the look or dashboard are run in the context of user account that owns the scheduled plan.  When `run_as_recipient` is `false` or not specified, the queries that provide the data for the look or dashboard are run in the context of user account that owns the scheduled plan.  When `run_as_recipient` is `true` and all the email recipients are Looker user accounts, the queries are run in the context of each recipient, so different recipients may see different data from the same scheduled render of a look or dashboard. For more details, see [Run As Recipient](https://looker.com/docs/r/admin/run-as-recipient).  Admins can create and modify scheduled plans on behalf of other users by specifying a user id. Non-admin users may not create or modify scheduled plans by or for other users.  #### Email Permissions:  For details about permissions required to schedule delivery to email and the safeguards Looker offers to protect against sending to unauthorized email destinations, see [Email Domain Whitelist for Scheduled Looks](https://docs.looker.com/r/api/embed-permissions).   #### Scheduled Plan Destination Formats  Scheduled plan destinations must specify the data format to produce and send to the destination.  Formats:  | format | Description | :-----------: | :--- | | json | A JSON object containing a `data` property which contains an array of JSON objects, one per row. No metadata. | json_detail | Row data plus metadata describing the fields, pivots, table calcs, and other aspects of the query | inline_json | Same as the JSON format, except that the `data` property is a string containing JSON-escaped row data. Additional properties describe the data operation. This format is primarily used to send data to web hooks so that the web hook doesn't have to re-encode the JSON row data in order to pass it on to its ultimate destination. | csv | Comma separated values with a header | txt | Tab separated values with a header | html | Simple html | xlsx | MS Excel spreadsheet | wysiwyg_pdf | Dashboard rendered in a tiled layout to produce a PDF document | assembled_pdf | Dashboard rendered in a single column layout to produce a PDF document | wysiwyg_png | Dashboard rendered in a tiled layout to produce a PNG image ||  Valid formats vary by destination type and source object. `wysiwyg_pdf` is only valid for dashboards, for example.   

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**ScheduledPlan**](ScheduledPlan.md)| Scheduled Plan | 

### Return type

[**ScheduledPlan**](ScheduledPlan.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteScheduledPlan

> string DeleteScheduledPlan(ctx, scheduledPlanId)

Delete Scheduled Plan

### Delete a Scheduled Plan  Normal users can only delete their own scheduled plans. Admins can delete other users' scheduled plans. This delete cannot be undone. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scheduledPlanId** | **int64**| Scheduled Plan Id | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScheduledPlan

> ScheduledPlan ScheduledPlan(ctx, scheduledPlanId, optional)

Get Scheduled Plan

### Get Information About a Scheduled Plan  Admins can fetch information about other users' Scheduled Plans. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scheduledPlanId** | **int64**| Scheduled Plan Id | 
 **optional** | ***ScheduledPlanOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ScheduledPlanOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Requested fields. | 

### Return type

[**ScheduledPlan**](ScheduledPlan.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScheduledPlanRunOnce

> ScheduledPlan ScheduledPlanRunOnce(ctx, body)

Run Scheduled Plan Once

### Run a Scheduled Plan Immediately  Create a scheduled plan that runs only once, and immediately.  This can be useful for testing a Scheduled Plan before committing to a production schedule.  Admins can create scheduled plans on behalf of other users by specifying a user id.  This API is rate limited to prevent it from being used for relay spam or DoS attacks  #### Email Permissions:  For details about permissions required to schedule delivery to email and the safeguards Looker offers to protect against sending to unauthorized email destinations, see [Email Domain Whitelist for Scheduled Looks](https://docs.looker.com/r/api/embed-permissions).   #### Scheduled Plan Destination Formats  Scheduled plan destinations must specify the data format to produce and send to the destination.  Formats:  | format | Description | :-----------: | :--- | | json | A JSON object containing a `data` property which contains an array of JSON objects, one per row. No metadata. | json_detail | Row data plus metadata describing the fields, pivots, table calcs, and other aspects of the query | inline_json | Same as the JSON format, except that the `data` property is a string containing JSON-escaped row data. Additional properties describe the data operation. This format is primarily used to send data to web hooks so that the web hook doesn't have to re-encode the JSON row data in order to pass it on to its ultimate destination. | csv | Comma separated values with a header | txt | Tab separated values with a header | html | Simple html | xlsx | MS Excel spreadsheet | wysiwyg_pdf | Dashboard rendered in a tiled layout to produce a PDF document | assembled_pdf | Dashboard rendered in a single column layout to produce a PDF document | wysiwyg_png | Dashboard rendered in a tiled layout to produce a PNG image ||  Valid formats vary by destination type and source object. `wysiwyg_pdf` is only valid for dashboards, for example.   

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**ScheduledPlan**](ScheduledPlan.md)| Scheduled Plan | 

### Return type

[**ScheduledPlan**](ScheduledPlan.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScheduledPlanRunOnceById

> ScheduledPlan ScheduledPlanRunOnceById(ctx, scheduledPlanId, optional)

Run Scheduled Plan Once by Id

### Run a Scheduled Plan By Id Immediately This function creates a run-once schedule plan based on an existing scheduled plan, applies modifications (if any) to the new scheduled plan, and runs the new schedule plan immediately. This can be useful for testing modifications to an existing scheduled plan before committing to a production schedule.  This function internally performs the following operations: 1. Copies the properties of the existing scheduled plan into a new scheduled plan 2. Copies any properties passed in the JSON body of this request into the new scheduled plan (replacing the original values) 3. Creates the new scheduled plan 4. Runs the new scheduled plan  The original scheduled plan is not modified by this operation. Admins can create, modify, and run scheduled plans on behalf of other users by specifying a user id. Non-admins can only create, modify, and run their own scheduled plans.  #### Email Permissions:  For details about permissions required to schedule delivery to email and the safeguards Looker offers to protect against sending to unauthorized email destinations, see [Email Domain Whitelist for Scheduled Looks](https://docs.looker.com/r/api/embed-permissions).   #### Scheduled Plan Destination Formats  Scheduled plan destinations must specify the data format to produce and send to the destination.  Formats:  | format | Description | :-----------: | :--- | | json | A JSON object containing a `data` property which contains an array of JSON objects, one per row. No metadata. | json_detail | Row data plus metadata describing the fields, pivots, table calcs, and other aspects of the query | inline_json | Same as the JSON format, except that the `data` property is a string containing JSON-escaped row data. Additional properties describe the data operation. This format is primarily used to send data to web hooks so that the web hook doesn't have to re-encode the JSON row data in order to pass it on to its ultimate destination. | csv | Comma separated values with a header | txt | Tab separated values with a header | html | Simple html | xlsx | MS Excel spreadsheet | wysiwyg_pdf | Dashboard rendered in a tiled layout to produce a PDF document | assembled_pdf | Dashboard rendered in a single column layout to produce a PDF document | wysiwyg_png | Dashboard rendered in a tiled layout to produce a PNG image ||  Valid formats vary by destination type and source object. `wysiwyg_pdf` is only valid for dashboards, for example.    This API is rate limited to prevent it from being used for relay spam or DoS attacks  

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scheduledPlanId** | **int64**| Id of schedule plan to copy and run | 
 **optional** | ***ScheduledPlanRunOnceByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ScheduledPlanRunOnceByIdOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of WriteScheduledPlan**](WriteScheduledPlan.md)| Property values to apply to the newly copied scheduled plan before running it | 

### Return type

[**ScheduledPlan**](ScheduledPlan.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScheduledPlansForDashboard

> []ScheduledPlan ScheduledPlansForDashboard(ctx, dashboardId, optional)

Scheduled Plans for Dashboard

### Get Scheduled Plans for a Dashboard  Returns all scheduled plans for a dashboard which belong to the caller or given user.  If no user_id is provided, this function returns the scheduled plans owned by the caller.   To list all schedules for all users, pass `all_users=true`.   The caller must have `see_schedules` permission to see other users' scheduled plans.   

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dashboardId** | **int64**| Dashboard Id | 
 **optional** | ***ScheduledPlansForDashboardOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ScheduledPlansForDashboardOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userId** | **optional.Int64**| User Id (default is requesting user if not specified) | 
 **allUsers** | **optional.Bool**| Return scheduled plans belonging to all users for the dashboard | 
 **fields** | **optional.String**| Requested fields. | 

### Return type

[**[]ScheduledPlan**](ScheduledPlan.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScheduledPlansForLook

> []ScheduledPlan ScheduledPlansForLook(ctx, lookId, optional)

Scheduled Plans for Look

### Get Scheduled Plans for a Look  Returns all scheduled plans for a look which belong to the caller or given user.  If no user_id is provided, this function returns the scheduled plans owned by the caller.   To list all schedules for all users, pass `all_users=true`.   The caller must have `see_schedules` permission to see other users' scheduled plans.   

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lookId** | **int64**| Look Id | 
 **optional** | ***ScheduledPlansForLookOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ScheduledPlansForLookOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userId** | **optional.Int64**| User Id (default is requesting user if not specified) | 
 **fields** | **optional.String**| Requested fields. | 
 **allUsers** | **optional.Bool**| Return scheduled plans belonging to all users for the look | 

### Return type

[**[]ScheduledPlan**](ScheduledPlan.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScheduledPlansForLookmlDashboard

> []ScheduledPlan ScheduledPlansForLookmlDashboard(ctx, lookmlDashboardId, optional)

Scheduled Plans for LookML Dashboard

### Get Scheduled Plans for a LookML Dashboard  Returns all scheduled plans for a LookML Dashboard which belong to the caller or given user.  If no user_id is provided, this function returns the scheduled plans owned by the caller.   To list all schedules for all users, pass `all_users=true`.   The caller must have `see_schedules` permission to see other users' scheduled plans.   

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lookmlDashboardId** | **int64**| LookML Dashboard Id | 
 **optional** | ***ScheduledPlansForLookmlDashboardOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ScheduledPlansForLookmlDashboardOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userId** | **optional.Int64**| User Id (default is requesting user if not specified) | 
 **fields** | **optional.String**| Requested fields. | 
 **allUsers** | **optional.Bool**| Return scheduled plans belonging to all users for the dashboard | 

### Return type

[**[]ScheduledPlan**](ScheduledPlan.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScheduledPlansForSpace

> []ScheduledPlan ScheduledPlansForSpace(ctx, spaceId, optional)

Scheduled Plans for Space

### Get Scheduled Plans for a Space  Returns scheduled plans owned by the caller for a given space id. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **int64**| Space Id | 
 **optional** | ***ScheduledPlansForSpaceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ScheduledPlansForSpaceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Requested fields. | 

### Return type

[**[]ScheduledPlan**](ScheduledPlan.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateScheduledPlan

> ScheduledPlan UpdateScheduledPlan(ctx, scheduledPlanId, body)

Update Scheduled Plan

### Update a Scheduled Plan  Admins can update other users' Scheduled Plans.  Note: Any scheduled plan destinations specified in an update will **replace** all scheduled plan destinations currently defined for the scheduled plan.  For Example: If a scheduled plan has destinations A, B, and C, and you call update on this scheduled plan specifying only B in the destinations, then destinations A and C will be deleted by the update.  Updating a scheduled plan to assign null or an empty array to the scheduled_plan_destinations property is an error, as a scheduled plan must always have at least one destination.  If you omit the scheduled_plan_destinations property from the object passed to update, then the destinations defined on the original scheduled plan will remain unchanged.  #### Email Permissions:  For details about permissions required to schedule delivery to email and the safeguards Looker offers to protect against sending to unauthorized email destinations, see [Email Domain Whitelist for Scheduled Looks](https://docs.looker.com/r/api/embed-permissions).   #### Scheduled Plan Destination Formats  Scheduled plan destinations must specify the data format to produce and send to the destination.  Formats:  | format | Description | :-----------: | :--- | | json | A JSON object containing a `data` property which contains an array of JSON objects, one per row. No metadata. | json_detail | Row data plus metadata describing the fields, pivots, table calcs, and other aspects of the query | inline_json | Same as the JSON format, except that the `data` property is a string containing JSON-escaped row data. Additional properties describe the data operation. This format is primarily used to send data to web hooks so that the web hook doesn't have to re-encode the JSON row data in order to pass it on to its ultimate destination. | csv | Comma separated values with a header | txt | Tab separated values with a header | html | Simple html | xlsx | MS Excel spreadsheet | wysiwyg_pdf | Dashboard rendered in a tiled layout to produce a PDF document | assembled_pdf | Dashboard rendered in a single column layout to produce a PDF document | wysiwyg_png | Dashboard rendered in a tiled layout to produce a PNG image ||  Valid formats vary by destination type and source object. `wysiwyg_pdf` is only valid for dashboards, for example.   

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scheduledPlanId** | **int64**| Scheduled Plan Id | 
**body** | [**ScheduledPlan**](ScheduledPlan.md)| Scheduled Plan | 

### Return type

[**ScheduledPlan**](ScheduledPlan.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

