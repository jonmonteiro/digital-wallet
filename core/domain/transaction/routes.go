package transaction

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	middelware "github.com/jmonteiro/picpay-like/core/middleware/auth"
	"github.com/jmonteiro/picpay-like/core/types"
	"github.com/jmonteiro/picpay-like/core/utils"
)

type Handler struct {
	transactionService *TransactionService
	userService        types.UserStore
}

func NewHandler(transactionService *TransactionService, userService types.UserStore) *Handler {
	return &Handler{
		transactionService: transactionService,
		userService:        userService,
	}
}

func (h *Handler) RegisterRoutes(r chi.Router) {
	r.Route("/transactions", func(r chi.Router) {
		r.Post("/", middelware.WithJWTAuth(h.handleMakeTransaction, h.userService))
	})
}

func (h *Handler) handleMakeTransaction(w http.ResponseWriter, r *http.Request) {
	userID := middelware.GetUserIDFromContext(r.Context())
	if userID == -1 {
		utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("user not authenticated"))
		return
	}

	var payload types.TransactionPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", err))
		return
	}

	err := h.transactionService.MakeTransaction(userID, payload)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, map[string]string{"message": "transaction completed successfully"})
}
