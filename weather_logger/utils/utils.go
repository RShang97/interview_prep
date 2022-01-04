package utils

/*
db schema

| id | city_name | zip code | state | temperature F | Temperature C | Time | is_forecast |

*/

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const (
	API_KEY                         = "f78dc6e0d19016d3f833c82c97e2d6fb"
	GET_WEATHER_FORECAST_URL_PREFIX = "pro.openweathermap.org/data/2.5/forecast/hourly?"
	ZIP_PARAM                       = "zip"
	APP_ID_PARAM                    = "appid"
)

type GetWeatherForecastRequest struct {
	Zip         string
	CityName    string
	StateCode   string
	CountryCode string
}

type GetWeatherForecastResponse map[string]string

func GetWeatherForecast(req *GetWeatherForecastRequest) (*GetWeatherForecastResponse, error) {
	values := url.Values{
		ZIP_PARAM:    {req.Zip},
		APP_ID_PARAM: {API_KEY},
	}
	var res GetWeatherForecastResponse
	err := CallGetAndUnmarshalJSON(GET_WEATHER_FORECAST_URL_PREFIX, values, res)

	return &res, err
}

func CallGetAndUnmarshalJSON(url string, values url.Values, response_holder interface{}) error {
	url += values.Encode()
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("error in http request\n", err)
		return err
	}

	fmt.Println("raw response: \n", response)

	err = json.NewDecoder(response.Body).Decode(response_holder)
	return err
}
