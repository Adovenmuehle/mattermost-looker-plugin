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
// Integration struct for Integration
type Integration struct {
	// Operations the current user is able to perform on this object
	Can map[string]bool `json:"can,omitempty"`
	// ID of the integration.
	Id string `json:"id,omitempty"`
	// ID of the integration hub.
	IntegrationHubId int64 `json:"integration_hub_id,omitempty"`
	// Label for the integration.
	Label string `json:"label,omitempty"`
	// Description of the integration.
	Description string `json:"description,omitempty"`
	// Whether the integration is available to users.
	Enabled bool `json:"enabled,omitempty"`
	// Array of params for the integration.
	Params []IntegrationParam `json:"params,omitempty"`
	// A list of data formats the integration supports. If unspecified, the default is all data formats. Valid values are: \"txt\", \"csv\", \"inline_json\", \"json\", \"json_label\", \"json_detail\", \"json_detail_lite_stream\", \"xlsx\", \"html\", \"wysiwyg_pdf\", \"assembled_pdf\", \"wysiwyg_png\", \"csv_zip\".
	SupportedFormats []string `json:"supported_formats,omitempty"`
	// A list of action types the integration supports. Valid values are: \"cell\", \"query\", \"dashboard\".
	SupportedActionTypes []string `json:"supported_action_types,omitempty"`
	// A list of formatting options the integration supports. If unspecified, defaults to all formats. Valid values are: \"formatted\", \"unformatted\".
	SupportedFormattings []string `json:"supported_formattings,omitempty"`
	// A list of visualization formatting options the integration supports. If unspecified, defaults to all formats. Valid values are: \"apply\", \"noapply\".
	SupportedVisualizationFormattings []string `json:"supported_visualization_formattings,omitempty"`
	// A list of all the download mechanisms the integration supports. The order of values is not significant: Looker will select the most appropriate supported download mechanism for a given query. The integration must ensure it can handle any of the mechanisms it claims to support. If unspecified, this defaults to all download setting values. Valid values are: \"push\", \"url\".
	SupportedDownloadSettings []string `json:"supported_download_settings,omitempty"`
	// URL to an icon for the integration.
	IconUrl string `json:"icon_url,omitempty"`
	// Whether the integration uses oauth.
	UsesOauth bool `json:"uses_oauth,omitempty"`
	// A list of descriptions of required fields that this integration is compatible with. If there are multiple entries in this list, the integration requires more than one field. If unspecified, no fields will be required.
	RequiredFields []IntegrationRequiredField `json:"required_fields,omitempty"`
	// Whether the integration uses delegate oauth, which allows federation between an integration installation scope specific entity (like org, group, and team, etc.) and Looker.
	DelegateOauth bool `json:"delegate_oauth,omitempty"`
	// Whether the integration is available to users.
	InstalledDelegateOauthTargets []string `json:"installed_delegate_oauth_targets,omitempty"`
}
