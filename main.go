// This package is used to test the Loggly package
package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	loggly "github.com/jamespearly/loggly"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func main() {
	var arg string
	timeLength := 1000

	//Get the length of time for the ticker to run at
	if len(os.Args) >= 2 {
		arg = os.Args[1]
		timeLength, _ = strconv.Atoi(arg)
	}

	fmt.Printf("Starting agent, time length (in ms): " + fmt.Sprint(timeLength) + "\n")

	//https://www.tutorialspoint.com/how-to-use-tickers-in-golang
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

		responsePoller, _ := http.Get(requestUrl)
		contentLength := responsePoller.ContentLength

		if responsePoller.StatusCode == http.StatusOK {
			defer responsePoller.Body.Close()

			//Used this: https://stackoverflow.com/questions/17156371/how-to-get-json-response-from-http-get
			responseBytes, responseError := io.ReadAll(responsePoller.Body)
			if responseError != nil {
				err := logglyClient.EchoSend("error", "Got an error while reading response body from polling.")
				if err != nil {
					print("Got error: ", err)
				}
			} else {
				// In ApiStructs.go
				var satelliteInfoBase SatelliteInfo
				json.Unmarshal(responseBytes, &satelliteInfoBase)

				satelliteInfo := transformApiResponse(satelliteInfoBase)

				logglyClient.EchoSend("info", "Content Length: "+fmt.Sprint(contentLength))

				// Initialize a session that the SDK will use to load
				// credentials from the shared credentials file ~/.aws/credentials
				// and region from the shared configuration file ~/.aws/config.
				sess := session.Must(session.NewSessionWithOptions(session.Options{
					SharedConfigState: session.SharedConfigEnable,
				}))

				sess.Config.Region = aws.String("us-east-1")
				// Create DynamoDB client
				svc := dynamodb.New(sess)

				av, _ := dynamodbattribute.MarshalMap(satelliteInfo)

				input := &dynamodb.PutItemInput{
					Item:      av,
					TableName: aws.String("vlockwoo-satellites"),
				}

				response, _ := svc.PutItem(input)
				result, _ := json.Marshal(response)
				fmt.Printf(string(result))
			}
		} else {
			logglyClient.EchoSend("error", "Got unexpected response while polling")
		}
	}
}
