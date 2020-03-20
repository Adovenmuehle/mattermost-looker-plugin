# UserAttributeGroupValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Can** | **map[string]bool** | Operations the current user is able to perform on this object | [optional] [readonly] 
**Id** | **int64** | Unique Id of this group-attribute relation | [optional] [readonly] 
**GroupId** | **int64** | Id of group | [optional] [readonly] 
**UserAttributeId** | **int64** | Id of user attribute | [optional] [readonly] 
**ValueIsHidden** | **bool** | If true, the \&quot;value\&quot; field will be null, because the attribute settings block access to this value | [optional] [readonly] 
**Rank** | **int64** | Precedence for resolving value for user | [optional] [readonly] 
**Value** | **string** | Value of user attribute for group | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


