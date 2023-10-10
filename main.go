// This package is used to test the Loggly package
package main

import (
	"encoding/json"
	"fmt"
	loggly "github.com/jamespearly/loggly"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

// Used this: https://transform.tools/json-to-go
type SatelliteInfo struct {
	Info struct {
		Satname           string `json:"satname"`
		Satid             int    `json:"satid"`
		Transactionscount int    `json:"transactionscount"`
	} `json:"info"`
	Positions []struct {
		Satlatitude  float64 `json:"satlatitude"`
		Satlongitude float64 `json:"satlongitude"`
		Sataltitude  float64 `json:"sataltitude"`
		Azimuth      float64 `json:"azimuth"`
		Elevation    float64 `json:"elevation"`
		Ra           float64 `json:"ra"`
		Dec          float64 `json:"dec"`
		Timestamp    int     `json:"timestamp"`
		Eclipsed     bool    `json:"eclipsed"`
	} `json:"positions"`
}

func main() {
	var arg string
	timeLength := 1000

	//Get the length of time for the ticker to run at
	if len(os.Args) >= 2 {
		arg = os.Args[1]
		timeLength, _ = strconv.Atoi(arg)
	}

	//https://www.tutorialspoint.com/how-to-use-tickers-in-golang
	//TODO: Figure out how to run it once, since it waits for a tick before running
	ticker := time.NewTicker(time.Duration(timeLength) * time.Millisecond)

	//Loop through the ticker interval until program stops
	for _ = range ticker.C {

		var tag string
		var n2yoKey string
		var requestUrl string

		// Instantiate the loggly_client
		logglyClient := loggly.New(tag)

		tag = "Poller"
		n2yoKey = os.Getenv("N2YO_KEY")
		requestUrl = "https://api.n2yo.com/rest/v1/satellite/positions/25544/41.702/-76.014/0/2/&apiKey=" + n2yoKey

		responsePoller, errorPoller := http.Get(requestUrl)

		if errorPoller != nil {
			err := logglyClient.EchoSend("error", "Got an error while polling.")
			if err != nil {
				print("Got error: ", err)
			}
		} else if responsePoller.StatusCode == http.StatusOK {
			defer responsePoller.Body.Close()

			//Used this: https://stackoverflow.com/questions/17156371/how-to-get-json-response-from-http-get
			responseBytes, responseError := io.ReadAll(responsePoller.Body)
			if responseError != nil {
				err := logglyClient.EchoSend("error", "Got an error while reading response body from polling.")
				if err != nil {
					print("Got error: ", err)
				}
			} else {
				var satelliteInfo SatelliteInfo
				json.Unmarshal(responseBytes, &satelliteInfo)

				err := logglyClient.EchoSend("info", satelliteInfo.Info.Satname+" Latitude: "+fmt.Sprint(satelliteInfo.Positions[0].Satlatitude)+" Longitude: "+fmt.Sprint(satelliteInfo.Positions[0].Satlongitude))
				if err != nil {
					print("Got error while attempting to log response")
				}
			}
		} else {
			err := logglyClient.EchoSend("error", "Got unexpected response while polling")
			if err != nil {
				print("Got error: ", err)
			}
		}
	}
}
