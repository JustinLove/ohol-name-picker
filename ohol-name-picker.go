package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

const numberOfNames int = 10

type Names struct {
	Firsts [numberOfNames]string `json:"firsts"`
	Lasts  [numberOfNames]string `json:"lasts"`
}

type Response struct {
	//IsBase64Encoded bool              `json:"isBase64Encoded"`
	StatusCode int `json:"statusCode"`
	//Headers         map[string]string `json:"headers"`
	Body string `json:"body"`
}

var errorResponse Response = Response{
	//IsBase64Encoded: false,
	StatusCode: 500,
	//Headers:         make(map[string]string),
	Body: "",
}

var firstNames []string
var lastNames []string

func init() {
	rand.Seed(time.Now().UnixNano())
	firstNames = loadNames("firstNames.txt")
	lastNames = loadNames("lastNames.txt")
}

func pickNames() Names {
	var names Names

	for i := 0; i < numberOfNames; i++ {
		names.Firsts[i] = firstNames[rand.Intn(len(firstNames))]
		names.Lasts[i] = lastNames[rand.Intn(len(lastNames))]
	}

	return names
}

func LambdaHandler() (Response, error) {
	names := pickNames()

	jsonString, err := json.Marshal(names)
	if err != nil {
		return errorResponse, err
	}

	response := Response{
		//IsBase64Encoded: true,
		StatusCode: 200,
		//Headers:         make(map[string]string),
		Body: string(jsonString),
	}

	return response, nil
}

func printNames() {
	names := pickNames()

	for i := 0; i < numberOfNames; i++ {
		fmt.Printf("%s %s\n", names.Firsts[i], names.Lasts[i])
	}
}

func printResponse() {
	response, lambdaErr := LambdaHandler()

	jsonString, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err)
		return
	}

	os.Stdout.Write(jsonString)
	log.Println(lambdaErr)
}

func main() {
	lambda.Start(LambdaHandler)
	//return

	//printNames()
	//printResponse()
}

func loadNames(filename string) []string {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	text := string(content)
	lines := strings.Split(text, "\n")

	return lines
}
