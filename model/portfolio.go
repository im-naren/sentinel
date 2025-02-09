package model

import (
	"encoding/csv"
	"os"
	"strconv"
)

type PortfolioItem struct {
	Symbol          string
	Quantity        int
	AverageCost     float64
	LastTradedPrice float64
	CurrentValue    float64
	PNL             float64
	NetChange       float64
	DayChange       float64
}

type Portfolio struct {
	Items []PortfolioItem
}

func ReadPortfolioFromFile(filePath string) (Portfolio, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return Portfolio{}, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return Portfolio{}, err
	}

	var portfolio Portfolio
	for _, record := range records {
		quantity, _ := strconv.Atoi(record[1])
		averageCost, _ := strconv.ParseFloat(record[2], 64)
		lastTradedPrice, _ := strconv.ParseFloat(record[3], 64)
		currentValue, _ := strconv.ParseFloat(record[4], 64)
		pnl, _ := strconv.ParseFloat(record[5], 64)
		netChange, _ := strconv.ParseFloat(record[6], 64)
		dayChange, _ := strconv.ParseFloat(record[7], 64)

		item := PortfolioItem{
			Symbol:          record[0],
			Quantity:        quantity,
			AverageCost:     averageCost,
			LastTradedPrice: lastTradedPrice,
			CurrentValue:    currentValue,
			PNL:             pnl,
			NetChange:       netChange,
			DayChange:       dayChange,
		}
		portfolio.Items = append(portfolio.Items, item)
	}

	return portfolio, nil
}
