# LookWithQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Can** | **map[string]bool** | Operations the current user is able to perform on this object | [optional] [readonly] 
**ContentMetadataId** | **int64** | Id of content metadata | [optional] [readonly] 
**Id** | **int64** | Unique Id | [optional] [readonly] 
**Title** | **string** | Look Title | [optional] 
**ContentFavoriteId** | **int64** | Content Favorite Id | [optional] [readonly] 
**CreatedAt** | [**time.Time**](time.Time.md) | Time that the Look was created. | [optional] [readonly] 
**Deleted** | **bool** | Whether or not a look is &#39;soft&#39; deleted. | [optional] 
**DeletedAt** | [**time.Time**](time.Time.md) | Time that the Look was deleted. | [optional] [readonly] 
**DeleterId** | **int64** | Id of User that deleted the look. | [optional] [readonly] 
**Description** | **string** | Description | [optional] 
**EmbedUrl** | **string** | Embed Url | [optional] [readonly] 
**ExcelFileUrl** | **string** | Excel File Url | [optional] [readonly] 
**FavoriteCount** | **int64** | Number of times favorited | [optional] [readonly] 
**GoogleSpreadsheetFormula** | **string** | Google Spreadsheet Formula | [optional] [readonly] 
**ImageEmbedUrl** | **string** | Image Embed Url | [optional] [readonly] 
**IsRunOnLoad** | **bool** | auto-run query when Look viewed | [optional] 
**LastAccessedAt** | [**time.Time**](time.Time.md) | Time that the Look was last accessed by any user | [optional] [readonly] 
**LastUpdaterId** | **int64** | Id of User that last updated the look. | [optional] [readonly] 
**LastViewedAt** | [**time.Time**](time.Time.md) | Time last viewed in the Looker web UI | [optional] [readonly] 
**Model** | [**LookModel**](LookModel.md) |  | [optional] 
**Public** | **bool** | Is Public | [optional] 
**PublicSlug** | **string** | Public Slug | [optional] [readonly] 
**PublicUrl** | **string** | Public Url | [optional] [readonly] 
**QueryId** | **int64** | Query Id | [optional] 
**ShortUrl** | **string** | Short Url | [optional] [readonly] 
**Space** | [**SpaceBase**](SpaceBase.md) |  | [optional] 
**Folder** | [**FolderBase**](FolderBase.md) |  | [optional] 
**SpaceId** | **string** | Space Id | [optional] 
**FolderId** | **string** | Folder Id | [optional] 
**UpdatedAt** | [**time.Time**](time.Time.md) | Time that the Look was updated. | [optional] [readonly] 
**UserId** | **int64** | User Id | [optional] 
**ViewCount** | **int64** | Number of times viewed in the Looker web UI | [optional] [readonly] 
**User** | [**UserIdOnly**](UserIdOnly.md) |  | [optional] 
**Query** | [**Query**](Query.md) |  | [optional] 
**Url** | **string** | Url | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


