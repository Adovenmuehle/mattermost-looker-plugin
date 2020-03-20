# \HomepageApi

All URIs are relative to *https://mattermost.looker.com:19999/api/3.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AllHomepageItems**](HomepageApi.md#AllHomepageItems) | **Get** /homepage_items | Get All Homepage Items
[**AllHomepageSections**](HomepageApi.md#AllHomepageSections) | **Get** /homepage_sections | Get All Homepage sections
[**AllHomepages**](HomepageApi.md#AllHomepages) | **Get** /homepages | Get All Homepages
[**CreateHomepage**](HomepageApi.md#CreateHomepage) | **Post** /homepages | Create Homepage
[**CreateHomepageItem**](HomepageApi.md#CreateHomepageItem) | **Post** /homepage_items | Create Homepage Item
[**CreateHomepageSection**](HomepageApi.md#CreateHomepageSection) | **Post** /homepage_sections | Create Homepage section
[**DeleteHomepage**](HomepageApi.md#DeleteHomepage) | **Delete** /homepages/{homepage_id} | Delete Homepage
[**DeleteHomepageItem**](HomepageApi.md#DeleteHomepageItem) | **Delete** /homepage_items/{homepage_item_id} | Delete Homepage Item
[**DeleteHomepageSection**](HomepageApi.md#DeleteHomepageSection) | **Delete** /homepage_sections/{homepage_section_id} | Delete Homepage section
[**Homepage**](HomepageApi.md#Homepage) | **Get** /homepages/{homepage_id} | Get Homepage
[**HomepageItem**](HomepageApi.md#HomepageItem) | **Get** /homepage_items/{homepage_item_id} | Get Homepage Item
[**HomepageSection**](HomepageApi.md#HomepageSection) | **Get** /homepage_sections/{homepage_section_id} | Get Homepage section
[**SearchHomepages**](HomepageApi.md#SearchHomepages) | **Get** /homepages/search | Search Homepages
[**UpdateHomepage**](HomepageApi.md#UpdateHomepage) | **Patch** /homepages/{homepage_id} | Update Homepage
[**UpdateHomepageItem**](HomepageApi.md#UpdateHomepageItem) | **Patch** /homepage_items/{homepage_item_id} | Update Homepage Item
[**UpdateHomepageSection**](HomepageApi.md#UpdateHomepageSection) | **Patch** /homepage_sections/{homepage_section_id} | Update Homepage section



## AllHomepageItems

> []HomepageItem AllHomepageItems(ctx, optional)

Get All Homepage Items

### Get information about all homepage items. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AllHomepageItemsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AllHomepageItemsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **optional.String**| Requested fields. | 
 **sorts** | **optional.String**| Fields to sort by. | 
 **homepageSectionId** | **optional.String**| Filter to a specific homepage section | 

### Return type

[**[]HomepageItem**](HomepageItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AllHomepageSections

> []HomepageSection AllHomepageSections(ctx, optional)

Get All Homepage sections

### Get information about all homepage sections. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AllHomepageSectionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AllHomepageSectionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **optional.String**| Requested fields. | 
 **sorts** | **optional.String**| Fields to sort by. | 

### Return type

[**[]HomepageSection**](HomepageSection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AllHomepages

> []Homepage AllHomepages(ctx, optional)

Get All Homepages

### Get information about all homepages. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AllHomepagesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AllHomepagesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **optional.String**| Requested fields. | 

### Return type

[**[]Homepage**](Homepage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateHomepage

> Homepage CreateHomepage(ctx, body, optional)

Create Homepage

### Create a new homepage. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**Homepage**](Homepage.md)| Homepage | 
 **optional** | ***CreateHomepageOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateHomepageOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Requested fields. | 

### Return type

[**Homepage**](Homepage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateHomepageItem

> HomepageItem CreateHomepageItem(ctx, body, optional)

Create Homepage Item

### Create a new homepage item. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**HomepageItem**](HomepageItem.md)| Homepage Item | 
 **optional** | ***CreateHomepageItemOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateHomepageItemOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Requested fields. | 

### Return type

[**HomepageItem**](HomepageItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateHomepageSection

> HomepageSection CreateHomepageSection(ctx, body, optional)

Create Homepage section

### Create a new homepage section. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**HomepageSection**](HomepageSection.md)| Homepage section | 
 **optional** | ***CreateHomepageSectionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateHomepageSectionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Requested fields. | 

### Return type

[**HomepageSection**](HomepageSection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteHomepage

> string DeleteHomepage(ctx, homepageId)

Delete Homepage

### Delete a homepage. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**homepageId** | **int64**| Id of homepage | 

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


## DeleteHomepageItem

> string DeleteHomepageItem(ctx, homepageItemId)

Delete Homepage Item

### Delete a homepage item. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**homepageItemId** | **int64**| Id of homepage_item | 

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


## DeleteHomepageSection

> string DeleteHomepageSection(ctx, homepageSectionId)

Delete Homepage section

### Delete a homepage section. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**homepageSectionId** | **int64**| Id of homepage_section | 

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


## Homepage

> Homepage Homepage(ctx, homepageId, optional)

Get Homepage

### Get information about a homepage. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**homepageId** | **int64**| Id of homepage | 
 **optional** | ***HomepageOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a HomepageOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Requested fields. | 

### Return type

[**Homepage**](Homepage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HomepageItem

> HomepageItem HomepageItem(ctx, homepageItemId, optional)

Get Homepage Item

### Get information about a homepage item. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**homepageItemId** | **int64**| Id of homepage item | 
 **optional** | ***HomepageItemOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a HomepageItemOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Requested fields. | 

### Return type

[**HomepageItem**](HomepageItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HomepageSection

> HomepageSection HomepageSection(ctx, homepageSectionId, optional)

Get Homepage section

### Get information about a homepage section. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**homepageSectionId** | **int64**| Id of homepage section | 
 **optional** | ***HomepageSectionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a HomepageSectionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Requested fields. | 

### Return type

[**HomepageSection**](HomepageSection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchHomepages

> []Homepage SearchHomepages(ctx, optional)

Search Homepages

### Search Homepages  If multiple search params are given and `filter_or` is FALSE or not specified, search params are combined in a logical AND operation. Only rows that match *all* search param criteria will be returned.  If `filter_or` is TRUE, multiple search params are combined in a logical OR operation. Results will include rows that match **any** of the search criteria.  String search params use case-insensitive matching. String search params can contain `%` and '_' as SQL LIKE pattern match wildcard expressions. example=\"dan%\" will match \"danger\" and \"Danzig\" but not \"David\" example=\"D_m%\" will match \"Damage\" and \"dump\"  Integer search params can accept a single value or a comma separated list of values. The multiple values will be combined under a logical OR operation - results will match at least one of the given values.  Most search params can accept \"IS NULL\" and \"NOT NULL\" as special expressions to match or exclude (respectively) rows where the column is null.  Boolean search params accept only \"true\" and \"false\" as values.  

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SearchHomepagesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SearchHomepagesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **title** | **optional.String**| Matches homepage title. | 
 **createdAt** | **optional.String**| Matches the timestamp for when the homepage was created. | 
 **firstName** | **optional.String**| The first name of the user who created this homepage. | 
 **lastName** | **optional.String**| The last name of the user who created this homepage. | 
 **fields** | **optional.String**| Requested fields. | 
 **favorited** | **optional.Bool**| Return favorited homepages when true. | 
 **creatorId** | **optional.String**| Filter on homepages created by a particular user. | 
 **sorts** | **optional.String**| The fields to sort the results by | 
 **page** | **optional.Int64**| The page to return. | 
 **perPage** | **optional.Int64**| The number of items in the returned page. | 
 **offset** | **optional.Int64**| The number of items to skip before returning any. (used with limit and takes priority over page and per_page) | 
 **limit** | **optional.Int64**| The maximum number of items to return. (used with offset and takes priority over page and per_page) | 
 **filterOr** | **optional.Bool**| Combine given search criteria in a boolean OR expression | 

### Return type

[**[]Homepage**](Homepage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateHomepage

> Homepage UpdateHomepage(ctx, homepageId, body, optional)

Update Homepage

### Update a homepage definition. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**homepageId** | **int64**| Id of homepage | 
**body** | [**Homepage**](Homepage.md)| Homepage | 
 **optional** | ***UpdateHomepageOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateHomepageOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **optional.String**| Requested fields. | 

### Return type

[**Homepage**](Homepage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateHomepageItem

> HomepageItem UpdateHomepageItem(ctx, homepageItemId, body, optional)

Update Homepage Item

### Update a homepage item definition. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**homepageItemId** | **int64**| Id of homepage item | 
**body** | [**HomepageItem**](HomepageItem.md)| Homepage Item | 
 **optional** | ***UpdateHomepageItemOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateHomepageItemOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **optional.String**| Requested fields. | 

### Return type

[**HomepageItem**](HomepageItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateHomepageSection

> HomepageSection UpdateHomepageSection(ctx, homepageSectionId, body, optional)

Update Homepage section

### Update a homepage section definition. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**homepageSectionId** | **int64**| Id of homepage section | 
**body** | [**HomepageSection**](HomepageSection.md)| Homepage section | 
 **optional** | ***UpdateHomepageSectionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateHomepageSectionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **optional.String**| Requested fields. | 

### Return type

[**HomepageSection**](HomepageSection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

