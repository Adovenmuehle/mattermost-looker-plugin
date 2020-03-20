# \UserApi

All URIs are relative to *https://mattermost.looker.com:19999/api/3.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AllUserCredentialsApi3s**](UserApi.md#AllUserCredentialsApi3s) | **Get** /users/{user_id}/credentials_api3 | Get All API 3 Credentials
[**AllUserCredentialsEmbeds**](UserApi.md#AllUserCredentialsEmbeds) | **Get** /users/{user_id}/credentials_embed | Get All Embedding Credentials
[**AllUserSessions**](UserApi.md#AllUserSessions) | **Get** /users/{user_id}/sessions | Get All Web Login Sessions
[**AllUsers**](UserApi.md#AllUsers) | **Get** /users | Get All Users
[**CreateUser**](UserApi.md#CreateUser) | **Post** /users | Create User
[**CreateUserCredentialsApi3**](UserApi.md#CreateUserCredentialsApi3) | **Post** /users/{user_id}/credentials_api3 | Create API 3 Credential
[**CreateUserCredentialsEmail**](UserApi.md#CreateUserCredentialsEmail) | **Post** /users/{user_id}/credentials_email | Create Email/Password Credential
[**CreateUserCredentialsEmailPasswordReset**](UserApi.md#CreateUserCredentialsEmailPasswordReset) | **Post** /users/{user_id}/credentials_email/password_reset | Create Password Reset Token
[**CreateUserCredentialsTotp**](UserApi.md#CreateUserCredentialsTotp) | **Post** /users/{user_id}/credentials_totp | Create Two-Factor Credential
[**DeleteUser**](UserApi.md#DeleteUser) | **Delete** /users/{user_id} | Delete User
[**DeleteUserAttributeUserValue**](UserApi.md#DeleteUserAttributeUserValue) | **Delete** /users/{user_id}/attribute_values/{user_attribute_id} | Delete User Attribute User Value
[**DeleteUserCredentialsApi3**](UserApi.md#DeleteUserCredentialsApi3) | **Delete** /users/{user_id}/credentials_api3/{credentials_api3_id} | Delete API 3 Credential
[**DeleteUserCredentialsEmail**](UserApi.md#DeleteUserCredentialsEmail) | **Delete** /users/{user_id}/credentials_email | Delete Email/Password Credential
[**DeleteUserCredentialsEmbed**](UserApi.md#DeleteUserCredentialsEmbed) | **Delete** /users/{user_id}/credentials_embed/{credentials_embed_id} | Delete Embedding Credential
[**DeleteUserCredentialsGoogle**](UserApi.md#DeleteUserCredentialsGoogle) | **Delete** /users/{user_id}/credentials_google | Delete Google Auth Credential
[**DeleteUserCredentialsLdap**](UserApi.md#DeleteUserCredentialsLdap) | **Delete** /users/{user_id}/credentials_ldap | Delete LDAP Credential
[**DeleteUserCredentialsLookerOpenid**](UserApi.md#DeleteUserCredentialsLookerOpenid) | **Delete** /users/{user_id}/credentials_looker_openid | Delete Looker OpenId Credential
[**DeleteUserCredentialsOidc**](UserApi.md#DeleteUserCredentialsOidc) | **Delete** /users/{user_id}/credentials_oidc | Delete OIDC Auth Credential
[**DeleteUserCredentialsSaml**](UserApi.md#DeleteUserCredentialsSaml) | **Delete** /users/{user_id}/credentials_saml | Delete Saml Auth Credential
[**DeleteUserCredentialsTotp**](UserApi.md#DeleteUserCredentialsTotp) | **Delete** /users/{user_id}/credentials_totp | Delete Two-Factor Credential
[**DeleteUserSession**](UserApi.md#DeleteUserSession) | **Delete** /users/{user_id}/sessions/{session_id} | Delete Web Login Session
[**Me**](UserApi.md#Me) | **Get** /user | Get Current User
[**SearchUsers**](UserApi.md#SearchUsers) | **Get** /users/search | Search Users
[**SearchUsersNames**](UserApi.md#SearchUsersNames) | **Get** /users/search/names/{pattern} | Search User Names
[**SetUserAttributeUserValue**](UserApi.md#SetUserAttributeUserValue) | **Patch** /users/{user_id}/attribute_values/{user_attribute_id} | Set User Attribute User Value
[**SetUserRoles**](UserApi.md#SetUserRoles) | **Put** /users/{user_id}/roles | Set User Roles
[**UpdateUser**](UserApi.md#UpdateUser) | **Patch** /users/{user_id} | Update User
[**UpdateUserCredentialsEmail**](UserApi.md#UpdateUserCredentialsEmail) | **Patch** /users/{user_id}/credentials_email | Update Email/Password Credential
[**User**](UserApi.md#User) | **Get** /users/{user_id} | Get User by Id
[**UserAttributeUserValues**](UserApi.md#UserAttributeUserValues) | **Get** /users/{user_id}/attribute_values | Get User Attribute Values
[**UserCredentialsApi3**](UserApi.md#UserCredentialsApi3) | **Get** /users/{user_id}/credentials_api3/{credentials_api3_id} | Get API 3 Credential
[**UserCredentialsEmail**](UserApi.md#UserCredentialsEmail) | **Get** /users/{user_id}/credentials_email | Get Email/Password Credential
[**UserCredentialsEmbed**](UserApi.md#UserCredentialsEmbed) | **Get** /users/{user_id}/credentials_embed/{credentials_embed_id} | Get Embedding Credential
[**UserCredentialsGoogle**](UserApi.md#UserCredentialsGoogle) | **Get** /users/{user_id}/credentials_google | Get Google Auth Credential
[**UserCredentialsLdap**](UserApi.md#UserCredentialsLdap) | **Get** /users/{user_id}/credentials_ldap | Get LDAP Credential
[**UserCredentialsLookerOpenid**](UserApi.md#UserCredentialsLookerOpenid) | **Get** /users/{user_id}/credentials_looker_openid | Get Looker OpenId Credential
[**UserCredentialsOidc**](UserApi.md#UserCredentialsOidc) | **Get** /users/{user_id}/credentials_oidc | Get OIDC Auth Credential
[**UserCredentialsSaml**](UserApi.md#UserCredentialsSaml) | **Get** /users/{user_id}/credentials_saml | Get Saml Auth Credential
[**UserCredentialsTotp**](UserApi.md#UserCredentialsTotp) | **Get** /users/{user_id}/credentials_totp | Get Two-Factor Credential
[**UserForCredential**](UserApi.md#UserForCredential) | **Get** /users/credential/{credential_type}/{credential_id} | Get User by Credential Id
[**UserRoles**](UserApi.md#UserRoles) | **Get** /users/{user_id}/roles | Get User Roles
[**UserSession**](UserApi.md#UserSession) | **Get** /users/{user_id}/sessions/{session_id} | Get Web Login Session



## AllUserCredentialsApi3s

> []CredentialsApi3 AllUserCredentialsApi3s(ctx, userId, optional)

Get All API 3 Credentials

### API 3 login information for the specified user. This is for the newer API keys that can be added for any user.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int64**| id of user | 
 **optional** | ***AllUserCredentialsApi3sOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AllUserCredentialsApi3sOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Requested fields. | 

### Return type

[**[]CredentialsApi3**](CredentialsApi3.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AllUserCredentialsEmbeds

> []CredentialsEmbed AllUserCredentialsEmbeds(ctx, userId, optional)

Get All Embedding Credentials

### Embed login information for the specified user.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int64**| id of user | 
 **optional** | ***AllUserCredentialsEmbedsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AllUserCredentialsEmbedsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Requested fields. | 

### Return type

[**[]CredentialsEmbed**](CredentialsEmbed.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AllUserSessions

> []Session AllUserSessions(ctx, userId, optional)

Get All Web Login Sessions

### Web login session for the specified user.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int64**| id of user | 
 **optional** | ***AllUserSessionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AllUserSessionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Requested fields. | 

### Return type

[**[]Session**](Session.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AllUsers

> []User AllUsers(ctx, optional)

Get All Users

### Get information about all users. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AllUsersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AllUsersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **optional.String**| Requested fields. | 
 **page** | **optional.Int64**| Requested page. | 
 **perPage** | **optional.Int64**| Results per page. | 
 **sorts** | **optional.String**| Fields to sort by. | 
 **ids** | [**optional.Interface of []int64**](int64.md)| Optional list of ids to get specific users. | 

### Return type

[**[]User**](User.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUser

> User CreateUser(ctx, optional)

Create User

### Create a user with the specified information. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateUserOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateUserOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **optional.String**| Requested fields. | 
 **body** | [**optional.Interface of User**](User.md)| User | 

### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUserCredentialsApi3

> CredentialsApi3 CreateUserCredentialsApi3(ctx, userId, optional)

Create API 3 Credential

### API 3 login information for the specified user. This is for the newer API keys that can be added for any user.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int64**| id of user | 
 **optional** | ***CreateUserCredentialsApi3Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateUserCredentialsApi3Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Requested fields. | 
 **body** | [**optional.Interface of CredentialsApi3**](CredentialsApi3.md)| API 3 Credential | 

### Return type

[**CredentialsApi3**](CredentialsApi3.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUserCredentialsEmail

> CredentialsEmail CreateUserCredentialsEmail(ctx, userId, body, optional)

Create Email/Password Credential

### Email/password login information for the specified user.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int64**| id of user | 
**body** | [**CredentialsEmail**](CredentialsEmail.md)| Email/Password Credential | 
 **optional** | ***CreateUserCredentialsEmailOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateUserCredentialsEmailOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **optional.String**| Requested fields. | 

### Return type

[**CredentialsEmail**](CredentialsEmail.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUserCredentialsEmailPasswordReset

> CredentialsEmail CreateUserCredentialsEmailPasswordReset(ctx, userId, optional)

Create Password Reset Token

### Create a password reset token. This will create a cryptographically secure random password reset token for the user. If the user already has a password reset token then this invalidates the old token and creates a new one. The token is expressed as the 'password_reset_url' of the user's email/password credential object. This takes an optional 'expires' param to indicate if the new token should be an expiring token. Tokens that expire are typically used for self-service password resets for existing users. Invitation emails for new users typically are not set to expire. The expire period is always 60 minutes when expires is enabled. This method can be called with an empty body. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int64**| Id of user | 
 **optional** | ***CreateUserCredentialsEmailPasswordResetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateUserCredentialsEmailPasswordResetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expires** | **optional.Bool**| Expiring token. | 
 **fields** | **optional.String**| Requested fields. | 

### Return type

[**CredentialsEmail**](CredentialsEmail.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUserCredentialsTotp

> CredentialsTotp CreateUserCredentialsTotp(ctx, userId, optional)

Create Two-Factor Credential

### Two-factor login information for the specified user.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int64**| id of user | 
 **optional** | ***CreateUserCredentialsTotpOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateUserCredentialsTotpOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Requested fields. | 
 **body** | [**optional.Interface of CredentialsTotp**](CredentialsTotp.md)| Two-Factor Credential | 

### Return type

[**CredentialsTotp**](CredentialsTotp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUser

> string DeleteUser(ctx, userId)

Delete User

### Delete the user with a specific id.  **DANGER** this will delete the user and all looks and other information owned by the user. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int64**| Id of user | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUserAttributeUserValue

> DeleteUserAttributeUserValue(ctx, userId, userAttributeId)

Delete User Attribute User Value

### Delete a user attribute value from a user's account settings.  After the user attribute value is deleted from the user's account settings, subsequent requests for the user attribute value for this user will draw from the user's groups or the default value of the user attribute. See [Get User Attribute Values](#!/User/user_attribute_user_values) for more information about how user attribute values are resolved. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int64**| Id of user | 
**userAttributeId** | **int64**| Id of user attribute | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUserCredentialsApi3

> string DeleteUserCredentialsApi3(ctx, userId, credentialsApi3Id)

Delete API 3 Credential

### API 3 login information for the specified user. This is for the newer API keys that can be added for any user.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int64**| id of user | 
**credentialsApi3Id** | **int64**| id of API 3 Credential | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUserCredentialsEmail

> string DeleteUserCredentialsEmail(ctx, userId)

Delete Email/Password Credential

### Email/password login information for the specified user.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int64**| id of user | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUserCredentialsEmbed

> string DeleteUserCredentialsEmbed(ctx, userId, credentialsEmbedId)

Delete Embedding Credential

### Embed login information for the specified user.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int64**| id of user | 
**credentialsEmbedId** | **int64**| id of Embedding Credential | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUserCredentialsGoogle

> string DeleteUserCredentialsGoogle(ctx, userId)

Delete Google Auth Credential

### Google authentication login information for the specified user.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int64**| id of user | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUserCredentialsLdap

> string DeleteUserCredentialsLdap(ctx, userId)

Delete LDAP Credential

### LDAP login information for the specified user.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int64**| id of user | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUserCredentialsLookerOpenid

> string DeleteUserCredentialsLookerOpenid(ctx, userId)

Delete Looker OpenId Credential

### Looker Openid login information for the specified user. Used by Looker Analysts.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int64**| id of user | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUserCredentialsOidc

> string DeleteUserCredentialsOidc(ctx, userId)

Delete OIDC Auth Credential

### OpenID Connect (OIDC) authentication login information for the specified user.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int64**| id of user | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUserCredentialsSaml

> string DeleteUserCredentialsSaml(ctx, userId)

Delete Saml Auth Credential

### Saml authentication login information for the specified user.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int64**| id of user | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUserCredentialsTotp

> string DeleteUserCredentialsTotp(ctx, userId)

Delete Two-Factor Credential

### Two-factor login information for the specified user.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int64**| id of user | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUserSession

> string DeleteUserSession(ctx, userId, sessionId)

Delete Web Login Session

### Web login session for the specified user.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int64**| id of user | 
**sessionId** | **int64**| id of Web Login Session | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Me

> User Me(ctx, optional)

Get Current User

### Get information about the current user; i.e. the user account currently calling the API. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MeOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MeOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **optional.String**| Requested fields. | 

### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchUsers

> []User SearchUsers(ctx, optional)

Search Users

### Search users  Returns all<sup>*</sup> user records that match the given search criteria.  If multiple search params are given and `filter_or` is FALSE or not specified, search params are combined in a logical AND operation. Only rows that match *all* search param criteria will be returned.  If `filter_or` is TRUE, multiple search params are combined in a logical OR operation. Results will include rows that match **any** of the search criteria.  String search params use case-insensitive matching. String search params can contain `%` and '_' as SQL LIKE pattern match wildcard expressions. example=\"dan%\" will match \"danger\" and \"Danzig\" but not \"David\" example=\"D_m%\" will match \"Damage\" and \"dump\"  Integer search params can accept a single value or a comma separated list of values. The multiple values will be combined under a logical OR operation - results will match at least one of the given values.  Most search params can accept \"IS NULL\" and \"NOT NULL\" as special expressions to match or exclude (respectively) rows where the column is null.  Boolean search params accept only \"true\" and \"false\" as values.   (<sup>*</sup>) Results are always filtered to the level of information the caller is permitted to view. Looker admins can see all user details; normal users in an open system can see names of other users but no details; normal users in a closed system can only see names of other users who are members of the same group as the user.  

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SearchUsersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SearchUsersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **optional.String**| Include only these fields in the response | 
 **page** | **optional.Int64**| Return only page N of paginated results | 
 **perPage** | **optional.Int64**| Return N rows of data per page | 
 **sorts** | **optional.String**| Fields to sort by. | 
 **id** | **optional.Int64**| Match User Id. | 
 **firstName** | **optional.String**| Match First name. | 
 **lastName** | **optional.String**| Match Last name. | 
 **verifiedLookerEmployee** | **optional.Bool**| Search for user accounts associated with Looker employees | 
 **email** | **optional.String**| Search for the user with this email address | 
 **isDisabled** | **optional.Bool**| Search for disabled user accounts | 
 **filterOr** | **optional.Bool**| Combine given search criteria in a boolean OR expression | 
 **contentMetadataId** | **optional.Int64**| Search for users who have access to this content_metadata item | 
 **groupId** | **optional.Int64**| Search for users who are direct members of this group | 

### Return type

[**[]User**](User.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchUsersNames

> []User SearchUsersNames(ctx, pattern, optional)

Search User Names

### Search for user accounts by name  Returns all user accounts where `first_name` OR `last_name` OR `email` field values match a pattern. The pattern can contain `%` and `_` wildcards as in SQL LIKE expressions.  Any additional search params will be combined into a logical AND expression. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pattern** | **string**| Pattern to match | 
 **optional** | ***SearchUsersNamesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SearchUsersNamesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Include only these fields in the response | 
 **page** | **optional.Int64**| Return only page N of paginated results | 
 **perPage** | **optional.Int64**| Return N rows of data per page | 
 **sorts** | **optional.String**| Fields to sort by | 
 **id** | **optional.Int64**| Match User Id | 
 **firstName** | **optional.String**| Match First name | 
 **lastName** | **optional.String**| Match Last name | 
 **verifiedLookerEmployee** | **optional.Bool**| Match Verified Looker employee | 
 **email** | **optional.String**| Match Email Address | 
 **isDisabled** | **optional.Bool**| Include or exclude disabled accounts in the results | 

### Return type

[**[]User**](User.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetUserAttributeUserValue

> UserAttributeWithValue SetUserAttributeUserValue(ctx, userId, userAttributeId, body)

Set User Attribute User Value

### Store a custom value for a user attribute in a user's account settings.  Per-user user attribute values take precedence over group or default values. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int64**| Id of user | 
**userAttributeId** | **int64**| Id of user attribute | 
**body** | [**UserAttributeWithValue**](UserAttributeWithValue.md)| New attribute value for user. | 

### Return type

[**UserAttributeWithValue**](UserAttributeWithValue.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetUserRoles

> []Role SetUserRoles(ctx, userId, body, optional)

Set User Roles

### Set roles of the user with a specific id. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int64**| id of user | 
**body** | [**[]int64**](int64.md)| array of roles ids for user | 
 **optional** | ***SetUserRolesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SetUserRolesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **optional.String**| Requested fields. | 

### Return type

[**[]Role**](Role.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUser

> User UpdateUser(ctx, userId, body, optional)

Update User

### Update information about the user with a specific id. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int64**| Id of user | 
**body** | [**User**](User.md)| User | 
 **optional** | ***UpdateUserOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateUserOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **optional.String**| Requested fields. | 

### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserCredentialsEmail

> CredentialsEmail UpdateUserCredentialsEmail(ctx, userId, body, optional)

Update Email/Password Credential

### Email/password login information for the specified user.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int64**| id of user | 
**body** | [**CredentialsEmail**](CredentialsEmail.md)| Email/Password Credential | 
 **optional** | ***UpdateUserCredentialsEmailOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateUserCredentialsEmailOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **optional.String**| Requested fields. | 

### Return type

[**CredentialsEmail**](CredentialsEmail.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## User

> User User(ctx, userId, optional)

Get User by Id

### Get information about the user with a specific id.  If the caller is an admin or the caller is the user being specified, then full user information will be returned. Otherwise, a minimal 'public' variant of the user information will be returned. This contains The user name and avatar url, but no sensitive information. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int64**| Id of user | 
 **optional** | ***UserOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UserOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Requested fields. | 

### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserAttributeUserValues

> []UserAttributeWithValue UserAttributeUserValues(ctx, userId, optional)

Get User Attribute Values

### Get user attribute values for a given user.  Returns the values of specified user attributes (or all user attributes) for a certain user.  A value for each user attribute is searched for in the following locations, in this order: 1. in the user's account information 1. in groups that the user is a member of 1. the default value of the user attribute  If more than one group has a value defined for a user attribute, the group with the lowest rank wins.  The response will only include user attributes for which values were found. Use `include_unset=true` to include empty records for user attributes with no value.  The value of all hidden user attributes will be blank. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int64**| Id of user | 
 **optional** | ***UserAttributeUserValuesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UserAttributeUserValuesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Requested fields. | 
 **userAttributeIds** | [**optional.Interface of []int64**](int64.md)| Specific user attributes to request. Omit or leave blank to request all user attributes. | 
 **allValues** | **optional.Bool**| If true, returns all values in the search path instead of just the first value found. Useful for debugging group precedence. | 
 **includeUnset** | **optional.Bool**| If true, returns an empty record for each requested attribute that has no user, group, or default value. | 

### Return type

[**[]UserAttributeWithValue**](UserAttributeWithValue.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserCredentialsApi3

> CredentialsApi3 UserCredentialsApi3(ctx, userId, credentialsApi3Id, optional)

Get API 3 Credential

### API 3 login information for the specified user. This is for the newer API keys that can be added for any user.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int64**| Id of user | 
**credentialsApi3Id** | **int64**| Id of API 3 Credential | 
 **optional** | ***UserCredentialsApi3Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UserCredentialsApi3Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **optional.String**| Requested fields. | 

### Return type

[**CredentialsApi3**](CredentialsApi3.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserCredentialsEmail

> CredentialsEmail UserCredentialsEmail(ctx, userId, optional)

Get Email/Password Credential

### Email/password login information for the specified user.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int64**| id of user | 
 **optional** | ***UserCredentialsEmailOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UserCredentialsEmailOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Requested fields. | 

### Return type

[**CredentialsEmail**](CredentialsEmail.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserCredentialsEmbed

> CredentialsEmbed UserCredentialsEmbed(ctx, userId, credentialsEmbedId, optional)

Get Embedding Credential

### Embed login information for the specified user.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int64**| Id of user | 
**credentialsEmbedId** | **int64**| Id of Embedding Credential | 
 **optional** | ***UserCredentialsEmbedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UserCredentialsEmbedOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **optional.String**| Requested fields. | 

### Return type

[**CredentialsEmbed**](CredentialsEmbed.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserCredentialsGoogle

> CredentialsGoogle UserCredentialsGoogle(ctx, userId, optional)

Get Google Auth Credential

### Google authentication login information for the specified user.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int64**| id of user | 
 **optional** | ***UserCredentialsGoogleOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UserCredentialsGoogleOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Requested fields. | 

### Return type

[**CredentialsGoogle**](CredentialsGoogle.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserCredentialsLdap

> CredentialsLdap UserCredentialsLdap(ctx, userId, optional)

Get LDAP Credential

### LDAP login information for the specified user.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int64**| id of user | 
 **optional** | ***UserCredentialsLdapOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UserCredentialsLdapOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Requested fields. | 

### Return type

[**CredentialsLdap**](CredentialsLDAP.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserCredentialsLookerOpenid

> CredentialsLookerOpenid UserCredentialsLookerOpenid(ctx, userId, optional)

Get Looker OpenId Credential

### Looker Openid login information for the specified user. Used by Looker Analysts.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int64**| id of user | 
 **optional** | ***UserCredentialsLookerOpenidOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UserCredentialsLookerOpenidOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Requested fields. | 

### Return type

[**CredentialsLookerOpenid**](CredentialsLookerOpenid.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserCredentialsOidc

> CredentialsOidc UserCredentialsOidc(ctx, userId, optional)

Get OIDC Auth Credential

### OpenID Connect (OIDC) authentication login information for the specified user.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int64**| id of user | 
 **optional** | ***UserCredentialsOidcOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UserCredentialsOidcOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Requested fields. | 

### Return type

[**CredentialsOidc**](CredentialsOIDC.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserCredentialsSaml

> CredentialsSaml UserCredentialsSaml(ctx, userId, optional)

Get Saml Auth Credential

### Saml authentication login information for the specified user.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int64**| id of user | 
 **optional** | ***UserCredentialsSamlOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UserCredentialsSamlOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Requested fields. | 

### Return type

[**CredentialsSaml**](CredentialsSaml.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserCredentialsTotp

> CredentialsTotp UserCredentialsTotp(ctx, userId, optional)

Get Two-Factor Credential

### Two-factor login information for the specified user.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int64**| id of user | 
 **optional** | ***UserCredentialsTotpOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UserCredentialsTotpOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Requested fields. | 

### Return type

[**CredentialsTotp**](CredentialsTotp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserForCredential

> User UserForCredential(ctx, credentialType, credentialId, optional)

Get User by Credential Id

### Get information about the user with a credential of given type with specific id.  This is used to do things like find users by their embed external_user_id. Or, find the user with a given api3 client_id, etc. The 'credential_type' matchs the 'type' name of the various credential types. It must be one of the values listed in the table below. The 'credential_id' is your unique Id for the user and is specific to each type of credential.  An example using the Ruby sdk might look like:  `sdk.user_for_credential('embed', 'customer-4959425')`  This table shows the supported 'Credential Type' strings. The right column is for reference; it shows which field in the given credential type is actually searched when finding a user with the supplied 'credential_id'.  | Credential Types | Id Field Matched | | ---------------- | ---------------- | | email            | email            | | google           | google_user_id   | | saml             | saml_user_id     | | oidc             | oidc_user_id     | | ldap             | ldap_id          | | api              | token            | | api3             | client_id        | | embed            | external_user_id | | looker_openid    | email            |  NOTE: The 'api' credential type was only used with the legacy Looker query API and is no longer supported. The credential type for API you are currently looking at is 'api3'.  

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**credentialType** | **string**| Type name of credential | 
**credentialId** | **string**| Id of credential | 
 **optional** | ***UserForCredentialOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UserForCredentialOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **optional.String**| Requested fields. | 

### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserRoles

> []Role UserRoles(ctx, userId, optional)

Get User Roles

### Get information about roles of a given user 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int64**| id of user | 
 **optional** | ***UserRolesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UserRolesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Requested fields. | 
 **directAssociationOnly** | **optional.Bool**| Get only roles associated directly with the user: exclude those only associated through groups. | 

### Return type

[**[]Role**](Role.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserSession

> Session UserSession(ctx, userId, sessionId, optional)

Get Web Login Session

### Web login session for the specified user.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int64**| Id of user | 
**sessionId** | **int64**| Id of Web Login Session | 
 **optional** | ***UserSessionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UserSessionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **optional.String**| Requested fields. | 

### Return type

[**Session**](Session.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

