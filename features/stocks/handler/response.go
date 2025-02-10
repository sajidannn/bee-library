package handler

import "bee-library/features/stocks/entity"

type StockResponse struct {
	ID             uint   `json:"id"`
	BookID         uint   `json:"book_id"`
	TotalStock     int    `json:"total_stock"`
	AvailableStock int    `json:"available_stock"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}

func ToStockResponse(stock entity.Stock) StockResponse {
	return StockResponse{
		ID:             stock.ID,
		BookID:         stock.BookID,
		TotalStock:     stock.TotalStock,
		AvailableStock: stock.AvailableStock,
		CreatedAt:      stock.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:      stock.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}

func ToStockResponseList(stocks []entity.Stock) []StockResponse {
	var responseList []StockResponse
	for _, stock := range stocks {
		responseList = append(responseList, ToStockResponse(stock))
	}
	return responseList
}
