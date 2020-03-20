# \DashboardApi

All URIs are relative to *https://mattermost.looker.com:19999/api/3.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AllDashboards**](DashboardApi.md#AllDashboards) | **Get** /dashboards | Get All Dashboards
[**CreateDashboard**](DashboardApi.md#CreateDashboard) | **Post** /dashboards | Create Dashboard
[**CreateDashboardElement**](DashboardApi.md#CreateDashboardElement) | **Post** /dashboard_elements | Create DashboardElement
[**CreateDashboardFilter**](DashboardApi.md#CreateDashboardFilter) | **Post** /dashboard_filters | Create Dashboard Filter
[**CreateDashboardLayout**](DashboardApi.md#CreateDashboardLayout) | **Post** /dashboard_layouts | Create DashboardLayout
[**Dashboard**](DashboardApi.md#Dashboard) | **Get** /dashboards/{dashboard_id} | Get Dashboard
[**DashboardDashboardElements**](DashboardApi.md#DashboardDashboardElements) | **Get** /dashboards/{dashboard_id}/dashboard_elements | Get All DashboardElements
[**DashboardDashboardFilters**](DashboardApi.md#DashboardDashboardFilters) | **Get** /dashboards/{dashboard_id}/dashboard_filters | Get All Dashboard Filters
[**DashboardDashboardLayouts**](DashboardApi.md#DashboardDashboardLayouts) | **Get** /dashboards/{dashboard_id}/dashboard_layouts | Get All DashboardLayouts
[**DashboardElement**](DashboardApi.md#DashboardElement) | **Get** /dashboard_elements/{dashboard_element_id} | Get DashboardElement
[**DashboardFilter**](DashboardApi.md#DashboardFilter) | **Get** /dashboard_filters/{dashboard_filter_id} | Get Dashboard Filter
[**DashboardLayout**](DashboardApi.md#DashboardLayout) | **Get** /dashboard_layouts/{dashboard_layout_id} | Get DashboardLayout
[**DashboardLayoutComponent**](DashboardApi.md#DashboardLayoutComponent) | **Get** /dashboard_layout_components/{dashboard_layout_component_id} | Get DashboardLayoutComponent
[**DashboardLayoutDashboardLayoutComponents**](DashboardApi.md#DashboardLayoutDashboardLayoutComponents) | **Get** /dashboard_layouts/{dashboard_layout_id}/dashboard_layout_components | Get All DashboardLayoutComponents
[**DashboardLookml**](DashboardApi.md#DashboardLookml) | **Get** /dashboards/lookml/{dashboard_id} | Get lookml of a UDD
[**DeleteDashboard**](DashboardApi.md#DeleteDashboard) | **Delete** /dashboards/{dashboard_id} | Delete Dashboard
[**DeleteDashboardElement**](DashboardApi.md#DeleteDashboardElement) | **Delete** /dashboard_elements/{dashboard_element_id} | Delete DashboardElement
[**DeleteDashboardFilter**](DashboardApi.md#DeleteDashboardFilter) | **Delete** /dashboard_filters/{dashboard_filter_id} | Delete Dashboard Filter
[**DeleteDashboardLayout**](DashboardApi.md#DeleteDashboardLayout) | **Delete** /dashboard_layouts/{dashboard_layout_id} | Delete DashboardLayout
[**ImportLookmlDashboard**](DashboardApi.md#ImportLookmlDashboard) | **Post** /dashboards/{lookml_dashboard_id}/import/{space_id} | Import LookML Dashboard
[**SearchDashboardElements**](DashboardApi.md#SearchDashboardElements) | **Get** /dashboard_elements/search | Search Dashboard Elements
[**SearchDashboards**](DashboardApi.md#SearchDashboards) | **Get** /dashboards/search | Search Dashboards
[**SyncLookmlDashboard**](DashboardApi.md#SyncLookmlDashboard) | **Patch** /dashboards/{lookml_dashboard_id}/sync | Sync LookML Dashboard
[**UpdateDashboard**](DashboardApi.md#UpdateDashboard) | **Patch** /dashboards/{dashboard_id} | Update Dashboard
[**UpdateDashboardElement**](DashboardApi.md#UpdateDashboardElement) | **Patch** /dashboard_elements/{dashboard_element_id} | Update DashboardElement
[**UpdateDashboardFilter**](DashboardApi.md#UpdateDashboardFilter) | **Patch** /dashboard_filters/{dashboard_filter_id} | Update Dashboard Filter
[**UpdateDashboardLayout**](DashboardApi.md#UpdateDashboardLayout) | **Patch** /dashboard_layouts/{dashboard_layout_id} | Update DashboardLayout
[**UpdateDashboardLayoutComponent**](DashboardApi.md#UpdateDashboardLayoutComponent) | **Patch** /dashboard_layout_components/{dashboard_layout_component_id} | Update DashboardLayoutComponent



## AllDashboards

> []DashboardBase AllDashboards(ctx, optional)

Get All Dashboards

### Get information about all active dashboards.  Returns an array of **abbreviated dashboard objects**. Dashboards marked as deleted are excluded from this list.  Get the **full details** of a specific dashboard by id with [dashboard()](#!/Dashboard/dashboard)  Find **deleted dashboards** with [search_dashboards()](#!/Dashboard/search_dashboards) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AllDashboardsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AllDashboardsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **optional.String**| Requested fields. | 

### Return type

[**[]DashboardBase**](DashboardBase.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDashboard

> Dashboard CreateDashboard(ctx, body)

Create Dashboard

### Create a new dashboard  Creates a new dashboard object and returns the details of the newly created dashboard.  `Title`, `user_id`, and `space_id` are all required fields. `Space_id` and `user_id` must contain the id of an existing space or user, respectively. A dashboard's `title` must be unique within the space in which it resides.  If you receive a 422 error response when creating a dashboard, be sure to look at the response body for information about exactly which fields are missing or contain invalid data.  You can **update** an existing dashboard with [update_dashboard()](#!/Dashboard/update_dashboard)  You can **permanently delete** an existing dashboard with [delete_dashboard()](#!/Dashboard/delete_dashboard) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**Dashboard**](Dashboard.md)| Dashboard | 

### Return type

[**Dashboard**](Dashboard.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDashboardElement

> DashboardElement CreateDashboardElement(ctx, body, optional)

Create DashboardElement

### Create a dashboard element on the dashboard with a specific id.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**DashboardElement**](DashboardElement.md)| DashboardElement | 
 **optional** | ***CreateDashboardElementOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateDashboardElementOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Requested fields. | 

### Return type

[**DashboardElement**](DashboardElement.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDashboardFilter

> DashboardFilter CreateDashboardFilter(ctx, body, optional)

Create Dashboard Filter

### Create a dashboard filter on the dashboard with a specific id.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**CreateDashboardFilter**](CreateDashboardFilter.md)| Dashboard Filter | 
 **optional** | ***CreateDashboardFilterOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateDashboardFilterOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Requested fields | 

### Return type

[**DashboardFilter**](DashboardFilter.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDashboardLayout

> DashboardLayout CreateDashboardLayout(ctx, body, optional)

Create DashboardLayout

### Create a dashboard layout on the dashboard with a specific id.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**DashboardLayout**](DashboardLayout.md)| DashboardLayout | 
 **optional** | ***CreateDashboardLayoutOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateDashboardLayoutOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Requested fields. | 

### Return type

[**DashboardLayout**](DashboardLayout.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Dashboard

> Dashboard Dashboard(ctx, dashboardId, optional)

Get Dashboard

### Get information about a dashboard  Returns the full details of the identified dashboard object  Get a **summary list** of all active dashboards with [all_dashboards()](#!/Dashboard/all_dashboards)  You can **Search** for dashboards with [search_dashboards()](#!/Dashboard/search_dashboards) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dashboardId** | **string**| Id of dashboard | 
 **optional** | ***DashboardOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DashboardOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Requested fields. | 

### Return type

[**Dashboard**](Dashboard.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DashboardDashboardElements

> []DashboardElement DashboardDashboardElements(ctx, dashboardId, optional)

Get All DashboardElements

### Get information about all the dashboard elements on a dashboard with a specific id.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dashboardId** | **string**| Id of dashboard | 
 **optional** | ***DashboardDashboardElementsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DashboardDashboardElementsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Requested fields. | 

### Return type

[**[]DashboardElement**](DashboardElement.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DashboardDashboardFilters

> []DashboardFilter DashboardDashboardFilters(ctx, dashboardId, optional)

Get All Dashboard Filters

### Get information about all the dashboard filters on a dashboard with a specific id.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dashboardId** | **string**| Id of dashboard | 
 **optional** | ***DashboardDashboardFiltersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DashboardDashboardFiltersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Requested fields. | 

### Return type

[**[]DashboardFilter**](DashboardFilter.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DashboardDashboardLayouts

> []DashboardLayout DashboardDashboardLayouts(ctx, dashboardId, optional)

Get All DashboardLayouts

### Get information about all the dashboard elements on a dashboard with a specific id.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dashboardId** | **string**| Id of dashboard | 
 **optional** | ***DashboardDashboardLayoutsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DashboardDashboardLayoutsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Requested fields. | 

### Return type

[**[]DashboardLayout**](DashboardLayout.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DashboardElement

> DashboardElement DashboardElement(ctx, dashboardElementId, optional)

Get DashboardElement

### Get information about the dashboard element with a specific id.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dashboardElementId** | **string**| Id of dashboard element | 
 **optional** | ***DashboardElementOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DashboardElementOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Requested fields. | 

### Return type

[**DashboardElement**](DashboardElement.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DashboardFilter

> DashboardFilter DashboardFilter(ctx, dashboardFilterId, optional)

Get Dashboard Filter

### Get information about the dashboard filters with a specific id.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dashboardFilterId** | **string**| Id of dashboard filters | 
 **optional** | ***DashboardFilterOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DashboardFilterOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Requested fields. | 

### Return type

[**DashboardFilter**](DashboardFilter.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DashboardLayout

> DashboardLayout DashboardLayout(ctx, dashboardLayoutId, optional)

Get DashboardLayout

### Get information about the dashboard layouts with a specific id.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dashboardLayoutId** | **string**| Id of dashboard layouts | 
 **optional** | ***DashboardLayoutOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DashboardLayoutOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Requested fields. | 

### Return type

[**DashboardLayout**](DashboardLayout.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DashboardLayoutComponent

> DashboardLayoutComponent DashboardLayoutComponent(ctx, dashboardLayoutComponentId, optional)

Get DashboardLayoutComponent

### Get information about the dashboard elements with a specific id.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dashboardLayoutComponentId** | **string**| Id of dashboard layout component | 
 **optional** | ***DashboardLayoutComponentOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DashboardLayoutComponentOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Requested fields. | 

### Return type

[**DashboardLayoutComponent**](DashboardLayoutComponent.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DashboardLayoutDashboardLayoutComponents

> []DashboardLayoutComponent DashboardLayoutDashboardLayoutComponents(ctx, dashboardLayoutId, optional)

Get All DashboardLayoutComponents

### Get information about all the dashboard layout components for a dashboard layout with a specific id.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dashboardLayoutId** | **string**| Id of dashboard layout component | 
 **optional** | ***DashboardLayoutDashboardLayoutComponentsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DashboardLayoutDashboardLayoutComponentsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Requested fields. | 

### Return type

[**[]DashboardLayoutComponent**](DashboardLayoutComponent.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DashboardLookml

> DashboardLookml DashboardLookml(ctx, dashboardId)

Get lookml of a UDD

### Get lookml of a UDD  Returns a JSON object that contains the dashboard id and the full lookml  

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dashboardId** | **string**| Id of dashboard | 

### Return type

[**DashboardLookml**](DashboardLookml.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDashboard

> string DeleteDashboard(ctx, dashboardId)

Delete Dashboard

### Delete the dashboard with the specified id  Permanently **deletes** a dashboard. (The dashboard cannot be recovered after this operation.)  \"Soft\" delete or hide a dashboard by setting its `deleted` status to `True` with [update_dashboard()](#!/Dashboard/update_dashboard).  Note: When a dashboard is deleted in the UI, it is soft deleted. Use this API call to permanently remove it, if desired. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dashboardId** | **string**| Id of dashboard | 

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


## DeleteDashboardElement

> string DeleteDashboardElement(ctx, dashboardElementId)

Delete DashboardElement

### Delete a dashboard element with a specific id.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dashboardElementId** | **string**| Id of dashboard element | 

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


## DeleteDashboardFilter

> string DeleteDashboardFilter(ctx, dashboardFilterId)

Delete Dashboard Filter

### Delete a dashboard filter with a specific id.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dashboardFilterId** | **string**| Id of dashboard filter | 

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


## DeleteDashboardLayout

> string DeleteDashboardLayout(ctx, dashboardLayoutId)

Delete DashboardLayout

### Delete a dashboard layout with a specific id.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dashboardLayoutId** | **string**| Id of dashboard layout | 

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


## ImportLookmlDashboard

> Dashboard ImportLookmlDashboard(ctx, lookmlDashboardId, spaceId, optional)

Import LookML Dashboard

### Import a LookML dashboard to a space as a UDD Creates a UDD (a dashboard which exists in the Looker database rather than as a LookML file) from the LookML dashboard and puts it in the space specified. The created UDD will have a lookml_link_id which links to the original LookML dashboard.  To give the imported dashboard specify a (e.g. title: \"my title\") in the body of your request, otherwise the imported dashboard will have the same title as the original LookML dashboard.  For this operation to succeed the user must have permission to see the LookML dashboard in question, and have permission to create content in the space the dashboard is being imported to.  **Sync** a linked UDD with [sync_lookml_dashboard()](#!/Dashboard/sync_lookml_dashboard) **Unlink** a linked UDD by setting lookml_link_id to null with [update_dashboard()](#!/Dashboard/update_dashboard) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lookmlDashboardId** | **string**| Id of LookML dashboard | 
**spaceId** | **string**| Id of space to import the dashboard to | 
 **optional** | ***ImportLookmlDashboardOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ImportLookmlDashboardOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **rawLocale** | **optional.Bool**| If true, and this dashboard is localized, export it with the raw keys, not localized. | 
 **body** | [**optional.Interface of Dashboard**](Dashboard.md)| Dashboard | 

### Return type

[**Dashboard**](Dashboard.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchDashboardElements

> []DashboardElement SearchDashboardElements(ctx, optional)

Search Dashboard Elements

### Search Dashboard Elements  Returns an **array of DashboardElement objects** that match the specified search criteria.  If multiple search params are given and `filter_or` is FALSE or not specified, search params are combined in a logical AND operation. Only rows that match *all* search param criteria will be returned.  If `filter_or` is TRUE, multiple search params are combined in a logical OR operation. Results will include rows that match **any** of the search criteria.  String search params use case-insensitive matching. String search params can contain `%` and '_' as SQL LIKE pattern match wildcard expressions. example=\"dan%\" will match \"danger\" and \"Danzig\" but not \"David\" example=\"D_m%\" will match \"Damage\" and \"dump\"  Integer search params can accept a single value or a comma separated list of values. The multiple values will be combined under a logical OR operation - results will match at least one of the given values.  Most search params can accept \"IS NULL\" and \"NOT NULL\" as special expressions to match or exclude (respectively) rows where the column is null.  Boolean search params accept only \"true\" and \"false\" as values.  

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SearchDashboardElementsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SearchDashboardElementsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dashboardId** | **optional.Int64**| Select elements that refer to a given dashboard id | 
 **lookId** | **optional.Int64**| Select elements that refer to a given look id | 
 **title** | **optional.String**| Match the title of element | 
 **deleted** | **optional.Bool**| Select soft-deleted dashboard elements | 
 **fields** | **optional.String**| Requested fields. | 
 **filterOr** | **optional.Bool**| Combine given search criteria in a boolean OR expression | 
 **sorts** | **optional.String**| Fields to sort by. Sortable fields: [:look_id, :dashboard_id, :deleted, :title] | 

### Return type

[**[]DashboardElement**](DashboardElement.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchDashboards

> []Dashboard SearchDashboards(ctx, optional)

Search Dashboards

### Search Dashboards  Returns an **array of dashboard objects** that match the specified search criteria.  If multiple search params are given and `filter_or` is FALSE or not specified, search params are combined in a logical AND operation. Only rows that match *all* search param criteria will be returned.  If `filter_or` is TRUE, multiple search params are combined in a logical OR operation. Results will include rows that match **any** of the search criteria.  String search params use case-insensitive matching. String search params can contain `%` and '_' as SQL LIKE pattern match wildcard expressions. example=\"dan%\" will match \"danger\" and \"Danzig\" but not \"David\" example=\"D_m%\" will match \"Damage\" and \"dump\"  Integer search params can accept a single value or a comma separated list of values. The multiple values will be combined under a logical OR operation - results will match at least one of the given values.  Most search params can accept \"IS NULL\" and \"NOT NULL\" as special expressions to match or exclude (respectively) rows where the column is null.  Boolean search params accept only \"true\" and \"false\" as values.   The parameters `limit`, and `offset` are recommended for fetching results in page-size chunks.  Get a **single dashboard** by id with [dashboard()](#!/Dashboard/dashboard) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SearchDashboardsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SearchDashboardsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.Int64**| Match dashboard id. | 
 **slug** | **optional.String**| Match dashboard slug. | 
 **title** | **optional.String**| Match Dashboard title. | 
 **description** | **optional.String**| Match Dashboard description. | 
 **contentFavoriteId** | **optional.Int64**| Filter on a content favorite id. | 
 **spaceId** | **optional.String**| Filter on a particular space. | 
 **deleted** | **optional.String**| Filter on dashboards deleted status. | 
 **userId** | **optional.String**| Filter on dashboards created by a particular user. | 
 **viewCount** | **optional.String**| Filter on a particular value of view_count | 
 **contentMetadataId** | **optional.Int64**| Filter on a content favorite id. | 
 **curate** | **optional.Bool**| Exclude items that exist only in personal spaces other than the users | 
 **fields** | **optional.String**| Requested fields. | 
 **page** | **optional.Int64**| Requested page. | 
 **perPage** | **optional.Int64**| Results per page. | 
 **limit** | **optional.Int64**| Number of results to return. (used with offset and takes priority over page and per_page) | 
 **offset** | **optional.Int64**| Number of results to skip before returning any. (used with limit and takes priority over page and per_page) | 
 **sorts** | **optional.String**| One or more fields to sort by. Sortable fields: [:title, :user_id, :id, :created_at, :space_id, :description, :view_count, :favorite_count, :slug, :content_favorite_id, :content_metadata_id, :deleted, :deleted_at, :last_viewed_at, :last_accessed_at] | 
 **filterOr** | **optional.Bool**| Combine given search criteria in a boolean OR expression | 

### Return type

[**[]Dashboard**](Dashboard.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SyncLookmlDashboard

> []int64 SyncLookmlDashboard(ctx, lookmlDashboardId, body, optional)

Sync LookML Dashboard

### Update all linked dashboards to match the specified LookML dashboard.  Any UDD (a dashboard which exists in the Looker database rather than as a LookML file) which has a `lookml_link_id` property value referring to a LookML dashboard's id (model::dashboardname) will be updated so that it matches the current state of the LookML dashboard.  For this operation to succeed the user must have permission to view the LookML dashboard, and only linked dashboards that the user has permission to update will be synced.  To **link** or **unlink** a UDD set the `lookml_link_id` property with [update_dashboard()](#!/Dashboard/update_dashboard) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lookmlDashboardId** | **string**| Id of LookML dashboard, in the form &#39;model::dashboardname&#39; | 
**body** | [**Dashboard**](Dashboard.md)| Dashboard | 
 **optional** | ***SyncLookmlDashboardOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SyncLookmlDashboardOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **rawLocale** | **optional.Bool**| If true, and this dashboard is localized, export it with the raw keys, not localized. | 

### Return type

**[]int64**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDashboard

> Dashboard UpdateDashboard(ctx, dashboardId, body)

Update Dashboard

### Update a dashboard  You can use this function to change the string and integer properties of a dashboard. Nested objects such as filters, dashboard elements, or dashboard layout components cannot be modified by this function - use the update functions for the respective nested object types (like [update_dashboard_filter()](#!/3.1/Dashboard/update_dashboard_filter) to change a filter) to modify nested objects referenced by a dashboard.  If you receive a 422 error response when updating a dashboard, be sure to look at the response body for information about exactly which fields are missing or contain invalid data. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dashboardId** | **string**| Id of dashboard | 
**body** | [**Dashboard**](Dashboard.md)| Dashboard | 

### Return type

[**Dashboard**](Dashboard.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDashboardElement

> DashboardElement UpdateDashboardElement(ctx, dashboardElementId, body, optional)

Update DashboardElement

### Update the dashboard element with a specific id.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dashboardElementId** | **string**| Id of dashboard element | 
**body** | [**DashboardElement**](DashboardElement.md)| DashboardElement | 
 **optional** | ***UpdateDashboardElementOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateDashboardElementOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **optional.String**| Requested fields. | 

### Return type

[**DashboardElement**](DashboardElement.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDashboardFilter

> DashboardFilter UpdateDashboardFilter(ctx, dashboardFilterId, body, optional)

Update Dashboard Filter

### Update the dashboard filter with a specific id.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dashboardFilterId** | **string**| Id of dashboard filter | 
**body** | [**DashboardFilter**](DashboardFilter.md)| Dashboard Filter | 
 **optional** | ***UpdateDashboardFilterOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateDashboardFilterOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **optional.String**| Requested fields. | 

### Return type

[**DashboardFilter**](DashboardFilter.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDashboardLayout

> DashboardLayout UpdateDashboardLayout(ctx, dashboardLayoutId, body, optional)

Update DashboardLayout

### Update the dashboard layout with a specific id.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dashboardLayoutId** | **string**| Id of dashboard layout | 
**body** | [**DashboardLayout**](DashboardLayout.md)| DashboardLayout | 
 **optional** | ***UpdateDashboardLayoutOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateDashboardLayoutOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **optional.String**| Requested fields. | 

### Return type

[**DashboardLayout**](DashboardLayout.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDashboardLayoutComponent

> DashboardLayoutComponent UpdateDashboardLayoutComponent(ctx, dashboardLayoutComponentId, body, optional)

Update DashboardLayoutComponent

### Update the dashboard element with a specific id.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dashboardLayoutComponentId** | **string**| Id of dashboard layout component | 
**body** | [**DashboardLayoutComponent**](DashboardLayoutComponent.md)| DashboardLayoutComponent | 
 **optional** | ***UpdateDashboardLayoutComponentOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateDashboardLayoutComponentOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **optional.String**| Requested fields. | 

### Return type

[**DashboardLayoutComponent**](DashboardLayoutComponent.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

