package main

import (
	"fmt"

	"github.com/RShang97/projects/weather_logger/utils"
	"github.com/mattn/go-sqlite3"
)

func main() {
	req := utils.GetWeatherForecastRequest{
		Zip: "98109",
	}
	res, err := utils.GetWeatherForecast(&req)
	fmt.Println("marshalled response: \n", *res)
	fmt.Println("ERROR:", err)
	go-sqlite3.SQLITE_BLOB

}
