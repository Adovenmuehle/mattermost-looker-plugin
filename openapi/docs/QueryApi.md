# \QueryApi

All URIs are relative to *https://mattermost.looker.com:19999/api/3.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AllRunningQueries**](QueryApi.md#AllRunningQueries) | **Get** /running_queries | Get All Running Queries
[**CreateMergeQuery**](QueryApi.md#CreateMergeQuery) | **Post** /merge_queries | Create Merge Query
[**CreateQuery**](QueryApi.md#CreateQuery) | **Post** /queries | Create Query
[**CreateQueryTask**](QueryApi.md#CreateQueryTask) | **Post** /query_tasks | Run Query Async
[**CreateSqlQuery**](QueryApi.md#CreateSqlQuery) | **Post** /sql_queries | Create SQL Runner Query
[**KillQuery**](QueryApi.md#KillQuery) | **Delete** /running_queries/{query_task_id} | Kill Running Query
[**MergeQuery**](QueryApi.md#MergeQuery) | **Get** /merge_queries/{merge_query_id} | Get Merge Query
[**Query**](QueryApi.md#Query) | **Get** /queries/{query_id} | Get Query
[**QueryForSlug**](QueryApi.md#QueryForSlug) | **Get** /queries/slug/{slug} | Get Query for Slug
[**QueryTask**](QueryApi.md#QueryTask) | **Get** /query_tasks/{query_task_id} | Get Async Query Info
[**QueryTaskMultiResults**](QueryApi.md#QueryTaskMultiResults) | **Get** /query_tasks/multi_results | Get Multiple Async Query Results
[**QueryTaskResults**](QueryApi.md#QueryTaskResults) | **Get** /query_tasks/{query_task_id}/results | Get Async Query Results
[**RunInlineQuery**](QueryApi.md#RunInlineQuery) | **Post** /queries/run/{result_format} | Run Inline Query
[**RunQuery**](QueryApi.md#RunQuery) | **Get** /queries/{query_id}/run/{result_format} | Run Query
[**RunSqlQuery**](QueryApi.md#RunSqlQuery) | **Post** /sql_queries/{slug}/run/{result_format} | Run SQL Runner Query
[**RunUrlEncodedQuery**](QueryApi.md#RunUrlEncodedQuery) | **Get** /queries/models/{model_name}/views/{view_name}/run/{result_format} | Run Url Encoded Query
[**SqlQuery**](QueryApi.md#SqlQuery) | **Get** /sql_queries/{slug} | Get SQL Runner Query



## AllRunningQueries

> []RunningQueries AllRunningQueries(ctx, )

Get All Running Queries

Get information about all running queries. 

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]RunningQueries**](RunningQueries.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMergeQuery

> MergeQuery CreateMergeQuery(ctx, optional)

Create Merge Query

### Create Merge Query  Creates a new merge query object.  A merge query takes the results of one or more queries and combines (merges) the results according to field mapping definitions. The result is similar to a SQL left outer join.  A merge query can merge results of queries from different SQL databases.  The order that queries are defined in the source_queries array property is significant. The first query in the array defines the primary key into which the results of subsequent queries will be merged.  Like model/view query objects, merge queries are immutable and have structural identity - if you make a request to create a new merge query that is identical to an existing merge query, the existing merge query will be returned instead of creating a duplicate. Conversely, any change to the contents of a merge query will produce a new object with a new id. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateMergeQueryOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateMergeQueryOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **optional.String**| Requested fields | 
 **body** | [**optional.Interface of MergeQuery**](MergeQuery.md)| Merge Query | 

### Return type

[**MergeQuery**](MergeQuery.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateQuery

> Query CreateQuery(ctx, body, optional)

Create Query

### Create a query.  This allows you to create a new query that you can later run. Looker queries are immutable once created and are not deleted. If you create a query that is exactly like an existing query then the existing query will be returned and no new query will be created. Whether a new query is created or not, you can use the 'id' in the returned query with the 'run' method.  The query parameters are passed as json in the body of the request.  

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**Query**](Query.md)| Query | 
 **optional** | ***CreateQueryOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateQueryOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Requested fields. | 

### Return type

[**Query**](Query.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateQueryTask

> QueryTask CreateQueryTask(ctx, body, optional)

Run Query Async

### Create an async query task  Creates a query task (job) to run a previously created query asynchronously. Returns a Query Task ID.  Use [query_task(query_task_id)](#!/Query/query_task) to check the execution status of the query task. After the query task status reaches \"Complete\", use [query_task_results(query_task_id)](#!/Query/query_task_results) to fetch the results of the query. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**CreateQueryTask**](CreateQueryTask.md)| Query parameters | 
 **optional** | ***CreateQueryTaskOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateQueryTaskOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int64**| Row limit (may override the limit in the saved query). | 
 **applyFormatting** | **optional.Bool**| Apply model-specified formatting to each result. | 
 **applyVis** | **optional.Bool**| Apply visualization options to results. | 
 **cache** | **optional.Bool**| Get results from cache if available. | 
 **imageWidth** | **optional.Int64**| Render width for image formats. | 
 **imageHeight** | **optional.Int64**| Render height for image formats. | 
 **generateDrillLinks** | **optional.Bool**| Generate drill links (only applicable to &#39;json_detail&#39; format. | 
 **forceProduction** | **optional.Bool**| Force use of production models even if the user is in development mode. | 
 **cacheOnly** | **optional.Bool**| Retrieve any results from cache even if the results have expired. | 
 **pathPrefix** | **optional.String**| Prefix to use for drill links (url encoded). | 
 **rebuildPdts** | **optional.Bool**| Rebuild PDTS used in query. | 
 **serverTableCalcs** | **optional.Bool**| Perform table calculations on query results | 
 **fields** | **optional.String**| Requested fields | 

### Return type

[**QueryTask**](QueryTask.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSqlQuery

> SqlQuery CreateSqlQuery(ctx, body)

Create SQL Runner Query

### Create a SQL Runner Query  Either the `connection_name` or `model_name` parameter MUST be provided. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**SqlQueryCreate**](SqlQueryCreate.md)| SQL Runner Query | 

### Return type

[**SqlQuery**](SqlQuery.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KillQuery

> string KillQuery(ctx, queryTaskId)

Kill Running Query

Kill a query with a specific query_task_id. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**queryTaskId** | **string**| Query task id. | 

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


## MergeQuery

> MergeQuery MergeQuery(ctx, mergeQueryId, optional)

Get Merge Query

### Get Merge Query  Returns a merge query object given its id. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mergeQueryId** | **string**| Merge Query Id | 
 **optional** | ***MergeQueryOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MergeQueryOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Requested fields | 

### Return type

[**MergeQuery**](MergeQuery.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Query

> Query Query(ctx, queryId, optional)

Get Query

### Get a previously created query by id.  A Looker query object includes the various parameters that define a database query that has been run or could be run in the future. These parameters include: model, view, fields, filters, pivots, etc. Query *results* are not part of the query object.  Query objects are unique and immutable. Query objects are created automatically in Looker as users explore data. Looker does not delete them; they become part of the query history. When asked to create a query for any given set of parameters, Looker will first try to find an existing query object with matching parameters and will only create a new object when an appropriate object can not be found.  This 'get' method is used to get the details about a query for a given id. See the other methods here to 'create' and 'run' queries.  Note that some fields like 'filter_config' and 'vis_config' etc are specific to how the Looker UI builds queries and visualizations and are not generally useful for API use. They are not required when creating new queries and can usually just be ignored.  

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**queryId** | **int64**| Id of query | 
 **optional** | ***QueryOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a QueryOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Requested fields. | 

### Return type

[**Query**](Query.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryForSlug

> Query QueryForSlug(ctx, slug, optional)

Get Query for Slug

### Get the query for a given query slug.  This returns the query for the 'slug' in a query share URL.  The 'slug' is a randomly chosen short string that is used as an alternative to the query's id value for use in URLs etc. This method exists as a convenience to help you use the API to 'find' queries that have been created using the Looker UI.  You can use the Looker explore page to build a query and then choose the 'Share' option to show the share url for the query. Share urls generally look something like 'https://looker.yourcompany/x/vwGSbfc'. The trailing 'vwGSbfc' is the share slug. You can pass that string to this api method to get details about the query. Those details include the 'id' that you can use to run the query. Or, you can copy the query body (perhaps with your own modification) and use that as the basis to make/run new queries.  This will also work with slugs from Looker explore urls like 'https://looker.yourcompany/explore/ecommerce/orders?qid=aogBgL6o3cKK1jN3RoZl5s'. In this case 'aogBgL6o3cKK1jN3RoZl5s' is the slug. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string**| Slug of query | 
 **optional** | ***QueryForSlugOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a QueryForSlugOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Requested fields. | 

### Return type

[**Query**](Query.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryTask

> QueryTask QueryTask(ctx, queryTaskId, optional)

Get Async Query Info

### Get Query Task details  Use this function to check the status of an async query task. After the status reaches \"Complete\", you can call [query_task_results(query_task_id)](#!/Query/query_task_results) to retrieve the results of the query.  Use [create_query_task()](#!/Query/create_query_task) to create an async query task. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**queryTaskId** | **string**| ID of the Query Task | 
 **optional** | ***QueryTaskOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a QueryTaskOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Requested fields. | 

### Return type

[**QueryTask**](QueryTask.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryTaskMultiResults

> map[string]string QueryTaskMultiResults(ctx, queryTaskIds)

Get Multiple Async Query Results

### Fetch results of multiple async queries  Returns the results of multiple async queries in one request.  For Query Tasks that are not completed, the response will include the execution status of the Query Task but will not include query results. Query Tasks whose results have expired will have a status of 'expired'. If the user making the API request does not have sufficient privileges to view a Query Task result, the result will have a status of 'missing' 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**queryTaskIds** | [**[]string**](string.md)| List of Query Task IDs | 

### Return type

**map[string]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryTaskResults

> string QueryTaskResults(ctx, queryTaskId)

Get Async Query Results

### Get Async Query Results  Returns the results of an async query task if the query has completed.  If the query task is still running or waiting to run, this function returns 204 No Content.  If the query task ID is invalid or the cached results of the query task have expired, this function returns 404 Not Found.  Use [query_task(query_task_id)](#!/Query/query_task) to check the execution status of the query task Call query_task_results only after the query task status reaches \"Complete\".  You can also use [query_task_multi_results()](#!/Query/query_task_multi_results) retrieve the results of multiple async query tasks at the same time.  #### SQL Error Handling: If the query fails due to a SQL db error, how this is communicated depends on the result_format you requested in `create_query_task()`.  For `json_detail` result_format: `query_task_results()` will respond with HTTP status '200 OK' and db SQL error info will be in the `errors` property of the response object. The 'data' property will be empty.  For all other result formats: `query_task_results()` will respond with HTTP status `400 Bad Request` and some db SQL error info will be in the message of the 400 error response, but not as detailed as expressed in `json_detail.errors`. These data formats can only carry row data, and error info is not row data. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**queryTaskId** | **string**| ID of the Query Task | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RunInlineQuery

> string RunInlineQuery(ctx, resultFormat, body, optional)

Run Inline Query

### Run the query that is specified inline in the posted body.  This allows running a query as defined in json in the posted body. This combines the two actions of posting & running a query into one step.  Here is an example body in json: ``` {   \"model\":\"thelook\",   \"view\":\"inventory_items\",   \"fields\":[\"category.name\",\"inventory_items.days_in_inventory_tier\",\"products.count\"],   \"filters\":{\"category.name\":\"socks\"},   \"sorts\":[\"products.count desc 0\"],   \"limit\":\"500\",   \"query_timezone\":\"America/Los_Angeles\" } ```  When using the Ruby SDK this would be passed as a Ruby hash like: ``` {  :model=>\"thelook\",  :view=>\"inventory_items\",  :fields=>   [\"category.name\",    \"inventory_items.days_in_inventory_tier\",    \"products.count\"],  :filters=>{:\"category.name\"=>\"socks\"},  :sorts=>[\"products.count desc 0\"],  :limit=>\"500\",  :query_timezone=>\"America/Los_Angeles\", } ```  This will return the result of running the query in the format specified by the 'result_format' parameter.  Supported formats:  | result_format | Description | :-----------: | :--- | | json | Plain json | json_detail | Row data plus metadata describing the fields, pivots, table calcs, and other aspects of the query | csv | Comma separated values with a header | txt | Tab separated values with a header | html | Simple html | md | Simple markdown | xlsx | MS Excel spreadsheet | sql | Returns the generated SQL rather than running the query | png | A PNG image of the visualization of the query | jpg | A JPG image of the visualization of the query   

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resultFormat** | **string**| Format of result | 
**body** | [**Query**](Query.md)| inline query | 
 **optional** | ***RunInlineQueryOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a RunInlineQueryOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **optional.Int64**| Row limit (may override the limit in the saved query). | 
 **applyFormatting** | **optional.Bool**| Apply model-specified formatting to each result. | 
 **applyVis** | **optional.Bool**| Apply visualization options to results. | 
 **cache** | **optional.Bool**| Get results from cache if available. | 
 **imageWidth** | **optional.Int64**| Render width for image formats. | 
 **imageHeight** | **optional.Int64**| Render height for image formats. | 
 **generateDrillLinks** | **optional.Bool**| Generate drill links (only applicable to &#39;json_detail&#39; format. | 
 **forceProduction** | **optional.Bool**| Force use of production models even if the user is in development mode. | 
 **cacheOnly** | **optional.Bool**| Retrieve any results from cache even if the results have expired. | 
 **pathPrefix** | **optional.String**| Prefix to use for drill links (url encoded). | 
 **rebuildPdts** | **optional.Bool**| Rebuild PDTS used in query. | 
 **serverTableCalcs** | **optional.Bool**| Perform table calculations on query results | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text, application/json, image/png, image/jpeg

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RunQuery

> string RunQuery(ctx, queryId, resultFormat, optional)

Run Query

### Run a saved query.  This runs a previously saved query. You can use this on a query that was generated in the Looker UI or one that you have explicitly created using the API. You can also use a query 'id' from a saved 'Look'.  The 'result_format' parameter specifies the desired structure and format of the response.  Supported formats:  | result_format | Description | :-----------: | :--- | | json | Plain json | json_detail | Row data plus metadata describing the fields, pivots, table calcs, and other aspects of the query | csv | Comma separated values with a header | txt | Tab separated values with a header | html | Simple html | md | Simple markdown | xlsx | MS Excel spreadsheet | sql | Returns the generated SQL rather than running the query | png | A PNG image of the visualization of the query | jpg | A JPG image of the visualization of the query   

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**queryId** | **int64**| Id of query | 
**resultFormat** | **string**| Format of result | 
 **optional** | ***RunQueryOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a RunQueryOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **optional.Int64**| Row limit (may override the limit in the saved query). | 
 **applyFormatting** | **optional.Bool**| Apply model-specified formatting to each result. | 
 **applyVis** | **optional.Bool**| Apply visualization options to results. | 
 **cache** | **optional.Bool**| Get results from cache if available. | 
 **imageWidth** | **optional.Int64**| Render width for image formats. | 
 **imageHeight** | **optional.Int64**| Render height for image formats. | 
 **generateDrillLinks** | **optional.Bool**| Generate drill links (only applicable to &#39;json_detail&#39; format. | 
 **forceProduction** | **optional.Bool**| Force use of production models even if the user is in development mode. | 
 **cacheOnly** | **optional.Bool**| Retrieve any results from cache even if the results have expired. | 
 **pathPrefix** | **optional.String**| Prefix to use for drill links (url encoded). | 
 **rebuildPdts** | **optional.Bool**| Rebuild PDTS used in query. | 
 **serverTableCalcs** | **optional.Bool**| Perform table calculations on query results | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text, application/json, image/png, image/jpeg

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RunSqlQuery

> string RunSqlQuery(ctx, slug, resultFormat, optional)

Run SQL Runner Query

Execute a SQL Runner query in a given result_format.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string**| slug of query | 
**resultFormat** | **string**| Format of result, options are: [\&quot;json\&quot;, \&quot;json_detail\&quot;, \&quot;json_fe\&quot;, \&quot;csv\&quot;, \&quot;html\&quot;, \&quot;md\&quot;, \&quot;txt\&quot;, \&quot;xlsx\&quot;, \&quot;gsxml\&quot;, \&quot;json_label\&quot;] | 
 **optional** | ***RunSqlQueryOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a RunSqlQueryOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **download** | **optional.String**| Defaults to false. If set to true, the HTTP response will have content-disposition and other headers set to make the HTTP response behave as a downloadable attachment instead of as inline content. | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text, application/json, image/png, image/jpeg

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RunUrlEncodedQuery

> string RunUrlEncodedQuery(ctx, modelName, viewName, resultFormat)

Run Url Encoded Query

### Run an URL encoded query.  This requires the caller to encode the specifiers for the query into the URL query part using Looker-specific syntax as explained below.  Generally, you would want to use one of the methods that takes the parameters as json in the POST body for creating and/or running queries. This method exists for cases where one really needs to encode the parameters into the URL of a single 'GET' request. This matches the way that the Looker UI formats 'explore' URLs etc.  The parameters here are very similar to the json body formatting except that the filter syntax is tricky. Unfortunately, this format makes this method not currently callible via the 'Try it out!' button in this documentation page. But, this is callable  when creating URLs manually or when using the Looker SDK.  Here is an example inline query URL:  ``` https://looker.mycompany.com:19999/api/3.0/queries/models/thelook/views/inventory_items/run/json?fields=category.name,inventory_items.days_in_inventory_tier,products.count&f[category.name]=socks&sorts=products.count+desc+0&limit=500&query_timezone=America/Los_Angeles ```  When invoking this endpoint with the Ruby SDK, pass the query parameter parts as a hash. The hash to match the above would look like:  ```ruby query_params = {   :fields => \"category.name,inventory_items.days_in_inventory_tier,products.count\",   :\"f[category.name]\" => \"socks\",   :sorts => \"products.count desc 0\",   :limit => \"500\",   :query_timezone => \"America/Los_Angeles\" } response = ruby_sdk.run_url_encoded_query('thelook','inventory_items','json', query_params)  ```  Again, it is generally easier to use the variant of this method that passes the full query in the POST body. This method is available for cases where other alternatives won't fit the need.  Supported formats:  | result_format | Description | :-----------: | :--- | | json | Plain json | json_detail | Row data plus metadata describing the fields, pivots, table calcs, and other aspects of the query | csv | Comma separated values with a header | txt | Tab separated values with a header | html | Simple html | md | Simple markdown | xlsx | MS Excel spreadsheet | sql | Returns the generated SQL rather than running the query | png | A PNG image of the visualization of the query | jpg | A JPG image of the visualization of the query   

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**modelName** | **string**| Model name | 
**viewName** | **string**| View name | 
**resultFormat** | **string**| Format of result | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text, application/json, image/png, image/jpeg

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SqlQuery

> SqlQuery SqlQuery(ctx, slug)

Get SQL Runner Query

Get a SQL Runner query.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string**| slug of query | 

### Return type

[**SqlQuery**](SqlQuery.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

