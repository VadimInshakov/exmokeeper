package main

import (
	"flag"
	"fmt"
	"github.com/vadiminshakov/exmo"
	"math/big"
)

var (
	api     exmo.Exmo
	buysym  *string
	sellsym *string
	offset  *int
	limit   *int
	pub     *string
	secret  *string
)

func Init() {
	buysym = flag.String("buysymbol", "BTC", "first symbol in currency pair")
	sellsym = flag.String("sellsymbol", "RUB", "second symbol in currency pair")
	offset = flag.Int("offset", 0, "last deal offset")
	limit = flag.Int("limit", 100, "limit the number of displayed positions (default: 100, max: 1000)")
	pub = flag.String("pub", "", "public API key")
	secret = flag.String("secret", "", "private API key")
	flag.Parse()

	api = exmo.Api(*pub, *secret)
}

func main() {
	Calculate(*buysym, *sellsym, *offset, *limit)
}

func Calculate(buySymbol, sellSymbol string, offset, limit int) {
	usertrades, err := api.GetUserTrades(fmt.Sprintf("%s_%s", buySymbol, sellSymbol), offset, limit)
	bought := big.NewFloat(0)
	sold := big.NewFloat(0)
	quantityBought := big.NewFloat(0)
	quantitySold := big.NewFloat(0)
	if err != nil {
		fmt.Printf("api error: %s\n", err)
	} else {
		for _, allTrades := range usertrades {

			for _, trade := range allTrades.([]interface{}) {
				tradeData := trade.(map[string]interface{})
				fmt.Println(tradeData["type"])
				switch tradeData["type"] {
				case "buy":

					tradeDataBig, _ := new(big.Float).SetString(tradeData["amount"].(string))
					bought.Add(bought, tradeDataBig)

					quantityDataBig, _ := new(big.Float).SetString(tradeData["quantity"].(string))
					quantityBought.Add(quantityBought, quantityDataBig)

				case "sell":

					tradeDataBig, _ := new(big.Float).SetString(tradeData["amount"].(string))
					sold.Add(sold, tradeDataBig)

					quantityDataBig, _ := new(big.Float).SetString(tradeData["quantity"].(string))
					quantitySold.Add(quantitySold, quantityDataBig)

				}

			}
		}
	}
	fmt.Printf("For %s_%s pair you:\n", buySymbol, sellSymbol)
	fmt.Printf("bought %s %s for %s %s\n", quantityBought.String(), buySymbol, bought.String(), sellSymbol)
	fmt.Printf("sold %s %s for %s %s\n\n", quantitySold.String(), buySymbol, sold.String(), sellSymbol)
	fmt.Printf("YOU HAVE NOW: %s %s\n", quantityBought.Sub(quantityBought, quantitySold).String(), buySymbol)
	fmt.Printf("YOU SPENT TOTAL: %s %s\n", bought.Sub(bought, sold).String(), sellSymbol)
}
