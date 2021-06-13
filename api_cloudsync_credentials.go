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

// CloudsyncCredentialsApiService CloudsyncCredentialsApi service
type CloudsyncCredentialsApiService service

type ApiCloudsyncCredentialsGetRequest struct {
	ctx _context.Context
	ApiService *CloudsyncCredentialsApiService
	limit *int32
	offset *int32
	count *bool
	sort *string
}

func (r ApiCloudsyncCredentialsGetRequest) Limit(limit int32) ApiCloudsyncCredentialsGetRequest {
	r.limit = &limit
	return r
}
func (r ApiCloudsyncCredentialsGetRequest) Offset(offset int32) ApiCloudsyncCredentialsGetRequest {
	r.offset = &offset
	return r
}
func (r ApiCloudsyncCredentialsGetRequest) Count(count bool) ApiCloudsyncCredentialsGetRequest {
	r.count = &count
	return r
}
func (r ApiCloudsyncCredentialsGetRequest) Sort(sort string) ApiCloudsyncCredentialsGetRequest {
	r.sort = &sort
	return r
}

func (r ApiCloudsyncCredentialsGetRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.CloudsyncCredentialsGetExecute(r)
}

/*
 * CloudsyncCredentialsGet Method for CloudsyncCredentialsGet
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiCloudsyncCredentialsGetRequest
 */
