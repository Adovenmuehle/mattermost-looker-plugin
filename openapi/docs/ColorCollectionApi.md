# \ColorCollectionApi

All URIs are relative to *https://mattermost.looker.com:19999/api/3.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AllColorCollections**](ColorCollectionApi.md#AllColorCollections) | **Get** /color_collections | Get all Color Collections
[**ColorCollection**](ColorCollectionApi.md#ColorCollection) | **Get** /color_collections/{collection_id} | Get Color Collection by ID
[**ColorCollectionsCustom**](ColorCollectionApi.md#ColorCollectionsCustom) | **Get** /color_collections/custom | Get all Custom Color Collections
[**ColorCollectionsStandard**](ColorCollectionApi.md#ColorCollectionsStandard) | **Get** /color_collections/standard | Get all Standard Color Collections
[**CreateColorCollection**](ColorCollectionApi.md#CreateColorCollection) | **Post** /color_collections | Create ColorCollection
[**DefaultColorCollection**](ColorCollectionApi.md#DefaultColorCollection) | **Get** /color_collections/default | Get Default Color Collection
[**DeleteColorCollection**](ColorCollectionApi.md#DeleteColorCollection) | **Delete** /color_collections/{collection_id} | Delete ColorCollection
[**SetDefaultColorCollection**](ColorCollectionApi.md#SetDefaultColorCollection) | **Put** /color_collections/default | Set Default Color Collection
[**UpdateColorCollection**](ColorCollectionApi.md#UpdateColorCollection) | **Patch** /color_collections/{collection_id} | Update Custom Color collection



## AllColorCollections

> []ColorCollection AllColorCollections(ctx, optional)

Get all Color Collections

### Get an array of all existing Color Collections Get a **single** color collection by id with [ColorCollection](#!/ColorCollection/color_collection)  Get all **standard** color collections with [ColorCollection](#!/ColorCollection/color_collections_standard)  Get all **custom** color collections with [ColorCollection](#!/ColorCollection/color_collections_custom)  **Note**: Only an API user with the Admin role can call this endpoint. Unauthorized requests will return `Not Found` (404) errors.  

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AllColorCollectionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AllColorCollectionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **optional.String**| Requested fields. | 

### Return type

[**[]ColorCollection**](ColorCollection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ColorCollection

> ColorCollection ColorCollection(ctx, collectionId, optional)

Get Color Collection by ID

### Get a Color Collection by ID  Use this to retrieve a specific Color Collection. Get a **single** color collection by id with [ColorCollection](#!/ColorCollection/color_collection)  Get all **standard** color collections with [ColorCollection](#!/ColorCollection/color_collections_standard)  Get all **custom** color collections with [ColorCollection](#!/ColorCollection/color_collections_custom)  **Note**: Only an API user with the Admin role can call this endpoint. Unauthorized requests will return `Not Found` (404) errors.  

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string**| Id of Color Collection | 
 **optional** | ***ColorCollectionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ColorCollectionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Requested fields. | 

### Return type

[**ColorCollection**](ColorCollection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ColorCollectionsCustom

> []ColorCollection ColorCollectionsCustom(ctx, optional)

Get all Custom Color Collections

### Get an array of all existing **Custom** Color Collections Get a **single** color collection by id with [ColorCollection](#!/ColorCollection/color_collection)  Get all **standard** color collections with [ColorCollection](#!/ColorCollection/color_collections_standard)  **Note**: Only an API user with the Admin role can call this endpoint. Unauthorized requests will return `Not Found` (404) errors.  

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ColorCollectionsCustomOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ColorCollectionsCustomOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **optional.String**| Requested fields. | 

### Return type

[**[]ColorCollection**](ColorCollection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ColorCollectionsStandard

> []ColorCollection ColorCollectionsStandard(ctx, optional)

Get all Standard Color Collections

### Get an array of all existing **Standard** Color Collections Get a **single** color collection by id with [ColorCollection](#!/ColorCollection/color_collection)  Get all **custom** color collections with [ColorCollection](#!/ColorCollection/color_collections_custom)  **Note**: Only an API user with the Admin role can call this endpoint. Unauthorized requests will return `Not Found` (404) errors.  

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ColorCollectionsStandardOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ColorCollectionsStandardOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **optional.String**| Requested fields. | 

### Return type

[**[]ColorCollection**](ColorCollection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateColorCollection

> ColorCollection CreateColorCollection(ctx, body)

Create ColorCollection

### Create a custom color collection with the specified information  Creates a new custom color collection object, returning the details, including the created id.  **Update** an existing color collection with [Update Color Collection](#!/ColorCollection/update_color_collection)  **Permanently delete** an existing custom color collection with [Delete Color Collection](#!/ColorCollection/delete_color_collection)  **Note**: Only an API user with the Admin role can call this endpoint. Unauthorized requests will return `Not Found` (404) errors.  

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**ColorCollection**](ColorCollection.md)| ColorCollection | 

### Return type

[**ColorCollection**](ColorCollection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DefaultColorCollection

> ColorCollection DefaultColorCollection(ctx, )

Get Default Color Collection

### Get the default color collection  Use this to retrieve the default Color Collection.  Set the default color collection with [ColorCollection](#!/ColorCollection/set_default_color_collection) 

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**ColorCollection**](ColorCollection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteColorCollection

> DeleteColorCollection(ctx, collectionId)

Delete ColorCollection

### Delete a custom color collection by id  This operation permanently deletes the identified **Custom** color collection.  **Standard** color collections cannot be deleted  Because multiple color collections can have the same label, they must be deleted by ID, not name. **Note**: Only an API user with the Admin role can call this endpoint. Unauthorized requests will return `Not Found` (404) errors.  

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string**| Id of Color Collection | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetDefaultColorCollection

> ColorCollection SetDefaultColorCollection(ctx, collectionId)

Set Default Color Collection

### Set the global default Color Collection by ID  Returns the new specified default Color Collection object. **Note**: Only an API user with the Admin role can call this endpoint. Unauthorized requests will return `Not Found` (404) errors.  

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string**| ID of color collection to set as default | 

### Return type

[**ColorCollection**](ColorCollection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateColorCollection

> ColorCollection UpdateColorCollection(ctx, collectionId, body)

Update Custom Color collection

### Update a custom color collection by id. **Note**: Only an API user with the Admin role can call this endpoint. Unauthorized requests will return `Not Found` (404) errors.  

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string**| Id of Custom Color Collection | 
**body** | [**ColorCollection**](ColorCollection.md)| ColorCollection | 

### Return type

[**ColorCollection**](ColorCollection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

