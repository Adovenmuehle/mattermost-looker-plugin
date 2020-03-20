# \LookmlModelApi

All URIs are relative to *https://mattermost.looker.com:19999/api/3.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AllLookmlModels**](LookmlModelApi.md#AllLookmlModels) | **Get** /lookml_models | Get All LookML Models
[**CreateLookmlModel**](LookmlModelApi.md#CreateLookmlModel) | **Post** /lookml_models | Create LookML Model
[**DeleteLookmlModel**](LookmlModelApi.md#DeleteLookmlModel) | **Delete** /lookml_models/{lookml_model_name} | Delete LookML Model
[**LookmlModel**](LookmlModelApi.md#LookmlModel) | **Get** /lookml_models/{lookml_model_name} | Get LookML Model
[**LookmlModelExplore**](LookmlModelApi.md#LookmlModelExplore) | **Get** /lookml_models/{lookml_model_name}/explores/{explore_name} | Get LookML Model Explore
[**UpdateLookmlModel**](LookmlModelApi.md#UpdateLookmlModel) | **Patch** /lookml_models/{lookml_model_name} | Update LookML Model



## AllLookmlModels

> []LookmlModel AllLookmlModels(ctx, optional)

Get All LookML Models

### Get information about all lookml models. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AllLookmlModelsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AllLookmlModelsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **optional.String**| Requested fields. | 

### Return type

[**[]LookmlModel**](LookmlModel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateLookmlModel

> LookmlModel CreateLookmlModel(ctx, body)

Create LookML Model

### Create a lookml model using the specified configuration. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**LookmlModel**](LookmlModel.md)| LookML Model | 

### Return type

[**LookmlModel**](LookmlModel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLookmlModel

> string DeleteLookmlModel(ctx, lookmlModelName)

Delete LookML Model

### Delete a lookml model. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lookmlModelName** | **string**| Name of lookml model. | 

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


## LookmlModel

> LookmlModel LookmlModel(ctx, lookmlModelName, optional)

Get LookML Model

### Get information about a lookml model. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lookmlModelName** | **string**| Name of lookml model. | 
 **optional** | ***LookmlModelOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a LookmlModelOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Requested fields. | 

### Return type

[**LookmlModel**](LookmlModel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LookmlModelExplore

> LookmlModelExplore LookmlModelExplore(ctx, lookmlModelName, exploreName, optional)

Get LookML Model Explore

### Get information about a lookml model explore. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lookmlModelName** | **string**| Name of lookml model. | 
**exploreName** | **string**| Name of explore. | 
 **optional** | ***LookmlModelExploreOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a LookmlModelExploreOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **optional.String**| Requested fields. | 

### Return type

[**LookmlModelExplore**](LookmlModelExplore.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLookmlModel

> LookmlModel UpdateLookmlModel(ctx, lookmlModelName, body)

Update LookML Model

### Update a lookml model using the specified configuration. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lookmlModelName** | **string**| Name of lookml model. | 
**body** | [**LookmlModel**](LookmlModel.md)| LookML Model | 

### Return type

[**LookmlModel**](LookmlModel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

