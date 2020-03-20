# Theme

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Can** | **map[string]bool** | Operations the current user is able to perform on this object | [optional] [readonly] 
**BeginAt** | [**time.Time**](time.Time.md) | Timestamp for when this theme becomes active. Null&#x3D;always | [optional] 
**EndAt** | [**time.Time**](time.Time.md) | Timestamp for when this theme expires. Null&#x3D;never | [optional] 
**Id** | **int64** | Unique Id | [optional] [readonly] 
**Name** | **string** | Name of theme. Can only be alphanumeric and underscores. | [optional] 
**Settings** | [**ThemeSettings**](ThemeSettings.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


