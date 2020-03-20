# \RoleApi

All URIs are relative to *https://mattermost.looker.com:19999/api/3.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AllModelSets**](RoleApi.md#AllModelSets) | **Get** /model_sets | Get All Model Sets
[**AllPermissionSets**](RoleApi.md#AllPermissionSets) | **Get** /permission_sets | Get All Permission Sets
[**AllPermissions**](RoleApi.md#AllPermissions) | **Get** /permissions | Get All Permissions
[**AllRoles**](RoleApi.md#AllRoles) | **Get** /roles | Get All Roles
[**CreateModelSet**](RoleApi.md#CreateModelSet) | **Post** /model_sets | Create Model Set
[**CreatePermissionSet**](RoleApi.md#CreatePermissionSet) | **Post** /permission_sets | Create Permission Set
[**CreateRole**](RoleApi.md#CreateRole) | **Post** /roles | Create Role
[**DeleteModelSet**](RoleApi.md#DeleteModelSet) | **Delete** /model_sets/{model_set_id} | Delete Model Set
[**DeletePermissionSet**](RoleApi.md#DeletePermissionSet) | **Delete** /permission_sets/{permission_set_id} | Delete Permission Set
[**DeleteRole**](RoleApi.md#DeleteRole) | **Delete** /roles/{role_id} | Delete Role
[**ModelSet**](RoleApi.md#ModelSet) | **Get** /model_sets/{model_set_id} | Get Model Set
[**PermissionSet**](RoleApi.md#PermissionSet) | **Get** /permission_sets/{permission_set_id} | Get Permission Set
[**Role**](RoleApi.md#Role) | **Get** /roles/{role_id} | Get Role
[**RoleGroups**](RoleApi.md#RoleGroups) | **Get** /roles/{role_id}/groups | Get Role Groups
[**RoleUsers**](RoleApi.md#RoleUsers) | **Get** /roles/{role_id}/users | Get Role Users
[**SearchModelSets**](RoleApi.md#SearchModelSets) | **Get** /model_sets/search | Search Model Sets
[**SearchPermissionSets**](RoleApi.md#SearchPermissionSets) | **Get** /permission_sets/search | Search Permission Sets
[**SearchRoles**](RoleApi.md#SearchRoles) | **Get** /roles/search | Search Roles
[**SetRoleGroups**](RoleApi.md#SetRoleGroups) | **Put** /roles/{role_id}/groups | Update Role Groups
[**SetRoleUsers**](RoleApi.md#SetRoleUsers) | **Put** /roles/{role_id}/users | Update Role Users
[**UpdateModelSet**](RoleApi.md#UpdateModelSet) | **Patch** /model_sets/{model_set_id} | Update Model Set
[**UpdatePermissionSet**](RoleApi.md#UpdatePermissionSet) | **Patch** /permission_sets/{permission_set_id} | Update Permission Set
[**UpdateRole**](RoleApi.md#UpdateRole) | **Patch** /roles/{role_id} | Update Role



## AllModelSets

> []ModelSet AllModelSets(ctx, optional)

Get All Model Sets

### Get information about all model sets. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AllModelSetsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AllModelSetsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **optional.String**| Requested fields. | 

### Return type

[**[]ModelSet**](ModelSet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AllPermissionSets

> []PermissionSet AllPermissionSets(ctx, optional)

Get All Permission Sets

### Get information about all permission sets. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AllPermissionSetsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AllPermissionSetsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **optional.String**| Requested fields. | 

### Return type

[**[]PermissionSet**](PermissionSet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AllPermissions

> []Permission AllPermissions(ctx, )

Get All Permissions

### Get all supported permissions. 

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]Permission**](Permission.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AllRoles

> []Role AllRoles(ctx, optional)

Get All Roles

### Get information about all roles. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AllRolesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AllRolesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **optional.String**| Requested fields. | 
 **ids** | [**optional.Interface of []int64**](int64.md)| Optional list of ids to get specific roles. | 

### Return type

[**[]Role**](Role.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateModelSet

> ModelSet CreateModelSet(ctx, body)

Create Model Set

### Create a model set with the specified information. Model sets are used by Roles. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**ModelSet**](ModelSet.md)| ModelSet | 

### Return type

[**ModelSet**](ModelSet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePermissionSet

> PermissionSet CreatePermissionSet(ctx, body)

Create Permission Set

### Create a permission set with the specified information. Permission sets are used by Roles. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**PermissionSet**](PermissionSet.md)| Permission Set | 

### Return type

[**PermissionSet**](PermissionSet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRole

> Role CreateRole(ctx, body)

Create Role

### Create a role with the specified information. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**Role**](Role.md)| Role | 

### Return type

[**Role**](Role.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteModelSet

> string DeleteModelSet(ctx, modelSetId)

Delete Model Set

### Delete the model set with a specific id. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**modelSetId** | **int64**| id of model set | 

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


## DeletePermissionSet

> string DeletePermissionSet(ctx, permissionSetId)

Delete Permission Set

### Delete the permission set with a specific id. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**permissionSetId** | **int64**| Id of permission set | 

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


## DeleteRole

> string DeleteRole(ctx, roleId)

Delete Role

### Delete the role with a specific id. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleId** | **int64**| id of role | 

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


## ModelSet

> ModelSet ModelSet(ctx, modelSetId, optional)

Get Model Set

### Get information about the model set with a specific id. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**modelSetId** | **int64**| Id of model set | 
 **optional** | ***ModelSetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ModelSetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Requested fields. | 

### Return type

[**ModelSet**](ModelSet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PermissionSet

> PermissionSet PermissionSet(ctx, permissionSetId, optional)

Get Permission Set

### Get information about the permission set with a specific id. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**permissionSetId** | **int64**| Id of permission set | 
 **optional** | ***PermissionSetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PermissionSetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Requested fields. | 

### Return type

[**PermissionSet**](PermissionSet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Role

> Role Role(ctx, roleId)

Get Role

### Get information about the role with a specific id. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleId** | **int64**| id of role | 

### Return type

[**Role**](Role.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RoleGroups

> []Group RoleGroups(ctx, roleId, optional)

Get Role Groups

### Get information about all the groups with the role that has a specific id. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleId** | **int64**| id of role | 
 **optional** | ***RoleGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a RoleGroupsOpts struct


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


## RoleUsers

> []User RoleUsers(ctx, roleId, optional)

Get Role Users

### Get information about all the users with the role that has a specific id. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleId** | **int64**| id of user | 
 **optional** | ***RoleUsersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a RoleUsersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Requested fields. | 
 **directAssociationOnly** | **optional.Bool**| Get only users associated directly with the role: exclude those only associated through groups. | 

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


## SearchModelSets

> []ModelSet SearchModelSets(ctx, optional)

Search Model Sets

### Search model sets Returns all model set records that match the given search criteria. If multiple search params are given and `filter_or` is FALSE or not specified, search params are combined in a logical AND operation. Only rows that match *all* search param criteria will be returned.  If `filter_or` is TRUE, multiple search params are combined in a logical OR operation. Results will include rows that match **any** of the search criteria.  String search params use case-insensitive matching. String search params can contain `%` and '_' as SQL LIKE pattern match wildcard expressions. example=\"dan%\" will match \"danger\" and \"Danzig\" but not \"David\" example=\"D_m%\" will match \"Damage\" and \"dump\"  Integer search params can accept a single value or a comma separated list of values. The multiple values will be combined under a logical OR operation - results will match at least one of the given values.  Most search params can accept \"IS NULL\" and \"NOT NULL\" as special expressions to match or exclude (respectively) rows where the column is null.  Boolean search params accept only \"true\" and \"false\" as values.  

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SearchModelSetsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SearchModelSetsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **optional.String**| Requested fields. | 
 **limit** | **optional.Int64**| Number of results to return (used with &#x60;offset&#x60;). | 
 **offset** | **optional.Int64**| Number of results to skip before returning any (used with &#x60;limit&#x60;). | 
 **sorts** | **optional.String**| Fields to sort by. | 
 **id** | **optional.Int64**| Match model set id. | 
 **name** | **optional.String**| Match model set name. | 
 **allAccess** | **optional.Bool**| Match model sets by all_access status. | 
 **builtIn** | **optional.Bool**| Match model sets by built_in status. | 
 **filterOr** | **optional.Bool**| Combine given search criteria in a boolean OR expression. | 

### Return type

[**[]ModelSet**](ModelSet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchPermissionSets

> []PermissionSet SearchPermissionSets(ctx, optional)

Search Permission Sets

### Search permission sets Returns all permission set records that match the given search criteria. If multiple search params are given and `filter_or` is FALSE or not specified, search params are combined in a logical AND operation. Only rows that match *all* search param criteria will be returned.  If `filter_or` is TRUE, multiple search params are combined in a logical OR operation. Results will include rows that match **any** of the search criteria.  String search params use case-insensitive matching. String search params can contain `%` and '_' as SQL LIKE pattern match wildcard expressions. example=\"dan%\" will match \"danger\" and \"Danzig\" but not \"David\" example=\"D_m%\" will match \"Damage\" and \"dump\"  Integer search params can accept a single value or a comma separated list of values. The multiple values will be combined under a logical OR operation - results will match at least one of the given values.  Most search params can accept \"IS NULL\" and \"NOT NULL\" as special expressions to match or exclude (respectively) rows where the column is null.  Boolean search params accept only \"true\" and \"false\" as values.  

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SearchPermissionSetsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SearchPermissionSetsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **optional.String**| Requested fields. | 
 **limit** | **optional.Int64**| Number of results to return (used with &#x60;offset&#x60;). | 
 **offset** | **optional.Int64**| Number of results to skip before returning any (used with &#x60;limit&#x60;). | 
 **sorts** | **optional.String**| Fields to sort by. | 
 **id** | **optional.Int64**| Match permission set id. | 
 **name** | **optional.String**| Match permission set name. | 
 **allAccess** | **optional.Bool**| Match permission sets by all_access status. | 
 **builtIn** | **optional.Bool**| Match permission sets by built_in status. | 
 **filterOr** | **optional.Bool**| Combine given search criteria in a boolean OR expression. | 

### Return type

[**[]PermissionSet**](PermissionSet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchRoles

> []Role SearchRoles(ctx, optional)

Search Roles

### Search roles  Returns all role records that match the given search criteria.  If multiple search params are given and `filter_or` is FALSE or not specified, search params are combined in a logical AND operation. Only rows that match *all* search param criteria will be returned.  If `filter_or` is TRUE, multiple search params are combined in a logical OR operation. Results will include rows that match **any** of the search criteria.  String search params use case-insensitive matching. String search params can contain `%` and '_' as SQL LIKE pattern match wildcard expressions. example=\"dan%\" will match \"danger\" and \"Danzig\" but not \"David\" example=\"D_m%\" will match \"Damage\" and \"dump\"  Integer search params can accept a single value or a comma separated list of values. The multiple values will be combined under a logical OR operation - results will match at least one of the given values.  Most search params can accept \"IS NULL\" and \"NOT NULL\" as special expressions to match or exclude (respectively) rows where the column is null.  Boolean search params accept only \"true\" and \"false\" as values.  

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SearchRolesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SearchRolesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **optional.String**| Requested fields. | 
 **limit** | **optional.Int64**| Number of results to return (used with &#x60;offset&#x60;). | 
 **offset** | **optional.Int64**| Number of results to skip before returning any (used with &#x60;limit&#x60;). | 
 **sorts** | **optional.String**| Fields to sort by. | 
 **id** | **optional.Int64**| Match role id. | 
 **name** | **optional.String**| Match role name. | 
 **builtIn** | **optional.Bool**| Match roles by built_in status. | 
 **filterOr** | **optional.Bool**| Combine given search criteria in a boolean OR expression. | 

### Return type

[**[]Role**](Role.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetRoleGroups

> []Group SetRoleGroups(ctx, roleId, body)

Update Role Groups

### Set all groups for a role, removing all existing group associations from that role. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleId** | **int64**| Id of Role | 
**body** | [**[]int64**](int64.md)| Array of Group Ids | 

### Return type

[**[]Group**](Group.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetRoleUsers

> []User SetRoleUsers(ctx, roleId, body)

Update Role Users

### Set all the users of the role with a specific id. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleId** | **int64**| id of role | 
**body** | [**[]int64**](int64.md)| array of user ids for role | 

### Return type

[**[]User**](User.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateModelSet

> ModelSet UpdateModelSet(ctx, modelSetId, body)

Update Model Set

### Update information about the model set with a specific id. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**modelSetId** | **int64**| id of model set | 
**body** | [**ModelSet**](ModelSet.md)| ModelSet | 

### Return type

[**ModelSet**](ModelSet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePermissionSet

> PermissionSet UpdatePermissionSet(ctx, permissionSetId, body)

Update Permission Set

### Update information about the permission set with a specific id. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**permissionSetId** | **int64**| id of permission set | 
**body** | [**PermissionSet**](PermissionSet.md)| Permission Set | 

### Return type

[**PermissionSet**](PermissionSet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRole

> Role UpdateRole(ctx, roleId, body)

Update Role

### Update information about the role with a specific id. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleId** | **int64**| id of role | 
**body** | [**Role**](Role.md)| Role | 

### Return type

[**Role**](Role.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

