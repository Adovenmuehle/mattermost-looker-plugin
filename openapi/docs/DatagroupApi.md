# \DatagroupApi

All URIs are relative to *https://mattermost.looker.com:19999/api/3.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AllDatagroups**](DatagroupApi.md#AllDatagroups) | **Get** /datagroups | Get All Datagroups
[**Datagroup**](DatagroupApi.md#Datagroup) | **Get** /datagroups/{datagroup_id} | Get Datagroup
[**UpdateDatagroup**](DatagroupApi.md#UpdateDatagroup) | **Patch** /datagroups/{datagroup_id} | Update Datagroup



## AllDatagroups

> []Datagroup AllDatagroups(ctx, )

Get All Datagroups

### Get information about all datagroups. 

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]Datagroup**](Datagroup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Datagroup

> Datagroup Datagroup(ctx, datagroupId)

Get Datagroup

### Get information about a datagroup. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datagroupId** | **string**| ID of datagroup. | 

### Return type

[**Datagroup**](Datagroup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDatagroup

> Datagroup UpdateDatagroup(ctx, datagroupId, body)

Update Datagroup

### Update a datagroup using the specified params. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datagroupId** | **string**| ID of datagroup. | 
**body** | [**Datagroup**](Datagroup.md)| Datagroup | 

### Return type

[**Datagroup**](Datagroup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

