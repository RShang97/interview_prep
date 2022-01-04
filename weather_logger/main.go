package main

import (
	"fmt"
	"github.com/RShang97/projects/weather_logger/utils"
)

func main() {
	req := utils.GetWeatherForecastRequest{
		Zip: "98109",
	}
	res, err := utils.GetWeatherForecast(&req)
	fmt.Println("marshalled response: \n", res)
	fmt.Println("error:", err)
}
