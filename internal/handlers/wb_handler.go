package handlers

import (
	"encoding/json"
	"net/http"

	"wb-api-integration/internal/client"
	"wb-api-integration/internal/config"
	"wb-api-integration/pkg/logger"
)

// ProductHandler обрабатывает HTTP-запросы, связанные с продуктами и заказами
type ProductHandler struct {
	wbClient *client.WbClient
}

// NewProductHandler создает новый экземпляр ProductHandler
func NewProductHandler(cfg *config.Config) *ProductHandler {
	return &ProductHandler{
		wbClient: client.NewWbClient(cfg.StandardToken, cfg.AdvToken, cfg.StatToken),
	}
}

// GetProducts обрабатывает запрос на получение списка продуктов
func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	logger.Info("Getting products from WB API")

	data, err := h.wbClient.GetProducts()
	if err != nil {
		logger.Error("Failed to get products: " + err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

// GetOrders обрабатывает запрос на получение списка заказов
func (h *ProductHandler) GetOrders(w http.ResponseWriter, r *http.Request) {
	logger.Info("Getting orders from WB API")

	data, err := h.wbClient.GetOrders()
	if err != nil {
		logger.Error("Failed to get orders: " + err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

// UpdateStocks обрабатывает запрос на обновление запасов продуктов
func (h *ProductHandler) UpdateStocks(w http.ResponseWriter, r *http.Request) {
	logger.Info("Updating stocks in WB API")

	var stocks interface{}
	if err := json.NewDecoder(r.Body).Decode(&stocks); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	data, err := h.wbClient.UpdateStocks(stocks)
	if err != nil {
		logger.Error("Failed to update stocks: " + err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

// GetAnalytics обрабатывает запрос на получение аналитических данных
func (h *ProductHandler) GetAnalytics(w http.ResponseWriter, r *http.Request) {
	logger.Info("Getting analytics from WB API")

	data, err := h.wbClient.GetAnalytics()
	if err != nil {
		logger.Error("Failed to get analytics: " + err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
