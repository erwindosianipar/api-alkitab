package handler

import (
	"api-alkitab/list"
	"api-alkitab/utils"
	"net/http"

	"github.com/gorilla/mux"
)

// ListHandler use to implement list service
type ListHandler struct {
	listService list.ListService
}

// SetupHandler exported to initialize list handler
func SetupHandler(r *mux.Router, listService list.ListService) {
	listHandler := ListHandler{listService}

	r.HandleFunc("/v2/passage/list", utils.HandleCORS(listHandler.listPassage)).Methods(http.MethodGet)
}

func (h *ListHandler) listPassage(w http.ResponseWriter, r *http.Request) {
	response, err := h.listService.ListPassage()
	if err != nil {
		utils.ResponseMessage(w, http.StatusInternalServerError, utils.Message(http.StatusInternalServerError, err.Error()))
		return
	}

	mapResponse := map[string]interface{}{
		"passage_list": response,
	}

	utils.ResponseData(w, http.StatusOK, mapResponse)
	return
}
