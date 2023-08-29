# CategoryApiApi

All URIs are relative to *http://localhost:3000/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**categoriesCategoryIdDelete**](CategoryApiApi.md#categoriesCategoryIdDelete) | **DELETE** /categories/{categoryId} | Delete Category by Id
[**categoriesCategoryIdGet**](CategoryApiApi.md#categoriesCategoryIdGet) | **GET** /categories/{categoryId} | Get Category by Id
[**categoriesCategoryIdPut**](CategoryApiApi.md#categoriesCategoryIdPut) | **PUT** /categories/{categoryId} | Update category by Id
[**categoriesGet**](CategoryApiApi.md#categoriesGet) | **GET** /categories | List All Categories
[**categoriesPost**](CategoryApiApi.md#categoriesPost) | **POST** /categories | Create new a Category


<a name="categoriesCategoryIdDelete"></a>
# **categoriesCategoryIdDelete**
> InlineResponse2001 categoriesCategoryIdDelete(UNKNOWN_PARAMETER_NAME)

Delete Category by Id

Delete Category by Id

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.models.*;
import org.openapitools.client.api.CategoryApiApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://localhost:3000/api");

    CategoryApiApi apiInstance = new CategoryApiApi(defaultClient);
     UNKNOWN_PARAMETER_NAME = new null(); //  | Category Id
    try {
      InlineResponse2001 result = apiInstance.categoriesCategoryIdDelete(UNKNOWN_PARAMETER_NAME);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling CategoryApiApi#categoriesCategoryIdDelete");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **UNKNOWN_PARAMETER_NAME** | [****](.md)| Category Id |

### Return type

[**InlineResponse2001**](InlineResponse2001.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Success delete category |  -  |

<a name="categoriesCategoryIdGet"></a>
# **categoriesCategoryIdGet**
> InlineResponse200 categoriesCategoryIdGet(UNKNOWN_PARAMETER_NAME)

Get Category by Id

Get Category by Id

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.models.*;
import org.openapitools.client.api.CategoryApiApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://localhost:3000/api");

    CategoryApiApi apiInstance = new CategoryApiApi(defaultClient);
     UNKNOWN_PARAMETER_NAME = new null(); //  | Category Id
    try {
      InlineResponse200 result = apiInstance.categoriesCategoryIdGet(UNKNOWN_PARAMETER_NAME);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling CategoryApiApi#categoriesCategoryIdGet");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **UNKNOWN_PARAMETER_NAME** | [****](.md)| Category Id |

### Return type

[**InlineResponse200**](InlineResponse200.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Success get category |  -  |

<a name="categoriesCategoryIdPut"></a>
# **categoriesCategoryIdPut**
> InlineResponse200 categoriesCategoryIdPut(UNKNOWN_PARAMETER_NAME, createorUpdateCategory)

Update category by Id

Update category by Id

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.models.*;
import org.openapitools.client.api.CategoryApiApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://localhost:3000/api");

    CategoryApiApi apiInstance = new CategoryApiApi(defaultClient);
     UNKNOWN_PARAMETER_NAME = new null(); //  | Category Id
    CreateorUpdateCategory createorUpdateCategory = new CreateorUpdateCategory(); // CreateorUpdateCategory | 
    try {
      InlineResponse200 result = apiInstance.categoriesCategoryIdPut(UNKNOWN_PARAMETER_NAME, createorUpdateCategory);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling CategoryApiApi#categoriesCategoryIdPut");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **UNKNOWN_PARAMETER_NAME** | [****](.md)| Category Id |
 **createorUpdateCategory** | [**CreateorUpdateCategory**](CreateorUpdateCategory.md)|  | [optional]

### Return type

[**InlineResponse200**](InlineResponse200.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Success get category |  -  |

<a name="categoriesGet"></a>
# **categoriesGet**
> InlineResponse200 categoriesGet()

List All Categories

List All of Categories

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.models.*;
import org.openapitools.client.api.CategoryApiApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://localhost:3000/api");

    CategoryApiApi apiInstance = new CategoryApiApi(defaultClient);
    try {
      InlineResponse200 result = apiInstance.categoriesGet();
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling CategoryApiApi#categoriesGet");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse200**](InlineResponse200.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Success Get All Categories |  -  |

<a name="categoriesPost"></a>
# **categoriesPost**
> InlineResponse200 categoriesPost(createorUpdateCategory)

Create new a Category

Create new a Category

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.models.*;
import org.openapitools.client.api.CategoryApiApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://localhost:3000/api");

    CategoryApiApi apiInstance = new CategoryApiApi(defaultClient);
    CreateorUpdateCategory createorUpdateCategory = new CreateorUpdateCategory(); // CreateorUpdateCategory | 
    try {
      InlineResponse200 result = apiInstance.categoriesPost(createorUpdateCategory);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling CategoryApiApi#categoriesPost");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createorUpdateCategory** | [**CreateorUpdateCategory**](CreateorUpdateCategory.md)|  | [optional]

### Return type

[**InlineResponse200**](InlineResponse200.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Success Create Category |  -  |

