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
	GET_WEATHER_FORECAST_URL_PREFIX = "https://api.openweathermap.org/data/2.5/weather?"
	ZIP_PARAM                       = "zip"
	APP_ID_PARAM                    = "appid"
)

type GetWeatherForecastRequest struct {
	Zip         string
	CityName    string
	StateCode   string
	CountryCode string
}

type GetWeatherForecastResponse struct {
	Base            string
	Clouds          map[string]int
	Cod             int
	Coord           map[string]float64
	Dt              float64
	ID              int
	TemperatureInfo map[string]float64 `json:"main"`
	Name            string
	LocationInfo    map[string]interface{} `json:"sys"`
	Timezone        int
	Visibility      int
	WeatherInfo     []map[string]interface{} `json:"weather"`
	Wind            map[string]float32
}

func GetWeatherForecast(req *GetWeatherForecastRequest) (*GetWeatherForecastResponse, error) {
	values := url.Values{
		ZIP_PARAM:    []string{req.Zip},
		APP_ID_PARAM: []string{API_KEY},
	}
	// var response GetWeatherForecastResponse
	var response_holder GetWeatherForecastResponse

	err := CallGetAndUnmarshalJSON(GET_WEATHER_FORECAST_URL_PREFIX, values, &response_holder)

	return &response_holder, err
}

func CallGetAndUnmarshalJSON(urlPrefix string, values url.Values, response_holder interface{}) error {
	urlPrefix += values.Encode()
	fmt.Println(urlPrefix)
	response, err := http.Get(urlPrefix)
	if err != nil {
		fmt.Println("error in http request\n", err)
		return err
	}
	defer response.Body.Close()

	fmt.Println("raw response: ", *response, "\n")
	fmt.Println("response body: ", response.Body, "\n")

	err = json.NewDecoder(response.Body).Decode(response_holder)
	if err != nil {
		fmt.Println("ERROR: CallGetAndUnmarshallJson, error decoding http response")
	}
	return err
}
