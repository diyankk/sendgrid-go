package main

import (
	"fmt"
	"github.com/sendgrid/sendgrid-go"
	"log"
)

///////////////////////////////////////////////////
// Retrieve global email statistics
// GET /stats

func Retrieveglobalemailstatistics() {
	apiKey := "JUST_A_TEST_KEY"
	host := "http://localhost:4010"
	request := sendgrid.GetRequest(apiKey, "/v3/stats", host)
	request.Method = "GET"
	queryParams := make(map[string]string)
	queryParams["aggregated_by"] = "day"
	queryParams["limit"] = "1"
	queryParams["start_date"] = "2016-01-01"
	queryParams["end_date"] = "2016-04-01"
	queryParams["offset"] = "1"
	request.QueryParams = queryParams
	response, err := sendgrid.API(request)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

func main() {
	Retrieveglobalemailstatistics()
}
