# \UserAttributeApi

All URIs are relative to *https://mattermost.looker.com:19999/api/3.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AllUserAttributeGroupValues**](UserAttributeApi.md#AllUserAttributeGroupValues) | **Get** /user_attributes/{user_attribute_id}/group_values | Get User Attribute Group Values
[**AllUserAttributes**](UserAttributeApi.md#AllUserAttributes) | **Get** /user_attributes | Get All User Attributes
[**CreateUserAttribute**](UserAttributeApi.md#CreateUserAttribute) | **Post** /user_attributes | Create User Attribute
[**DeleteUserAttribute**](UserAttributeApi.md#DeleteUserAttribute) | **Delete** /user_attributes/{user_attribute_id} | Delete User Attribute
[**SetUserAttributeGroupValues**](UserAttributeApi.md#SetUserAttributeGroupValues) | **Post** /user_attributes/{user_attribute_id}/group_values | Set User Attribute Group Values
[**UpdateUserAttribute**](UserAttributeApi.md#UpdateUserAttribute) | **Patch** /user_attributes/{user_attribute_id} | Update User Attribute
[**UserAttribute**](UserAttributeApi.md#UserAttribute) | **Get** /user_attributes/{user_attribute_id} | Get User Attribute



## AllUserAttributeGroupValues

> []UserAttributeGroupValue AllUserAttributeGroupValues(ctx, userAttributeId, optional)

Get User Attribute Group Values

### Returns all values of a user attribute defined by user groups, in precedence order.  A user may be a member of multiple groups which define different values for a given user attribute. The order of group-values in the response determines precedence for selecting which group-value applies to a given user.  For more information, see [Set User Attribute Group Values](#!/UserAttribute/set_user_attribute_group_values).  Results will only include groups that the caller's user account has permission to see. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userAttributeId** | **int64**| Id of user attribute | 
 **optional** | ***AllUserAttributeGroupValuesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AllUserAttributeGroupValuesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Requested fields. | 

### Return type

[**[]UserAttributeGroupValue**](UserAttributeGroupValue.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AllUserAttributes

> []UserAttribute AllUserAttributes(ctx, optional)

Get All User Attributes

### Get information about all user attributes. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AllUserAttributesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AllUserAttributesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **optional.String**| Requested fields. | 
 **sorts** | **optional.String**| Fields to order the results by. Sortable fields include: name, label | 

### Return type

[**[]UserAttribute**](UserAttribute.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUserAttribute

> UserAttribute CreateUserAttribute(ctx, body, optional)

Create User Attribute

### Create a new user attribute  Permission information for a user attribute is conveyed through the `can` and `user_can_edit` fields. The `user_can_edit` field indicates whether an attribute is user-editable _anywhere_ in the application. The `can` field gives more granular access information, with the `set_value` child field indicating whether an attribute's value can be set by [Setting the User Attribute User Value](#!/User/set_user_attribute_user_value).  Note: `name` and `label` fields must be unique across all user attributes in the Looker instance. Attempting to create a new user attribute with a name or label that duplicates an existing user attribute will fail with a 422 error. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**UserAttribute**](UserAttribute.md)| User Attribute | 
 **optional** | ***CreateUserAttributeOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateUserAttributeOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Requested fields. | 

### Return type

[**UserAttribute**](UserAttribute.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUserAttribute

> string DeleteUserAttribute(ctx, userAttributeId)

Delete User Attribute

### Delete a user attribute (admin only). 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userAttributeId** | **int64**| Id of user_attribute | 

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


## SetUserAttributeGroupValues

> []UserAttributeGroupValue SetUserAttributeGroupValues(ctx, userAttributeId, body)

Set User Attribute Group Values

### Define values for a user attribute across a set of groups, in priority order.  This function defines all values for a user attribute defined by user groups. This is a global setting, potentially affecting all users in the system. This function replaces any existing group value definitions for the indicated user attribute.  The value of a user attribute for a given user is determined by searching the following locations, in this order:  1. the user's account settings 2. the groups that the user is a member of 3. the default value of the user attribute, if any  The user may be a member of multiple groups which define different values for that user attribute. The order of items in the group_values parameter determines which group takes priority for that user. Lowest array index wins.  An alternate method to indicate the selection precedence of group-values is to assign numbers to the 'rank' property of each group-value object in the array. Lowest 'rank' value wins. If you use this technique, you must assign a rank value to every group-value object in the array.    To set a user attribute value for a single user, see [Set User Attribute User Value](#!/User/set_user_attribute_user_value). To set a user attribute value for all members of a group, see [Set User Attribute Group Value](#!/Group/update_user_attribute_group_value). 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userAttributeId** | **int64**| Id of user attribute | 
**body** | [**[]UserAttributeGroupValue**](UserAttributeGroupValue.md)| Array of group values. | 

### Return type

[**[]UserAttributeGroupValue**](UserAttributeGroupValue.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserAttribute

> UserAttribute UpdateUserAttribute(ctx, userAttributeId, body, optional)

Update User Attribute

### Update a user attribute definition. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userAttributeId** | **int64**| Id of user attribute | 
**body** | [**UserAttribute**](UserAttribute.md)| User Attribute | 
 **optional** | ***UpdateUserAttributeOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateUserAttributeOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **optional.String**| Requested fields. | 

### Return type

[**UserAttribute**](UserAttribute.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserAttribute

> UserAttribute UserAttribute(ctx, userAttributeId, optional)

Get User Attribute

### Get information about a user attribute. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userAttributeId** | **int64**| Id of user attribute | 
 **optional** | ***UserAttributeOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UserAttributeOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Requested fields. | 

### Return type

[**UserAttribute**](UserAttribute.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

