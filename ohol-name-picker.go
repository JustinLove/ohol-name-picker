package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"io/ioutil"
	"log"
	"math/rand"
	"strings"
	"time"
)

const numberOfNames int = 10

type Names struct {
	Firsts [numberOfNames]string `json:"firsts"`
	Lasts  [numberOfNames]string `json:"lasts"`
}

var firstNames []string
var lastNames []string

func init() {
	rand.Seed(time.Now().UnixNano())
	firstNames = loadNames("firstNames.txt")
	lastNames = loadNames("lastNames.txt")
}

func LambdaHandler() (Names, error) {
	var response Names

	for i := 0; i < numberOfNames; i++ {
		response.Firsts[i] = firstNames[rand.Intn(len(firstNames))]
		response.Lasts[i] = lastNames[rand.Intn(len(lastNames))]
	}

	return response, nil
}

func main() {
	lambda.Start(LambdaHandler)
	return
	response, _ := LambdaHandler()

	for i := 0; i < numberOfNames; i++ {
		fmt.Printf("%s %s\n", response.Firsts[i], response.Lasts[i])
	}
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
