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
)

// Linger please
var (
	_ _context.Context
)

// NetworkConfigurationApiService NetworkConfigurationApi service
type NetworkConfigurationApiService service

type ApiNetworkConfigurationGetRequest struct {
	ctx _context.Context
	ApiService *NetworkConfigurationApiService
}


func (r ApiNetworkConfigurationGetRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.NetworkConfigurationGetExecute(r)
}

/*
 * NetworkConfigurationGet Method for NetworkConfigurationGet
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiNetworkConfigurationGetRequest
 */
func (a *NetworkConfigurationApiService) NetworkConfigurationGet(ctx _context.Context) ApiNetworkConfigurationGetRequest {
	return ApiNetworkConfigurationGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 */
func (a *NetworkConfigurationApiService) NetworkConfigurationGetExecute(r ApiNetworkConfigurationGetRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NetworkConfigurationApiService.NetworkConfigurationGet")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/network/configuration"

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

type ApiNetworkConfigurationPutRequest struct {
	ctx _context.Context
	ApiService *NetworkConfigurationApiService
	networkConfigurationUpdate0 *NetworkConfigurationUpdate0
}

func (r ApiNetworkConfigurationPutRequest) NetworkConfigurationUpdate0(networkConfigurationUpdate0 NetworkConfigurationUpdate0) ApiNetworkConfigurationPutRequest {
	r.networkConfigurationUpdate0 = &networkConfigurationUpdate0
	return r
}

func (r ApiNetworkConfigurationPutRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.NetworkConfigurationPutExecute(r)
}

/*
 * NetworkConfigurationPut Method for NetworkConfigurationPut
 * Update Network Configuration Service configuration.

`ipv4gateway` if set is used instead of the default gateway provided by DHCP.

`nameserver1` is primary DNS server.

`nameserver2` is secondary DNS server.

`nameserver3` is tertiary DNS server.

`httpproxy` attribute must be provided if a proxy is to be used for network operations.

`netwait_enabled` is a boolean attribute which when set indicates that network services will not start at
boot unless they are able to ping the addresses listed in `netwait_ip` list.

`service_announcement` determines the broadcast protocols that will be used to advertise the server.
`netbios` enables the NetBIOS name server (NBNS), which starts concurrently with the SMB service. SMB clients
will only perform NBNS lookups if SMB1 is enabled. NBNS may be required for legacy SMB clients.
`mdns` enables multicast DNS service announcements for enabled services. `wsd` enables Web Service
Discovery support.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiNetworkConfigurationPutRequest
 */
func (a *NetworkConfigurationApiService) NetworkConfigurationPut(ctx _context.Context) ApiNetworkConfigurationPutRequest {
	return ApiNetworkConfigurationPutRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 */
func (a *NetworkConfigurationApiService) NetworkConfigurationPutExecute(r ApiNetworkConfigurationPutRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NetworkConfigurationApiService.NetworkConfigurationPut")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/network/configuration"

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
	localVarPostBody = r.networkConfigurationUpdate0
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
