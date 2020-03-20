# ContentMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Can** | **map[string]bool** | Operations the current user is able to perform on this object | [optional] [readonly] 
**Id** | **int64** | Unique Id | [optional] [readonly] 
**Name** | **string** | Name or title of underlying content | [optional] [readonly] 
**ParentId** | **int64** | Id of Parent Content | [optional] [readonly] 
**SpaceId** | **string** | Id of associated space when content_type is \&quot;space\&quot; | [optional] [readonly] 
**FolderId** | **string** | Id of associated folder when content_type is \&quot;space\&quot; | [optional] [readonly] 
**DashboardId** | **string** | Id of associated dashboard when content_type is \&quot;dashboard\&quot; | [optional] [readonly] 
**LookId** | **int64** | Id of associated look when content_type is \&quot;look\&quot; | [optional] [readonly] 
**ContentType** | **string** | Content Type (\&quot;dashboard\&quot;, \&quot;look\&quot;, or \&quot;space\&quot;) | [optional] [readonly] 
**Inherits** | **bool** | Whether content inherits its access levels from parent | [optional] 
**InheritingId** | **int64** | Id of Inherited Content | [optional] [readonly] 
**Slug** | **string** | Content Slug | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


