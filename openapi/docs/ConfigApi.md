# \ConfigApi

All URIs are relative to *https://mattermost.looker.com:19999/api/3.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AllLegacyFeatures**](ConfigApi.md#AllLegacyFeatures) | **Get** /legacy_features | Get All Legacy Features
[**AllLocales**](ConfigApi.md#AllLocales) | **Get** /locales | Get All Locales
[**AllTimezones**](ConfigApi.md#AllTimezones) | **Get** /timezones | Get All Timezones
[**BackupConfiguration**](ConfigApi.md#BackupConfiguration) | **Get** /backup_configuration | Get Backup Configuration
[**CustomWelcomeEmail**](ConfigApi.md#CustomWelcomeEmail) | **Get** /custom_welcome_email | Get Custom Welcome Email
[**InternalHelpResources**](ConfigApi.md#InternalHelpResources) | **Get** /internal_help_resources_enabled | Get Internal Help Resources
[**InternalHelpResourcesContent**](ConfigApi.md#InternalHelpResourcesContent) | **Get** /internal_help_resources_content | Get Internal Help Resources Content
[**LegacyFeature**](ConfigApi.md#LegacyFeature) | **Get** /legacy_features/{legacy_feature_id} | Get Legacy Feature
[**UpdateBackupConfiguration**](ConfigApi.md#UpdateBackupConfiguration) | **Patch** /backup_configuration | Update Backup Configuration
[**UpdateCustomWelcomeEmail**](ConfigApi.md#UpdateCustomWelcomeEmail) | **Patch** /custom_welcome_email | Update Custom Welcome Email Content
[**UpdateCustomWelcomeEmailTest**](ConfigApi.md#UpdateCustomWelcomeEmailTest) | **Put** /custom_welcome_email_test | Send a test welcome email to the currently logged in user with the supplied content 
[**UpdateInternalHelpResources**](ConfigApi.md#UpdateInternalHelpResources) | **Patch** /internal_help_resources | Update internal help resources configuration
[**UpdateInternalHelpResourcesContent**](ConfigApi.md#UpdateInternalHelpResourcesContent) | **Patch** /internal_help_resources_content | Update internal help resources content
[**UpdateLegacyFeature**](ConfigApi.md#UpdateLegacyFeature) | **Patch** /legacy_features/{legacy_feature_id} | Update Legacy Feature
[**UpdateWhitelabelConfiguration**](ConfigApi.md#UpdateWhitelabelConfiguration) | **Put** /whitelabel_configuration | Update Whitelabel configuration
[**Versions**](ConfigApi.md#Versions) | **Get** /versions | Get ApiVersion
[**WhitelabelConfiguration**](ConfigApi.md#WhitelabelConfiguration) | **Get** /whitelabel_configuration | Get Whitelabel configuration



## AllLegacyFeatures

> []LegacyFeature AllLegacyFeatures(ctx, )

Get All Legacy Features

### Get all legacy features. 

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]LegacyFeature**](LegacyFeature.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AllLocales

> []Locale AllLocales(ctx, )

Get All Locales

### Get a list of locales that Looker supports. 

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]Locale**](Locale.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AllTimezones

> []Timezone AllTimezones(ctx, )

Get All Timezones

### Get a list of timezones that Looker supports (e.g. useful for scheduling tasks). 

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]Timezone**](Timezone.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackupConfiguration

> BackupConfiguration BackupConfiguration(ctx, )

Get Backup Configuration

### Get the current Looker internal database backup configuration. 

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**BackupConfiguration**](BackupConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomWelcomeEmail

> CustomWelcomeEmail CustomWelcomeEmail(ctx, )

Get Custom Welcome Email

### Get the current status and content of custom welcome emails 

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**CustomWelcomeEmail**](CustomWelcomeEmail.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InternalHelpResources

> InternalHelpResources InternalHelpResources(ctx, )

Get Internal Help Resources

### Get and set the options for internal help resources 

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**InternalHelpResources**](InternalHelpResources.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InternalHelpResourcesContent

> InternalHelpResourcesContent InternalHelpResourcesContent(ctx, )

Get Internal Help Resources Content

### Set the menu item name and content for internal help resources 

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**InternalHelpResourcesContent**](InternalHelpResourcesContent.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LegacyFeature

> LegacyFeature LegacyFeature(ctx, legacyFeatureId)

Get Legacy Feature

### Get information about the legacy feature with a specific id. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**legacyFeatureId** | **int64**| id of legacy feature | 

### Return type

[**LegacyFeature**](LegacyFeature.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBackupConfiguration

> BackupConfiguration UpdateBackupConfiguration(ctx, body)

Update Backup Configuration

### Update the Looker internal database backup configuration. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**BackupConfiguration**](BackupConfiguration.md)| Options for Backup Configuration | 

### Return type

[**BackupConfiguration**](BackupConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCustomWelcomeEmail

> CustomWelcomeEmail UpdateCustomWelcomeEmail(ctx, body, optional)

Update Custom Welcome Email Content

Update custom welcome email setting and values. Optionally send a test email with the new content to the currently logged in user. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**CustomWelcomeEmail**](CustomWelcomeEmail.md)| Custom Welcome Email setting and value to save | 
 **optional** | ***UpdateCustomWelcomeEmailOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateCustomWelcomeEmailOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sendTestWelcomeEmail** | **optional.Bool**| If true a test email with the content from the request will be sent to the current user after saving | 

### Return type

[**CustomWelcomeEmail**](CustomWelcomeEmail.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCustomWelcomeEmailTest

> WelcomeEmailTest UpdateCustomWelcomeEmailTest(ctx, body)

Send a test welcome email to the currently logged in user with the supplied content 

Requests to this endpoint will send a welcome email with the custom content provided in the body to the currently logged in user. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**WelcomeEmailTest**](WelcomeEmailTest.md)| Body of the email to be sent. | 

### Return type

[**WelcomeEmailTest**](WelcomeEmailTest.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateInternalHelpResources

> InternalHelpResources UpdateInternalHelpResources(ctx, body)

Update internal help resources configuration

Update internal help resources settings 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**InternalHelpResources**](InternalHelpResources.md)| Custom Welcome Email | 

### Return type

[**InternalHelpResources**](InternalHelpResources.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateInternalHelpResourcesContent

> InternalHelpResourcesContent UpdateInternalHelpResourcesContent(ctx, body)

Update internal help resources content

Update internal help resources content 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**InternalHelpResourcesContent**](InternalHelpResourcesContent.md)| Internal Help Resources Content | 

### Return type

[**InternalHelpResourcesContent**](InternalHelpResourcesContent.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLegacyFeature

> LegacyFeature UpdateLegacyFeature(ctx, legacyFeatureId, body)

Update Legacy Feature

### Update information about the legacy feature with a specific id. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**legacyFeatureId** | **int64**| id of legacy feature | 
**body** | [**LegacyFeature**](LegacyFeature.md)| Legacy Feature | 

### Return type

[**LegacyFeature**](LegacyFeature.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWhitelabelConfiguration

> WhitelabelConfiguration UpdateWhitelabelConfiguration(ctx, body)

Update Whitelabel configuration

### Update the whitelabel configuration 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**WhitelabelConfiguration**](WhitelabelConfiguration.md)| Whitelabel configuration | 

### Return type

[**WhitelabelConfiguration**](WhitelabelConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Versions

> ApiVersion Versions(ctx, optional)

Get ApiVersion

### Get information about all API versions supported by this Looker instance. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***VersionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a VersionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **optional.String**| Requested fields. | 

### Return type

[**ApiVersion**](ApiVersion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WhitelabelConfiguration

> WhitelabelConfiguration WhitelabelConfiguration(ctx, optional)

Get Whitelabel configuration

### This feature is enabled only by special license. ### Gets the whitelabel configuration, which includes hiding documentation links, custom favicon uploading, etc. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***WhitelabelConfigurationOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a WhitelabelConfigurationOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **optional.String**| Requested fields. | 

### Return type

[**WhitelabelConfiguration**](WhitelabelConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