func (a *CloudsyncCredentialsApiService) CloudsyncCredentialsGet(ctx _context.Context) ApiCloudsyncCredentialsGetRequest {
	return ApiCloudsyncCredentialsGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 */
func (a *CloudsyncCredentialsApiService) CloudsyncCredentialsGetExecute(r ApiCloudsyncCredentialsGetRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CloudsyncCredentialsApiService.CloudsyncCredentialsGet")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/cloudsync/credentials"

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

type ApiCloudsyncCredentialsIdIdDeleteRequest struct {
	ctx _context.Context
	ApiService *CloudsyncCredentialsApiService
	id int32
}


func (r ApiCloudsyncCredentialsIdIdDeleteRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.CloudsyncCredentialsIdIdDeleteExecute(r)
}

/*
 * CloudsyncCredentialsIdIdDelete Method for CloudsyncCredentialsIdIdDelete
 * Delete Cloud Sync Credentials of `id`.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id
 * @return ApiCloudsyncCredentialsIdIdDeleteRequest
 */
func (a *CloudsyncCredentialsApiService) CloudsyncCredentialsIdIdDelete(ctx _context.Context, id int32) ApiCloudsyncCredentialsIdIdDeleteRequest {
	return ApiCloudsyncCredentialsIdIdDeleteRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

/*
 * Execute executes the request
 */
func (a *CloudsyncCredentialsApiService) CloudsyncCredentialsIdIdDeleteExecute(r ApiCloudsyncCredentialsIdIdDeleteRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CloudsyncCredentialsApiService.CloudsyncCredentialsIdIdDelete")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/cloudsync/credentials/id/{id}"
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

type ApiCloudsyncCredentialsIdIdGetRequest struct {
	ctx _context.Context
	ApiService *CloudsyncCredentialsApiService
	id []interface{}
	limit *int32
	offset *int32
	count *bool
	sort *string
}

func (r ApiCloudsyncCredentialsIdIdGetRequest) Limit(limit int32) ApiCloudsyncCredentialsIdIdGetRequest {
	r.limit = &limit
	return r
}
func (r ApiCloudsyncCredentialsIdIdGetRequest) Offset(offset int32) ApiCloudsyncCredentialsIdIdGetRequest {
	r.offset = &offset
	return r
}
func (r ApiCloudsyncCredentialsIdIdGetRequest) Count(count bool) ApiCloudsyncCredentialsIdIdGetRequest {
	r.count = &count
	return r
}
func (r ApiCloudsyncCredentialsIdIdGetRequest) Sort(sort string) ApiCloudsyncCredentialsIdIdGetRequest {
	r.sort = &sort
	return r
}

func (r ApiCloudsyncCredentialsIdIdGetRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.CloudsyncCredentialsIdIdGetExecute(r)
}

/*
 * CloudsyncCredentialsIdIdGet Method for CloudsyncCredentialsIdIdGet
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id
 * @return ApiCloudsyncCredentialsIdIdGetRequest
 */
func (a *CloudsyncCredentialsApiService) CloudsyncCredentialsIdIdGet(ctx _context.Context, id []interface{}) ApiCloudsyncCredentialsIdIdGetRequest {
	return ApiCloudsyncCredentialsIdIdGetRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

/*
 * Execute executes the request
 */
func (a *CloudsyncCredentialsApiService) CloudsyncCredentialsIdIdGetExecute(r ApiCloudsyncCredentialsIdIdGetRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CloudsyncCredentialsApiService.CloudsyncCredentialsIdIdGet")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/cloudsync/credentials/id/{id}"
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

type ApiCloudsyncCredentialsIdIdPutRequest struct {
	ctx _context.Context
	ApiService *CloudsyncCredentialsApiService
	id int32
	cloudsyncCredentialsUpdate1 *CloudsyncCredentialsUpdate1
}

func (r ApiCloudsyncCredentialsIdIdPutRequest) CloudsyncCredentialsUpdate1(cloudsyncCredentialsUpdate1 CloudsyncCredentialsUpdate1) ApiCloudsyncCredentialsIdIdPutRequest {
	r.cloudsyncCredentialsUpdate1 = &cloudsyncCredentialsUpdate1
	return r
}

func (r ApiCloudsyncCredentialsIdIdPutRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.CloudsyncCredentialsIdIdPutExecute(r)
}

/*
 * CloudsyncCredentialsIdIdPut Method for CloudsyncCredentialsIdIdPut
 * Update Cloud Sync Credentials of `id`.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id
 * @return ApiCloudsyncCredentialsIdIdPutRequest
 */
func (a *CloudsyncCredentialsApiService) CloudsyncCredentialsIdIdPut(ctx _context.Context, id int32) ApiCloudsyncCredentialsIdIdPutRequest {
	return ApiCloudsyncCredentialsIdIdPutRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

/*
 * Execute executes the request
 */
func (a *CloudsyncCredentialsApiService) CloudsyncCredentialsIdIdPutExecute(r ApiCloudsyncCredentialsIdIdPutRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CloudsyncCredentialsApiService.CloudsyncCredentialsIdIdPut")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/cloudsync/credentials/id/{id}"
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
	localVarPostBody = r.cloudsyncCredentialsUpdate1
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

type ApiCloudsyncCredentialsPostRequest struct {
	ctx _context.Context
	ApiService *CloudsyncCredentialsApiService
	cloudsyncCredentialsCreate0 *CloudsyncCredentialsCreate0
}

func (r ApiCloudsyncCredentialsPostRequest) CloudsyncCredentialsCreate0(cloudsyncCredentialsCreate0 CloudsyncCredentialsCreate0) ApiCloudsyncCredentialsPostRequest {
	r.cloudsyncCredentialsCreate0 = &cloudsyncCredentialsCreate0
	return r
}

func (r ApiCloudsyncCredentialsPostRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.CloudsyncCredentialsPostExecute(r)
}

/*
 * CloudsyncCredentialsPost Method for CloudsyncCredentialsPost
 * Create Cloud Sync Credentials.

`attributes` is a dictionary of valid values which will be used to authorize with the `provider`.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiCloudsyncCredentialsPostRequest
 */
func (a *CloudsyncCredentialsApiService) CloudsyncCredentialsPost(ctx _context.Context) ApiCloudsyncCredentialsPostRequest {
	return ApiCloudsyncCredentialsPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 */
func (a *CloudsyncCredentialsApiService) CloudsyncCredentialsPostExecute(r ApiCloudsyncCredentialsPostRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CloudsyncCredentialsApiService.CloudsyncCredentialsPost")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/cloudsync/credentials"

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
	localVarPostBody = r.cloudsyncCredentialsCreate0
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

type ApiCloudsyncCredentialsVerifyPostRequest struct {
	ctx _context.Context
	ApiService *CloudsyncCredentialsApiService
	cloudsyncCredentialsVerify0 *CloudsyncCredentialsVerify0
}

func (r ApiCloudsyncCredentialsVerifyPostRequest) CloudsyncCredentialsVerify0(cloudsyncCredentialsVerify0 CloudsyncCredentialsVerify0) ApiCloudsyncCredentialsVerifyPostRequest {
	r.cloudsyncCredentialsVerify0 = &cloudsyncCredentialsVerify0
	return r
}

func (r ApiCloudsyncCredentialsVerifyPostRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.CloudsyncCredentialsVerifyPostExecute(r)
}

/*
 * CloudsyncCredentialsVerifyPost Method for CloudsyncCredentialsVerifyPost
 * Verify if `attributes` provided for `provider` are authorized by the `provider`.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiCloudsyncCredentialsVerifyPostRequest
 */
func (a *CloudsyncCredentialsApiService) CloudsyncCredentialsVerifyPost(ctx _context.Context) ApiCloudsyncCredentialsVerifyPostRequest {
	return ApiCloudsyncCredentialsVerifyPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 */
func (a *CloudsyncCredentialsApiService) CloudsyncCredentialsVerifyPostExecute(r ApiCloudsyncCredentialsVerifyPostRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CloudsyncCredentialsApiService.CloudsyncCredentialsVerifyPost")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/cloudsync/credentials/verify"

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
	localVarPostBody = r.cloudsyncCredentialsVerify0
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
