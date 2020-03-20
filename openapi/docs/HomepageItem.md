# HomepageItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Can** | **map[string]bool** | Operations the current user is able to perform on this object | [optional] [readonly] 
**ContentCreatedBy** | **string** | Name of user who created the content this item is based on | [optional] [readonly] 
**ContentFavoriteId** | **int64** | Content favorite id associated with the item this content is based on | [optional] [readonly] 
**ContentMetadataId** | **int64** | Content metadata id associated with the item this content is based on | [optional] [readonly] 
**ContentUpdatedAt** | **string** | Last time the content that this item is based on was updated | [optional] [readonly] 
**CustomDescription** | **string** | Custom description entered by the user, if present | [optional] 
**CustomImageDataBase64** | **string** | (Write-Only) base64 encoded image data | [optional] 
**CustomImageUrl** | **string** | Custom image_url entered by the user, if present | [optional] [readonly] 
**CustomTitle** | **string** | Custom title entered by the user, if present | [optional] 
**CustomUrl** | **string** | Custom url entered by the user, if present | [optional] 
**DashboardId** | **int64** | Dashboard to base this item on | [optional] 
**Description** | **string** | The actual description for display | [optional] [readonly] 
**FavoriteCount** | **int64** | Number of times content has been favorited, if present | [optional] [readonly] 
**HomepageSectionId** | **string** | Associated Homepage Section | [optional] 
**Id** | **string** | Unique Id | [optional] [readonly] 
**ImageUrl** | **string** | The actual image_url for display | [optional] [readonly] 
**Location** | **string** | The container folder name of the content | [optional] [readonly] 
**LookId** | **int64** | Look to base this item on | [optional] 
**LookmlDashboardId** | **string** | LookML Dashboard to base this item on | [optional] 
**Order** | **int64** | An arbitrary integer representing the sort order within the section | [optional] 
**SectionFetchTime** | **float32** | Number of seconds it took to fetch the section this item is in | [optional] [readonly] 
**Title** | **string** | The actual title for display | [optional] [readonly] 
**Url** | **string** | The actual url for display | [optional] [readonly] 
**UseCustomDescription** | **bool** | Whether the custom description should be used instead of the content description, if the item is associated with content | [optional] 
**UseCustomImage** | **bool** | Whether the custom image should be used instead of the content image, if the item is associated with content | [optional] 
**UseCustomTitle** | **bool** | Whether the custom title should be used instead of the content title, if the item is associated with content | [optional] 
**UseCustomUrl** | **bool** | Whether the custom url should be used instead of the content url, if the item is associated with content | [optional] 
**ViewCount** | **int64** | Number of times content has been viewed, if present | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


