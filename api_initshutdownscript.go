/*
 * TrueNAS RESTful API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v2.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package truenas

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// InitshutdownscriptApiService InitshutdownscriptApi service
type InitshutdownscriptApiService service

type ApiInitshutdownscriptGetRequest struct {
	ctx _context.Context
	ApiService *InitshutdownscriptApiService
	limit *int32
	offset *int32
	count *bool
	sort *string
}

func (r ApiInitshutdownscriptGetRequest) Limit(limit int32) ApiInitshutdownscriptGetRequest {
	r.limit = &limit
	return r
}
func (r ApiInitshutdownscriptGetRequest) Offset(offset int32) ApiInitshutdownscriptGetRequest {
	r.offset = &offset
	return r
}
func (r ApiInitshutdownscriptGetRequest) Count(count bool) ApiInitshutdownscriptGetRequest {
	r.count = &count
	return r
}
func (r ApiInitshutdownscriptGetRequest) Sort(sort string) ApiInitshutdownscriptGetRequest {
	r.sort = &sort
	return r
}

func (r ApiInitshutdownscriptGetRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.InitshutdownscriptGetExecute(r)
}

/*
 * InitshutdownscriptGet Method for InitshutdownscriptGet
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiInitshutdownscriptGetRequest
 */
func (a *InitshutdownscriptApiService) InitshutdownscriptGet(ctx _context.Context) ApiInitshutdownscriptGetRequest {
	return ApiInitshutdownscriptGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 */
func (a *InitshutdownscriptApiService) InitshutdownscriptGetExecute(r ApiInitshutdownscriptGetRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InitshutdownscriptApiService.InitshutdownscriptGet")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/initshutdownscript"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.offset != nil {
		localVarQueryParams.Add("offset", parameterToString(*r.offset, ""))
	}
	if r.count != nil {
		localVarQueryParams.Add("count", parameterToString(*r.count, ""))
	}
	if r.sort != nil {
		localVarQueryParams.Add("sort", parameterToString(*r.sort, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiInitshutdownscriptIdIdDeleteRequest struct {
	ctx _context.Context
	ApiService *InitshutdownscriptApiService
	id int32
}


func (r ApiInitshutdownscriptIdIdDeleteRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.InitshutdownscriptIdIdDeleteExecute(r)
}

/*
 * InitshutdownscriptIdIdDelete Method for InitshutdownscriptIdIdDelete
 * Delete init/shutdown task of `id`.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id
 * @return ApiInitshutdownscriptIdIdDeleteRequest
 */
func (a *InitshutdownscriptApiService) InitshutdownscriptIdIdDelete(ctx _context.Context, id int32) ApiInitshutdownscriptIdIdDeleteRequest {
	return ApiInitshutdownscriptIdIdDeleteRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

/*
 * Execute executes the request
 */
func (a *InitshutdownscriptApiService) InitshutdownscriptIdIdDeleteExecute(r ApiInitshutdownscriptIdIdDeleteRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InitshutdownscriptApiService.InitshutdownscriptIdIdDelete")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/initshutdownscript/id/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiInitshutdownscriptIdIdGetRequest struct {
	ctx _context.Context
	ApiService *InitshutdownscriptApiService
	id []interface{}
	limit *int32
	offset *int32
	count *bool
	sort *string
}

func (r ApiInitshutdownscriptIdIdGetRequest) Limit(limit int32) ApiInitshutdownscriptIdIdGetRequest {
	r.limit = &limit
	return r
}
func (r ApiInitshutdownscriptIdIdGetRequest) Offset(offset int32) ApiInitshutdownscriptIdIdGetRequest {
	r.offset = &offset
	return r
}
func (r ApiInitshutdownscriptIdIdGetRequest) Count(count bool) ApiInitshutdownscriptIdIdGetRequest {
	r.count = &count
	return r
}
func (r ApiInitshutdownscriptIdIdGetRequest) Sort(sort string) ApiInitshutdownscriptIdIdGetRequest {
	r.sort = &sort
	return r
}

func (r ApiInitshutdownscriptIdIdGetRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.InitshutdownscriptIdIdGetExecute(r)
}

/*
 * InitshutdownscriptIdIdGet Method for InitshutdownscriptIdIdGet
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id
 * @return ApiInitshutdownscriptIdIdGetRequest
 */
func (a *InitshutdownscriptApiService) InitshutdownscriptIdIdGet(ctx _context.Context, id []interface{}) ApiInitshutdownscriptIdIdGetRequest {
	return ApiInitshutdownscriptIdIdGetRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

/*
 * Execute executes the request
 */
func (a *InitshutdownscriptApiService) InitshutdownscriptIdIdGetExecute(r ApiInitshutdownscriptIdIdGetRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InitshutdownscriptApiService.InitshutdownscriptIdIdGet")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/initshutdownscript/id/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(r.id, "csv")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.offset != nil {
		localVarQueryParams.Add("offset", parameterToString(*r.offset, ""))
	}
	if r.count != nil {
		localVarQueryParams.Add("count", parameterToString(*r.count, ""))
	}
	if r.sort != nil {
		localVarQueryParams.Add("sort", parameterToString(*r.sort, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiInitshutdownscriptIdIdPutRequest struct {
	ctx _context.Context
	ApiService *InitshutdownscriptApiService
	id int32
	initshutdownscriptUpdate1 *InitshutdownscriptUpdate1
}

func (r ApiInitshutdownscriptIdIdPutRequest) InitshutdownscriptUpdate1(initshutdownscriptUpdate1 InitshutdownscriptUpdate1) ApiInitshutdownscriptIdIdPutRequest {
	r.initshutdownscriptUpdate1 = &initshutdownscriptUpdate1
	return r
}

func (r ApiInitshutdownscriptIdIdPutRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.InitshutdownscriptIdIdPutExecute(r)
}

/*
 * InitshutdownscriptIdIdPut Method for InitshutdownscriptIdIdPut
 * Update initshutdown script task of `id`.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id
 * @return ApiInitshutdownscriptIdIdPutRequest
 */
func (a *InitshutdownscriptApiService) InitshutdownscriptIdIdPut(ctx _context.Context, id int32) ApiInitshutdownscriptIdIdPutRequest {
	return ApiInitshutdownscriptIdIdPutRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

/*
 * Execute executes the request
 */
func (a *InitshutdownscriptApiService) InitshutdownscriptIdIdPutExecute(r ApiInitshutdownscriptIdIdPutRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InitshutdownscriptApiService.InitshutdownscriptIdIdPut")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/initshutdownscript/id/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.initshutdownscriptUpdate1
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiInitshutdownscriptPostRequest struct {
	ctx _context.Context
	ApiService *InitshutdownscriptApiService
	initshutdownscriptCreate0 *InitshutdownscriptCreate0
}

func (r ApiInitshutdownscriptPostRequest) InitshutdownscriptCreate0(initshutdownscriptCreate0 InitshutdownscriptCreate0) ApiInitshutdownscriptPostRequest {
	r.initshutdownscriptCreate0 = &initshutdownscriptCreate0
	return r
}

func (r ApiInitshutdownscriptPostRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.InitshutdownscriptPostExecute(r)
}

/*
 * InitshutdownscriptPost Method for InitshutdownscriptPost
 * Create an initshutdown script task.

`type` indicates if a command or script should be executed at `when`.

There are three choices for `when`:

1) PREINIT - This is early in the boot process before all the services / rc scripts have started
2) POSTINIT - This is late in the boot process when most of the services / rc scripts have started
3) SHUTDOWN - This is on shutdown

`timeout` is an integer value which indicates time in seconds which the system should wait for the execution
of script/command. It should be noted that a hard limit for a timeout is configured by the base OS, so when
a script/command is set to execute on SHUTDOWN, the hard limit configured by the base OS is changed adding
the timeout specified by script/command so it can be ensured that it executes as desired and is not interrupted
by the base OS's limit.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiInitshutdownscriptPostRequest
 */
func (a *InitshutdownscriptApiService) InitshutdownscriptPost(ctx _context.Context) ApiInitshutdownscriptPostRequest {
	return ApiInitshutdownscriptPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 */
func (a *InitshutdownscriptApiService) InitshutdownscriptPostExecute(r ApiInitshutdownscriptPostRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InitshutdownscriptApiService.InitshutdownscriptPost")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/initshutdownscript"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.initshutdownscriptCreate0
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
