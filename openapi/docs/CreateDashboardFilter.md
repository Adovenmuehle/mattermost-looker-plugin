# CreateDashboardFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique Id | [optional] [readonly] 
**DashboardId** | **string** | Id of Dashboard | 
**Name** | **string** | Name of filter | 
**Title** | **string** | Title of filter | 
**Type** | **string** | Type of filter: one of date, number, string, or field | 
**DefaultValue** | **string** | Default value of filter | [optional] 
**Model** | **string** | Model of filter (required if type &#x3D; field) | [optional] 
**Explore** | **string** | Explore of filter (required if type &#x3D; field) | [optional] 
**Dimension** | **string** | Dimension of filter (required if type &#x3D; field) | [optional] 
**Field** | **map[string]string** | Field information | [optional] [readonly] 
**Row** | **int64** | Display order of this filter relative to other filters | [optional] 
**ListensToFilters** | **[]string** | Array of listeners for faceted filters | [optional] 
**AllowMultipleValues** | **bool** | Whether the filter allows multiple filter values | [optional] 
**Required** | **bool** | Whether the filter requires a value to run the dashboard | [optional] 
**UiConfig** | **map[string]string** | The visual configuration for this filter. Used to set up how the UI for this filter should appear. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


