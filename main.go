package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
  "os"
  "github.com/joho/godotenv"
)

type ExchangeRates struct {
	Rates map[string]float64 `json:"rates"`
	Base  string             `json:"base"`
}

func getExchangeRates(baseCurrency string) (*ExchangeRates, error) {
	
  envErr := godotenv.Load()

  if envErr != nil {
    log.Fatalf("Error loading .env file: %v", envErr)
  }

  apiKey := os.Getenv("API_KEY")

  if apiKey == ""{
    log.Fatalf("API_KEY not found in the environment")
  }

	host := "currency-conversion-and-exchange-rates.p.rapidapi.com"
	url := fmt.Sprintf("https://currency-conversion-and-exchange-rates.p.rapidapi.com/latest?from=%s", baseCurrency)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %v", err)
	}

	req.Header.Add("x-rapidapi-key", apiKey)
	req.Header.Add("x-rapidapi-host", host)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make HTTP request: %v", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %v", err)
	}

	var exchangeRates ExchangeRates
	if err = json.Unmarshal(body, &exchangeRates); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	return &exchangeRates, nil
}

func convertCurrency(amount float64, sourceCurrency, targetCurrency string) (float64, error) {
	exchangeRates, err := getExchangeRates(sourceCurrency)
	if err != nil {
		return 0, err
	}

	targetRate, ok := exchangeRates.Rates[targetCurrency]
	if !ok {
		return 0, fmt.Errorf("target currency %s not found", targetCurrency)
	}

	convertedAmount := amount * targetRate
	return convertedAmount, nil
}

func main() {
	var amount float64
	var sourceCurrency, targetCurrency string

	fmt.Print("Enter the amount: ")
	_, err := fmt.Scan(&amount)
	if err != nil {
		log.Fatal("Error reading amount:", err)
	}

	fmt.Print("Enter the source currency (e.g., USD): ")
	_, err = fmt.Scan(&sourceCurrency)
	if err != nil {
		log.Fatal("Error reading source currency:", err)
	}

	fmt.Print("Enter the target currency (e.g., EUR): ")
	_, err = fmt.Scan(&targetCurrency)
	if err != nil {
		log.Fatal("Error reading target currency:", err)
	}

	convertedAmount, err := convertCurrency(amount, sourceCurrency, targetCurrency)
	if err != nil {
		log.Fatal("Error converting currency:", err)
	}

	fmt.Printf("%.2f %s is equivalent to %.2f %s\n", amount, sourceCurrency, convertedAmount, targetCurrency)
}

