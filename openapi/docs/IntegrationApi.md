# \IntegrationApi

All URIs are relative to *https://mattermost.looker.com:19999/api/3.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AcceptIntegrationHubLegalAgreement**](IntegrationApi.md#AcceptIntegrationHubLegalAgreement) | **Post** /integration_hubs/{integration_hub_id}/accept_legal_agreement | Accept Integration Hub Legal Agreement
[**AllIntegrationHubs**](IntegrationApi.md#AllIntegrationHubs) | **Get** /integration_hubs | Get All Integration Hubs
[**AllIntegrations**](IntegrationApi.md#AllIntegrations) | **Get** /integrations | Get All Integrations
[**CreateIntegrationHub**](IntegrationApi.md#CreateIntegrationHub) | **Post** /integration_hubs | Create Integration Hub
[**DeleteIntegrationHub**](IntegrationApi.md#DeleteIntegrationHub) | **Delete** /integration_hubs/{integration_hub_id} | Delete Integration Hub
[**FetchIntegrationForm**](IntegrationApi.md#FetchIntegrationForm) | **Post** /integrations/{integration_id}/form | Fetch Remote Integration Form
[**Integration**](IntegrationApi.md#Integration) | **Get** /integrations/{integration_id} | Get Integration
[**IntegrationHub**](IntegrationApi.md#IntegrationHub) | **Get** /integration_hubs/{integration_hub_id} | Get Integration Hub
[**TestIntegration**](IntegrationApi.md#TestIntegration) | **Post** /integrations/{integration_id}/test | Test integration
[**UpdateIntegration**](IntegrationApi.md#UpdateIntegration) | **Patch** /integrations/{integration_id} | Update Integration
[**UpdateIntegrationHub**](IntegrationApi.md#UpdateIntegrationHub) | **Patch** /integration_hubs/{integration_hub_id} | Update Integration Hub



## AcceptIntegrationHubLegalAgreement

> IntegrationHub AcceptIntegrationHubLegalAgreement(ctx, integrationHubId)

Accept Integration Hub Legal Agreement

Accepts the legal agreement for a given integration hub. This only works for integration hubs that have legal_agreement_required set to true and legal_agreement_signed set to false.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationHubId** | **int64**| Id of integration_hub | 

### Return type

[**IntegrationHub**](IntegrationHub.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AllIntegrationHubs

> []IntegrationHub AllIntegrationHubs(ctx, optional)

Get All Integration Hubs

### Get information about all Integration Hubs. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AllIntegrationHubsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AllIntegrationHubsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **optional.String**| Requested fields. | 

### Return type

[**[]IntegrationHub**](IntegrationHub.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AllIntegrations

> []Integration AllIntegrations(ctx, optional)

Get All Integrations

### Get information about all Integrations. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AllIntegrationsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AllIntegrationsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **optional.String**| Requested fields. | 
 **integrationHubId** | **optional.String**| Filter to a specific provider | 

### Return type

[**[]Integration**](Integration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateIntegrationHub

> IntegrationHub CreateIntegrationHub(ctx, body, optional)

Create Integration Hub

### Create a new Integration Hub.  This API is rate limited to prevent it from being used for SSRF attacks 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**IntegrationHub**](IntegrationHub.md)| Integration Hub | 
 **optional** | ***CreateIntegrationHubOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateIntegrationHubOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Requested fields. | 

### Return type

[**IntegrationHub**](IntegrationHub.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIntegrationHub

> string DeleteIntegrationHub(ctx, integrationHubId)

Delete Integration Hub

### Delete a Integration Hub. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationHubId** | **int64**| Id of integration_hub | 

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


## FetchIntegrationForm

> DataActionForm FetchIntegrationForm(ctx, integrationId)

Fetch Remote Integration Form

Returns the Integration form for presentation to the user.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationId** | **int64**| Id of integration | 

### Return type

[**DataActionForm**](DataActionForm.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Integration

> Integration Integration(ctx, integrationId, optional)

Get Integration

### Get information about a Integration. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationId** | **int64**| Id of integration | 
 **optional** | ***IntegrationOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a IntegrationOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Requested fields. | 

### Return type

[**Integration**](Integration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationHub

> IntegrationHub IntegrationHub(ctx, integrationHubId, optional)

Get Integration Hub

### Get information about a Integration Hub. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationHubId** | **int64**| Id of Integration Hub | 
 **optional** | ***IntegrationHubOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a IntegrationHubOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Requested fields. | 

### Return type

[**IntegrationHub**](IntegrationHub.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestIntegration

> IntegrationTestResult TestIntegration(ctx, integrationId)

Test integration

Tests the integration to make sure all the settings are working.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationId** | **int64**| Id of integration | 

### Return type

[**IntegrationTestResult**](IntegrationTestResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIntegration

> Integration UpdateIntegration(ctx, integrationId, body, optional)

Update Integration

### Update parameters on a Integration. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationId** | **int64**| Id of integration | 
**body** | [**Integration**](Integration.md)| Integration | 
 **optional** | ***UpdateIntegrationOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateIntegrationOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **optional.String**| Requested fields. | 

### Return type

[**Integration**](Integration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIntegrationHub

> IntegrationHub UpdateIntegrationHub(ctx, integrationHubId, body, optional)

Update Integration Hub

### Update a Integration Hub definition.  This API is rate limited to prevent it from being used for SSRF attacks 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationHubId** | **int64**| Id of Integration Hub | 
**body** | [**IntegrationHub**](IntegrationHub.md)| Integration Hub | 
 **optional** | ***UpdateIntegrationHubOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateIntegrationHubOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **optional.String**| Requested fields. | 

### Return type

[**IntegrationHub**](IntegrationHub.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

