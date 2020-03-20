# ResultMakerWithIdVisConfigAndDynamicFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Unique Id. | [optional] [readonly] 
**DynamicFields** | **string** | JSON string of dynamic field information. | [optional] [readonly] 
**Filterables** | [**[]ResultMakerFilterables**](ResultMakerFilterables.md) | array of items that can be filtered and information about them. | [optional] [readonly] 
**Sorts** | **[]string** | Sorts of the constituent Look, Query, or Merge Query | [optional] [readonly] 
**MergeResultId** | **string** | ID of merge result if this is a merge_result. | [optional] [readonly] 
**Total** | **bool** | Total of the constituent Look, Query, or Merge Query | [optional] [readonly] 
**QueryId** | **int64** | ID of query if this is a query. | [optional] [readonly] 
**Query** | [**Query**](Query.md) |  | [optional] 
**VisConfig** | **map[string]string** | Vis config of the constituent Query, or Merge Query. | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


