/*
 * Looker API 3.1 Reference
 *
 * ### Authorization  The Looker API uses Looker **API3** credentials for authorization and access control. Looker admins can create API3 credentials on Looker's **Admin/Users** page. Pass API3 credentials to the **_/login** endpoint to obtain a temporary access_token. Include that access_token in the Authorization header of Looker API requests. For details, see [Looker API Authorization](https://looker.com/docs/r/api/authorization)  ### Client SDKs  The Looker API is a RESTful system that should be usable by any programming language capable of making HTTPS requests. Client SDKs for a variety of programming languages can be generated from the Looker API's Swagger JSON metadata to streamline use of the Looker API in your applications. A client SDK for Ruby is available as an example. For more information, see [Looker API Client SDKs](https://looker.com/docs/r/api/client_sdks)  ### Try It Out!  The 'api-docs' page served by the Looker instance includes 'Try It Out!' buttons for each API method. After logging in with API3 credentials, you can use the \"Try It Out!\" buttons to call the API directly from the documentation page to interactively explore API features and responses.  Note! With great power comes great responsibility: The \"Try It Out!\" button makes API calls to your live Looker instance. Be especially careful with destructive API operations such as `delete_user` or similar. There is no \"undo\" for API operations.  ### Versioning  Future releases of Looker will expand this API release-by-release to securely expose more and more of the core power of Looker to API client applications. API endpoints marked as \"beta\" may receive breaking changes without warning (but we will try to avoid doing that). Stable (non-beta) API endpoints should not receive breaking changes in future releases. For more information, see [Looker API Versioning](https://looker.com/docs/r/api/versioning)  ### In This Release  The following are a few examples of noteworthy items that have changed between API 3.0 and API 3.1. For more comprehensive coverage of API changes, please see the release notes for your Looker release.  ### Examples of new things added in API 3.1 (compared to API 3.0):  * [Dashboard construction](#!/3.1/Dashboard/) APIs * [Themes](#!/3.1/Theme/) and [custom color collections](#!/3.1/ColorCollection) APIs * Create and run [SQL Runner](#!/3.1/Query/run_sql_query) queries * Create and run [merged results](#!/3.1/Query/create_merge_query) queries * Create and modify [dashboard filters](#!/3.1/Dashboard/create_dashboard_filter) * Create and modify [password requirements](#!/3.1/Auth/password_config)  ### Deprecated in API 3.0  The following functions and properties have been deprecated in API 3.0.  They continue to exist and work in API 3.0 for the next several Looker releases but they have not been carried forward to API 3.1:  * Dashboard Prefetch functions * User access_filter functions * User API 1.0 credentials functions * Space.is_root and Space.is_user_root properties. Use Space.is_shared_root and Space.is_users_root instead.  ### Semantic changes in API 3.1:  * [all_looks()](#!/3.1/Look/all_looks) no longer includes soft-deleted looks, matching [all_dashboards()](#!/3.1/Dashboard/all_dashboards) behavior. You can find soft-deleted looks using [search_looks()](#!/3.1/Look/search_looks) with the `deleted` param set to True. * [all_spaces()](#!/3.1/Space/all_spaces) no longer includes duplicate items * [search_users()](#!/3.1/User/search_users) no longer accepts Y,y,1,0,N,n for Boolean params, only \"true\" and \"false\". * For greater client and network compatibility, [render_task_results](#!/3.1/RenderTask/render_task_results) now returns HTTP status **202 Accepted** instead of HTTP status **102 Processing** * [all_running_queries()](#!/3.1/Query/all_running_queries) and [kill_query](#!/3.1/Query/kill_query) functions have moved into the [Query](#!/3.1/Query/) function group.   If you have application code which relies on the old behavior of the APIs above, you may continue using the API 3.0 functions in this Looker release. We strongly suggest you update your code to use API 3.1 analogs as soon as possible.  
 *
 * API version: 3.1.0
 * Contact: support@looker.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// Query struct for Query
type Query struct {
	// Operations the current user is able to perform on this object
	Can map[string]bool `json:"can,omitempty"`
	// Unique Id
	Id int64 `json:"id,omitempty"`
	// Model
	Model string `json:"model"`
	// Explore Name
	View string `json:"view"`
	// Fields
	Fields []string `json:"fields,omitempty"`
	// Pivots
	Pivots []string `json:"pivots,omitempty"`
	// Fill Fields
	FillFields []string `json:"fill_fields,omitempty"`
	// Filters
	Filters map[string]string `json:"filters,omitempty"`
	// Filter Expression
	FilterExpression string `json:"filter_expression,omitempty"`
	// Sorting for the query results. Use the format `[\"view.field\", ...]` to sort on fields in ascending order. Use the format `[\"view.field desc\", ...]` to sort on fields in descending order. Use `[\"__UNSORTED__\"]` (2 underscores before and after) to disable sorting entirely. Empty sorts `[]` will trigger a default sort.
	Sorts []string `json:"sorts,omitempty"`
	// Limit
	Limit string `json:"limit,omitempty"`
	// Column Limit
	ColumnLimit string `json:"column_limit,omitempty"`
	// Total
	Total bool `json:"total,omitempty"`
	// Raw Total
	RowTotal string `json:"row_total,omitempty"`
	// Fields on which to run subtotals
	Subtotals []string `json:"subtotals,omitempty"`
	// Runtime
	Runtime float64 `json:"runtime,omitempty"`
	// Visualization configuration properties. These properties are typically opaque and differ based on the type of visualization used. There is no specified set of allowed keys. The values can be any type supported by JSON. A \"type\" key with a string value is often present, and is used by Looker to determine which visualization to present. Visualizations ignore unknown vis_config properties.
	VisConfig map[string]string `json:"vis_config,omitempty"`
	// The filter_config represents the state of the filter UI on the explore page for a given query. When running a query via the Looker UI, this parameter takes precedence over \"filters\". When creating a query or modifying an existing query, \"filter_config\" should be set to null. Setting it to any other value could cause unexpected filtering behavior. The format should be considered opaque.
	FilterConfig map[string]string `json:"filter_config,omitempty"`
	// Visible UI Sections
	VisibleUiSections string `json:"visible_ui_sections,omitempty"`
	// Slug
	Slug string `json:"slug,omitempty"`
	// Dynamic Fields
	DynamicFields string `json:"dynamic_fields,omitempty"`
	// Client Id: used to generate shortened explore URLs. If set by client, must be a unique 22 character alphanumeric string. Otherwise one will be generated.
	ClientId string `json:"client_id,omitempty"`
	// Share Url
	ShareUrl string `json:"share_url,omitempty"`
	// Expanded Share Url
	ExpandedShareUrl string `json:"expanded_share_url,omitempty"`
	// Expanded Url
	Url string `json:"url,omitempty"`
	// Query Timezone
	QueryTimezone string `json:"query_timezone,omitempty"`
	// Has Table Calculations
	HasTableCalculations bool `json:"has_table_calculations,omitempty"`
}
