package main

import (
	"encoding/csv"
	// "encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

type Birthday struct{
	date string
	name string
	greetingName string 
	phone string
}

func main(){
	lambda.Start(sendBirthdayWish)
	// sendBirthdayWish()
}

func sendBirthdayWish(){
	t := time.Now()

	date := fmt.Sprintf("%02d/%02d/%02d", int(t.Month()), int(t.Day()), int(t.Year()))

	fmt.Println(date)

	file, err := os.Open("birthdaysList.csv") // for read access
	if err != nil{
		log.Fatal(err)
	}

	defer file.Close()

	// reading through CSV file
	r:= csv.NewReader(file)

	var bd Birthday

	for{
			line, err := r.Read()
			if err != nil{
				fmt.Println(err)
			}

			if line[0] == date{
				bd = Birthday{
					date: line[0],
					name: line[1],
					greetingName: line[2], 
					phone: line[3], 
				}
				break
			}
		}

		text := "Happy Birthday, " + bd.greetingName + " from Omkar."

		fmt.Println(text)
		// send the text message

		accountSid := os.Getenv("TWILIO_ACCOUNT_SID")
		authToken := os.Getenv("TWILIO_AUTH_TOKEN")

		client := twilio.NewRestClientWithParams(twilio.ClientParams{
			Username: accountSid,
			Password: authToken,
		})

		params := &twilioApi.CreateMessageParams{}
		params.SetBody(text)
		params.SetFrom("+12052893550")
		params.SetTo("+91" + bd.phone)

		resp, err := client.Api.CreateMessage(params)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			if resp.Sid != nil {
				fmt.Println(*resp.Sid)
			} else {
				fmt.Println(resp.Sid)
			}
		}
}
