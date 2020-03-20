# \GroupApi

All URIs are relative to *https://mattermost.looker.com:19999/api/3.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddGroupGroup**](GroupApi.md#AddGroupGroup) | **Post** /groups/{group_id}/groups | Add a Group to Group
[**AddGroupUser**](GroupApi.md#AddGroupUser) | **Post** /groups/{group_id}/users | Add a User to Group
[**AllGroupGroups**](GroupApi.md#AllGroupGroups) | **Get** /groups/{group_id}/groups | Get All Groups in Group
[**AllGroupUsers**](GroupApi.md#AllGroupUsers) | **Get** /groups/{group_id}/users | Get All Users in Group
[**AllGroups**](GroupApi.md#AllGroups) | **Get** /groups | Get All Groups
[**CreateGroup**](GroupApi.md#CreateGroup) | **Post** /groups | Create Group
[**DeleteGroup**](GroupApi.md#DeleteGroup) | **Delete** /groups/{group_id} | Delete Group
[**DeleteGroupFromGroup**](GroupApi.md#DeleteGroupFromGroup) | **Delete** /groups/{group_id}/groups/{deleting_group_id} | Deletes a Group from Group
[**DeleteGroupUser**](GroupApi.md#DeleteGroupUser) | **Delete** /groups/{group_id}/users/{user_id} | Remove a User from Group
[**DeleteUserAttributeGroupValue**](GroupApi.md#DeleteUserAttributeGroupValue) | **Delete** /groups/{group_id}/attribute_values/{user_attribute_id} | Delete User Attribute Group Value
[**Group**](GroupApi.md#Group) | **Get** /groups/{group_id} | Get Group
[**SearchGroups**](GroupApi.md#SearchGroups) | **Get** /groups/search | Search Groups
[**UpdateGroup**](GroupApi.md#UpdateGroup) | **Patch** /groups/{group_id} | Update Group
[**UpdateUserAttributeGroupValue**](GroupApi.md#UpdateUserAttributeGroupValue) | **Patch** /groups/{group_id}/attribute_values/{user_attribute_id} | Set User Attribute Group Value



## AddGroupGroup

> Group AddGroupGroup(ctx, groupId, body)

Add a Group to Group

### Adds a new group to a group. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **int64**| Id of group | 
**body** | [**GroupIdForGroupInclusion**](GroupIdForGroupInclusion.md)| Group id to add | 

### Return type

[**Group**](Group.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddGroupUser

> User AddGroupUser(ctx, groupId, body)

Add a User to Group

### Adds a new user to a group. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **int64**| Id of group | 
**body** | [**GroupIdForGroupUserInclusion**](GroupIdForGroupUserInclusion.md)| User id to add | 

### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AllGroupGroups

> []Group AllGroupGroups(ctx, groupId, optional)

Get All Groups in Group

### Get information about all the groups in a group 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **int64**| Id of group | 
 **optional** | ***AllGroupGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AllGroupGroupsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Requested fields. | 

### Return type

[**[]Group**](Group.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AllGroupUsers

> []User AllGroupUsers(ctx, groupId, optional)

Get All Users in Group

### Get information about all the users directly included in a group. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **int64**| Id of group | 
 **optional** | ***AllGroupUsersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AllGroupUsersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Requested fields. | 
 **page** | **optional.Int64**| Requested page. | 
 **perPage** | **optional.Int64**| Results per page. | 
 **sorts** | **optional.String**| Fields to sort by. | 

### Return type

[**[]User**](User.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AllGroups

> []Group AllGroups(ctx, optional)

Get All Groups

### Get information about all groups. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AllGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AllGroupsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **optional.String**| Requested fields. | 
 **page** | **optional.Int64**| Requested page. | 
 **perPage** | **optional.Int64**| Results per page. | 
 **sorts** | **optional.String**| Fields to sort by. | 
 **ids** | [**optional.Interface of []int64**](int64.md)| Optional of ids to get specific groups. | 
 **contentMetadataId** | **optional.Int64**| Id of content metadata to which groups must have access. | 
 **canAddToContentMetadata** | **optional.Bool**| Select only groups that either can/cannot be given access to content. | 

### Return type

[**[]Group**](Group.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateGroup

> Group CreateGroup(ctx, body, optional)

Create Group

### Creates a new group (admin only). 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**Group**](Group.md)| Group | 
 **optional** | ***CreateGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateGroupOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Requested fields. | 

### Return type

[**Group**](Group.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGroup

> string DeleteGroup(ctx, groupId)

Delete Group

### Deletes a group (admin only). 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **int64**| Id of group | 

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


## DeleteGroupFromGroup

> DeleteGroupFromGroup(ctx, groupId, deletingGroupId)

Deletes a Group from Group

### Removes a group from a group. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **int64**| Id of group | 
**deletingGroupId** | **int64**| Id of group to delete | 

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


## DeleteGroupUser

> DeleteGroupUser(ctx, groupId, userId)

Remove a User from Group

### Removes a user from a group. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **int64**| Id of group | 
**userId** | **int64**| Id of user to remove from group | 

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


## DeleteUserAttributeGroupValue

> DeleteUserAttributeGroupValue(ctx, groupId, userAttributeId)

Delete User Attribute Group Value

### Remove a user attribute value from a group. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **int64**| Id of group | 
**userAttributeId** | **int64**| Id of user attribute | 

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


## Group

> Group Group(ctx, groupId, optional)

Get Group

### Get information about a group. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **int64**| Id of group | 
 **optional** | ***GroupOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Requested fields. | 

### Return type

[**Group**](Group.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchGroups

> []Group SearchGroups(ctx, optional)

Search Groups

### Search groups  Returns all group records that match the given search criteria.  If multiple search params are given and `filter_or` is FALSE or not specified, search params are combined in a logical AND operation. Only rows that match *all* search param criteria will be returned.  If `filter_or` is TRUE, multiple search params are combined in a logical OR operation. Results will include rows that match **any** of the search criteria.  String search params use case-insensitive matching. String search params can contain `%` and '_' as SQL LIKE pattern match wildcard expressions. example=\"dan%\" will match \"danger\" and \"Danzig\" but not \"David\" example=\"D_m%\" will match \"Damage\" and \"dump\"  Integer search params can accept a single value or a comma separated list of values. The multiple values will be combined under a logical OR operation - results will match at least one of the given values.  Most search params can accept \"IS NULL\" and \"NOT NULL\" as special expressions to match or exclude (respectively) rows where the column is null.  Boolean search params accept only \"true\" and \"false\" as values.  

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SearchGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SearchGroupsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **optional.String**| Requested fields. | 
 **limit** | **optional.Int64**| Number of results to return (used with &#x60;offset&#x60;). | 
 **offset** | **optional.Int64**| Number of results to skip before returning any (used with &#x60;limit&#x60;). | 
 **sorts** | **optional.String**| Fields to sort by. | 
 **filterOr** | **optional.Bool**| Combine given search criteria in a boolean OR expression | 
 **id** | **optional.Int64**| Match group id. | 
 **name** | **optional.String**| Match group name. | 
 **externalGroupId** | **optional.Bool**| Match group external_group_id. | 
 **externallyManaged** | **optional.Bool**| Match group externally_managed. | 
 **externallyOrphaned** | **optional.Bool**| Match group externally_orphaned. | 

### Return type

[**[]Group**](Group.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGroup

> Group UpdateGroup(ctx, groupId, body, optional)

Update Group

### Updates the a group (admin only).

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **int64**| Id of group | 
**body** | [**Group**](Group.md)| Group | 
 **optional** | ***UpdateGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateGroupOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **optional.String**| Requested fields. | 

### Return type

[**Group**](Group.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserAttributeGroupValue

> UserAttributeGroupValue UpdateUserAttributeGroupValue(ctx, groupId, userAttributeId, body)

Set User Attribute Group Value

### Set the value of a user attribute for a group.  For information about how user attribute values are calculated, see [Set User Attribute Group Values](#!/UserAttribute/set_user_attribute_group_values). 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **int64**| Id of group | 
**userAttributeId** | **int64**| Id of user attribute | 
**body** | [**UserAttributeGroupValue**](UserAttributeGroupValue.md)| New value for group. | 

### Return type

[**UserAttributeGroupValue**](UserAttributeGroupValue.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

