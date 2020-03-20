# CreateQueryTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Can** | **map[string]bool** | Operations the current user is able to perform on this object | [optional] [readonly] 
**QueryId** | **int64** | Id of query to run | 
**ResultFormat** | **string** | Desired async query result format. Valid values are: \&quot;json\&quot;, \&quot;json_detail\&quot;, \&quot;json_fe\&quot;, \&quot;csv\&quot;, \&quot;html\&quot;, \&quot;md\&quot;, \&quot;txt\&quot;, \&quot;xlsx\&quot;, \&quot;gsxml\&quot;. | 
**Source** | **string** | Source of query task | [optional] 
**Deferred** | **bool** | Create the task but defer execution | [optional] 
**LookId** | **int64** | Id of look associated with query. | [optional] 
**DashboardId** | **string** | Id of dashboard associated with query. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


