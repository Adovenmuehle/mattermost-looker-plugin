# \ThemeApi

All URIs are relative to *https://mattermost.looker.com:19999/api/3.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActiveThemes**](ThemeApi.md#ActiveThemes) | **Get** /themes/active | Get Active Themes
[**AllThemes**](ThemeApi.md#AllThemes) | **Get** /themes | Get All Themes
[**CreateTheme**](ThemeApi.md#CreateTheme) | **Post** /themes | Create Theme
[**DefaultTheme**](ThemeApi.md#DefaultTheme) | **Get** /themes/default | Get Default Theme
[**DeleteTheme**](ThemeApi.md#DeleteTheme) | **Delete** /themes/{theme_id} | Delete Theme
[**SearchThemes**](ThemeApi.md#SearchThemes) | **Get** /themes/search | Search Themes
[**SetDefaultTheme**](ThemeApi.md#SetDefaultTheme) | **Put** /themes/default | Set Default Theme
[**Theme**](ThemeApi.md#Theme) | **Get** /themes/{theme_id} | Get Theme
[**ThemeOrDefault**](ThemeApi.md#ThemeOrDefault) | **Get** /themes/theme_or_default | Get Theme or Default
[**UpdateTheme**](ThemeApi.md#UpdateTheme) | **Patch** /themes/{theme_id} | Update Theme
[**ValidateTheme**](ThemeApi.md#ValidateTheme) | **Post** /themes/validate | Validate Theme



## ActiveThemes

> []Theme ActiveThemes(ctx, optional)

Get Active Themes

### Get active themes  Returns an array of active themes.  If the `name` parameter is specified, it will return an array with one theme if it's active and found.  The optional `ts` parameter can specify a different timestamp than \"now.\"  **Note**: Custom themes needs to be enabled by Looker. Unless custom themes are enabled, only the automatically generated default theme can be used. Please contact your Account Manager or support@looker.com to update your license for this feature.   

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ActiveThemesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ActiveThemesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**| Name of theme | 
 **ts** | **optional.Time**| Timestamp representing the target datetime for the active period. Defaults to &#39;now&#39; | 
 **fields** | **optional.String**| Requested fields. | 

### Return type

[**[]Theme**](Theme.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AllThemes

> []Theme AllThemes(ctx, optional)

Get All Themes

### Get an array of all existing themes  Get a **single theme** by id with [Theme](#!/Theme/theme)  This method returns an array of all existing themes. The active time for the theme is not considered.  **Note**: Custom themes needs to be enabled by Looker. Unless custom themes are enabled, only the automatically generated default theme can be used. Please contact your Account Manager or support@looker.com to update your license for this feature.  

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AllThemesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AllThemesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **optional.String**| Requested fields. | 

### Return type

[**[]Theme**](Theme.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTheme

> Theme CreateTheme(ctx, body)

Create Theme

### Create a theme  Creates a new theme object, returning the theme details, including the created id.  If `settings` are not specified, the default theme settings will be copied into the new theme.  The theme `name` can only contain alphanumeric characters or underscores. Theme names should not contain any confidential information, such as customer names.  **Update** an existing theme with [Update Theme](#!/Theme/update_theme)  **Permanently delete** an existing theme with [Delete Theme](#!/Theme/delete_theme)  For more information, see [Creating and Applying Themes](https://looker.com/docs/r/admin/themes).  **Note**: Custom themes needs to be enabled by Looker. Unless custom themes are enabled, only the automatically generated default theme can be used. Please contact your Account Manager or support@looker.com to update your license for this feature.  

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**Theme**](Theme.md)| Theme | 

### Return type

[**Theme**](Theme.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DefaultTheme

> Theme DefaultTheme(ctx, optional)

Get Default Theme

### Get the default theme  Returns the active theme object set as the default.  The **default** theme name can be set in the UI on the Admin|Theme UI page  The optional `ts` parameter can specify a different timestamp than \"now.\" If specified, it returns the default theme at the time indicated. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DefaultThemeOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DefaultThemeOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ts** | **optional.Time**| Timestamp representing the target datetime for the active period. Defaults to &#39;now&#39; | 

### Return type

[**Theme**](Theme.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTheme

> string DeleteTheme(ctx, themeId)

Delete Theme

### Delete a specific theme by id  This operation permanently deletes the identified theme from the database.  Because multiple themes can have the same name (with different activation time spans) themes can only be deleted by ID.  All IDs associated with a theme name can be retrieved by searching for the theme name with [Theme Search](#!/Theme/search).  **Note**: Custom themes needs to be enabled by Looker. Unless custom themes are enabled, only the automatically generated default theme can be used. Please contact your Account Manager or support@looker.com to update your license for this feature.  

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**themeId** | **string**| Id of theme | 

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


## SearchThemes

> []Theme SearchThemes(ctx, optional)

Search Themes

### Search all themes for matching criteria.  Returns an **array of theme objects** that match the specified search criteria.  | Search Parameters | Description | :-------------------: | :------ | | `begin_at` only | Find themes active at or after `begin_at` | `end_at` only | Find themes active at or before `end_at` | both set | Find themes with an active inclusive period between `begin_at` and `end_at`  Note: Range matching requires boolean AND logic. When using `begin_at` and `end_at` together, do not use `filter_or`=TRUE  If multiple search params are given and `filter_or` is FALSE or not specified, search params are combined in a logical AND operation. Only rows that match *all* search param criteria will be returned.  If `filter_or` is TRUE, multiple search params are combined in a logical OR operation. Results will include rows that match **any** of the search criteria.  String search params use case-insensitive matching. String search params can contain `%` and '_' as SQL LIKE pattern match wildcard expressions. example=\"dan%\" will match \"danger\" and \"Danzig\" but not \"David\" example=\"D_m%\" will match \"Damage\" and \"dump\"  Integer search params can accept a single value or a comma separated list of values. The multiple values will be combined under a logical OR operation - results will match at least one of the given values.  Most search params can accept \"IS NULL\" and \"NOT NULL\" as special expressions to match or exclude (respectively) rows where the column is null.  Boolean search params accept only \"true\" and \"false\" as values.   Get a **single theme** by id with [Theme](#!/Theme/theme)  **Note**: Custom themes needs to be enabled by Looker. Unless custom themes are enabled, only the automatically generated default theme can be used. Please contact your Account Manager or support@looker.com to update your license for this feature.  

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SearchThemesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SearchThemesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.Int64**| Match theme id. | 
 **name** | **optional.String**| Match theme name. | 
 **beginAt** | **optional.Time**| Timestamp for activation. | 
 **endAt** | **optional.Time**| Timestamp for expiration. | 
 **limit** | **optional.Int64**| Number of results to return (used with &#x60;offset&#x60;). | 
 **offset** | **optional.Int64**| Number of results to skip before returning any (used with &#x60;limit&#x60;). | 
 **sorts** | **optional.String**| Fields to sort by. | 
 **fields** | **optional.String**| Requested fields. | 
 **filterOr** | **optional.Bool**| Combine given search criteria in a boolean OR expression | 

### Return type

[**[]Theme**](Theme.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetDefaultTheme

> Theme SetDefaultTheme(ctx, name)

Set Default Theme

### Set the global default theme by theme name  Only Admin users can call this function.  Only an active theme with no expiration (`end_at` not set) can be assigned as the default theme. As long as a theme has an active record with no expiration, it can be set as the default.  [Create Theme](#!/Theme/create) has detailed information on rules for default and active themes  Returns the new specified default theme object.  **Note**: Custom themes needs to be enabled by Looker. Unless custom themes are enabled, only the automatically generated default theme can be used. Please contact your Account Manager or support@looker.com to update your license for this feature.  

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string**| Name of theme to set as default | 

### Return type

[**Theme**](Theme.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Theme

> Theme Theme(ctx, themeId, optional)

Get Theme

### Get a theme by ID  Use this to retrieve a specific theme, whether or not it's currently active.  **Note**: Custom themes needs to be enabled by Looker. Unless custom themes are enabled, only the automatically generated default theme can be used. Please contact your Account Manager or support@looker.com to update your license for this feature.  

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**themeId** | **string**| Id of theme | 
 **optional** | ***ThemeOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ThemeOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Requested fields. | 

### Return type

[**Theme**](Theme.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ThemeOrDefault

> Theme ThemeOrDefault(ctx, name, optional)

Get Theme or Default

### Get the named theme if it's active. Otherwise, return the default theme  The optional `ts` parameter can specify a different timestamp than \"now.\" Note: API users with `show` ability can call this function  **Note**: Custom themes needs to be enabled by Looker. Unless custom themes are enabled, only the automatically generated default theme can be used. Please contact your Account Manager or support@looker.com to update your license for this feature.  

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string**| Name of theme | 
 **optional** | ***ThemeOrDefaultOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ThemeOrDefaultOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ts** | **optional.Time**| Timestamp representing the target datetime for the active period. Defaults to &#39;now&#39; | 

### Return type

[**Theme**](Theme.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTheme

> Theme UpdateTheme(ctx, themeId, body)

Update Theme

### Update the theme by id.  **Note**: Custom themes needs to be enabled by Looker. Unless custom themes are enabled, only the automatically generated default theme can be used. Please contact your Account Manager or support@looker.com to update your license for this feature.  

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**themeId** | **string**| Id of theme | 
**body** | [**Theme**](Theme.md)| Theme | 

### Return type

[**Theme**](Theme.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateTheme

> ValidationError ValidateTheme(ctx, body)

Validate Theme

### Validate a theme with the specified information  Validates all values set for the theme, returning any errors encountered, or 200 OK if valid  See [Create Theme](#!/Theme/create_theme) for constraints  **Note**: Custom themes needs to be enabled by Looker. Unless custom themes are enabled, only the automatically generated default theme can be used. Please contact your Account Manager or support@looker.com to update your license for this feature.  

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**Theme**](Theme.md)| Theme | 

### Return type

[**ValidationError**](ValidationError.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

