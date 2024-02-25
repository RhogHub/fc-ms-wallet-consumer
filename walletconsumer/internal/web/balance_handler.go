package web

import (
	"encoding/json"
	"net/http"

	"github.com/RhogHub/fc-ms-wallet-consumer/internal/usecase/find_balance"
	"github.com/go-chi/chi/v5"
)

type WebBalanceHandler struct {
	FindBalanceUseCase find_balance.FindBalanceUseCase
}

func NewWebFindBalanceHandler(findBalanceUseCase find_balance.FindBalanceUseCase) *WebBalanceHandler {
	return &WebBalanceHandler{
		FindBalanceUseCase: findBalanceUseCase,
	}
}

func (h *WebBalanceHandler) FindBalance(w http.ResponseWriter, r *http.Request) {
	var dto find_balance.FindBalanceInputDTO
	dto.AccountID = chi.URLParam(r, "account_id")

	if dto.AccountID == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	output, err := h.FindBalanceUseCase.Execute(dto)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
