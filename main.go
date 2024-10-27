package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type ExchangeRateData struct {
	Base          string             `json:"base"`
	LastUpdated   int64              `json:"last_updated"`
	ExchangeRates map[string]float64 `json:"exchange_rates"`
}

func main() {
	if len(os.Args) != 3 {
		log.Fatalln("Incorrect number of arguments. Expecting base and target currencies")
	}

	MakeRequest(os.Args[1], os.Args[2])
}

func MakeRequest(base string, target string) {
	env, err := readDotfile(".env")
	if err != nil {
		log.Fatalln(err)
	}

	key := env["api_key"]
	if len(key) == 0 {
		log.Fatalln("missing api_key in .env file")
	}

	resp, err := http.Get(fmt.Sprintf("https://exchange-rates.abstractapi.com/v1/live/?api_key=%s&base=%s&target=%s", key, base, target))
	if err != nil {
		log.Fatalln(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(body))

	var data ExchangeRateData
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	lastUpdated := time.Unix(data.LastUpdated, 0).UTC().Local().Format(time.RFC822)

	fmt.Printf("Base Currency: %s\n", data.Base)
	fmt.Printf("Last Updated: %s\n", lastUpdated)
	fmt.Printf("EUR Exchange Rate: %.2f\n", data.ExchangeRates["EUR"])
}

func readDotfile(filename string) (map[string]string, error) {
	headers := make(map[string]string)

	// The file expects key/value pairs in the format
	//   api_key: 000000000000000000000000

	// Open the file and return on error
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Read the contents into a map, excluding blank and comment lines
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// Ignore empty lines and comment lines (starting with #)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		// Split into name and value, separated by the first colon
		parts := strings.SplitN(line, ":", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			headers[key] = value
		}
	}

	// Return any error from reading the file
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	// Return the name/value map
	return headers, nil
}
