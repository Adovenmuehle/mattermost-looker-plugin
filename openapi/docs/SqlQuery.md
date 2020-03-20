# SqlQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Can** | **map[string]bool** | Operations the current user is able to perform on this object | [optional] [readonly] 
**Slug** | **string** | The identifier of the SQL query | [optional] [readonly] 
**LastRuntime** | **float32** | Number of seconds this query took to run the most recent time it was run | [optional] [readonly] 
**RunCount** | **int64** | Number of times this query has been run | [optional] [readonly] 
**BrowserLimit** | **int64** | Maximum number of rows this query will display on the SQL Runner page | [optional] [readonly] 
**Sql** | **string** | SQL query text | [optional] [readonly] 
**LastRunAt** | **string** | The most recent time this query was run | [optional] [readonly] 
**Connection** | [**DbConnectionBase**](DBConnectionBase.md) |  | [optional] 
**ModelName** | **string** | Model name this query uses | [optional] [readonly] 
**Creator** | [**UserPublic**](UserPublic.md) |  | [optional] 
**ExploreUrl** | **string** | Explore page URL for this SQL query | [optional] [readonly] 
**Plaintext** | **bool** | Should this query be rendered as plain text | [optional] [readonly] 
**VisConfig** | **map[string]string** | Visualization configuration properties. These properties are typically opaque and differ based on the type of visualization used. There is no specified set of allowed keys. The values can be any type supported by JSON. A \&quot;type\&quot; key with a string value is often present, and is used by Looker to determine which visualization to present. Visualizations ignore unknown vis_config properties. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


