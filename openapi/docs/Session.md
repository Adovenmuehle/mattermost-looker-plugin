# Session

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Can** | **map[string]bool** | Operations the current user is able to perform on this object | [optional] [readonly] 
**Id** | **int64** | Unique Id | [optional] [readonly] 
**IpAddress** | **string** | IP address of user when this session was initiated | [optional] [readonly] 
**Browser** | **string** | User&#39;s browser type | [optional] [readonly] 
**OperatingSystem** | **string** | User&#39;s Operating System | [optional] [readonly] 
**City** | **string** | City component of user location (derived from IP address) | [optional] [readonly] 
**State** | **string** | State component of user location (derived from IP address) | [optional] [readonly] 
**Country** | **string** | Country component of user location (derived from IP address) | [optional] [readonly] 
**CredentialsType** | **string** | Type of credentials used for logging in this session | [optional] [readonly] 
**ExtendedAt** | **string** | Time when this session was last extended by the user | [optional] [readonly] 
**ExtendedCount** | **int64** | Number of times this session was extended | [optional] [readonly] 
**SudoUserId** | **int64** | Actual user in the case when this session represents one user sudo&#39;ing as another | [optional] [readonly] 
**CreatedAt** | **string** | Time when this session was initiated | [optional] [readonly] 
**ExpiresAt** | **string** | Time when this session will expire | [optional] [readonly] 
**Url** | **string** | Link to get this item | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


