package handler

import (
	"net/http"
	"strconv"

	"api-alkitab/passage"
	"api-alkitab/utils"

	"github.com/gorilla/mux"
)

// PassageHandler use to implement passage service
type PassageHandler struct {
	passageService passage.PassageService
}

// SetupHandler exported to initialize passage handler
func SetupHandler(r *mux.Router, passageService passage.PassageService) {
	passageHandler := PassageHandler{passageService}

	// MARK: - V1

	var v1 = r.PathPrefix("/v1").Subrouter()

	v1.HandleFunc("/passage/{passage}/{chapter}", utils.HandleCORS(passageHandler.passageChapter)).Methods(http.MethodGet)
	v1.HandleFunc("/passage/{passage}/{chapter}/{verse}", utils.HandleCORS(passageHandler.passageChapterVerse)).Methods(http.MethodGet)

	// MARK: - V2

	var v2 = r.PathPrefix("/v2").Subrouter()

	v2.HandleFunc("/passage/{passage}/{chapter}", utils.HandleCORS(passageHandler.passageChapterV2)).Methods(http.MethodGet)
	v2.HandleFunc("/passage/{passage}/{chapter}/{verse}", utils.HandleCORS(passageHandler.passageChapterVerseV2)).Methods(http.MethodGet)

	// MARK: - V3

	var v3 = r.PathPrefix("/v3").Subrouter()

	v3.HandleFunc("/passage/{passage}/{chapter}", utils.HandleCORS(passageHandler.passageChapterV3)).Methods(http.MethodGet)
}

func (h *PassageHandler) passageChapter(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	chapter, err := strconv.Atoi(params["chapter"])
	if err != nil {
		utils.ResponseMessage(w, http.StatusBadRequest, utils.Message(http.StatusBadRequest, utils.ChapterMustValid))
		return
	}

	response, err := h.passageService.PassageChapter(params["passage"], chapter)
	if err != nil {
		utils.ResponseMessage(w, http.StatusInternalServerError, utils.Message(http.StatusInternalServerError, err.Error()))
		return
	}

	utils.ResponseMessage(w, http.StatusOK, response)
}

func (h *PassageHandler) passageChapterVerse(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	chapter, err := strconv.Atoi(params["chapter"])
	if err != nil {
		utils.ResponseMessage(w, http.StatusBadRequest, utils.Message(http.StatusBadRequest, utils.ChapterMustValid))
		return
	}

	verse, err := strconv.Atoi(params["verse"])
	if err != nil {
		utils.ResponseMessage(w, http.StatusBadRequest, utils.Message(http.StatusBadRequest, utils.VerseMustValid))
		return
	}

	response, err := h.passageService.PassageChapterVerse(params["passage"], chapter, verse)
	if err != nil {
		utils.ResponseMessage(w, http.StatusInternalServerError, utils.Message(http.StatusInternalServerError, err.Error()))
		return
	}

	utils.ResponseMessage(w, http.StatusOK, response)
}

func (h *PassageHandler) passageChapterV2(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ver := r.URL.Query().Get("ver")

	if ver == "" || ver == "tb" || ver == "bis" || ver == "net" {
		chapter, err := strconv.Atoi(params["chapter"])
		if err != nil {
			utils.ResponseMessage(w, http.StatusBadRequest, utils.Message(http.StatusBadRequest, utils.ChapterMustValid))
			return
		}

		response, err := h.passageService.PassageChapterV2(params["passage"], chapter, ver)
		if err != nil {
			utils.ResponseMessage(w, http.StatusInternalServerError, utils.Message(http.StatusInternalServerError, err.Error()))
			return
		}

		utils.ResponseData(w, http.StatusOK, response)
		return
	}

	utils.ResponseMessage(w, http.StatusBadRequest, utils.Message(http.StatusBadRequest, utils.ErrorVersionType))
}

func (h *PassageHandler) passageChapterVerseV2(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ver := r.URL.Query().Get("ver")

	if ver == "" || ver == "tb" || ver == "bis" || ver == "net" {
		chapter, err := strconv.Atoi(params["chapter"])
		if err != nil {
			utils.ResponseMessage(w, http.StatusBadRequest, utils.Message(http.StatusBadRequest, utils.ChapterMustValid))
			return
		}

		verse, err := strconv.Atoi(params["verse"])
		if err != nil {
			utils.ResponseMessage(w, http.StatusBadRequest, utils.Message(http.StatusBadRequest, utils.VerseMustValid))
			return
		}

		response, err := h.passageService.PassageChapterVerseV2(params["passage"], chapter, verse, ver)
		if err != nil {
			utils.ResponseMessage(w, http.StatusInternalServerError, utils.Message(http.StatusInternalServerError, err.Error()))
			return
		}

		utils.ResponseData(w, http.StatusOK, response)
		return
	}

	utils.ResponseMessage(w, http.StatusBadRequest, utils.Message(http.StatusBadRequest, utils.ErrorVersionType))
}

func (h *PassageHandler) passageChapterV3(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ver := r.URL.Query().Get("ver")

	if ver == "" || ver == "tb" || ver == "bis" || ver == "net" {
		chapter, err := strconv.Atoi(params["chapter"])
		if err != nil {
			utils.ResponseMessage(w, http.StatusBadRequest, utils.Message(http.StatusBadRequest, utils.ChapterMustValid))
			return
		}

		response, err := h.passageService.PassageChapterV3(params["passage"], chapter, ver)
		if err != nil {
			utils.ResponseMessage(w, http.StatusInternalServerError, utils.Message(http.StatusInternalServerError, err.Error()))
			return
		}

		utils.ResponseData(w, http.StatusOK, response)
		return
	}

	utils.ResponseMessage(w, http.StatusBadRequest, utils.Message(http.StatusBadRequest, utils.ErrorVersionType))
}
