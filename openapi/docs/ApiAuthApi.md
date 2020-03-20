# \ApiAuthApi

All URIs are relative to *https://mattermost.looker.com:19999/api/3.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Login**](ApiAuthApi.md#Login) | **Post** /login | Login
[**LoginUser**](ApiAuthApi.md#LoginUser) | **Post** /login/{user_id} | Login user
[**Logout**](ApiAuthApi.md#Logout) | **Delete** /logout | Logout



## Login

> AccessToken Login(ctx, optional)

Login

### Present client credentials to obtain an authorization token  Looker API implements the OAuth2 [Resource Owner Password Credentials Grant](https://looker.com/docs/r/api/outh2_resource_owner_pc) pattern. The client credentials required for this login must be obtained by creating an API3 key on a user account in the Looker Admin console. The API3 key consists of a public `client_id` and a private `client_secret`.  The access token returned by `login` must be used in the HTTP Authorization header of subsequent API requests, like this: ``` Authorization: token 4QDkCyCtZzYgj4C2p2cj3csJH7zqS5RzKs2kTnG4 ``` Replace \"4QDkCy...\" with the `access_token` value returned by `login`. The word `token` is a string literal and must be included exactly as shown.  This function can accept `client_id` and `client_secret` parameters as URL query params or as www-form-urlencoded params in the body of the HTTP request. Since there is a small risk that URL parameters may be visible to intermediate nodes on the network route (proxies, routers, etc), passing credentials in the body of the request is considered more secure than URL params.  Example of passing credentials in the HTTP request body: ```` POST HTTP /login Content-Type: application/x-www-form-urlencoded  client_id=CGc9B7v7J48dQSJvxxx&client_secret=nNVS9cSS3xNpSC9JdsBvvvvv ````  ### Best Practice: Always pass credentials in body params. Pass credentials in URL query params **only** when you cannot pass body params due to application, tool, or other limitations.  For more information and detailed examples of Looker API authorization, see [How to Authenticate to Looker API3](https://github.com/looker/looker-sdk-ruby/blob/master/authentication.md). 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LoginOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a LoginOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **optional.String**| client_id part of API3 Key. | 
 **clientSecret** | **optional.String**| client_secret part of API3 Key. | 

### Return type

[**AccessToken**](AccessToken.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LoginUser

> AccessToken LoginUser(ctx, userId)

Login user

### Create an access token for a given user.  This can only be called by an authenticated admin user. It allows that admin to generate a new authentication token for the user with the given user id. That token can then be used for subsequent API calls - which are then performed *as* that target user.  The target user does *not* need to have a pre-existing API client_id/client_secret pair. And, no such credentials are created by this call.  This allows for building systems where api user authentication for an arbitrary number of users is done outside of Looker and funneled through a single 'service account' with admin permissions. Note that a new access token is generated on each call. If target users are going to be making numerous API calls in a short period then it is wise to cache this authentication token rather than call this before each of those API calls.  See 'login' for more detail on the access token and how to use it. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int64**| Id of user. | 

### Return type

[**AccessToken**](AccessToken.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Logout

> string Logout(ctx, )

Logout

### Logout of the API and invalidate the current access token. 

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

