package service

import "github.com/im-naren/sentinel/model"

type IPortfolioService interface {
	GetListOfStocks() (model.Portfolio, error)
	GetTotalProfitsAndLoss() (float64, error)
	GetTopGainers() (model.Portfolio, error)
	GetTopLosers() (model.Portfolio, error)
	// Add other functionalities as needed
}

