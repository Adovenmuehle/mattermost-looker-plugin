# \AuthApi

All URIs are relative to *https://mattermost.looker.com:19999/api/3.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AllUserLoginLockouts**](AuthApi.md#AllUserLoginLockouts) | **Get** /user_login_lockouts | Get All User Login Lockouts
[**CreateOidcTestConfig**](AuthApi.md#CreateOidcTestConfig) | **Post** /oidc_test_configs | Create OIDC Test Configuration
[**CreateSamlTestConfig**](AuthApi.md#CreateSamlTestConfig) | **Post** /saml_test_configs | Create SAML Test Configuration
[**CreateSsoEmbedUrl**](AuthApi.md#CreateSsoEmbedUrl) | **Post** /embed/sso_url | Create SSO Embed Url
[**DeleteOidcTestConfig**](AuthApi.md#DeleteOidcTestConfig) | **Delete** /oidc_test_configs/{test_slug} | Delete OIDC Test Configuration
[**DeleteSamlTestConfig**](AuthApi.md#DeleteSamlTestConfig) | **Delete** /saml_test_configs/{test_slug} | Delete SAML Test Configuration
[**DeleteUserLoginLockout**](AuthApi.md#DeleteUserLoginLockout) | **Delete** /user_login_lockout/{key} | Delete User Login Lockout
[**FetchAndParseSamlIdpMetadata**](AuthApi.md#FetchAndParseSamlIdpMetadata) | **Post** /fetch_and_parse_saml_idp_metadata | Parse SAML IdP Url
[**ForcePasswordResetAtNextLoginForAllUsers**](AuthApi.md#ForcePasswordResetAtNextLoginForAllUsers) | **Put** /password_config/force_password_reset_at_next_login_for_all_users | Force password reset
[**LdapConfig**](AuthApi.md#LdapConfig) | **Get** /ldap_config | Get LDAP Configuration
[**OidcConfig**](AuthApi.md#OidcConfig) | **Get** /oidc_config | Get OIDC Configuration
[**OidcTestConfig**](AuthApi.md#OidcTestConfig) | **Get** /oidc_test_configs/{test_slug} | Get OIDC Test Configuration
[**ParseSamlIdpMetadata**](AuthApi.md#ParseSamlIdpMetadata) | **Post** /parse_saml_idp_metadata | Parse SAML IdP XML
[**PasswordConfig**](AuthApi.md#PasswordConfig) | **Get** /password_config | Get Password Config
[**SamlConfig**](AuthApi.md#SamlConfig) | **Get** /saml_config | Get SAML Configuration
[**SamlTestConfig**](AuthApi.md#SamlTestConfig) | **Get** /saml_test_configs/{test_slug} | Get SAML Test Configuration
[**SearchUserLoginLockouts**](AuthApi.md#SearchUserLoginLockouts) | **Get** /user_login_lockouts/search | Search User Login Lockouts
[**SessionConfig**](AuthApi.md#SessionConfig) | **Get** /session_config | Get Session Config
[**TestLdapConfigAuth**](AuthApi.md#TestLdapConfigAuth) | **Put** /ldap_config/test_auth | Test LDAP Auth
[**TestLdapConfigConnection**](AuthApi.md#TestLdapConfigConnection) | **Put** /ldap_config/test_connection | Test LDAP Connection
[**TestLdapConfigUserAuth**](AuthApi.md#TestLdapConfigUserAuth) | **Put** /ldap_config/test_user_auth | Test LDAP User Auth
[**TestLdapConfigUserInfo**](AuthApi.md#TestLdapConfigUserInfo) | **Put** /ldap_config/test_user_info | Test LDAP User Info
[**UpdateLdapConfig**](AuthApi.md#UpdateLdapConfig) | **Patch** /ldap_config | Update LDAP Configuration
[**UpdateOidcConfig**](AuthApi.md#UpdateOidcConfig) | **Patch** /oidc_config | Update OIDC Configuration
[**UpdatePasswordConfig**](AuthApi.md#UpdatePasswordConfig) | **Patch** /password_config | Update Password Config
[**UpdateSamlConfig**](AuthApi.md#UpdateSamlConfig) | **Patch** /saml_config | Update SAML Configuration
[**UpdateSessionConfig**](AuthApi.md#UpdateSessionConfig) | **Patch** /session_config | Update Session Config



## AllUserLoginLockouts

> []UserLoginLockout AllUserLoginLockouts(ctx, optional)

Get All User Login Lockouts

### Get currently locked-out users. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AllUserLoginLockoutsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AllUserLoginLockoutsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **optional.String**| Include only these fields in the response | 

### Return type

[**[]UserLoginLockout**](UserLoginLockout.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOidcTestConfig

> OidcConfig CreateOidcTestConfig(ctx, body)

Create OIDC Test Configuration

### Create a OIDC test configuration. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**OidcConfig**](OidcConfig.md)| OIDC test config | 

### Return type

[**OidcConfig**](OIDCConfig.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSamlTestConfig

> SamlConfig CreateSamlTestConfig(ctx, body)

Create SAML Test Configuration

### Create a SAML test configuration. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**SamlConfig**](SamlConfig.md)| SAML test config | 

### Return type

[**SamlConfig**](SamlConfig.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSsoEmbedUrl

> EmbedSsoUrl CreateSsoEmbedUrl(ctx, body)

Create SSO Embed Url

### Create SSO Embed URL  Creates an SSO embed URL and cryptographically signs it with an embed secret. This signed URL can then be used to instantiate a Looker embed session in a PBL web application. Do not make any modifications to this URL - any change may invalidate the signature and cause the URL to fail to load a Looker embed session.  A signed SSO embed URL can only be used once. After it has been used to request a page from the Looker server, the URL is invalid. Future requests using the same URL will fail. This is to prevent 'replay attacks'.  The `target_url` property must be a complete URL of a Looker UI page - scheme, hostname, path and query params. To load a dashboard with id 56 and with a filter of `Date=1 years`, the looker URL would look like `https:/myname.looker.com/dashboards/56?Date=1%20years`. The best way to obtain this target_url is to navigate to the desired Looker page in your web browser, copy the URL shown in the browser address bar and paste it into the `target_url` property as a quoted string value in this API request.  Permissions for the embed user are defined by the groups in which the embed user is a member (group_ids property) and the lists of models and permissions assigned to the embed user. At a minimum, you must provide values for either the group_ids property, or both the models and permissions properties. These properties are additive; an embed user can be a member of certain groups AND be granted access to models and permissions.  The embed user's access is the union of permissions granted by the group_ids, models, and permissions properties.  This function does not strictly require all group_ids, user attribute names, or model names to exist at the moment the SSO embed url is created. Unknown group_id, user attribute names or model names will be passed through to the output URL. To diagnose potential problems with an SSO embed URL, you can copy the signed URL into the Embed URI Validator text box in `<your looker instance>/admin/embed`.  The `secret_id` parameter is optional. If specified, its value must be the id of an active secret defined in the Looker instance. if not specified, the URL will be signed using the newest active secret defined in the Looker instance.  #### Security Note Protect this signed URL as you would an access token or password credentials - do not write it to disk, do not pass it to a third party, and only pass it through a secure HTTPS encrypted transport. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**EmbedSsoUrlParams**](EmbedSsoUrlParams.md)| SSO parameters | 

### Return type

[**EmbedSsoUrl**](EmbedSsoUrl.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOidcTestConfig

> string DeleteOidcTestConfig(ctx, testSlug)

Delete OIDC Test Configuration

### Delete a OIDC test configuration. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**testSlug** | **string**| Slug of test config | 

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


## DeleteSamlTestConfig

> string DeleteSamlTestConfig(ctx, testSlug)

Delete SAML Test Configuration

### Delete a SAML test configuration. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**testSlug** | **string**| Slug of test config | 

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


## DeleteUserLoginLockout

> string DeleteUserLoginLockout(ctx, key)

Delete User Login Lockout

### Removes login lockout for the associated user. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string**| The key associated with the locked user | 

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


## FetchAndParseSamlIdpMetadata

> SamlMetadataParseResult FetchAndParseSamlIdpMetadata(ctx, body)

Parse SAML IdP Url

### Fetch the given url and parse it as a SAML IdP metadata document and return the result. Note that this requires that the url be public or at least at a location where the Looker instance can fetch it without requiring any special authentication. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | **string**| SAML IdP metadata public url | 

### Return type

[**SamlMetadataParseResult**](SamlMetadataParseResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: text/plain
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ForcePasswordResetAtNextLoginForAllUsers

> string ForcePasswordResetAtNextLoginForAllUsers(ctx, )

Force password reset

### Force all credentials_email users to reset their login passwords upon their next login. 

### Required Parameters

This endpoint does not need any parameter.

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


## LdapConfig

> LdapConfig LdapConfig(ctx, )

Get LDAP Configuration

### Get the LDAP configuration.  Looker can be optionally configured to authenticate users against an Active Directory or other LDAP directory server. LDAP setup requires coordination with an administrator of that directory server.  Only Looker administrators can read and update the LDAP configuration.  Configuring LDAP impacts authentication for all users. This configuration should be done carefully.  Looker maintains a single LDAP configuration. It can be read and updated.       Updates only succeed if the new state will be valid (in the sense that all required fields are populated);       it is up to you to ensure that the configuration is appropriate and correct).  LDAP is enabled or disabled for Looker using the **enabled** field.  Looker will never return an **auth_password** field. That value can be set, but never retrieved.  See the [Looker LDAP docs](https://www.looker.com/docs/r/api/ldap_setup) for additional information. 

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**LdapConfig**](LDAPConfig.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OidcConfig

> OidcConfig OidcConfig(ctx, )

Get OIDC Configuration

### Get the OIDC configuration.  Looker can be optionally configured to authenticate users against an OpenID Connect (OIDC) authentication server. OIDC setup requires coordination with an administrator of that server.  Only Looker administrators can read and update the OIDC configuration.  Configuring OIDC impacts authentication for all users. This configuration should be done carefully.  Looker maintains a single OIDC configuation. It can be read and updated.       Updates only succeed if the new state will be valid (in the sense that all required fields are populated);       it is up to you to ensure that the configuration is appropriate and correct).  OIDC is enabled or disabled for Looker using the **enabled** field. 

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**OidcConfig**](OIDCConfig.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OidcTestConfig

> OidcConfig OidcTestConfig(ctx, testSlug)

Get OIDC Test Configuration

### Get a OIDC test configuration by test_slug. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**testSlug** | **string**| Slug of test config | 

### Return type

[**OidcConfig**](OIDCConfig.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ParseSamlIdpMetadata

> SamlMetadataParseResult ParseSamlIdpMetadata(ctx, body)

Parse SAML IdP XML

### Parse the given xml as a SAML IdP metadata document and return the result. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | **string**| SAML IdP metadata xml | 

### Return type

[**SamlMetadataParseResult**](SamlMetadataParseResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: text/plain
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PasswordConfig

> PasswordConfig PasswordConfig(ctx, )

Get Password Config

### Get password config. 

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**PasswordConfig**](PasswordConfig.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SamlConfig

> SamlConfig SamlConfig(ctx, )

Get SAML Configuration

### Get the SAML configuration.  Looker can be optionally configured to authenticate users against a SAML authentication server. SAML setup requires coordination with an administrator of that server.  Only Looker administrators can read and update the SAML configuration.  Configuring SAML impacts authentication for all users. This configuration should be done carefully.  Looker maintains a single SAML configuation. It can be read and updated.       Updates only succeed if the new state will be valid (in the sense that all required fields are populated);       it is up to you to ensure that the configuration is appropriate and correct).  SAML is enabled or disabled for Looker using the **enabled** field. 

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**SamlConfig**](SamlConfig.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SamlTestConfig

> SamlConfig SamlTestConfig(ctx, testSlug)

Get SAML Test Configuration

### Get a SAML test configuration by test_slug. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**testSlug** | **string**| Slug of test config | 

### Return type

[**SamlConfig**](SamlConfig.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchUserLoginLockouts

> []UserLoginLockout SearchUserLoginLockouts(ctx, optional)

Search User Login Lockouts

### Search currently locked-out users. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SearchUserLoginLockoutsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SearchUserLoginLockoutsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **optional.String**| Include only these fields in the response | 
 **page** | **optional.Int64**| Return only page N of paginated results | 
 **perPage** | **optional.Int64**| Return N rows of data per page | 
 **sorts** | **optional.String**| Fields to sort by. | 
 **authType** | **optional.String**| Auth type user is locked out for (email, ldap, totp, api) | 
 **fullName** | **optional.String**| Match name | 
 **email** | **optional.String**| Match email | 
 **remoteId** | **optional.String**| Match remote LDAP ID | 
 **filterOr** | **optional.Bool**| Combine given search criteria in a boolean OR expression | 

### Return type

[**[]UserLoginLockout**](UserLoginLockout.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SessionConfig

> SessionConfig SessionConfig(ctx, )

Get Session Config

### Get session config. 

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**SessionConfig**](SessionConfig.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestLdapConfigAuth

> LdapConfigTestResult TestLdapConfigAuth(ctx, body)

Test LDAP Auth

### Test the connection authentication settings for an LDAP configuration.  This tests that the connection is possible and that a 'server' account to be used by Looker can       authenticate to the LDAP server given connection and authentication information.  **connection_host**, **connection_port**, and **auth_username**, are required.       **connection_tls** and **auth_password** are optional.  Example: ```json {   \"connection_host\": \"ldap.example.com\",   \"connection_port\": \"636\",   \"connection_tls\": true,   \"auth_username\": \"cn=looker,dc=example,dc=com\",   \"auth_password\": \"secret\" } ```  Looker will never return an **auth_password**. If this request omits the **auth_password** field, then       the **auth_password** value from the active config (if present) will be used for the test.  The active LDAP settings are not modified.  

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**LdapConfig**](LdapConfig.md)| LDAP Config | 

### Return type

[**LdapConfigTestResult**](LDAPConfigTestResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestLdapConfigConnection

> LdapConfigTestResult TestLdapConfigConnection(ctx, body)

Test LDAP Connection

### Test the connection settings for an LDAP configuration.  This tests that the connection is possible given a connection_host and connection_port.  **connection_host** and **connection_port** are required. **connection_tls** is optional.  Example: ```json {   \"connection_host\": \"ldap.example.com\",   \"connection_port\": \"636\",   \"connection_tls\": true } ```  No authentication to the LDAP server is attempted.  The active LDAP settings are not modified. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**LdapConfig**](LdapConfig.md)| LDAP Config | 

### Return type

[**LdapConfigTestResult**](LDAPConfigTestResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestLdapConfigUserAuth

> LdapConfigTestResult TestLdapConfigUserAuth(ctx, body)

Test LDAP User Auth

### Test the user authentication settings for an LDAP configuration.  This test accepts a full LDAP configuration along with a username/password pair and attempts to       authenticate the user with the LDAP server. The configuration is validated before attempting the       authentication.  Looker will never return an **auth_password**. If this request omits the **auth_password** field, then       the **auth_password** value from the active config (if present) will be used for the test.  **test_ldap_user** and **test_ldap_password** are required.  The active LDAP settings are not modified.  

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**LdapConfig**](LdapConfig.md)| LDAP Config | 

### Return type

[**LdapConfigTestResult**](LDAPConfigTestResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestLdapConfigUserInfo

> LdapConfigTestResult TestLdapConfigUserInfo(ctx, body)

Test LDAP User Info

### Test the user authentication settings for an LDAP configuration without authenticating the user.  This test will let you easily test the mapping for user properties and roles for any user without      needing to authenticate as that user.  This test accepts a full LDAP configuration along with a username and attempts to find the full info      for the user from the LDAP server without actually authenticating the user. So, user password is not      required.The configuration is validated before attempting to contact the server.  **test_ldap_user** is required.  The active LDAP settings are not modified.  

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**LdapConfig**](LdapConfig.md)| LDAP Config | 

### Return type

[**LdapConfigTestResult**](LDAPConfigTestResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLdapConfig

> LdapConfig UpdateLdapConfig(ctx, body)

Update LDAP Configuration

### Update the LDAP configuration.  Configuring LDAP impacts authentication for all users. This configuration should be done carefully.  Only Looker administrators can read and update the LDAP configuration.  LDAP is enabled or disabled for Looker using the **enabled** field.  It is **highly** recommended that any LDAP setting changes be tested using the APIs below before being set globally.  See the [Looker LDAP docs](https://www.looker.com/docs/r/api/ldap_setup) for additional information. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**LdapConfig**](LdapConfig.md)| LDAP Config | 

### Return type

[**LdapConfig**](LDAPConfig.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOidcConfig

> OidcConfig UpdateOidcConfig(ctx, body)

Update OIDC Configuration

### Update the OIDC configuration.  Configuring OIDC impacts authentication for all users. This configuration should be done carefully.  Only Looker administrators can read and update the OIDC configuration.  OIDC is enabled or disabled for Looker using the **enabled** field.  It is **highly** recommended that any OIDC setting changes be tested using the APIs below before being set globally. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**OidcConfig**](OidcConfig.md)| OIDC Config | 

### Return type

[**OidcConfig**](OIDCConfig.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePasswordConfig

> PasswordConfig UpdatePasswordConfig(ctx, body)

Update Password Config

### Update password config. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**PasswordConfig**](PasswordConfig.md)| Password Config | 

### Return type

[**PasswordConfig**](PasswordConfig.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSamlConfig

> SamlConfig UpdateSamlConfig(ctx, body)

Update SAML Configuration

### Update the SAML configuration.  Configuring SAML impacts authentication for all users. This configuration should be done carefully.  Only Looker administrators can read and update the SAML configuration.  SAML is enabled or disabled for Looker using the **enabled** field.  It is **highly** recommended that any SAML setting changes be tested using the APIs below before being set globally. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**SamlConfig**](SamlConfig.md)| SAML Config | 

### Return type

[**SamlConfig**](SamlConfig.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSessionConfig

> SessionConfig UpdateSessionConfig(ctx, body)

Update Session Config

### Update session config. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**SessionConfig**](SessionConfig.md)| Session Config | 

### Return type

[**SessionConfig**](SessionConfig.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

