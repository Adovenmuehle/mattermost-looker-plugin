# \RenderTaskApi

All URIs are relative to *https://mattermost.looker.com:19999/api/3.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDashboardRenderTask**](RenderTaskApi.md#CreateDashboardRenderTask) | **Post** /render_tasks/dashboards/{dashboard_id}/{result_format} | Create Dashboard Render Task
[**CreateLookRenderTask**](RenderTaskApi.md#CreateLookRenderTask) | **Post** /render_tasks/looks/{look_id}/{result_format} | Create Look Render Task
[**CreateLookmlDashboardRenderTask**](RenderTaskApi.md#CreateLookmlDashboardRenderTask) | **Post** /render_tasks/lookml_dashboards/{dashboard_id}/{result_format} | Create Lookml Dashboard Render Task
[**CreateQueryRenderTask**](RenderTaskApi.md#CreateQueryRenderTask) | **Post** /render_tasks/queries/{query_id}/{result_format} | Create Query Render Task
[**RenderTask**](RenderTaskApi.md#RenderTask) | **Get** /render_tasks/{render_task_id} | Get Render Task
[**RenderTaskResults**](RenderTaskApi.md#RenderTaskResults) | **Get** /render_tasks/{render_task_id}/results | Render Task Results



## CreateDashboardRenderTask

> RenderTask CreateDashboardRenderTask(ctx, dashboardId, resultFormat, width, height, body, optional)

Create Dashboard Render Task

### Create a new task to render a dashboard to a document or image.  Returns a render task object. To check the status of a render task, pass the render_task.id to [Get Render Task](#!/RenderTask/get_render_task). Once the render task is complete, you can download the resulting document or image using [Get Render Task Results](#!/RenderTask/get_render_task_results).  

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dashboardId** | **int64**| Id of dashboard to render | 
**resultFormat** | **string**| Output type: pdf, png, or jpg | 
**width** | **int64**| Output width in pixels | 
**height** | **int64**| Output height in pixels | 
**body** | [**CreateDashboardRenderTask**](CreateDashboardRenderTask.md)| Dashboard render task parameters | 
 **optional** | ***CreateDashboardRenderTaskOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateDashboardRenderTaskOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **fields** | **optional.String**| Requested fields. | 
 **pdfPaperSize** | **optional.String**| Paper size for pdf | 
 **pdfLandscape** | **optional.Bool**| Whether to render pdf in landscape | 

### Return type

[**RenderTask**](RenderTask.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateLookRenderTask

> RenderTask CreateLookRenderTask(ctx, lookId, resultFormat, width, height, optional)

Create Look Render Task

### Create a new task to render a look to an image.  Returns a render task object. To check the status of a render task, pass the render_task.id to [Get Render Task](#!/RenderTask/get_render_task). Once the render task is complete, you can download the resulting document or image using [Get Render Task Results](#!/RenderTask/get_render_task_results).  

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lookId** | **int64**| Id of look to render | 
**resultFormat** | **string**| Output type: png, or jpg | 
**width** | **int64**| Output width in pixels | 
**height** | **int64**| Output height in pixels | 
 **optional** | ***CreateLookRenderTaskOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateLookRenderTaskOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **fields** | **optional.String**| Requested fields. | 

### Return type

[**RenderTask**](RenderTask.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateLookmlDashboardRenderTask

> RenderTask CreateLookmlDashboardRenderTask(ctx, dashboardId, resultFormat, width, height, body, optional)

Create Lookml Dashboard Render Task

### Create a new task to render a lookml dashboard to a document or image.  Returns a render task object. To check the status of a render task, pass the render_task.id to [Get Render Task](#!/RenderTask/get_render_task). Once the render task is complete, you can download the resulting document or image using [Get Render Task Results](#!/RenderTask/get_render_task_results).  

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dashboardId** | **string**| Id of lookml dashboard to render | 
**resultFormat** | **string**| Output type: pdf, png, or jpg | 
**width** | **int64**| Output width in pixels | 
**height** | **int64**| Output height in pixels | 
**body** | [**CreateDashboardRenderTask**](CreateDashboardRenderTask.md)| Dashboard render task parameters | 
 **optional** | ***CreateLookmlDashboardRenderTaskOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateLookmlDashboardRenderTaskOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **fields** | **optional.String**| Requested fields. | 
 **pdfPaperSize** | **optional.String**| Paper size for pdf | 
 **pdfLandscape** | **optional.Bool**| Whether to render pdf in landscape | 

### Return type

[**RenderTask**](RenderTask.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateQueryRenderTask

> RenderTask CreateQueryRenderTask(ctx, queryId, resultFormat, width, height, optional)

Create Query Render Task

### Create a new task to render an existing query to an image.  Returns a render task object. To check the status of a render task, pass the render_task.id to [Get Render Task](#!/RenderTask/get_render_task). Once the render task is complete, you can download the resulting document or image using [Get Render Task Results](#!/RenderTask/get_render_task_results).  

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**queryId** | **int64**| Id of the query to render | 
**resultFormat** | **string**| Output type: png or jpg | 
**width** | **int64**| Output width in pixels | 
**height** | **int64**| Output height in pixels | 
 **optional** | ***CreateQueryRenderTaskOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateQueryRenderTaskOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **fields** | **optional.String**| Requested fields. | 

### Return type

[**RenderTask**](RenderTask.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RenderTask

> RenderTask RenderTask(ctx, renderTaskId, optional)

Get Render Task

### Get information about a render task.  Returns a render task object. To check the status of a render task, pass the render_task.id to [Get Render Task](#!/RenderTask/get_render_task). Once the render task is complete, you can download the resulting document or image using [Get Render Task Results](#!/RenderTask/get_render_task_results).  

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**renderTaskId** | **string**| Id of render task | 
 **optional** | ***RenderTaskOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a RenderTaskOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Requested fields. | 

### Return type

[**RenderTask**](RenderTask.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RenderTaskResults

> string RenderTaskResults(ctx, renderTaskId)

Render Task Results

### Get the document or image produced by a completed render task.  Note that the PDF or image result will be a binary blob in the HTTP response, as indicated by the Content-Type in the response headers. This may require specialized (or at least different) handling than text responses such as JSON. You may need to tell your HTTP client that the response is binary so that it does not attempt to parse the binary data as text.  If the render task exists but has not finished rendering the results, the response HTTP status will be **202 Accepted**, the response body will be empty, and the response will have a Retry-After header indicating that the caller should repeat the request at a later time.  Returns 404 if the render task cannot be found, if the cached result has expired, or if the caller does not have permission to view the results.  For detailed information about the status of the render task, use [Render Task](#!/RenderTask/render_task). Polling loops waiting for completion of a render task would be better served by polling **render_task(id)** until the task status reaches completion (or error) instead of polling **render_task_results(id)** alone. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**renderTaskId** | **string**| Id of render task | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/jpeg, image/png, application/pdf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

