package yahoo

import (
	"assetio/internal/port"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"strings"
)

// Yahoo struct
type yahoo struct {
	exchangeMap map[string]string
}

// New initializes Yahoo object
func New(exchangeMap map[string]string) port.Marketer {
	fmt.Println("Exchange Map:", exchangeMap)
	return &yahoo{
		exchangeMap: exchangeMap,
	}
}

// Convert symbol to Yahoo Finance format
func (y *yahoo) getSymbolSign(symbol, exchange string) string {
	if val, ok := y.exchangeMap[strings.ToLower(exchange)]; ok {
		return symbol + "." + val
	}
	return ""
}

// Query fetches stock data from Yahoo Finance API
func (y *yahoo) Query(symbol, exchange string) (port.MarketerData, error) {
	symbolSign := y.getSymbolSign(symbol, exchange)

	// Construct the Yahoo Finance API URL
	url := "https://query1.finance.yahoo.com/v8/finance/chart/" + symbolSign

	// Make the HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Parse the JSON response into chartData struct
	var data chartData
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// GetMarketPrice retrieves the market price from the chartData
func (y *chartData) GetMarketPrice() float64 {
	if len(y.Chart.Result) == 0 {
		return 0
	}
	return y.Chart.Result[0].Meta.RegularMarketPrice
}

// GetMarketChange retrieves the market change from the chartData
func (y *chartData) GetMarketChange() float64 {
	if len(y.Chart.Result) == 0 {
		return 0
	}

	return math.Floor((y.Chart.Result[0].Meta.RegularMarketPrice-y.Chart.Result[0].Meta.PreviousClose)*100) / 100
}

// GetMarketChangePercent retrieves the market change percentage from the chartData
func (y *chartData) GetMarketChangePercent() float64 {
	if len(y.Chart.Result) == 0 {
		return 0
	}

	// Calculate market change percentage
	currentPrice := y.Chart.Result[0].Meta.RegularMarketPrice
	previousClose := y.Chart.Result[0].Meta.PreviousClose

	// Return the percentage change
	if previousClose == 0 {
		// Avoid division by zero error
		return 0
	}

	changePercent := ((currentPrice - previousClose) / previousClose) * 100
	return math.Floor(changePercent*100) / 100

}
