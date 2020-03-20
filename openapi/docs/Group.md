# Group

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Can** | **map[string]bool** | Operations the current user is able to perform on this object | [optional] [readonly] 
**CanAddToContentMetadata** | **bool** | Group can be used in content access controls | [optional] 
**ContainsCurrentUser** | **bool** | Currently logged in user is group member | [optional] [readonly] 
**ExternalGroupId** | **string** | External Id group if embed group | [optional] [readonly] 
**ExternallyManaged** | **bool** | Group membership controlled outside of Looker | [optional] [readonly] 
**Id** | **int64** | Unique Id | [optional] [readonly] 
**IncludeByDefault** | **bool** | New users are added to this group by default | [optional] [readonly] 
**Name** | **string** | Name of group | [optional] 
**UserCount** | **int64** | Number of users included in this group | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


