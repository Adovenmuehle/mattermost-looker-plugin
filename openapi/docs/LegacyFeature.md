# LegacyFeature

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Can** | **map[string]bool** | Operations the current user is able to perform on this object | [optional] [readonly] 
**Id** | **string** | Unique Id | [optional] [readonly] 
**Name** | **string** | Name | [optional] [readonly] 
**Description** | **string** | Description | [optional] [readonly] 
**EnabledLocally** | **bool** | Whether this feature has been enabled by a user | [optional] 
**Enabled** | **bool** | Whether this feature is currently enabled | [optional] [readonly] 
**DisallowedAsOfVersion** | **string** | Looker version where this feature became a legacy feature | [optional] [readonly] 
**DisableOnUpgradeToVersion** | **string** | Looker version where this feature will be automatically disabled | [optional] [readonly] 
**EndOfLifeVersion** | **string** | Future Looker version where this feature will be removed | [optional] [readonly] 
**DocumentationUrl** | **string** | URL for documentation about this feature | [optional] [readonly] 
**ApproximateDisableDate** | [**time.Time**](time.Time.md) | Approximate date that this feature will be automatically disabled. | [optional] [readonly] 
**ApproximateEndOfLifeDate** | [**time.Time**](time.Time.md) | Approximate date that this feature will be removed. | [optional] [readonly] 
**HasDisabledOnUpgrade** | **bool** | Whether this legacy feature may have been automatically disabled when upgrading to the current version. | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


