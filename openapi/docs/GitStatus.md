# GitStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** | Git action: add, delete, etc | [optional] [readonly] 
**Conflict** | **bool** | When true, changes to the local file conflict with the remote repository | [optional] [readonly] 
**Revertable** | **bool** | When true, the file can be reverted to an earlier state | [optional] [readonly] 
**Text** | **string** | Git description of the action | [optional] [readonly] 
**Can** | **map[string]bool** | Operations the current user is able to perform on this object | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


