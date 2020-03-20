# Query

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Can** | **map[string]bool** | Operations the current user is able to perform on this object | [optional] [readonly] 
**Id** | **int64** | Unique Id | [optional] [readonly] 
**Model** | **string** | Model | 
**View** | **string** | Explore Name | 
**Fields** | **[]string** | Fields | [optional] 
**Pivots** | **[]string** | Pivots | [optional] 
**FillFields** | **[]string** | Fill Fields | [optional] 
**Filters** | **map[string]string** | Filters | [optional] 
**FilterExpression** | **string** | Filter Expression | [optional] 
**Sorts** | **[]string** | Sorting for the query results. Use the format &#x60;[\&quot;view.field\&quot;, ...]&#x60; to sort on fields in ascending order. Use the format &#x60;[\&quot;view.field desc\&quot;, ...]&#x60; to sort on fields in descending order. Use &#x60;[\&quot;__UNSORTED__\&quot;]&#x60; (2 underscores before and after) to disable sorting entirely. Empty sorts &#x60;[]&#x60; will trigger a default sort. | [optional] 
**Limit** | **string** | Limit | [optional] 
**ColumnLimit** | **string** | Column Limit | [optional] 
**Total** | **bool** | Total | [optional] 
**RowTotal** | **string** | Raw Total | [optional] 
**Subtotals** | **[]string** | Fields on which to run subtotals | [optional] 
**Runtime** | **float64** | Runtime | [optional] 
**VisConfig** | **map[string]string** | Visualization configuration properties. These properties are typically opaque and differ based on the type of visualization used. There is no specified set of allowed keys. The values can be any type supported by JSON. A \&quot;type\&quot; key with a string value is often present, and is used by Looker to determine which visualization to present. Visualizations ignore unknown vis_config properties. | [optional] 
**FilterConfig** | **map[string]string** | The filter_config represents the state of the filter UI on the explore page for a given query. When running a query via the Looker UI, this parameter takes precedence over \&quot;filters\&quot;. When creating a query or modifying an existing query, \&quot;filter_config\&quot; should be set to null. Setting it to any other value could cause unexpected filtering behavior. The format should be considered opaque. | [optional] 
**VisibleUiSections** | **string** | Visible UI Sections | [optional] 
**Slug** | **string** | Slug | [optional] [readonly] 
**DynamicFields** | **string** | Dynamic Fields | [optional] 
**ClientId** | **string** | Client Id: used to generate shortened explore URLs. If set by client, must be a unique 22 character alphanumeric string. Otherwise one will be generated. | [optional] 
**ShareUrl** | **string** | Share Url | [optional] [readonly] 
**ExpandedShareUrl** | **string** | Expanded Share Url | [optional] [readonly] 
**Url** | **string** | Expanded Url | [optional] [readonly] 
**QueryTimezone** | **string** | Query Timezone | [optional] 
**HasTableCalculations** | **bool** | Has Table Calculations | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


