// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package core

import (
	"fmt"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/oci/vendor-internal/github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListAppCatalogListingResourceVersionsRequest wrapper for the ListAppCatalogListingResourceVersions operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/core/ListAppCatalogListingResourceVersions.go.html to see an example of how to use ListAppCatalogListingResourceVersionsRequest.
type ListAppCatalogListingResourceVersionsRequest struct {

	// The OCID of the listing.
	ListingId *string `mandatory:"true" contributesTo:"path" name:"listingId"`

	// For list pagination. The maximum number of results per page, or items to return in a paginated
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from the previous "List"
	// call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`). The DISPLAYNAME sort order
	// is case sensitive.
	SortOrder ListAppCatalogListingResourceVersionsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAppCatalogListingResourceVersionsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAppCatalogListingResourceVersionsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAppCatalogListingResourceVersionsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAppCatalogListingResourceVersionsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListAppCatalogListingResourceVersionsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListAppCatalogListingResourceVersionsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListAppCatalogListingResourceVersionsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListAppCatalogListingResourceVersionsResponse wrapper for the ListAppCatalogListingResourceVersions operation
type ListAppCatalogListingResourceVersionsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []AppCatalogListingResourceVersionSummary instances
	Items []AppCatalogListingResourceVersionSummary `presentIn:"body"`

	// For list pagination. When this header appears in the response, additional pages
	// of results remain. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListAppCatalogListingResourceVersionsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAppCatalogListingResourceVersionsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAppCatalogListingResourceVersionsSortOrderEnum Enum with underlying type: string
type ListAppCatalogListingResourceVersionsSortOrderEnum string

// Set of constants representing the allowable values for ListAppCatalogListingResourceVersionsSortOrderEnum
const (
	ListAppCatalogListingResourceVersionsSortOrderAsc  ListAppCatalogListingResourceVersionsSortOrderEnum = "ASC"
	ListAppCatalogListingResourceVersionsSortOrderDesc ListAppCatalogListingResourceVersionsSortOrderEnum = "DESC"
)

var mappingListAppCatalogListingResourceVersionsSortOrderEnum = map[string]ListAppCatalogListingResourceVersionsSortOrderEnum{
	"ASC":  ListAppCatalogListingResourceVersionsSortOrderAsc,
	"DESC": ListAppCatalogListingResourceVersionsSortOrderDesc,
}

var mappingListAppCatalogListingResourceVersionsSortOrderEnumLowerCase = map[string]ListAppCatalogListingResourceVersionsSortOrderEnum{
	"asc":  ListAppCatalogListingResourceVersionsSortOrderAsc,
	"desc": ListAppCatalogListingResourceVersionsSortOrderDesc,
}

// GetListAppCatalogListingResourceVersionsSortOrderEnumValues Enumerates the set of values for ListAppCatalogListingResourceVersionsSortOrderEnum
func GetListAppCatalogListingResourceVersionsSortOrderEnumValues() []ListAppCatalogListingResourceVersionsSortOrderEnum {
	values := make([]ListAppCatalogListingResourceVersionsSortOrderEnum, 0)
	for _, v := range mappingListAppCatalogListingResourceVersionsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListAppCatalogListingResourceVersionsSortOrderEnumStringValues Enumerates the set of values in String for ListAppCatalogListingResourceVersionsSortOrderEnum
func GetListAppCatalogListingResourceVersionsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListAppCatalogListingResourceVersionsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAppCatalogListingResourceVersionsSortOrderEnum(val string) (ListAppCatalogListingResourceVersionsSortOrderEnum, bool) {
	enum, ok := mappingListAppCatalogListingResourceVersionsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
