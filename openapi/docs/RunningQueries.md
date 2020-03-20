# RunningQueries

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Can** | **map[string]bool** | Operations the current user is able to perform on this object | [optional] [readonly] 
**Id** | **int64** | Unique Id | [optional] [readonly] 
**User** | [**UserPublic**](UserPublic.md) |  | [optional] 
**Query** | [**Query**](Query.md) |  | [optional] 
**SqlQuery** | [**SqlQuery**](SqlQuery.md) |  | [optional] 
**Look** | [**LookBasic**](LookBasic.md) |  | [optional] 
**CreatedAt** | **string** | Date/Time Query was initiated | [optional] [readonly] 
**CompletedAt** | **string** | Date/Time Query was completed | [optional] [readonly] 
**QueryId** | **string** | Query Id | [optional] [readonly] 
**Source** | **string** | Source (look, dashboard, queryrunner, explore, etc.) | [optional] [readonly] 
**NodeId** | **string** | Node Id | [optional] [readonly] 
**Slug** | **string** | Slug | [optional] [readonly] 
**QueryTaskId** | **string** | ID of a Query Task | [optional] [readonly] 
**CacheKey** | **string** | Cache Key | [optional] [readonly] 
**ConnectionName** | **string** | Connection | [optional] [readonly] 
**Dialect** | **string** | Dialect | [optional] [readonly] 
**ConnectionId** | **string** | Connection ID | [optional] [readonly] 
**Message** | **string** | Additional Information(Error message or verbose status) | [optional] [readonly] 
**Status** | **string** | Status description | [optional] [readonly] 
**Runtime** | **float64** | Number of seconds elapsed running the Query | [optional] [readonly] 
**Sql** | **string** | SQL text of the query as run | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


