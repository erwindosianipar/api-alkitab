package handler

import (
	"net/http"
	"strconv"

	"api-alkitab/passage"
	"api-alkitab/utils"

	"github.com/gorilla/mux"
)

type PassageHandler struct {
	passageService passage.PassageService
}

func SetupHandler(r *mux.Router, passageService passage.PassageService) {
	passageHandler := PassageHandler{passageService}

	var v1 = r.PathPrefix("/v1").Subrouter()

	v1.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	})

	v1.HandleFunc("/passage/{passage}/{chapter}", passageHandler.passageChapter).Methods(http.MethodGet)
	v1.HandleFunc("/passage/{passage}/{chapter}/{number}", passageHandler.passageChapterNumber).Methods(http.MethodGet)
}

func (h *PassageHandler) passageChapter(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	chapter, err := strconv.Atoi(params["chapter"])
	if err != nil {
		utils.Response(w, http.StatusBadRequest, utils.Message(http.StatusBadRequest, "Chapter must be a valid number"))
		return
	}

	response, err := h.passageService.PassageChapter(params["passage"], chapter)
	if err != nil {
		utils.Response(w, http.StatusInternalServerError, utils.Message(http.StatusInternalServerError, err.Error()))
		return
	}

	utils.Response(w, http.StatusOK, response)
	return
}

func (h *PassageHandler) passageChapterNumber(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	chapter, err := strconv.Atoi(params["chapter"])
	if err != nil {
		utils.Response(w, http.StatusBadRequest, utils.Message(http.StatusBadRequest, "Chapter must be a valid number"))
		return
	}

	number, err := strconv.Atoi(params["number"])
	if err != nil {
		utils.Response(w, http.StatusBadRequest, utils.Message(http.StatusBadRequest, "Number must be a valid number"))
		return
	}

	response, err := h.passageService.PassageChapterNumber(params["passage"], chapter, number)
	if err != nil {
		utils.Response(w, http.StatusInternalServerError, utils.Message(http.StatusInternalServerError, err.Error()))
		return
	}

	utils.Response(w, http.StatusOK, response)
	return
}
