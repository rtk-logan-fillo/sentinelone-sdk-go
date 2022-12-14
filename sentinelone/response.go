/*
S1 MGMT API

SentinelOne Management Console API specification. Changes made to original schema (saved as sentinelone_original.json) - Converted to yaml and intially removed all paths, definitions, and tags - Added security definition APIToken - Added global security APIToken - Added tags Agents, Agent Actions - Added /agents path and required definitions - Added /agents/actions/disconnect and required definitions - Added /agents/actions/connect and required definitions - Added operation id for all paths (for more readable generation) - Removed minimum value attribute if value is of type string (to prevents validation warnings) - Removed firstFullModeTime parameter in GetAgents (breaks generation due to malformed datetime serverside)

API version: 2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sentinelone

import (
	"net/http"
)

// APIResponse stores the API response returned by the server.
type APIResponse struct {
	*http.Response `json:"-"`
	Message        string `json:"message,omitempty"`
	// Operation is the name of the OpenAPI operation.
	Operation string `json:"operation,omitempty"`
	// RequestURL is the request URL. This value is always available, even if the
	// embedded *http.Response is nil.
	RequestURL string `json:"url,omitempty"`
	// Method is the HTTP method used for the request.  This value is always
	// available, even if the embedded *http.Response is nil.
	Method string `json:"method,omitempty"`
	// Payload holds the contents of the response body (which may be nil or empty).
	// This is provided here as the raw response.Body() reader will have already
	// been drained.
	Payload []byte `json:"-"`
}

// NewAPIResponse returns a new APIResponse object.
func NewAPIResponse(r *http.Response) *APIResponse {

	response := &APIResponse{Response: r}
	return response
}

// NewAPIResponseWithError returns a new APIResponse object with the provided error message.
func NewAPIResponseWithError(errorMessage string) *APIResponse {

	response := &APIResponse{Message: errorMessage}
	return response
}
