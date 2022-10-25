# \AgentActionsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConnectAgents**](AgentActionsApi.md#ConnectAgents) | **Post** /web/api/v2.1/agents/actions/connect | Connect to Network
[**DisconnectAgents**](AgentActionsApi.md#DisconnectAgents) | **Post** /web/api/v2.1/agents/actions/disconnect | Disconnect from Network



## ConnectAgents

> AffectedResultsSchema200 ConnectAgents(ctx).Body(body).Execute()

Connect to Network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewAgentsSchemasAgentsActionSchema(*openapiclient.NewAgentsSchemasAgentsActionSchemaFilter()) // AgentsSchemasAgentsActionSchema |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AgentActionsApi.ConnectAgents(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgentActionsApi.ConnectAgents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConnectAgents`: AffectedResultsSchema200
    fmt.Fprintf(os.Stdout, "Response from `AgentActionsApi.ConnectAgents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConnectAgentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AgentsSchemasAgentsActionSchema**](AgentsSchemasAgentsActionSchema.md) |  | 

### Return type

[**AffectedResultsSchema200**](AffectedResultsSchema200.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DisconnectAgents

> AffectedResultsSchema200 DisconnectAgents(ctx).Body(body).Execute()

Disconnect from Network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewAgentsSchemasAgentsDangerousActionSchema(*openapiclient.NewAgentsSchemasAgentsDangerousActionSchemaFilter()) // AgentsSchemasAgentsDangerousActionSchema |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AgentActionsApi.DisconnectAgents(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgentActionsApi.DisconnectAgents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DisconnectAgents`: AffectedResultsSchema200
    fmt.Fprintf(os.Stdout, "Response from `AgentActionsApi.DisconnectAgents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDisconnectAgentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AgentsSchemasAgentsDangerousActionSchema**](AgentsSchemasAgentsDangerousActionSchema.md) |  | 

### Return type

[**AffectedResultsSchema200**](AffectedResultsSchema200.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

