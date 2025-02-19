# kfp_server_api_v1beta1.HealthzServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**get_healthz**](HealthzServiceApi.md#get_healthz) | **GET** /apis/v1beta1/healthz | Get healthz data.


# **get_healthz**
> V1beta1GetHealthzResponse get_healthz()

Get healthz data.

### Example

* Api Key Authentication (Bearer):
```python
from __future__ import print_function
import time
import kfp_server_api_v1beta1
from kfp_server_api_v1beta1.rest import ApiException
from pprint import pprint
# Defining the host is optional and defaults to http://localhost
# See configuration.py for a list of all supported configuration parameters.
configuration = kfp_server_api_v1beta1.Configuration(
    host = "http://localhost"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: Bearer
configuration = kfp_server_api_v1beta1.Configuration(
    host = "http://localhost",
    api_key = {
        'authorization': 'YOUR_API_KEY'
    }
)
# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['authorization'] = 'Bearer'

# Enter a context with an instance of the API client
with kfp_server_api_v1beta1.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = kfp_server_api_v1beta1.HealthzServiceApi(api_client)
    
    try:
        # Get healthz data.
        api_response = api_instance.get_healthz()
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling HealthzServiceApi->get_healthz: %s\n" % e)
```

### Parameters
This endpoint does not need any parameter.

### Return type

[**V1beta1GetHealthzResponse**](V1beta1GetHealthzResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A successful response. |  -  |
**0** |  |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

