package handler

type StockUpdateRequest struct {
	TotalStock     int `json:"total_stock" binding:"gte=0"`
	AvailableStock int `json:"available_stock" binding:"gte=0"`
}
