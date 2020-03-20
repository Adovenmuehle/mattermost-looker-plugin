# \LookApi

All URIs are relative to *https://mattermost.looker.com:19999/api/3.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AllLooks**](LookApi.md#AllLooks) | **Get** /looks | Get All Looks
[**CreateLook**](LookApi.md#CreateLook) | **Post** /looks | Create Look
[**DeleteLook**](LookApi.md#DeleteLook) | **Delete** /looks/{look_id} | Delete Look
[**Look**](LookApi.md#Look) | **Get** /looks/{look_id} | Get Look
[**RunLook**](LookApi.md#RunLook) | **Get** /looks/{look_id}/run/{result_format} | Run Look
[**SearchLooks**](LookApi.md#SearchLooks) | **Get** /looks/search | Search Looks
[**UpdateLook**](LookApi.md#UpdateLook) | **Patch** /looks/{look_id} | Update Look



## AllLooks

> []Look AllLooks(ctx, optional)

Get All Looks

### Get information about all active Looks  Returns an array of **abbreviated Look objects** describing all the looks that the caller has access to. Soft-deleted Looks are **not** included.  Get the **full details** of a specific look by id with [look(id)](#!/Look/look)  Find **soft-deleted looks** with [search_looks()](#!/Look/search_looks) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AllLooksOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AllLooksOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **optional.String**| Requested fields. | 

### Return type

[**[]Look**](Look.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateLook

> LookWithQuery CreateLook(ctx, body, optional)

Create Look

### Create a Look  To create a look to display query data, first create the query with [create_query()](#!/Query/create_query) then assign the query's id to the `query_id` property in the call to `create_look()`.  To place the look into a particular space, assign the space's id to the `space_id` property in the call to `create_look()`. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**LookWithQuery**](LookWithQuery.md)| Look | 
 **optional** | ***CreateLookOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateLookOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Requested fields. | 

### Return type

[**LookWithQuery**](LookWithQuery.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLook

> string DeleteLook(ctx, lookId)

Delete Look

### Permanently Delete a Look  This operation **permanently** removes a look from the Looker database.  NOTE: There is no \"undo\" for this kind of delete.  For information about soft-delete (which can be undone) see [update_look()](#!/Look/update_look). 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lookId** | **int64**| Id of look | 

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


## Look

> LookWithQuery Look(ctx, lookId, optional)

Get Look

### Get a Look.  Returns detailed information about a Look and its associated Query.  

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lookId** | **int64**| Id of look | 
 **optional** | ***LookOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a LookOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Requested fields. | 

### Return type

[**LookWithQuery**](LookWithQuery.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RunLook

> string RunLook(ctx, lookId, resultFormat, optional)

Run Look

### Run a Look  Runs a given look's query and returns the results in the requested format.  Supported formats:  | result_format | Description | :-----------: | :--- | | json | Plain json | json_detail | Row data plus metadata describing the fields, pivots, table calcs, and other aspects of the query | csv | Comma separated values with a header | txt | Tab separated values with a header | html | Simple html | md | Simple markdown | xlsx | MS Excel spreadsheet | sql | Returns the generated SQL rather than running the query | png | A PNG image of the visualization of the query | jpg | A JPG image of the visualization of the query   

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lookId** | **int64**| Id of look | 
**resultFormat** | **string**| Format of result | 
 **optional** | ***RunLookOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a RunLookOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **optional.Int64**| Row limit (may override the limit in the saved query). | 
 **applyFormatting** | **optional.Bool**| Apply model-specified formatting to each result. | 
 **applyVis** | **optional.Bool**| Apply visualization options to results. | 
 **cache** | **optional.Bool**| Get results from cache if available. | 
 **imageWidth** | **optional.Int64**| Render width for image formats. | 
 **imageHeight** | **optional.Int64**| Render height for image formats. | 
 **generateDrillLinks** | **optional.Bool**| Generate drill links (only applicable to &#39;json_detail&#39; format. | 
 **forceProduction** | **optional.Bool**| Force use of production models even if the user is in development mode. | 
 **cacheOnly** | **optional.Bool**| Retrieve any results from cache even if the results have expired. | 
 **pathPrefix** | **optional.String**| Prefix to use for drill links (url encoded). | 
 **rebuildPdts** | **optional.Bool**| Rebuild PDTS used in query. | 
 **serverTableCalcs** | **optional.Bool**| Perform table calculations on query results | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text, application/json, image/png, image/jpeg

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchLooks

> []Look SearchLooks(ctx, optional)

Search Looks

### Search Looks  Returns an **array of Look objects** that match the specified search criteria.  If multiple search params are given and `filter_or` is FALSE or not specified, search params are combined in a logical AND operation. Only rows that match *all* search param criteria will be returned.  If `filter_or` is TRUE, multiple search params are combined in a logical OR operation. Results will include rows that match **any** of the search criteria.  String search params use case-insensitive matching. String search params can contain `%` and '_' as SQL LIKE pattern match wildcard expressions. example=\"dan%\" will match \"danger\" and \"Danzig\" but not \"David\" example=\"D_m%\" will match \"Damage\" and \"dump\"  Integer search params can accept a single value or a comma separated list of values. The multiple values will be combined under a logical OR operation - results will match at least one of the given values.  Most search params can accept \"IS NULL\" and \"NOT NULL\" as special expressions to match or exclude (respectively) rows where the column is null.  Boolean search params accept only \"true\" and \"false\" as values.   Get a **single look** by id with [look(id)](#!/Look/look) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SearchLooksOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SearchLooksOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **title** | **optional.String**| Match Look title. | 
 **description** | **optional.String**| Match Look description. | 
 **contentFavoriteId** | **optional.Int64**| Select looks with a particular content favorite id | 
 **spaceId** | **optional.String**| Select looks in a particular space. | 
 **userId** | **optional.String**| Select looks created by a particular user. | 
 **viewCount** | **optional.String**| Select looks with particular view_count value | 
 **deleted** | **optional.Bool**| Select soft-deleted looks | 
 **queryId** | **optional.Int64**| Select looks that reference a particular query by query_id | 
 **curate** | **optional.Bool**| Exclude items that exist only in personal spaces other than the users | 
 **fields** | **optional.String**| Requested fields. | 
 **page** | **optional.Int64**| Requested page. | 
 **perPage** | **optional.Int64**| Results per page. | 
 **limit** | **optional.Int64**| Number of results to return. (used with offset and takes priority over page and per_page) | 
 **offset** | **optional.Int64**| Number of results to skip before returning any. (used with limit and takes priority over page and per_page) | 
 **sorts** | **optional.String**| One or more fields to sort results by. Sortable fields: [:title, :user_id, :id, :created_at, :space_id, :description, :updated_at, :last_updater_id, :view_count, :favorite_count, :content_favorite_id, :deleted, :deleted_at, :last_viewed_at, :last_accessed_at, :query_id] | 
 **filterOr** | **optional.Bool**| Combine given search criteria in a boolean OR expression | 

### Return type

[**[]Look**](Look.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLook

> LookWithQuery UpdateLook(ctx, lookId, body, optional)

Update Look

### Modify a Look  Use this function to modify parts of a look. Property values given in a call to `update_look` are applied to the existing look, so there's no need to include properties whose values are not changing. It's best to specify only the properties you want to change and leave everything else out of your `update_look` call. **Look properties marked 'read-only' will be ignored.**  When a user deletes a look in the Looker UI, the look data remains in the database but is marked with a deleted flag (\"soft-deleted\"). Soft-deleted looks can be undeleted (by an admin) if the delete was in error.  To soft-delete a look via the API, use [update_look()](#!/Look/update_look) to change the look's `deleted` property to `true`. You can undelete a look by calling `update_look` to change the look's `deleted` property to `false`.  Soft-deleted looks are excluded from the results of [all_looks()](#!/Look/all_looks) and [search_looks()](#!/Look/search_looks), so they essentially disappear from view even though they still reside in the db. In API 3.1 and later, you can pass `deleted: true` as a parameter to [search_looks()](#!/3.1/Look/search_looks) to list soft-deleted looks.  NOTE: [delete_look()](#!/Look/delete_look) performs a \"hard delete\" - the look data is removed from the Looker database and destroyed. There is no \"undo\" for `delete_look()`. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lookId** | **int64**| Id of look | 
**body** | [**LookWithQuery**](LookWithQuery.md)| Look | 
 **optional** | ***UpdateLookOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateLookOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **optional.String**| Requested fields. | 

### Return type

[**LookWithQuery**](LookWithQuery.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

