# Datagroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Can** | **map[string]bool** | Operations the current user is able to perform on this object | [optional] [readonly] 
**CreatedAt** | **int64** | UNIX timestamp at which this entry was created. | [optional] [readonly] 
**Id** | **string** | Unique ID of the datagroup | [optional] [readonly] 
**ModelName** | **string** | Name of the model containing the datagroup. Unique when combined with name. | [optional] [readonly] 
**Name** | **string** | Name of the datagroup. Unique when combined with model_name. | [optional] [readonly] 
**StaleBefore** | **int64** | UNIX timestamp before which cache entries are considered stale. Cannot be in the future. | [optional] 
**TriggerCheckAt** | **int64** | UNIX timestamp at which this entry trigger was last checked. | [optional] [readonly] 
**TriggerError** | **string** | The message returned with the error of the last trigger check. | [optional] [readonly] 
**TriggerValue** | **string** | The value of the trigger when last checked. | [optional] [readonly] 
**TriggeredAt** | **int64** | UNIX timestamp at which this entry became triggered. Cannot be in the future. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


