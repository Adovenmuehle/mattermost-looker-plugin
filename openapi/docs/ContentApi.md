# \ContentApi

All URIs are relative to *https://mattermost.looker.com:19999/api/3.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AllContentMetadataAccesses**](ContentApi.md#AllContentMetadataAccesses) | **Get** /content_metadata_access | Get All Content Metadata Accesses
[**AllContentMetadatas**](ContentApi.md#AllContentMetadatas) | **Get** /content_metadata | Get All Content Metadatas
[**ContentFavorite**](ContentApi.md#ContentFavorite) | **Get** /content_favorite/{content_favorite_id} | Get Favorite Content
[**ContentMetadata**](ContentApi.md#ContentMetadata) | **Get** /content_metadata/{content_metadata_id} | Get Content Metadata
[**ContentValidation**](ContentApi.md#ContentValidation) | **Get** /content_validation | Validate Content
[**CreateContentFavorite**](ContentApi.md#CreateContentFavorite) | **Post** /content_favorite | Create Favorite Content
[**CreateContentMetadataAccess**](ContentApi.md#CreateContentMetadataAccess) | **Post** /content_metadata_access | Create Content Metadata Access
[**DeleteContentFavorite**](ContentApi.md#DeleteContentFavorite) | **Delete** /content_favorite/{content_favorite_id} | Delete Favorite Content
[**DeleteContentMetadataAccess**](ContentApi.md#DeleteContentMetadataAccess) | **Delete** /content_metadata_access/{content_metadata_access_id} | Delete Content Metadata Access
[**SearchContentFavorites**](ContentApi.md#SearchContentFavorites) | **Get** /content_favorite/search | Search Favorite Contents
[**SearchContentViews**](ContentApi.md#SearchContentViews) | **Get** /content_view/search | Search Content Views
[**UpdateContentMetadata**](ContentApi.md#UpdateContentMetadata) | **Patch** /content_metadata/{content_metadata_id} | Update Content Metadata
[**UpdateContentMetadataAccess**](ContentApi.md#UpdateContentMetadataAccess) | **Put** /content_metadata_access/{content_metadata_access_id} | Update Content Metadata Access



## AllContentMetadataAccesses

> []ContentMetaGroupUser AllContentMetadataAccesses(ctx, contentMetadataId, optional)

Get All Content Metadata Accesses

### All content metadata access records for a content metadata item. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contentMetadataId** | **int64**| Id of content metadata | 
 **optional** | ***AllContentMetadataAccessesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AllContentMetadataAccessesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Requested fields. | 

### Return type

[**[]ContentMetaGroupUser**](ContentMetaGroupUser.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AllContentMetadatas

> []ContentMeta AllContentMetadatas(ctx, parentId, optional)

Get All Content Metadatas

### Get information about all content metadata in a space. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int64**| Parent space of content. | 
 **optional** | ***AllContentMetadatasOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AllContentMetadatasOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Requested fields. | 

### Return type

[**[]ContentMeta**](ContentMeta.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentFavorite

> ContentFavorite ContentFavorite(ctx, contentFavoriteId, optional)

Get Favorite Content

### Get favorite content by its id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contentFavoriteId** | **int64**| Id of favorite content | 
 **optional** | ***ContentFavoriteOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ContentFavoriteOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Requested fields. | 

### Return type

[**ContentFavorite**](ContentFavorite.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentMetadata

> ContentMeta ContentMetadata(ctx, contentMetadataId, optional)

Get Content Metadata

### Get information about an individual content metadata record. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contentMetadataId** | **int64**| Id of content metadata | 
 **optional** | ***ContentMetadataOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ContentMetadataOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Requested fields. | 

### Return type

[**ContentMeta**](ContentMeta.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContentValidation

> ContentValidation ContentValidation(ctx, optional)

Validate Content

### Validate All Content Requires Content Validation Labs Feature be enabled  Performs validation of all looks and dashboards Returns a list of errors found as well as metadata about the content validation run. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ContentValidationOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ContentValidationOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **optional.String**| Requested fields. | 

### Return type

[**ContentValidation**](ContentValidation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateContentFavorite

> ContentFavorite CreateContentFavorite(ctx, body)

Create Favorite Content

### Create favorite content

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**ContentFavorite**](ContentFavorite.md)| Favorite Content | 

### Return type

[**ContentFavorite**](ContentFavorite.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateContentMetadataAccess

> ContentMetaGroupUser CreateContentMetadataAccess(ctx, body, optional)

Create Content Metadata Access

### Create content metadata access. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**ContentMetaGroupUser**](ContentMetaGroupUser.md)| Content Metadata Access | 
 **optional** | ***CreateContentMetadataAccessOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateContentMetadataAccessOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sendBoardsNotificationEmail** | **optional.Bool**| Optionally sends notification email when granting access to a board. | 

### Return type

[**ContentMetaGroupUser**](ContentMetaGroupUser.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteContentFavorite

> string DeleteContentFavorite(ctx, contentFavoriteId)

Delete Favorite Content

### Delete favorite content

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contentFavoriteId** | **int64**| Id of favorite content | 

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


## DeleteContentMetadataAccess

> string DeleteContentMetadataAccess(ctx, contentMetadataAccessId)

Delete Content Metadata Access

### Remove content metadata access. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contentMetadataAccessId** | **int64**| Id of content metadata access | 

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


## SearchContentFavorites

> []ContentFavorite SearchContentFavorites(ctx, optional)

Search Favorite Contents

### Search Favorite Content  If multiple search params are given and `filter_or` is FALSE or not specified, search params are combined in a logical AND operation. Only rows that match *all* search param criteria will be returned.  If `filter_or` is TRUE, multiple search params are combined in a logical OR operation. Results will include rows that match **any** of the search criteria.  String search params use case-insensitive matching. String search params can contain `%` and '_' as SQL LIKE pattern match wildcard expressions. example=\"dan%\" will match \"danger\" and \"Danzig\" but not \"David\" example=\"D_m%\" will match \"Damage\" and \"dump\"  Integer search params can accept a single value or a comma separated list of values. The multiple values will be combined under a logical OR operation - results will match at least one of the given values.  Most search params can accept \"IS NULL\" and \"NOT NULL\" as special expressions to match or exclude (respectively) rows where the column is null.  Boolean search params accept only \"true\" and \"false\" as values.  

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SearchContentFavoritesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SearchContentFavoritesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.Int64**| Match content favorite id(s) | 
 **userId** | **optional.Int64**| Match user id(s) | 
 **contentMetadataId** | **optional.Int64**| Match content metadata id(s) | 
 **dashboardId** | **optional.Int64**| Match dashboard id(s) | 
 **lookId** | **optional.Int64**| Match look id(s) | 
 **limit** | **optional.Int64**| Number of results to return. (used with offset) | 
 **offset** | **optional.Int64**| Number of results to skip before returning any. (used with limit) | 
 **sorts** | **optional.String**| Fields to sort by. | 
 **fields** | **optional.String**| Requested fields. | 
 **filterOr** | **optional.Bool**| Combine given search criteria in a boolean OR expression | 

### Return type

[**[]ContentFavorite**](ContentFavorite.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchContentViews

> []ContentView SearchContentViews(ctx, optional)

Search Content Views

### Search Content Views  If multiple search params are given and `filter_or` is FALSE or not specified, search params are combined in a logical AND operation. Only rows that match *all* search param criteria will be returned.  If `filter_or` is TRUE, multiple search params are combined in a logical OR operation. Results will include rows that match **any** of the search criteria.  String search params use case-insensitive matching. String search params can contain `%` and '_' as SQL LIKE pattern match wildcard expressions. example=\"dan%\" will match \"danger\" and \"Danzig\" but not \"David\" example=\"D_m%\" will match \"Damage\" and \"dump\"  Integer search params can accept a single value or a comma separated list of values. The multiple values will be combined under a logical OR operation - results will match at least one of the given values.  Most search params can accept \"IS NULL\" and \"NOT NULL\" as special expressions to match or exclude (respectively) rows where the column is null.  Boolean search params accept only \"true\" and \"false\" as values.  

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SearchContentViewsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SearchContentViewsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **viewCount** | **optional.Int64**| Match view count | 
 **groupId** | **optional.Int64**| Match Group Id | 
 **lookId** | **optional.String**| Match look_id | 
 **dashboardId** | **optional.String**| Match dashboard_id | 
 **contentMetadataId** | **optional.Int64**| Match content metadata id | 
 **startOfWeekDate** | **optional.String**| Match start of week date | 
 **allTime** | **optional.Bool**| True if only all time view records should be returned | 
 **userId** | **optional.Int64**| Match user id | 
 **fields** | **optional.String**| Requested fields | 
 **limit** | **optional.Int64**| Number of results to return. Use with &#x60;offset&#x60; to manage pagination of results | 
 **offset** | **optional.Int64**| Number of results to skip before returning data | 
 **sorts** | **optional.String**| Fields to sort by | 
 **filterOr** | **optional.Bool**| Combine given search criteria in a boolean OR expression | 

### Return type

[**[]ContentView**](ContentView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateContentMetadata

> ContentMeta UpdateContentMetadata(ctx, contentMetadataId, body)

Update Content Metadata

### Move a piece of content. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contentMetadataId** | **int64**| Id of content metadata | 
**body** | [**ContentMeta**](ContentMeta.md)| Content Metadata | 

### Return type

[**ContentMeta**](ContentMeta.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateContentMetadataAccess

> ContentMetaGroupUser UpdateContentMetadataAccess(ctx, contentMetadataAccessId, body)

Update Content Metadata Access

### Update type of access for content metadata. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contentMetadataAccessId** | **int64**| Id of content metadata access | 
**body** | [**ContentMetaGroupUser**](ContentMetaGroupUser.md)| Content Metadata Access | 

### Return type

[**ContentMetaGroupUser**](ContentMetaGroupUser.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

