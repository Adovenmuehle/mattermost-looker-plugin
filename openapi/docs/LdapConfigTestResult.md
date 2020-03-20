# LdapConfigTestResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Details** | **string** | Additional details for error cases | [optional] [readonly] 
**Issues** | [**[]LdapConfigTestIssue**](LDAPConfigTestIssue.md) | Array of issues/considerations about the result | [optional] [readonly] 
**Message** | **string** | Short human readable test about the result | [optional] [readonly] 
**Status** | **string** | Test status code: always &#39;success&#39; or &#39;error&#39; | [optional] [readonly] 
**Trace** | **string** | A more detailed trace of incremental results during auth tests | [optional] [readonly] 
**User** | [**LdapUser**](LDAPUser.md) |  | [optional] 
**Url** | **string** | Link to ldap config | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


