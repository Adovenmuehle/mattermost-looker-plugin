# Project

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Can** | **map[string]bool** | Operations the current user is able to perform on this object | [optional] [readonly] 
**Id** | **string** | Project Id | [optional] [readonly] 
**Name** | **string** | Project display name | [optional] 
**UsesGit** | **bool** | If true the project is configured with a git repository | [optional] [readonly] 
**GitRemoteUrl** | **string** | Git remote repository url | [optional] 
**GitUsername** | **string** | Git username for HTTPS authentication. (For production only, if using user attributes.) | [optional] 
**GitPassword** | **string** | (Write-Only) Git password for HTTPS authentication. (For production only, if using user attributes.) | [optional] 
**GitUsernameUserAttribute** | **string** | User attribute name for username in per-user HTTPS authentication. | [optional] 
**GitPasswordUserAttribute** | **string** | User attribute name for password in per-user HTTPS authentication. | [optional] 
**GitServiceName** | **string** | Name of the git service provider | [optional] 
**DeploySecret** | **string** | (Write-Only) Optional secret token with which to authenticate requests to the webhook deploy endpoint. If not set, endpoint is unauthenticated. | [optional] 
**UnsetDeploySecret** | **bool** | (Write-Only) When true, unsets the deploy secret to allow unauthenticated access to the webhook deploy endpoint. | [optional] 
**PullRequestMode** | **string** | The git pull request policy for this project. Valid values are: \&quot;off\&quot;, \&quot;links\&quot;, \&quot;recommended\&quot;, \&quot;required\&quot;. | [optional] 
**ValidationRequired** | **bool** | Validation policy: If true, the project must pass validation checks before project changes can be committed to the git repository | [optional] 
**FoldersEnabled** | **bool** | If true, folders are enabled for this project | [optional] 
**AllowWarnings** | **bool** | Validation policy: If true, the project can be committed with warnings when &#x60;validation_required&#x60; is true. (&#x60;allow_warnings&#x60; does nothing if &#x60;validation_required&#x60; is false). | [optional] 
**IsExample** | **bool** | If true the project is an example project and cannot be modified | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


