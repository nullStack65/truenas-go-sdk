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

// IdmapApiService IdmapApi service
type IdmapApiService service

type ApiIdmapBackendChoicesGetRequest struct {
	ctx _context.Context
	ApiService *IdmapApiService
}


func (r ApiIdmapBackendChoicesGetRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.IdmapBackendChoicesGetExecute(r)
}

/*
 * IdmapBackendChoicesGet Method for IdmapBackendChoicesGet
 * Returns array of valid idmap backend choices per directory service.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiIdmapBackendChoicesGetRequest
 */
func (a *IdmapApiService) IdmapBackendChoicesGet(ctx _context.Context) ApiIdmapBackendChoicesGetRequest {
	return ApiIdmapBackendChoicesGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 */
func (a *IdmapApiService) IdmapBackendChoicesGetExecute(r ApiIdmapBackendChoicesGetRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IdmapApiService.IdmapBackendChoicesGet")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/idmap/backend_choices"

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

type ApiIdmapBackendOptionsGetRequest struct {
	ctx _context.Context
	ApiService *IdmapApiService
}


func (r ApiIdmapBackendOptionsGetRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.IdmapBackendOptionsGetExecute(r)
}

/*
 * IdmapBackendOptionsGet Method for IdmapBackendOptionsGet
 * This returns full information about idmap backend options. Not all
`options` are valid for every backend.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiIdmapBackendOptionsGetRequest
 */
func (a *IdmapApiService) IdmapBackendOptionsGet(ctx _context.Context) ApiIdmapBackendOptionsGetRequest {
	return ApiIdmapBackendOptionsGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 */
func (a *IdmapApiService) IdmapBackendOptionsGetExecute(r ApiIdmapBackendOptionsGetRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IdmapApiService.IdmapBackendOptionsGet")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/idmap/backend_options"

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

type ApiIdmapClearIdmapCacheGetRequest struct {
	ctx _context.Context
	ApiService *IdmapApiService
}


func (r ApiIdmapClearIdmapCacheGetRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.IdmapClearIdmapCacheGetExecute(r)
}

/*
 * IdmapClearIdmapCacheGet Method for IdmapClearIdmapCacheGet
 * Stop samba, remove the winbindd_cache.tdb file, start samba, flush samba's cache.
This should be performed after finalizing idmap changes.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiIdmapClearIdmapCacheGetRequest
 */
func (a *IdmapApiService) IdmapClearIdmapCacheGet(ctx _context.Context) ApiIdmapClearIdmapCacheGetRequest {
	return ApiIdmapClearIdmapCacheGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 */
func (a *IdmapApiService) IdmapClearIdmapCacheGetExecute(r ApiIdmapClearIdmapCacheGetRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IdmapApiService.IdmapClearIdmapCacheGet")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/idmap/clear_idmap_cache"

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

type ApiIdmapGetRequest struct {
	ctx _context.Context
	ApiService *IdmapApiService
	limit *int32
	offset *int32
	count *bool
	sort *string
}

func (r ApiIdmapGetRequest) Limit(limit int32) ApiIdmapGetRequest {
	r.limit = &limit
	return r
}
func (r ApiIdmapGetRequest) Offset(offset int32) ApiIdmapGetRequest {
	r.offset = &offset
	return r
}
func (r ApiIdmapGetRequest) Count(count bool) ApiIdmapGetRequest {
	r.count = &count
	return r
}
func (r ApiIdmapGetRequest) Sort(sort string) ApiIdmapGetRequest {
	r.sort = &sort
	return r
}

func (r ApiIdmapGetRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.IdmapGetExecute(r)
}

/*
 * IdmapGet Method for IdmapGet
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiIdmapGetRequest
 */
func (a *IdmapApiService) IdmapGet(ctx _context.Context) ApiIdmapGetRequest {
	return ApiIdmapGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 */
func (a *IdmapApiService) IdmapGetExecute(r ApiIdmapGetRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IdmapApiService.IdmapGet")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/idmap"

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

type ApiIdmapIdIdDeleteRequest struct {
	ctx _context.Context
	ApiService *IdmapApiService
	id int32
}


func (r ApiIdmapIdIdDeleteRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.IdmapIdIdDeleteExecute(r)
}

/*
 * IdmapIdIdDelete Method for IdmapIdIdDelete
 * Delete a domain by id. Deletion of default system domains is not permitted.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id
 * @return ApiIdmapIdIdDeleteRequest
 */
func (a *IdmapApiService) IdmapIdIdDelete(ctx _context.Context, id int32) ApiIdmapIdIdDeleteRequest {
	return ApiIdmapIdIdDeleteRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

/*
 * Execute executes the request
 */
func (a *IdmapApiService) IdmapIdIdDeleteExecute(r ApiIdmapIdIdDeleteRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IdmapApiService.IdmapIdIdDelete")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/idmap/id/{id}"
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

type ApiIdmapIdIdGetRequest struct {
	ctx _context.Context
	ApiService *IdmapApiService
	id []interface{}
	limit *int32
	offset *int32
	count *bool
	sort *string
}

func (r ApiIdmapIdIdGetRequest) Limit(limit int32) ApiIdmapIdIdGetRequest {
	r.limit = &limit
	return r
}
func (r ApiIdmapIdIdGetRequest) Offset(offset int32) ApiIdmapIdIdGetRequest {
	r.offset = &offset
	return r
}
func (r ApiIdmapIdIdGetRequest) Count(count bool) ApiIdmapIdIdGetRequest {
	r.count = &count
	return r
}
func (r ApiIdmapIdIdGetRequest) Sort(sort string) ApiIdmapIdIdGetRequest {
	r.sort = &sort
	return r
}

func (r ApiIdmapIdIdGetRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.IdmapIdIdGetExecute(r)
}

/*
 * IdmapIdIdGet Method for IdmapIdIdGet
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id
 * @return ApiIdmapIdIdGetRequest
 */
func (a *IdmapApiService) IdmapIdIdGet(ctx _context.Context, id []interface{}) ApiIdmapIdIdGetRequest {
	return ApiIdmapIdIdGetRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

/*
 * Execute executes the request
 */
func (a *IdmapApiService) IdmapIdIdGetExecute(r ApiIdmapIdIdGetRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IdmapApiService.IdmapIdIdGet")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/idmap/id/{id}"
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

type ApiIdmapIdIdPutRequest struct {
	ctx _context.Context
	ApiService *IdmapApiService
	id int32
	idmapUpdate1 *IdmapUpdate1
}

func (r ApiIdmapIdIdPutRequest) IdmapUpdate1(idmapUpdate1 IdmapUpdate1) ApiIdmapIdIdPutRequest {
	r.idmapUpdate1 = &idmapUpdate1
	return r
}

func (r ApiIdmapIdIdPutRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.IdmapIdIdPutExecute(r)
}

/*
 * IdmapIdIdPut Method for IdmapIdIdPut
 * Update a domain by id.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id
 * @return ApiIdmapIdIdPutRequest
 */
func (a *IdmapApiService) IdmapIdIdPut(ctx _context.Context, id int32) ApiIdmapIdIdPutRequest {
	return ApiIdmapIdIdPutRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

/*
 * Execute executes the request
 */
func (a *IdmapApiService) IdmapIdIdPutExecute(r ApiIdmapIdIdPutRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IdmapApiService.IdmapIdIdPut")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/idmap/id/{id}"
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
	localVarPostBody = r.idmapUpdate1
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

type ApiIdmapOptionsChoicesPostRequest struct {
	ctx _context.Context
	ApiService *IdmapApiService
	body *string
}

func (r ApiIdmapOptionsChoicesPostRequest) Body(body string) ApiIdmapOptionsChoicesPostRequest {
	r.body = &body
	return r
}

func (r ApiIdmapOptionsChoicesPostRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.IdmapOptionsChoicesPostExecute(r)
}

/*
 * IdmapOptionsChoicesPost Method for IdmapOptionsChoicesPost
 * Returns a list of supported keys for the specified idmap backend.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiIdmapOptionsChoicesPostRequest
 */
func (a *IdmapApiService) IdmapOptionsChoicesPost(ctx _context.Context) ApiIdmapOptionsChoicesPostRequest {
	return ApiIdmapOptionsChoicesPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 */
func (a *IdmapApiService) IdmapOptionsChoicesPostExecute(r ApiIdmapOptionsChoicesPostRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IdmapApiService.IdmapOptionsChoicesPost")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/idmap/options_choices"

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
	localVarPostBody = r.body
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

type ApiIdmapPostRequest struct {
	ctx _context.Context
	ApiService *IdmapApiService
	idmapCreate0 *IdmapCreate0
}

func (r ApiIdmapPostRequest) IdmapCreate0(idmapCreate0 IdmapCreate0) ApiIdmapPostRequest {
	r.idmapCreate0 = &idmapCreate0
	return r
}

func (r ApiIdmapPostRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.IdmapPostExecute(r)
}

/*
 * IdmapPost Method for IdmapPost
 * Create a new IDMAP domain. These domains must be unique. This table
will be automatically populated after joining an Active Directory domain
if "allow trusted domains" is set to True in the AD service configuration.
There are three default system domains: DS_TYPE_ACTIVEDIRECTORY, DS_TYPE_LDAP, DS_TYPE_DEFAULT_DOMAIN.
The system domains correspond with the idmap settings under Active Directory, LDAP, and SMB
respectively.

`name` the pre-windows 2000 domain name.

`DNS_domain_name` DNS name of the domain.

`idmap_backend` provides a plugin interface for Winbind to use varying
backends to store SID/uid/gid mapping tables. The correct setting
depends on the environment in which the NAS is deployed.

`range_low` and `range_high` specify the UID and GID range for which this backend is authoritative.

`certificate_id` references the certificate ID of the SSL certificate to use for certificate-based
authentication to a remote LDAP server. This parameter is not supported for all idmap backends as some
backends will generate SID to ID mappings algorithmically without causing network traffic.

`options` are additional parameters that are backend-dependent:

`AD` idmap backend options:
`unix_primary_group` If True, the primary group membership is fetched from the LDAP attributes (gidNumber).
If False, the primary group membership is calculated via the "primaryGroupID" LDAP attribute.

`unix_nss_info` if True winbind will retrieve the login shell and home directory from the LDAP attributes.
If False or if the AD LDAP entry lacks the SFU attributes the smb4.conf parameters `template shell` and `template homedir` are used.

`schema_mode` Defines the schema that idmap_ad should use when querying Active Directory regarding user and group information.
This can be either the RFC2307 schema support included in Windows 2003 R2 or the Service for Unix (SFU) schema.
For SFU 3.0 or 3.5 please choose "SFU", for SFU 2.0 please choose "SFU20". The behavior of primary group membership is
controlled by the unix_primary_group option.

`AUTORID` idmap backend options:
`readonly` sets the module to read-only mode. No new ranges will be allocated and new mappings
will not be created in the idmap pool.

`ignore_builtin` ignores mapping requests for the BUILTIN domain.

`LDAP` idmap backend options:
`ldap_base_dn` defines the directory base suffix to use for SID/uid/gid mapping entries.

`ldap_user_dn` defines the user DN to be used for authentication.

`ldap_url` specifies the LDAP server to use for SID/uid/gid map entries.

`ssl` specifies whether to encrypt the LDAP transport for the idmap backend.

`NSS` idmap backend options:
`linked_service` specifies the auxiliary directory service ID provider.

`RFC2307` idmap backend options:
`domain` specifies the domain for which the idmap backend is being created. Numeric id, short-form
domain name, or long-form DNS domain name of the domain may be specified. Entry must be entered as
it appears in `idmap.domain`.

`range_low` and `range_high` specify the UID and GID range for which this backend is authoritative.

`ldap_server` defines the type of LDAP server to use. This can either be an LDAP server provided
by the Active Directory Domain (ad) or a stand-alone LDAP server.

`bind_path_user` specfies the search base where user objects can be found in the LDAP server.

`bind_path_group` specifies the search base where group objects can be found in the LDAP server.

`user_cn` query cn attribute instead of uid attribute for the user name in LDAP.

`realm` append @realm to cn for groups (and users if user_cn is set) in LDAP queries.

`ldmap_domain` when using the LDAP server in the Active Directory server, this allows one to
specify the domain where to access the Active Directory server. This allows using trust relationships
while keeping all RFC 2307 records in one place. This parameter is optional, the default is to access
the AD server in the current domain to query LDAP records.

`ldap_url` when using a stand-alone LDAP server, this parameter specifies the LDAP URL for accessing the LDAP server.

`ldap_user_dn` defines the user DN to be used for authentication.

`ldap_user_dn_password` is the password to be used for LDAP authentication.

`realm` defines the realm to use in the user and group names. This is only required when using cn_realm together with
 a stand-alone ldap server.

`RID` backend options:
`sssd_compat` generate idmap low range based on same algorithm that SSSD uses by default.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiIdmapPostRequest
 */
func (a *IdmapApiService) IdmapPost(ctx _context.Context) ApiIdmapPostRequest {
	return ApiIdmapPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 */
func (a *IdmapApiService) IdmapPostExecute(r ApiIdmapPostRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IdmapApiService.IdmapPost")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/idmap"

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
	localVarPostBody = r.idmapCreate0
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
