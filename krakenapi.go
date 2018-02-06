package main

import (
	"fmt"
	"log"
	"github.com/beldur/kraken-go-api-client"
)

func main() {
	api := krakenapi.New("BlNtopq1GpooZAqqofMF6pxy7shaGss+sJlw3P1XtLW7eUd10frJyOix", "nBBwYFp5wNMXQq1sCzMth4ySaiVyb94CkkGCchKftT+2fKpPI39YH55fp31L4NnvHEJ+7sPhTp79nOJLZHlgEQ==")
	result, err := api.Query("Ticker", map[string]string{
		"pair": "XXBTZEUR",
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Result: %+v\n", result)

	// There are also some strongly typed methods available
	ticker, err := api.Ticker(krakenapi.XXBTZEUR)
	if err != nil {
		log.Fatal(err)
	}

	balance, err := api.Balance ()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(ticker.XXBTZEUR.OpeningPrice)

	fmt.Println(balance);
}