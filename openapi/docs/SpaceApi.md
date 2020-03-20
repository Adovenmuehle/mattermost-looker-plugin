# \SpaceApi

All URIs are relative to *https://mattermost.looker.com:19999/api/3.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AllSpaces**](SpaceApi.md#AllSpaces) | **Get** /spaces | Get All Spaces
[**CreateSpace**](SpaceApi.md#CreateSpace) | **Post** /spaces | Create Space
[**DeleteSpace**](SpaceApi.md#DeleteSpace) | **Delete** /spaces/{space_id} | Delete Space
[**SearchSpaces**](SpaceApi.md#SearchSpaces) | **Get** /spaces/search | Search Spaces
[**Space**](SpaceApi.md#Space) | **Get** /spaces/{space_id} | Get Space
[**SpaceAncestors**](SpaceApi.md#SpaceAncestors) | **Get** /spaces/{space_id}/ancestors | Get Space Ancestors
[**SpaceChildren**](SpaceApi.md#SpaceChildren) | **Get** /spaces/{space_id}/children | Get Space Children
[**SpaceChildrenSearch**](SpaceApi.md#SpaceChildrenSearch) | **Get** /spaces/{space_id}/children/search | Search Space Children
[**SpaceDashboards**](SpaceApi.md#SpaceDashboards) | **Get** /spaces/{space_id}/dashboards | Get Space Dashboards
[**SpaceLooks**](SpaceApi.md#SpaceLooks) | **Get** /spaces/{space_id}/looks | Get Space Looks
[**SpaceParent**](SpaceApi.md#SpaceParent) | **Get** /spaces/{space_id}/parent | Get Space Parent
[**UpdateSpace**](SpaceApi.md#UpdateSpace) | **Patch** /spaces/{space_id} | Update Space



## AllSpaces

> []SpaceBase AllSpaces(ctx, optional)

Get All Spaces

### Get information about all spaces.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AllSpacesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AllSpacesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **optional.String**| Requested fields. | 

### Return type

[**[]SpaceBase**](SpaceBase.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSpace

> Space CreateSpace(ctx, body)

Create Space

### Create a space with specified information.  Caller must have permission to edit the parent space and to create spaces, otherwise the request returns 404 Not Found. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**CreateSpace**](CreateSpace.md)| Create a new space | 

### Return type

[**Space**](Space.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSpace

> string DeleteSpace(ctx, spaceId)

Delete Space

### Delete the space with a specific id including any children spaces. **DANGER** this will delete all looks and dashboards in the space. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string**| Id of space | 

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


## SearchSpaces

> []Space SearchSpaces(ctx, optional)

Search Spaces

### Search Spaces    Returns an **array of space objects** that match the given search criteria.    If multiple search params are given and `filter_or` is FALSE or not specified, search params are combined in a logical AND operation. Only rows that match *all* search param criteria will be returned.  If `filter_or` is TRUE, multiple search params are combined in a logical OR operation. Results will include rows that match **any** of the search criteria.  String search params use case-insensitive matching. String search params can contain `%` and '_' as SQL LIKE pattern match wildcard expressions. example=\"dan%\" will match \"danger\" and \"Danzig\" but not \"David\" example=\"D_m%\" will match \"Damage\" and \"dump\"  Integer search params can accept a single value or a comma separated list of values. The multiple values will be combined under a logical OR operation - results will match at least one of the given values.  Most search params can accept \"IS NULL\" and \"NOT NULL\" as special expressions to match or exclude (respectively) rows where the column is null.  Boolean search params accept only \"true\" and \"false\" as values.     The parameters `limit`, and `offset` are recommended for fetching results in page-size chunks.    Get a **single space** by id with [Space](#!/Space/space) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SearchSpacesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SearchSpacesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **optional.String**| Requested fields. | 
 **page** | **optional.Int64**| Requested page. | 
 **perPage** | **optional.Int64**| Results per page. | 
 **limit** | **optional.Int64**| Number of results to return. (used with offset and takes priority over page and per_page) | 
 **offset** | **optional.Int64**| Number of results to skip before returning any. (used with limit and takes priority over page and per_page) | 
 **sorts** | **optional.String**| Fields to sort by. | 
 **name** | **optional.String**| Match Space title. | 
 **id** | **optional.Int64**| Match Space id | 
 **parentId** | **optional.String**| Filter on a children of a particular space. | 
 **creatorId** | **optional.String**| Filter on spaces created by a particular user. | 
 **filterOr** | **optional.Bool**| Combine given search criteria in a boolean OR expression | 

### Return type

[**[]Space**](Space.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Space

> Space Space(ctx, spaceId, optional)

Get Space

### Get information about the space with a specific id.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string**| Id of space | 
 **optional** | ***SpaceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SpaceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Requested fields. | 

### Return type

[**Space**](Space.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SpaceAncestors

> []Space SpaceAncestors(ctx, spaceId, optional)

Get Space Ancestors

### Get the ancestors of a space

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string**| Id of space | 
 **optional** | ***SpaceAncestorsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SpaceAncestorsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Requested fields. | 

### Return type

[**[]Space**](Space.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SpaceChildren

> []Space SpaceChildren(ctx, spaceId, optional)

Get Space Children

### Get the children of a space.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string**| Id of space | 
 **optional** | ***SpaceChildrenOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SpaceChildrenOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Requested fields. | 
 **page** | **optional.Int64**| Requested page. | 
 **perPage** | **optional.Int64**| Results per page. | 
 **sorts** | **optional.String**| Fields to sort by. | 

### Return type

[**[]Space**](Space.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SpaceChildrenSearch

> []Space SpaceChildrenSearch(ctx, spaceId, optional)

Search Space Children

### Search the children of a space

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string**| Id of space | 
 **optional** | ***SpaceChildrenSearchOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SpaceChildrenSearchOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Requested fields. | 
 **sorts** | **optional.String**| Fields to sort by. | 
 **name** | **optional.String**| Match Space name. | 

### Return type

[**[]Space**](Space.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SpaceDashboards

> []Dashboard SpaceDashboards(ctx, spaceId, optional)

Get Space Dashboards

### Get the dashboards in a space

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string**| Id of space | 
 **optional** | ***SpaceDashboardsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SpaceDashboardsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Requested fields. | 

### Return type

[**[]Dashboard**](Dashboard.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SpaceLooks

> []LookWithQuery SpaceLooks(ctx, spaceId, optional)

Get Space Looks

### Get the looks in a space

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string**| Id of space | 
 **optional** | ***SpaceLooksOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SpaceLooksOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Requested fields. | 

### Return type

[**[]LookWithQuery**](LookWithQuery.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SpaceParent

> Space SpaceParent(ctx, spaceId, optional)

Get Space Parent

### Get the parent of a space

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string**| Id of space | 
 **optional** | ***SpaceParentOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SpaceParentOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Requested fields. | 

### Return type

[**Space**](Space.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSpace

> Space UpdateSpace(ctx, spaceId, body)

Update Space

### Update the space with a specific id.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string**| Id of space | 
**body** | [**UpdateSpace**](UpdateSpace.md)| Space parameters | 

### Return type

[**Space**](Space.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

