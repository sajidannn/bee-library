package handler

// StockUpdateRequest digunakan untuk mengupdate stok secara manual
type StockUpdateRequest struct {
	TotalStock     int `json:"total_stock" binding:"required,gte=0"`
	AvailableStock int `json:"available_stock" binding:"required,gte=0"`
}
