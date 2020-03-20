# \ConnectionApi

All URIs are relative to *https://mattermost.looker.com:19999/api/3.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AllConnections**](ConnectionApi.md#AllConnections) | **Get** /connections | Get All Connections
[**AllDialectInfos**](ConnectionApi.md#AllDialectInfos) | **Get** /dialect_info | Get All Dialect Infos
[**Connection**](ConnectionApi.md#Connection) | **Get** /connections/{connection_name} | Get Connection
[**CreateConnection**](ConnectionApi.md#CreateConnection) | **Post** /connections | Create Connection
[**DeleteConnection**](ConnectionApi.md#DeleteConnection) | **Delete** /connections/{connection_name} | Delete Connection
[**DeleteConnectionOverride**](ConnectionApi.md#DeleteConnectionOverride) | **Delete** /connections/{connection_name}/connection_override/{override_context} | Delete Connection Override
[**TestConnection**](ConnectionApi.md#TestConnection) | **Put** /connections/{connection_name}/test | Test Connection
[**TestConnectionConfig**](ConnectionApi.md#TestConnectionConfig) | **Put** /connections/test | Test Connection Configuration
[**UpdateConnection**](ConnectionApi.md#UpdateConnection) | **Patch** /connections/{connection_name} | Update Connection



## AllConnections

> []DbConnection AllConnections(ctx, optional)

Get All Connections

### Get information about all connections. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AllConnectionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AllConnectionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **optional.String**| Requested fields. | 

### Return type

[**[]DbConnection**](DBConnection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AllDialectInfos

> []DialectInfo AllDialectInfos(ctx, optional)

Get All Dialect Infos

### Get information about all dialects. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AllDialectInfosOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AllDialectInfosOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **optional.String**| Requested fields. | 

### Return type

[**[]DialectInfo**](DialectInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Connection

> DbConnection Connection(ctx, connectionName, optional)

Get Connection

### Get information about a connection. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionName** | **string**| Name of connection | 
 **optional** | ***ConnectionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ConnectionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Requested fields. | 

### Return type

[**DbConnection**](DBConnection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateConnection

> DbConnection CreateConnection(ctx, body)

Create Connection

### Create a connection using the specified configuration. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**DbConnection**](DbConnection.md)| Connection | 

### Return type

[**DbConnection**](DBConnection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteConnection

> string DeleteConnection(ctx, connectionName)

Delete Connection

### Delete a connection. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionName** | **string**| Name of connection | 

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


## DeleteConnectionOverride

> string DeleteConnectionOverride(ctx, connectionName, overrideContext)

Delete Connection Override

### Delete a connection override. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionName** | **string**| Name of connection | 
**overrideContext** | **string**| Context of connection override | 

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


## TestConnection

> []DbConnectionTestResult TestConnection(ctx, connectionName, optional)

Test Connection

### Test an existing connection.  Note that a connection's 'dialect' property has a 'connection_tests' property that lists the specific types of tests that the connection supports.  This API is rate limited.  Unsupported tests in the request will be ignored. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionName** | **string**| Name of connection | 
 **optional** | ***TestConnectionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a TestConnectionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tests** | [**optional.Interface of []string**](string.md)| Array of names of tests to run | 

### Return type

[**[]DbConnectionTestResult**](DBConnectionTestResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestConnectionConfig

> []DbConnectionTestResult TestConnectionConfig(ctx, body, optional)

Test Connection Configuration

### Test a connection configuration.  Note that a connection's 'dialect' property has a 'connection_tests' property that lists the specific types of tests that the connection supports.  This API is rate limited.  Unsupported tests in the request will be ignored. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**DbConnection**](DbConnection.md)| Connection | 
 **optional** | ***TestConnectionConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a TestConnectionConfigOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tests** | [**optional.Interface of []string**](string.md)| Array of names of tests to run | 

### Return type

[**[]DbConnectionTestResult**](DBConnectionTestResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConnection

> DbConnection UpdateConnection(ctx, connectionName, body)

Update Connection

### Update a connection using the specified configuration. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionName** | **string**| Name of connection | 
**body** | [**DbConnection**](DbConnection.md)| Connection | 

### Return type

[**DbConnection**](DBConnection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

