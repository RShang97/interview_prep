package zillow_structs

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

/**
 * constants for the GetSearchResults endpoint, might want to move into the struct
 */
const (
	ZWSID_PARAM                      = "zws-id"
	ADDRESS_PARAM                    = "address"
	CITYSTATEZIP_PARAM               = "citystatezip"
	RENTZESTIMATE_PARAM              = "rentzestimate"
	ZILLOW_PREFIX_URL                = "http://www.zillow.com/webservice/"
	GET_SEARCH_RESULTS_ENDPOINT_NAME = "GetSearchResults"
)

type GetSearchResultsRequest struct {
	ZwsID         string
	Address       string
	CityStateZip  string
	RentZestimate bool
}

type GetSearchResultsResponse struct {
	Zpid              string
	Links             []string
	Address           map[string]string
	ZestimateData     map[string]string
	RentZestimateData map[string]string
	LocalRealEstate   map[string]string
	LimitWarning      bool
}

func callGetAndMarshallXML(apiEndpointName string, values url.Values, responseHolder interface{}) error {
	url := ZILLOW_PREFIX_URL + apiEndpointName + ".htm?" + values.Encode()
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	fmt.Println(response)
	err = xml.NewDecoder(response.Body).Decode(responseHolder)
	if err != nil {
		return err
	}
	return nil
}

func GetSearchResults(request *GetSearchResultsRequest) (GetSearchResultsResponse, error) {
	fmt.Println("calling GetSearchResults ", request)
	values := url.Values{
		ZWSID_PARAM:         {request.ZwsID},
		ADDRESS_PARAM:       {request.Address},
		CITYSTATEZIP_PARAM:  {request.CityStateZip},
		RENTZESTIMATE_PARAM: {strconv.FormatBool(request.RentZestimate)},
	}
	var response GetSearchResultsResponse
	err := callGetAndMarshallXML(GET_SEARCH_RESULTS_ENDPOINT_NAME, values, &response)
	return response, err
}
