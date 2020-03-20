# \ProjectApi

All URIs are relative to *https://mattermost.looker.com:19999/api/3.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AllGitBranches**](ProjectApi.md#AllGitBranches) | **Get** /projects/{project_id}/git_branches | Get All Git Branches
[**AllGitConnectionTests**](ProjectApi.md#AllGitConnectionTests) | **Get** /projects/{project_id}/git_connection_tests | Get All Git Connection Tests
[**AllLookmlTests**](ProjectApi.md#AllLookmlTests) | **Get** /projects/{project_id}/lookml_tests | Get All LookML Tests
[**AllProjectFiles**](ProjectApi.md#AllProjectFiles) | **Get** /projects/{project_id}/files | Get All Project Files
[**AllProjects**](ProjectApi.md#AllProjects) | **Get** /projects | Get All Projects
[**CreateGitBranch**](ProjectApi.md#CreateGitBranch) | **Post** /projects/{project_id}/git_branch | Checkout New Git Branch
[**CreateGitDeployKey**](ProjectApi.md#CreateGitDeployKey) | **Post** /projects/{project_id}/git/deploy_key | Create Deploy Key
[**CreateProject**](ProjectApi.md#CreateProject) | **Post** /projects | Create Project
[**DeleteGitBranch**](ProjectApi.md#DeleteGitBranch) | **Delete** /projects/{project_id}/git_branch/{branch_name} | Delete a Git Branch
[**DeleteRepositoryCredential**](ProjectApi.md#DeleteRepositoryCredential) | **Delete** /projects/{root_project_id}/credential/{credential_id} | Delete Repository Credential
[**DeployToProduction**](ProjectApi.md#DeployToProduction) | **Post** /projects/{project_id}/deploy_to_production | Deploy To Production
[**FindGitBranch**](ProjectApi.md#FindGitBranch) | **Get** /projects/{project_id}/git_branch/{branch_name} | Find a Git Branch
[**GetAllRepositoryCredentials**](ProjectApi.md#GetAllRepositoryCredentials) | **Get** /projects/{root_project_id}/credentials | Get All Repository Credentials
[**GitBranch**](ProjectApi.md#GitBranch) | **Get** /projects/{project_id}/git_branch | Get Active Git Branch
[**GitDeployKey**](ProjectApi.md#GitDeployKey) | **Get** /projects/{project_id}/git/deploy_key | Git Deploy Key
[**Manifest**](ProjectApi.md#Manifest) | **Get** /projects/{project_id}/manifest | Get Manifest
[**Project**](ProjectApi.md#Project) | **Get** /projects/{project_id} | Get Project
[**ProjectFile**](ProjectApi.md#ProjectFile) | **Get** /projects/{project_id}/files/file | Get Project File
[**ProjectValidationResults**](ProjectApi.md#ProjectValidationResults) | **Get** /projects/{project_id}/validate | Cached Project Validation Results
[**ProjectWorkspace**](ProjectApi.md#ProjectWorkspace) | **Get** /projects/{project_id}/current_workspace | Get Project Workspace
[**ResetProjectToProduction**](ProjectApi.md#ResetProjectToProduction) | **Post** /projects/{project_id}/reset_to_production | Reset To Production
[**ResetProjectToRemote**](ProjectApi.md#ResetProjectToRemote) | **Post** /projects/{project_id}/reset_to_remote | Reset To Remote
[**RunGitConnectionTest**](ProjectApi.md#RunGitConnectionTest) | **Get** /projects/{project_id}/git_connection_tests/{test_id} | Run Git Connection Test
[**RunLookmlTest**](ProjectApi.md#RunLookmlTest) | **Get** /projects/{project_id}/lookml_tests/run | Run LookML Test
[**UpdateGitBranch**](ProjectApi.md#UpdateGitBranch) | **Put** /projects/{project_id}/git_branch | Update Project Git Branch
[**UpdateProject**](ProjectApi.md#UpdateProject) | **Patch** /projects/{project_id} | Update Project
[**UpdateRepositoryCredential**](ProjectApi.md#UpdateRepositoryCredential) | **Put** /projects/{root_project_id}/credential/{credential_id} | Create Repository Credential
[**ValidateProject**](ProjectApi.md#ValidateProject) | **Post** /projects/{project_id}/validate | Validate Project



## AllGitBranches

> []GitBranch AllGitBranches(ctx, projectId)

Get All Git Branches

### Get All Git Branches  Returns a list of git branches in the project repository 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 

### Return type

[**[]GitBranch**](GitBranch.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AllGitConnectionTests

> []GitConnectionTest AllGitConnectionTests(ctx, projectId, optional)

Get All Git Connection Tests

### Get All Git Connection Tests  dev mode required.   - Call `update_session` to select the 'dev' workspace.  Returns a list of tests which can be run against a project's (or the dependency project for the provided remote_url) git connection. Call [Run Git Connection Test](#!/Project/run_git_connection_test) to execute each test in sequence.  Tests are ordered by increasing specificity. Tests should be run in the order returned because later tests require functionality tested by tests earlier in the test list.  For example, a late-stage test for write access is meaningless if connecting to the git server (an early test) is failing. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
 **optional** | ***AllGitConnectionTestsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AllGitConnectionTestsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **remoteUrl** | **optional.String**| (Optional: leave blank for root project) The remote url for remote dependency to test. | 

### Return type

[**[]GitConnectionTest**](GitConnectionTest.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AllLookmlTests

> []LookmlTest AllLookmlTests(ctx, projectId, optional)

Get All LookML Tests

### Get All LookML Tests  Returns a list of tests which can be run to validate a project's LookML code and/or the underlying data, optionally filtered by the file id. Call [Run LookML Test](#!/Project/run_lookml_test) to execute tests. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
 **optional** | ***AllLookmlTestsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AllLookmlTestsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fileId** | **optional.String**| File Id | 

### Return type

[**[]LookmlTest**](LookmlTest.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AllProjectFiles

> []ProjectFile AllProjectFiles(ctx, projectId, optional)

Get All Project Files

### Get All Project Files  Returns a list of the files in the project 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
 **optional** | ***AllProjectFilesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AllProjectFilesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Requested fields | 

### Return type

[**[]ProjectFile**](ProjectFile.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AllProjects

> []Project AllProjects(ctx, optional)

Get All Projects

### Get All Projects  Returns all projects visible to the current user 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AllProjectsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AllProjectsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **optional.String**| Requested fields | 

### Return type

[**[]Project**](Project.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateGitBranch

> GitBranch CreateGitBranch(ctx, projectId, body)

Checkout New Git Branch

### Create and Checkout a Git Branch  Creates and checks out a new branch in the given project repository Only allowed in development mode   - Call `update_session` to select the 'dev' workspace.  Optionally specify a branch name, tag name or commit SHA as the start point in the ref field.   If no ref is specified, HEAD of the current branch will be used as the start point for the new branch.  

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**body** | [**GitBranch**](GitBranch.md)| Git Branch | 

### Return type

[**GitBranch**](GitBranch.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateGitDeployKey

> string CreateGitDeployKey(ctx, projectId)

Create Deploy Key

### Create Git Deploy Key  Create a public/private key pair for authenticating ssh git requests from Looker to a remote git repository for a particular Looker project.  Returns the public key of the generated ssh key pair.  Copy this public key to your remote git repository's ssh keys configuration so that the remote git service can validate and accept git requests from the Looker server. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateProject

> Project CreateProject(ctx, body)

Create Project

### Create A Project  dev mode required. - Call `update_session` to select the 'dev' workspace.  `name` is required. `git_remote_url` is not allowed. To configure Git for the newly created project, follow the instructions in `update_project`.  

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**Project**](Project.md)| Project | 

### Return type

[**Project**](Project.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGitBranch

> string DeleteGitBranch(ctx, projectId, branchName)

Delete a Git Branch

### Delete the specified Git Branch  Delete git branch specified in branch_name path param from local and remote of specified project repository 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**branchName** | **string**| Branch Name | 

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


## DeleteRepositoryCredential

> string DeleteRepositoryCredential(ctx, rootProjectId, credentialId)

Delete Repository Credential

### Repository Credential for a remote dependency  Admin required.  `root_project_id` is required. `credential_id` is required. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rootProjectId** | **string**| Root Project Id | 
**credentialId** | **string**| Credential Id | 

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


## DeployToProduction

> string DeployToProduction(ctx, projectId)

Deploy To Production

### Deploy LookML from this Development Mode Project to Production  Git must have been configured, must be in dev mode and deploy permission required  Deploy is a two / three step process 1. Push commits in current branch of dev mode project to the production branch (origin/master).    Note a. This step is skipped in read-only projects.    Note b. If this step is unsuccessful for any reason (e.g. rejected non-fastforward because production branch has              commits not in current branch), subsequent steps will be skipped. 2. If this is the first deploy of this project, create the production project with git repository. 3. Pull the production branch into the production project.  

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Id of project | 

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


## FindGitBranch

> GitBranch FindGitBranch(ctx, projectId, branchName)

Find a Git Branch

### Get the specified Git Branch  Returns the git branch specified in branch_name path param if it exists in the given project repository 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**branchName** | **string**| Branch Name | 

### Return type

[**GitBranch**](GitBranch.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllRepositoryCredentials

> []RepositoryCredential GetAllRepositoryCredentials(ctx, rootProjectId)

Get All Repository Credentials

### Get all Repository Credentials for a project  `root_project_id` is required. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rootProjectId** | **string**| Root Project Id | 

### Return type

[**[]RepositoryCredential**](RepositoryCredential.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GitBranch

> GitBranch GitBranch(ctx, projectId)

Get Active Git Branch

### Get the Current Git Branch  Returns the git branch currently checked out in the given project repository 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 

### Return type

[**GitBranch**](GitBranch.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GitDeployKey

> string GitDeployKey(ctx, projectId)

Git Deploy Key

### Git Deploy Key  Returns the ssh public key previously created for a project's git repository. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Manifest

> Manifest Manifest(ctx, projectId)

Get Manifest

### Get A Projects Manifest object  Returns the project with the given project id 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 

### Return type

[**Manifest**](Manifest.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Project

> Project Project(ctx, projectId, optional)

Get Project

### Get A Project  Returns the project with the given project id 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
 **optional** | ***ProjectOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ProjectOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Requested fields | 

### Return type

[**Project**](Project.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectFile

> ProjectFile ProjectFile(ctx, projectId, fileId, optional)

Get Project File

### Get Project File Info  Returns information about a file in the project 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**fileId** | **string**| File Id | 
 **optional** | ***ProjectFileOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ProjectFileOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **optional.String**| Requested fields | 

### Return type

[**ProjectFile**](ProjectFile.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectValidationResults

> ProjectValidationCache ProjectValidationResults(ctx, projectId, optional)

Cached Project Validation Results

### Get Cached Project Validation Results  Returns the cached results of a previous project validation calculation, if any. Returns http status 204 No Content if no validation results exist.  Validating the content of all the files in a project can be computationally intensive for large projects. Use this API to simply fetch the results of the most recent project validation rather than revalidating the entire project from scratch.  A value of `\"stale\": true` in the response indicates that the project has changed since the cached validation results were computed. The cached validation results may no longer reflect the current state of the project. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
 **optional** | ***ProjectValidationResultsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ProjectValidationResultsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Requested fields | 

### Return type

[**ProjectValidationCache**](ProjectValidationCache.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectWorkspace

> ProjectWorkspace ProjectWorkspace(ctx, projectId, optional)

Get Project Workspace

### Get Project Workspace  Returns information about the state of the project files in the currently selected workspace 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
 **optional** | ***ProjectWorkspaceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ProjectWorkspaceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Requested fields | 

### Return type

[**ProjectWorkspace**](ProjectWorkspace.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetProjectToProduction

> string ResetProjectToProduction(ctx, projectId)

Reset To Production

### Reset a project to the revision of the project that is in production.  **DANGER** this will delete any changes that have not been pushed to a remote repository. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Id of project | 

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


## ResetProjectToRemote

> string ResetProjectToRemote(ctx, projectId)

Reset To Remote

### Reset a project development branch to the revision of the project that is on the remote.  **DANGER** this will delete any changes that have not been pushed to a remote repository. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Id of project | 

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


## RunGitConnectionTest

> GitConnectionTestResult RunGitConnectionTest(ctx, projectId, testId, optional)

Run Git Connection Test

### Run a git connection test  Run the named test on the git service used by this project (or the dependency project for the provided remote_url) and return the result. This is intended to help debug git connections when things do not work properly, to give more helpful information about why a git url is not working with Looker.  Tests should be run in the order they are returned by [Get All Git Connection Tests](#!/Project/all_git_connection_tests). 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**testId** | **string**| Test Id | 
 **optional** | ***RunGitConnectionTestOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a RunGitConnectionTestOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **remoteUrl** | **optional.String**| (Optional: leave blank for root project) The remote url for remote dependency to test. | 

### Return type

[**GitConnectionTestResult**](GitConnectionTestResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RunLookmlTest

> []LookmlTestResult RunLookmlTest(ctx, projectId, optional)

Run LookML Test

### Run LookML Tests  Runs all tests in the project, optionally filtered by file, test, and/or model. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
 **optional** | ***RunLookmlTestOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a RunLookmlTestOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fileId** | **optional.String**| File Name | 
 **test** | **optional.String**| Test Name | 
 **model** | **optional.String**| Model Name | 

### Return type

[**[]LookmlTestResult**](LookmlTestResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGitBranch

> GitBranch UpdateGitBranch(ctx, projectId, body)

Update Project Git Branch

### Checkout and/or reset --hard an existing Git Branch  Only allowed in development mode   - Call `update_session` to select the 'dev' workspace.  Checkout an existing branch if name field is different from the name of the currently checked out branch.  Optionally specify a branch name, tag name or commit SHA to which the branch should be reset.   **DANGER** hard reset will be force pushed to the remote. Unsaved changes and commits may be permanently lost.  

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**body** | [**GitBranch**](GitBranch.md)| Git Branch | 

### Return type

[**GitBranch**](GitBranch.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProject

> Project UpdateProject(ctx, projectId, body, optional)

Update Project

### Update Project Configuration  Apply changes to a project's configuration.   #### Configuring Git for a Project  To set up a Looker project with a remote git repository, follow these steps:  1. Call `update_session` to select the 'dev' workspace. 1. Call `create_git_deploy_key` to create a new deploy key for the project 1. Copy the deploy key text into the remote git repository's ssh key configuration 1. Call `update_project` to set project's `git_remote_url` ()and `git_service_name`, if necessary).  When you modify a project's `git_remote_url`, Looker connects to the remote repository to fetch metadata. The remote git repository MUST be configured with the Looker-generated deploy key for this project prior to setting the project's `git_remote_url`.  To set up a Looker project with a git repository residing on the Looker server (a 'bare' git repo): 1. Call `update_session` to select the 'dev' workspace. 1. Call `update_project` setting `git_remote_url` to nil and `git_service_name` to \"bare\".  

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
**body** | [**Project**](Project.md)| Project | 
 **optional** | ***UpdateProjectOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateProjectOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **optional.String**| Requested fields | 

### Return type

[**Project**](Project.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRepositoryCredential

> RepositoryCredential UpdateRepositoryCredential(ctx, rootProjectId, credentialId, body)

Create Repository Credential

### Configure Repository Credential for a remote dependency  Admin required.  `root_project_id` is required. `credential_id` is required.  

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rootProjectId** | **string**| Root Project Id | 
**credentialId** | **string**| Credential Id | 
**body** | [**RepositoryCredential**](RepositoryCredential.md)| Remote Project Information | 

### Return type

[**RepositoryCredential**](RepositoryCredential.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateProject

> ProjectValidation ValidateProject(ctx, projectId, optional)

Validate Project

### Validate Project  Performs lint validation of all lookml files in the project. Returns a list of errors found, if any.  Validating the content of all the files in a project can be computationally intensive for large projects. For best performance, call `validate_project(project_id)` only when you really want to recompute project validation. To quickly display the results of the most recent project validation (without recomputing), use `project_validation_results(project_id)` 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string**| Project Id | 
 **optional** | ***ValidateProjectOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ValidateProjectOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Requested fields | 

### Return type

[**ProjectValidation**](ProjectValidation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

