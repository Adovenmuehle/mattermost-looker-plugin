# SessionConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Can** | **map[string]bool** | Operations the current user is able to perform on this object | [optional] [readonly] 
**AllowPersistentSessions** | **bool** | Allow users to have persistent sessions when they login | [optional] 
**SessionMinutes** | **int64** | Number of minutes for user sessions.  Must be between 5 and 43200 | [optional] 
**UnlimitedSessionsPerUser** | **bool** | Allow users to have an unbounded number of concurrent sessions (otherwise, users will be limited to only one session at a time). | [optional] 
**UseInactivityBasedLogout** | **bool** | Enforce session logout for sessions that are inactive for 15 minutes. | [optional] 
**TrackSessionLocation** | **bool** | Track location of session when user logs in. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


