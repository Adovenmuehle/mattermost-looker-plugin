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
// LookmlModelExplore struct for LookmlModelExplore
type LookmlModelExplore struct {
	// Fully qualified name model plus explore name
	Id string `json:"id,omitempty"`
	// Explore name
	Name string `json:"name,omitempty"`
	// Description
	Description string `json:"description,omitempty"`
	// Label
	Label string `json:"label,omitempty"`
	// Scopes
	Scopes []string `json:"scopes,omitempty"`
	// Can Total
	CanTotal bool `json:"can_total,omitempty"`
	// Can Save
	CanSave bool `json:"can_save,omitempty"`
	// Can Explain
	CanExplain bool `json:"can_explain,omitempty"`
	// Can pivot in the DB
	CanPivotInDb bool `json:"can_pivot_in_db,omitempty"`
	// Can use subtotals
	CanSubtotal bool `json:"can_subtotal,omitempty"`
	// Has timezone support
	HasTimezoneSupport bool `json:"has_timezone_support,omitempty"`
	// Cost estimates supported
	SupportsCostEstimate bool `json:"supports_cost_estimate,omitempty"`
	// Connection name
	ConnectionName string `json:"connection_name,omitempty"`
	// How nulls are sorted, possible values are \"low\", \"high\", \"first\" and \"last\"
	NullSortTreatment string `json:"null_sort_treatment,omitempty"`
	// List of model source files
	Files []string `json:"files,omitempty"`
	// Primary source_file file
	SourceFile string `json:"source_file,omitempty"`
	// Name of project
	ProjectName string `json:"project_name,omitempty"`
	// Name of model
	ModelName string `json:"model_name,omitempty"`
	// Name of view
	ViewName string `json:"view_name,omitempty"`
	// Is hidden
	Hidden bool `json:"hidden,omitempty"`
	// A sql_table_name expression that defines what sql table the view/explore maps onto. Example: \"prod_orders2 AS orders\" in a view named orders.
	SqlTableName string `json:"sql_table_name,omitempty"`
	// (DEPRECATED) Array of access filter field names
	AccessFilterFields []string `json:"access_filter_fields,omitempty"`
	// Access filters
	AccessFilters []LookmlModelExploreAccessFilter `json:"access_filters,omitempty"`
	// Aliases
	Aliases []LookmlModelExploreAlias `json:"aliases,omitempty"`
	// Always filter
	AlwaysFilter []LookmlModelExploreAlwaysFilter `json:"always_filter,omitempty"`
	// Conditionally filter
	ConditionallyFilter []LookmlModelExploreConditionallyFilter `json:"conditionally_filter,omitempty"`
	// Array of index fields
	IndexFields []string `json:"index_fields,omitempty"`
	// Sets
	Sets []LookmlModelExploreSet `json:"sets,omitempty"`
	// An array of arbitrary string tags provided in the model for this explore.
	Tags []string `json:"tags,omitempty"`
	// Errors
	Errors []LookmlModelExploreError `json:"errors,omitempty"`
	Fields LookmlModelExploreFieldset `json:"fields,omitempty"`
	// Views joined into this explore
	Joins []LookmlModelExploreJoins `json:"joins,omitempty"`
	// Label used to group explores in the navigation menus
	GroupLabel string `json:"group_label,omitempty"`
	// An array of items describing which custom measure types are supported for creating a custom measure 'baed_on' each possible dimension type.
	SupportedMeasureTypes []LookmlModelExploreSupportedMeasureType `json:"supported_measure_types,omitempty"`
}
