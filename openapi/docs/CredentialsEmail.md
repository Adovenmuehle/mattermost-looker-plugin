# CredentialsEmail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Can** | **map[string]bool** | Operations the current user is able to perform on this object | [optional] [readonly] 
**CreatedAt** | **string** | Timestamp for the creation of this credential | [optional] [readonly] 
**Email** | **string** | EMail address used for user login | [optional] 
**ForcedPasswordResetAtNextLogin** | **bool** | Force the user to change their password upon their next login | [optional] 
**IsDisabled** | **bool** | Has this credential been disabled? | [optional] [readonly] 
**LoggedInAt** | **string** | Timestamp for most recent login using credential | [optional] [readonly] 
**PasswordResetUrl** | **string** | Url with one-time use secret token that the user can use to reset password | [optional] [readonly] 
**Type** | **string** | Short name for the type of this kind of credential | [optional] [readonly] 
**Url** | **string** | Link to get this item | [optional] [readonly] 
**UserUrl** | **string** | Link to get this user | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


