# \FolderApi

All URIs are relative to *https://mattermost.looker.com:19999/api/3.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AllFolders**](FolderApi.md#AllFolders) | **Get** /folders | Get All Folders
[**CreateFolder**](FolderApi.md#CreateFolder) | **Post** /folders | Create Folder
[**DeleteFolder**](FolderApi.md#DeleteFolder) | **Delete** /folders/{folder_id} | Delete Folder
[**Folder**](FolderApi.md#Folder) | **Get** /folders/{folder_id} | Get Folder
[**FolderAncestors**](FolderApi.md#FolderAncestors) | **Get** /folders/{folder_id}/ancestors | Get Folder Ancestors
[**FolderChildren**](FolderApi.md#FolderChildren) | **Get** /folders/{folder_id}/children | Get Folder Children
[**FolderChildrenSearch**](FolderApi.md#FolderChildrenSearch) | **Get** /folders/{folder_id}/children/search | Search Folder Children
[**FolderDashboards**](FolderApi.md#FolderDashboards) | **Get** /folders/{folder_id}/dashboards | Get Folder Dashboards
[**FolderLooks**](FolderApi.md#FolderLooks) | **Get** /folders/{folder_id}/looks | Get Folder Looks
[**FolderParent**](FolderApi.md#FolderParent) | **Get** /folders/{folder_id}/parent | Get Folder Parent
[**SearchFolders**](FolderApi.md#SearchFolders) | **Get** /folders/search | Search Folders
[**UpdateFolder**](FolderApi.md#UpdateFolder) | **Patch** /folders/{folder_id} | Update Folder



## AllFolders

> []Folder AllFolders(ctx, optional)

Get All Folders

### Get information about all folders.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AllFoldersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AllFoldersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **optional.String**| Requested fields. | 

### Return type

[**[]Folder**](Folder.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateFolder

> Folder CreateFolder(ctx, body)

Create Folder

### Create a folder with specified information.  Caller must have permission to edit the parent folder and to create folders, otherwise the request returns 404 Not Found. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**CreateFolder**](CreateFolder.md)| Folder parameters | 

### Return type

[**Folder**](Folder.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFolder

> string DeleteFolder(ctx, folderId)

Delete Folder

### Delete the folder with a specific id including any children folders. **DANGER** this will delete all looks and dashboards in the folder. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **string**| Id of folder | 

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


## Folder

> Folder Folder(ctx, folderId, optional)

Get Folder

### Get information about the folder with a specific id.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **string**| Id of folder | 
 **optional** | ***FolderOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FolderOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Requested fields. | 

### Return type

[**Folder**](Folder.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FolderAncestors

> []Folder FolderAncestors(ctx, folderId, optional)

Get Folder Ancestors

### Get the ancestors of a folder

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **string**| Id of folder | 
 **optional** | ***FolderAncestorsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FolderAncestorsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Requested fields. | 

### Return type

[**[]Folder**](Folder.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FolderChildren

> []Folder FolderChildren(ctx, folderId, optional)

Get Folder Children

### Get the children of a folder.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **string**| Id of folder | 
 **optional** | ***FolderChildrenOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FolderChildrenOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Requested fields. | 
 **page** | **optional.Int64**| Requested page. | 
 **perPage** | **optional.Int64**| Results per page. | 
 **sorts** | **optional.String**| Fields to sort by. | 

### Return type

[**[]Folder**](Folder.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FolderChildrenSearch

> []Folder FolderChildrenSearch(ctx, folderId, optional)

Search Folder Children

### Search the children of a folder

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **string**| Id of folder | 
 **optional** | ***FolderChildrenSearchOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FolderChildrenSearchOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Requested fields. | 
 **sorts** | **optional.String**| Fields to sort by. | 
 **name** | **optional.String**| Match folder name. | 

### Return type

[**[]Folder**](Folder.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FolderDashboards

> []Dashboard FolderDashboards(ctx, folderId, optional)

Get Folder Dashboards

### Get the dashboards in a folder

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **string**| Id of folder | 
 **optional** | ***FolderDashboardsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FolderDashboardsOpts struct


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


## FolderLooks

> []LookWithQuery FolderLooks(ctx, folderId, optional)

Get Folder Looks

### Get the looks in a folder

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **string**| Id of folder | 
 **optional** | ***FolderLooksOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FolderLooksOpts struct


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


## FolderParent

> Folder FolderParent(ctx, folderId, optional)

Get Folder Parent

### Get the parent of a folder

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **string**| Id of folder | 
 **optional** | ***FolderParentOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FolderParentOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Requested fields. | 

### Return type

[**Folder**](Folder.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchFolders

> []Folder SearchFolders(ctx, optional)

Search Folders

Search for folders by creator id, parent id, name, etc

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SearchFoldersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SearchFoldersOpts struct


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
 **parentId** | **optional.String**| Filter on a children of a particular folder. | 
 **creatorId** | **optional.String**| Filter on folder created by a particular user. | 
 **filterOr** | **optional.Bool**| Combine given search criteria in a boolean OR expression | 

### Return type

[**[]Folder**](Folder.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFolder

> Folder UpdateFolder(ctx, folderId, body)

Update Folder

### Update the folder with a specific id.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **string**| Id of folder | 
**body** | [**UpdateFolder**](UpdateFolder.md)| Folder parameters | 

### Return type

[**Folder**](Folder.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

